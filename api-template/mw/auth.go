package mw

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/consts"
	"github.com/MStation-io/api/model/errors"
	"github.com/MStation-io/api/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

type AuthMiddleware struct {
	cfg config.Config
}

// NewAuthMiddleware ...
func NewAuthMiddleware(cfg config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		cfg: cfg,
	}
}

func (a *AuthMiddleware) WithSignature(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		zap.L().Sugar().Infof("[mw.WithSignature] authenticattion challenge: %s, signature: %s", c.Request().Header.Get("Challenge"), c.Request().Header.Get("Signature"))
		err := authenticateChallenge(c)
		if err != nil {
			return util.HandleError(c, err)
		}

		err = verifySignature(c)
		if err != nil {
			return util.HandleError(c, err)
		}
		return next(c)
	}
}

func authenticateChallenge(c echo.Context) error {
	challenge, err := util.GetChallengeInfoFromContext(c)
	if err != nil {
		zap.L().Sugar().Errorf("[mw.authenticateChallenge] util.GetChallengeInfoFromContext()")
		return err
	}
	if challenge.ExpiresAt < time.Now().Unix() {
		return errors.ErrExpiredChallenge
	}
	return nil
}

func verifySignature(c echo.Context) error {
	challenge := c.Request().Header.Get("Challenge")
	signature := c.Request().Header.Get("Signature")
	challengeInfo, err := util.GetChallengeInfoFromContext(c)
	if err != nil {
		zap.L().Sugar().Errorf("[mw.verifySignature] util.GetChallengeInfoFromContext()")
		return err
	}

	signatureBytes, err := hex.DecodeString(signature[2:])
	if err != nil {
		zap.L().Sugar().Errorf("[mw.verifySignature]  hex.DecodeString()")
		return errors.ErrInvalidSignature
	}
	if len(signatureBytes) != consts.DefaultSignatureLengthWithRecoveryId {
		zap.L().Sugar().Errorf("[mw.verifySignature] len(signatureBytes) != 65")
		return errors.ErrInvalidSignature
	}
	//web3, etherjs: signature[64] = v instead of recoveryParam
	//more info: https://github.com/ethers-io/ethers.js/issues/823#issuecomment-625953096
	if signatureBytes[64] == 27 {
		signatureBytes[64] = 0
	} else if signatureBytes[64] == 28 {
		signatureBytes[64] = 1
	} else {
		zap.L().Sugar().Errorf("[mw.verifySignature] signatureBytes[64] != (27,28)")
		return errors.ErrInvalidSignature
	}
	//web3, etherjs sign message with prefix
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(challenge), challenge)

	hashedChallenge := crypto.Keccak256Hash([]byte(prefix))
	sigPublicKeyECDSA, err := crypto.SigToPub(hashedChallenge.Bytes(), signatureBytes)
	if err != nil {
		zap.L().Sugar().Errorf("[mw.verifySignature] scrypto.SigToPub()")
		return errors.ErrInvalidSignature
	}
	address := crypto.PubkeyToAddress(*sigPublicKeyECDSA)

	//validate address
	if address.Hex() != challengeInfo.Address {
		zap.L().Sugar().Errorf("[mw.verifySignature] wrong address")
		return errors.ErrInvalidSignature
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	signatureNoRecoverID := signatureBytes[:len(signatureBytes)-1]
	verified := crypto.VerifySignature(sigPublicKeyBytes, hashedChallenge.Bytes(), signatureNoRecoverID)
	if !verified {
		zap.L().Sugar().Errorf("[mw.verifySignature] crypto.VerifySignature()")
		return errors.ErrInvalidSignature
	}
	return nil
}
