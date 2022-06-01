package util

import (
	b64 "encoding/base64"
	"encoding/json"
	"regexp"

	"github.com/MStation-io/api/model"
	"github.com/MStation-io/api/model/errors"
	"github.com/labstack/echo/v4"
)

func GetChallengeInfoFromContext(c echo.Context) (*model.ChallengeInfo, error) {
	challengeRequest := c.Request().Header.Get("Challenge")
	if challengeRequest == "" {
		return nil, errors.ErrInvalidChallenge
	}
	rawChallenge, err := b64.StdEncoding.DecodeString(challengeRequest)
	if err != nil {
		return nil, errors.ErrInvalidChallenge
	}
	var challenge model.ChallengeInfo
	err = json.Unmarshal(rawChallenge, &challenge)
	if err != nil {
		return nil, errors.ErrInvalidChallenge
	}
	return &challenge, nil
}

func IsValidAddress(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}
