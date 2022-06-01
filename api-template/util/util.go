package util

import (
	"math/big"

	"github.com/MStation-io/api/model/errors"
	"go.uber.org/zap"
)

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func ParsePrice(priceStr string) (float64, error) {
	FloatBasePrice, ok := new(big.Float).SetString("1000000000000000000")
	if !ok {
		zap.L().Sugar().Errorf("[util.ParsePrice] new(big.Float).SetString() base price")
		return 0, errors.ErrInternalServerError
	}
	FloatPrice, ok := new(big.Float).SetString(priceStr)
	if !ok {
		zap.L().Sugar().Errorf("[util.ParsePrice] nnew(big.Float).SetString()")
		return 0, errors.ErrInvalidPrice
	}
	Price, _ := FloatPrice.Quo(FloatPrice, FloatBasePrice).Float64()
	return Price, nil
}

func WeiToUnit(wei *big.Int, decimals uint8) (float64, big.Accuracy) {
	floatValue := new(big.Float).SetInt(wei)
	i, exp := big.NewInt(10), big.NewInt(int64(decimals))
	i.Exp(i, exp, nil)
	floatExp := new(big.Float).SetInt(i)
	value, accuracy := floatValue.Quo(floatValue, floatExp).Float64()

	return value, accuracy
}
