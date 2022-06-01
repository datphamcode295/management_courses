// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package school

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SchoolMetaData contains all meta data concerning the School contract.
var SchoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newLevel\",\"type\":\"uint16\"}],\"name\":\"UpgradeLevel\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Upgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscdAddress\",\"outputs\":[{\"internalType\":\"contractERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscsAddress\",\"outputs\":[{\"internalType\":\"contractERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnNFTWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treausy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_burnWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_BSCSAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_BSCDAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amountBSCS\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountBSCD\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"rateFailure\",\"type\":\"uint16[]\"}],\"name\":\"initData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mstationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_burnNFTWallet\",\"type\":\"address\"}],\"name\":\"updateBurnWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountBSCS\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountBSCD\",\"type\":\"uint256[]\"},{\"internalType\":\"uint16[]\",\"name\":\"rateFailure\",\"type\":\"uint16[]\"}],\"name\":\"updateUpgradeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SchoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SchoolMetaData.ABI instead.
var SchoolABI = SchoolMetaData.ABI

// School is an auto generated Go binding around an Ethereum contract.
type School struct {
	SchoolCaller     // Read-only binding to the contract
	SchoolTransactor // Write-only binding to the contract
	SchoolFilterer   // Log filterer for contract events
}

// SchoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SchoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SchoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SchoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SchoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SchoolSession struct {
	Contract     *School           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SchoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SchoolCallerSession struct {
	Contract *SchoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SchoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SchoolTransactorSession struct {
	Contract     *SchoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SchoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SchoolRaw struct {
	Contract *School // Generic contract binding to access the raw methods on
}

// SchoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SchoolCallerRaw struct {
	Contract *SchoolCaller // Generic read-only contract binding to access the raw methods on
}

// SchoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SchoolTransactorRaw struct {
	Contract *SchoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSchool creates a new instance of School, bound to a specific deployed contract.
func NewSchool(address common.Address, backend bind.ContractBackend) (*School, error) {
	contract, err := bindSchool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &School{SchoolCaller: SchoolCaller{contract: contract}, SchoolTransactor: SchoolTransactor{contract: contract}, SchoolFilterer: SchoolFilterer{contract: contract}}, nil
}

// NewSchoolCaller creates a new read-only instance of School, bound to a specific deployed contract.
func NewSchoolCaller(address common.Address, caller bind.ContractCaller) (*SchoolCaller, error) {
	contract, err := bindSchool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SchoolCaller{contract: contract}, nil
}

// NewSchoolTransactor creates a new write-only instance of School, bound to a specific deployed contract.
func NewSchoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SchoolTransactor, error) {
	contract, err := bindSchool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SchoolTransactor{contract: contract}, nil
}

// NewSchoolFilterer creates a new log filterer instance of School, bound to a specific deployed contract.
func NewSchoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SchoolFilterer, error) {
	contract, err := bindSchool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SchoolFilterer{contract: contract}, nil
}

// bindSchool binds a generic wrapper to an already deployed contract.
func bindSchool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SchoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_School *SchoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _School.Contract.SchoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_School *SchoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _School.Contract.SchoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_School *SchoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _School.Contract.SchoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_School *SchoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _School.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_School *SchoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _School.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_School *SchoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _School.Contract.contract.Transact(opts, method, params...)
}

// BscdAddress is a free data retrieval call binding the contract method 0x2e9b0819.
//
// Solidity: function bscdAddress() view returns(address)
func (_School *SchoolCaller) BscdAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "bscdAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BscdAddress is a free data retrieval call binding the contract method 0x2e9b0819.
//
// Solidity: function bscdAddress() view returns(address)
func (_School *SchoolSession) BscdAddress() (common.Address, error) {
	return _School.Contract.BscdAddress(&_School.CallOpts)
}

// BscdAddress is a free data retrieval call binding the contract method 0x2e9b0819.
//
// Solidity: function bscdAddress() view returns(address)
func (_School *SchoolCallerSession) BscdAddress() (common.Address, error) {
	return _School.Contract.BscdAddress(&_School.CallOpts)
}

// BscsAddress is a free data retrieval call binding the contract method 0xb1acc0c6.
//
// Solidity: function bscsAddress() view returns(address)
func (_School *SchoolCaller) BscsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "bscsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BscsAddress is a free data retrieval call binding the contract method 0xb1acc0c6.
//
// Solidity: function bscsAddress() view returns(address)
func (_School *SchoolSession) BscsAddress() (common.Address, error) {
	return _School.Contract.BscsAddress(&_School.CallOpts)
}

