// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Faucet

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
	_ = abi.ConvertType
)

// FaucetMetaData contains all meta data concerning the Faucet contract.
var FaucetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveer\",\"type\":\"address\"}],\"name\":\"sendToken\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c060405267016345785d8a00006002556040516108fe3803806108fe833981810160405281019061003191906100fd565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff168152505050610128565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100cc826100a3565b9050919050565b6100dc816100c2565b81146100e6575f5ffd5b50565b5f815190506100f7816100d3565b92915050565b5f602082840312156101125761011161009f565b5b5f61011f848285016100e9565b91505092915050565b60805160a0516107a86101565f395f8181610248015261029e01525f818160cf015261015f01526107a85ff3fe60806040526004361061003e575f3560e01c806312065fe0146100425780633ccfd60b1461006c578063c575c2381461008a578063d0e30db0146100ba575b5f5ffd5b34801561004d575f5ffd5b506100566100c4565b6040516100639190610490565b60405180910390f35b6100746100cb565b6040516100819190610519565b60405180910390f35b6100a4600480360381019061009f9190610597565b61022d565b6040516100b19190610519565b60405180910390f35b6100c2610476565b005b5f47905090565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461015b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101529061061c565b60405180910390fd5b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16476040516101a19061068e565b5f6040518083038185875af1925050503d805f81146101db576040519150601f19603f3d011682016040523d82523d5f602084013e6101e0565b606091505b509150915081610225576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021c906106ec565b60405180910390fd5b809250505090565b60603373ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff161480156102d657508173ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1614155b610315576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030c90610754565b60405180910390fd5b5f15155f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615151461036c575f5ffd5b5f5f8373ffffffffffffffffffffffffffffffffffffffff166002546040516103949061068e565b5f6040518083038185875af1925050503d805f81146103ce576040519150601f19603f3d011682016040523d82523d5f602084013e6103d3565b606091505b509150915081610418576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040f906106ec565b60405180910390fd5b60015f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508092505050919050565b565b5f819050919050565b61048a81610478565b82525050565b5f6020820190506104a35f830184610481565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6104eb826104a9565b6104f581856104b3565b93506105058185602086016104c3565b61050e816104d1565b840191505092915050565b5f6020820190508181035f83015261053181846104e1565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6105668261053d565b9050919050565b6105768161055c565b8114610580575f5ffd5b50565b5f813590506105918161056d565b92915050565b5f602082840312156105ac576105ab610539565b5b5f6105b984828501610583565b91505092915050565b5f82825260208201905092915050565b7f4f6e6c79204f776e65722063616e2077697468647261770000000000000000005f82015250565b5f6106066017836105c2565b9150610611826105d2565b602082019050919050565b5f6020820190508181035f830152610633816105fa565b9050919050565b5f81905092915050565b7f5769746864726177205375636365737366756c000000000000000000000000005f82015250565b5f61067860138361063a565b915061068382610644565b601382019050919050565b5f6106988261066c565b9150819050919050565b7f4661696c656420746f20776974686472617700000000000000000000000000005f82015250565b5f6106d66012836105c2565b91506106e1826106a2565b602082019050919050565b5f6020820190508181035f830152610703816106ca565b9050919050565b7f4572726f720000000000000000000000000000000000000000000000000000005f82015250565b5f61073e6005836105c2565b91506107498261070a565b602082019050919050565b5f6020820190508181035f83015261076b81610732565b905091905056fea2646970667358221220c0169e00aba431d8455052a5f61ffd17aa60452f384cfded4516ea56491b0f6964736f6c634300081b0033",
}

// FaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use FaucetMetaData.ABI instead.
var FaucetABI = FaucetMetaData.ABI

// FaucetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FaucetMetaData.Bin instead.
var FaucetBin = FaucetMetaData.Bin

// DeployFaucet deploys a new Ethereum contract, binding an instance of Faucet to it.
func DeployFaucet(auth *bind.TransactOpts, backend bind.ContractBackend, _signer common.Address) (common.Address, *types.Transaction, *Faucet, error) {
	parsed, err := FaucetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FaucetBin), backend, _signer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// Faucet is an auto generated Go binding around an Ethereum contract.
type Faucet struct {
	FaucetCaller     // Read-only binding to the contract
	FaucetTransactor // Write-only binding to the contract
	FaucetFilterer   // Log filterer for contract events
}

// FaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetSession struct {
	Contract     *Faucet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetCallerSession struct {
	Contract *FaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTransactorSession struct {
	Contract     *FaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetRaw struct {
	Contract *Faucet // Generic contract binding to access the raw methods on
}

// FaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetCallerRaw struct {
	Contract *FaucetCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTransactorRaw struct {
	Contract *FaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucet creates a new instance of Faucet, bound to a specific deployed contract.
func NewFaucet(address common.Address, backend bind.ContractBackend) (*Faucet, error) {
	contract, err := bindFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// NewFaucetCaller creates a new read-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetCaller(address common.Address, caller bind.ContractCaller) (*FaucetCaller, error) {
	contract, err := bindFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetCaller{contract: contract}, nil
}

// NewFaucetTransactor creates a new write-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTransactor, error) {
	contract, err := bindFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTransactor{contract: contract}, nil
}

// NewFaucetFilterer creates a new log filterer instance of Faucet, bound to a specific deployed contract.
func NewFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetFilterer, error) {
	contract, err := bindFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetFilterer{contract: contract}, nil
}

// bindFaucet binds a generic wrapper to an already deployed contract.
func bindFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.FaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Faucet *FaucetCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Faucet *FaucetSession) GetBalance() (*big.Int, error) {
	return _Faucet.Contract.GetBalance(&_Faucet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Faucet *FaucetCallerSession) GetBalance() (*big.Int, error) {
	return _Faucet.Contract.GetBalance(&_Faucet.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Faucet *FaucetTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Faucet *FaucetSession) Deposit() (*types.Transaction, error) {
	return _Faucet.Contract.Deposit(&_Faucet.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Faucet *FaucetTransactorSession) Deposit() (*types.Transaction, error) {
	return _Faucet.Contract.Deposit(&_Faucet.TransactOpts)
}

// SendToken is a paid mutator transaction binding the contract method 0xc575c238.
//
// Solidity: function sendToken(address _receiveer) payable returns(bytes)
func (_Faucet *FaucetTransactor) SendToken(opts *bind.TransactOpts, _receiveer common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "sendToken", _receiveer)
}

// SendToken is a paid mutator transaction binding the contract method 0xc575c238.
//
// Solidity: function sendToken(address _receiveer) payable returns(bytes)
func (_Faucet *FaucetSession) SendToken(_receiveer common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SendToken(&_Faucet.TransactOpts, _receiveer)
}

// SendToken is a paid mutator transaction binding the contract method 0xc575c238.
//
// Solidity: function sendToken(address _receiveer) payable returns(bytes)
func (_Faucet *FaucetTransactorSession) SendToken(_receiveer common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SendToken(&_Faucet.TransactOpts, _receiveer)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns(bytes)
func (_Faucet *FaucetTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns(bytes)
func (_Faucet *FaucetSession) Withdraw() (*types.Transaction, error) {
	return _Faucet.Contract.Withdraw(&_Faucet.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns(bytes)
func (_Faucet *FaucetTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Faucet.Contract.Withdraw(&_Faucet.TransactOpts)
}
