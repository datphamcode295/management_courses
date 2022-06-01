package handler

import (
	b64 "encoding/base64"
	"encoding/json"
	inerr "errors"
	"net/http"
	"time"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type responseData struct {
	Message string `json:"message"`
}

type ChallengeData struct {
	Challenge string `json:"challenge"`
}
type getChallengeResponse struct {
	Data ChallengeData `json:"data"`
}
type upsertWalletResponse struct {
	Data responseData `json:"data"`
}

// UpsertWalletHandler
// @Summary upsert wallet address
// @Description upsert wallet address
// @Tags Authentication
// @Accept	json
// @Produce  json
// @Param walletAddress query string true "wallet address"
// @Security ApiKeyAuth
// @Success 200 {object} handler.upsertWalletResponse	"ok"
// @Router /api/v1/auth/upsert-wallet [post]
func (h *Handler) UpsertWalletHandler(c echo.Context) error {
	walletAddress := c.QueryParam("walletAddress")

	if !util.IsValidAddress(walletAddress) {
		zap.L().Sugar().Errorf("[handler.UpsertWalletHandler] util.IsValidAddress()")
		return util.HandleError(c, errors.ErrInvalidAddress)
	}

	// check if wallet address existed
	_, err := h.repo.User.GetByWalletAddress(h.store, walletAddress)
	if err != nil {
		if inerr.Is(err, gorm.ErrRecordNotFound) {
			userParam := model.User{
				WalletAddress: walletAddress,
				ReferredBy:    "",
			}

			_, err := h.repo.User.Create(h.store, userParam)
			if err != nil {
				zap.L().Sugar().Errorf("[handler.UpsertWalletHandler] User.Create()")
				return util.HandleError(c, err)
			}
		} else {
			zap.L().Sugar().Errorf("[handler.UpsertWalletHandler] User.CheckExistedByWalletAddress()")
			return util.HandleError(c, errors.ErrInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, &upsertWalletResponse{Data: responseData{Message: "Success"}})
}

// GetChallengeHandler
// @Summary get challenge
// @Description get challenge
// @Tags Authentication
// @Accept	json
// @Produce  json
// @Security ApiKeyAuth
// @Param address query string true "wallet address"
// @Success 200 {object} handler.getChallengeResponse	"ok"
// @Failure 400 {object} errors.Error
// @Router /api/v1/auth/challenge [get]
func (h *Handler) GetChallengeHandler(c echo.Context) error {
	address := c.QueryParam("address")

	user, err := h.repo.User.GetByWalletAddress(h.store, address)
	if err != nil {
		if inerr.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Sugar().Errorf("[handler.GetChallengeHandler] User.GetByWalletAddress() user not found")
			return util.HandleError(c, errors.ErrUserNotfound)
		}
		zap.L().Sugar().Errorf("[handler.GetChallengeHandler] User.GetByWalletAddress()")
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	expiredAt := time.Now().Add(time.Minute * 60).Unix()
	challenge := model.ChallengeInfo{
		ExpiresAt: expiredAt,
		Address:   user.WalletAddress,
	}
	challengeStr, err := json.Marshal(challenge)
	if err != nil {
		zap.L().Sugar().Errorf("[handler.GetChallengeHandler] json.Marshal()")
		return util.HandleError(c, errors.ErrInternalServerError)
	}
	encodedChallenge := b64.StdEncoding.EncodeToString(challengeStr)
	return c.JSON(http.StatusOK, &getChallengeResponse{Data: ChallengeData{Challenge: encodedChallenge}})
}