// BscsAddress is a free data retrieval call binding the contract method 0xb1acc0c6.
//
// Solidity: function bscsAddress() view returns(address)
func (_School *SchoolCallerSession) BscsAddress() (common.Address, error) {
	return _School.Contract.BscsAddress(&_School.CallOpts)
}

// BurnNFTWallet is a free data retrieval call binding the contract method 0xabe2f3d0.
//
// Solidity: function burnNFTWallet() view returns(address)
func (_School *SchoolCaller) BurnNFTWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "burnNFTWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnNFTWallet is a free data retrieval call binding the contract method 0xabe2f3d0.
//
// Solidity: function burnNFTWallet() view returns(address)
func (_School *SchoolSession) BurnNFTWallet() (common.Address, error) {
	return _School.Contract.BurnNFTWallet(&_School.CallOpts)
}

// BurnNFTWallet is a free data retrieval call binding the contract method 0xabe2f3d0.
//
// Solidity: function burnNFTWallet() view returns(address)
func (_School *SchoolCallerSession) BurnNFTWallet() (common.Address, error) {
	return _School.Contract.BurnNFTWallet(&_School.CallOpts)
}

// BurnWallet is a free data retrieval call binding the contract method 0x06228749.
//
// Solidity: function burnWallet() view returns(address)
func (_School *SchoolCaller) BurnWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "burnWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnWallet is a free data retrieval call binding the contract method 0x06228749.
//
// Solidity: function burnWallet() view returns(address)
func (_School *SchoolSession) BurnWallet() (common.Address, error) {
	return _School.Contract.BurnWallet(&_School.CallOpts)
}

// BurnWallet is a free data retrieval call binding the contract method 0x06228749.
//
// Solidity: function burnWallet() view returns(address)
func (_School *SchoolCallerSession) BurnWallet() (common.Address, error) {
	return _School.Contract.BurnWallet(&_School.CallOpts)
}

