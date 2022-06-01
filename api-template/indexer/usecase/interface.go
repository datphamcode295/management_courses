package usecase

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type IUsecase interface {
	// Topics return list of topics this usecase listening for
	Topics() []common.Hash
	//
	Process(log types.Log) error
	//
	ShouldProcessLog(log types.Log) bool
	//
	Name() string
	//
	ContractAddress() common.Address
}
