package log

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type UserError struct {
	Type          string      `json:"type"`
	Error         error       `json:"error"`
	WalletAddress string      `json:"walletAddress"`
	RequestBody   interface{} `json:"requestBody"`
	Header        http.Header `json:"header"`
}

func TraceUserError(context echo.Context, err error, address string, requestBody interface{}) string {
	userError := &UserError{
		Type:          "TraceUser",
		Error:         err,
		WalletAddress: address,
		RequestBody:   requestBody,
		Header:        context.Request().Header,
	}
	userErrorData, err := json.Marshal(userError)
	if err != nil {
		zap.L().Sugar().Errorf("json.Marshal(): %v", err)
		return ""
	}

	return string(userErrorData)
}