// GetLevel is a free data retrieval call binding the contract method 0x86481d40.
//
// Solidity: function getLevel(uint256 _tokenId) view returns(uint256)
func (_School *SchoolCaller) GetLevel(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "getLevel", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevel is a free data retrieval call binding the contract method 0x86481d40.
//
// Solidity: function getLevel(uint256 _tokenId) view returns(uint256)
func (_School *SchoolSession) GetLevel(_tokenId *big.Int) (*big.Int, error) {
	return _School.Contract.GetLevel(&_School.CallOpts, _tokenId)
}

// GetLevel is a free data retrieval call binding the contract method 0x86481d40.
//
// Solidity: function getLevel(uint256 _tokenId) view returns(uint256)
func (_School *SchoolCallerSession) GetLevel(_tokenId *big.Int) (*big.Int, error) {
	return _School.Contract.GetLevel(&_School.CallOpts, _tokenId)
}

// MstationAddress is a free data retrieval call binding the contract method 0x8be939da.
//
// Solidity: function mstationAddress() view returns(address)
func (_School *SchoolCaller) MstationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "mstationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MstationAddress is a free data retrieval call binding the contract method 0x8be939da.
//
// Solidity: function mstationAddress() view returns(address)
func (_School *SchoolSession) MstationAddress() (common.Address, error) {
	return _School.Contract.MstationAddress(&_School.CallOpts)
}

// MstationAddress is a free data retrieval call binding the contract method 0x8be939da.
//
// Solidity: function mstationAddress() view returns(address)
func (_School *SchoolCallerSession) MstationAddress() (common.Address, error) {
	return _School.Contract.MstationAddress(&_School.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_School *SchoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_School *SchoolSession) Owner() (common.Address, error) {
	return _School.Contract.Owner(&_School.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_School *SchoolCallerSession) Owner() (common.Address, error) {
	return _School.Contract.Owner(&_School.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_School *SchoolCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _School.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_School *SchoolSession) Treasury() (common.Address, error) {
	return _School.Contract.Treasury(&_School.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_School *SchoolCallerSession) Treasury() (common.Address, error) {
	return _School.Contract.Treasury(&_School.CallOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0x318d2be9.
//
// Solidity: function Upgrade(address _nftAddress, uint256 _tokenId) payable returns()
func (_School *SchoolTransactor) Upgrade(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "Upgrade", _nftAddress, _tokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x318d2be9.
//
// Solidity: function Upgrade(address _nftAddress, uint256 _tokenId) payable returns()
func (_School *SchoolSession) Upgrade(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _School.Contract.Upgrade(&_School.TransactOpts, _nftAddress, _tokenId)
}

// Upgrade is a paid mutator transaction binding the contract method 0x318d2be9.
//
// Solidity: function Upgrade(address _nftAddress, uint256 _tokenId) payable returns()
func (_School *SchoolTransactorSession) Upgrade(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _School.Contract.Upgrade(&_School.TransactOpts, _nftAddress, _tokenId)
}

// InitData is a paid mutator transaction binding the contract method 0x06c31f50.
//
// Solidity: function initData(address _treausy, address _burnWallet, address _BSCSAddress, address _BSCDAddress, uint256[] amountBSCS, uint256[] amountBSCD, uint16[] rateFailure) returns()
func (_School *SchoolTransactor) InitData(opts *bind.TransactOpts, _treausy common.Address, _burnWallet common.Address, _BSCSAddress common.Address, _BSCDAddress common.Address, amountBSCS []*big.Int, amountBSCD []*big.Int, rateFailure []uint16) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "initData", _treausy, _burnWallet, _BSCSAddress, _BSCDAddress, amountBSCS, amountBSCD, rateFailure)
}

// InitData is a paid mutator transaction binding the contract method 0x06c31f50.
//
// Solidity: function initData(address _treausy, address _burnWallet, address _BSCSAddress, address _BSCDAddress, uint256[] amountBSCS, uint256[] amountBSCD, uint16[] rateFailure) returns()
func (_School *SchoolSession) InitData(_treausy common.Address, _burnWallet common.Address, _BSCSAddress common.Address, _BSCDAddress common.Address, amountBSCS []*big.Int, amountBSCD []*big.Int, rateFailure []uint16) (*types.Transaction, error) {
	return _School.Contract.InitData(&_School.TransactOpts, _treausy, _burnWallet, _BSCSAddress, _BSCDAddress, amountBSCS, amountBSCD, rateFailure)
}

// InitData is a paid mutator transaction binding the contract method 0x06c31f50.
//
// Solidity: function initData(address _treausy, address _burnWallet, address _BSCSAddress, address _BSCDAddress, uint256[] amountBSCS, uint256[] amountBSCD, uint16[] rateFailure) returns()
func (_School *SchoolTransactorSession) InitData(_treausy common.Address, _burnWallet common.Address, _BSCSAddress common.Address, _BSCDAddress common.Address, amountBSCS []*big.Int, amountBSCD []*big.Int, rateFailure []uint16) (*types.Transaction, error) {
	return _School.Contract.InitData(&_School.TransactOpts, _treausy, _burnWallet, _BSCSAddress, _BSCDAddress, amountBSCS, amountBSCD, rateFailure)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_School *SchoolTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_School *SchoolSession) Initialize() (*types.Transaction, error) {
	return _School.Contract.Initialize(&_School.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_School *SchoolTransactorSession) Initialize() (*types.Transaction, error) {
	return _School.Contract.Initialize(&_School.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_School *SchoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_School *SchoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _School.Contract.RenounceOwnership(&_School.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_School *SchoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _School.Contract.RenounceOwnership(&_School.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_School *SchoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_School *SchoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _School.Contract.TransferOwnership(&_School.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_School *SchoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _School.Contract.TransferOwnership(&_School.TransactOpts, newOwner)
}

// UpdateBurnWallet is a paid mutator transaction binding the contract method 0xb3839389.
//
// Solidity: function updateBurnWallet(address _burnNFTWallet) returns()
func (_School *SchoolTransactor) UpdateBurnWallet(opts *bind.TransactOpts, _burnNFTWallet common.Address) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "updateBurnWallet", _burnNFTWallet)
}

// UpdateBurnWallet is a paid mutator transaction binding the contract method 0xb3839389.
//
// Solidity: function updateBurnWallet(address _burnNFTWallet) returns()
func (_School *SchoolSession) UpdateBurnWallet(_burnNFTWallet common.Address) (*types.Transaction, error) {
	return _School.Contract.UpdateBurnWallet(&_School.TransactOpts, _burnNFTWallet)
}

// UpdateBurnWallet is a paid mutator transaction binding the contract method 0xb3839389.
//
// Solidity: function updateBurnWallet(address _burnNFTWallet) returns()
func (_School *SchoolTransactorSession) UpdateBurnWallet(_burnNFTWallet common.Address) (*types.Transaction, error) {
	return _School.Contract.UpdateBurnWallet(&_School.TransactOpts, _burnNFTWallet)
}

// UpdateUpgradeFee is a paid mutator transaction binding the contract method 0xf8730894.
//
// Solidity: function updateUpgradeFee(uint256[] amountBSCS, uint256[] amountBSCD, uint16[] rateFailure) returns()
func (_School *SchoolTransactor) UpdateUpgradeFee(opts *bind.TransactOpts, amountBSCS []*big.Int, amountBSCD []*big.Int, rateFailure []uint16) (*types.Transaction, error) {
	return _School.contract.Transact(opts, "updateUpgradeFee", amountBSCS, amountBSCD, rateFailure)
}

// UpdateUpgradeFee is a paid mutator transaction binding the contract method 0xf8730894.
//
// Solidity: function updateUpgradeFee(uint256[] amountBSCS, uint256[] amountBSCD, uint16[] rateFailure) returns()
func (_School *SchoolSession) UpdateUpgradeFee(amountBSCS []*big.Int, amountBSCD []*big.Int, rateFailure []uint16) (*types.Transaction, error) {
	return _School.Contract.UpdateUpgradeFee(&_School.TransactOpts, amountBSCS, amountBSCD, rateFailure)
}

// UpdateUpgradeFee is a paid mutator transaction binding the contract method 0xf8730894.
//
// Solidity: function updateUpgradeFee(uint256[] amountBSCS, uint256[] amountBSCD, uint16[] rateFailure) returns()
func (_School *SchoolTransactorSession) UpdateUpgradeFee(amountBSCS []*big.Int, amountBSCD []*big.Int, rateFailure []uint16) (*types.Transaction, error) {
	return _School.Contract.UpdateUpgradeFee(&_School.TransactOpts, amountBSCS, amountBSCD, rateFailure)
}

// SchoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the School contract.
type SchoolOwnershipTransferredIterator struct {
	Event *SchoolOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SchoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchoolOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SchoolOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SchoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchoolOwnershipTransferred represents a OwnershipTransferred event raised by the School contract.
type SchoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_School *SchoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SchoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _School.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SchoolOwnershipTransferredIterator{contract: _School.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_School *SchoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SchoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _School.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchoolOwnershipTransferred)
				if err := _School.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_School *SchoolFilterer) ParseOwnershipTransferred(log types.Log) (*SchoolOwnershipTransferred, error) {
	event := new(SchoolOwnershipTransferred)
	if err := _School.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SchoolUpgradeLevelIterator is returned from FilterUpgradeLevel and is used to iterate over the raw logs and unpacked data for UpgradeLevel events raised by the School contract.
type SchoolUpgradeLevelIterator struct {
	Event *SchoolUpgradeLevel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SchoolUpgradeLevelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SchoolUpgradeLevel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SchoolUpgradeLevel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SchoolUpgradeLevelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SchoolUpgradeLevelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SchoolUpgradeLevel represents a UpgradeLevel event raised by the School contract.
type SchoolUpgradeLevel struct {
	NftAddress common.Address
	TokenId    *big.Int
	Result     bool
	NewLevel   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpgradeLevel is a free log retrieval operation binding the contract event 0x12d3017676fe53f22ca1d28f832b4f698fc63923a79ef7f9376a2edb83cb9b2a.
//
// Solidity: event UpgradeLevel(address _nftAddress, uint256 _tokenId, bool result, uint16 newLevel)
func (_School *SchoolFilterer) FilterUpgradeLevel(opts *bind.FilterOpts) (*SchoolUpgradeLevelIterator, error) {

	logs, sub, err := _School.contract.FilterLogs(opts, "UpgradeLevel")
	if err != nil {
		return nil, err
	}
	return &SchoolUpgradeLevelIterator{contract: _School.contract, event: "UpgradeLevel", logs: logs, sub: sub}, nil
}

// WatchUpgradeLevel is a free log subscription operation binding the contract event 0x12d3017676fe53f22ca1d28f832b4f698fc63923a79ef7f9376a2edb83cb9b2a.
//
// Solidity: event UpgradeLevel(address _nftAddress, uint256 _tokenId, bool result, uint16 newLevel)
func (_School *SchoolFilterer) WatchUpgradeLevel(opts *bind.WatchOpts, sink chan<- *SchoolUpgradeLevel) (event.Subscription, error) {

	logs, sub, err := _School.contract.WatchLogs(opts, "UpgradeLevel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SchoolUpgradeLevel)
				if err := _School.contract.UnpackLog(event, "UpgradeLevel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgradeLevel is a log parse operation binding the contract event 0x12d3017676fe53f22ca1d28f832b4f698fc63923a79ef7f9376a2edb83cb9b2a.
//
// Solidity: event UpgradeLevel(address _nftAddress, uint256 _tokenId, bool result, uint16 newLevel)
func (_School *SchoolFilterer) ParseUpgradeLevel(log types.Log) (*SchoolUpgradeLevel, error) {
	event := new(SchoolUpgradeLevel)
	if err := _School.contract.UnpackLog(event, "UpgradeLevel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
