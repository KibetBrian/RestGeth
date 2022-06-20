// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Election

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

// ElectionCandidate is an auto generated low-level Go binding around an user-defined struct.
type ElectionCandidate struct {
	Name string
	Addr common.Address
}

// ElectionVote is an auto generated low-level Go binding around an user-defined struct.
type ElectionVote struct {
	Addr common.Address
}

// ElectionVoter is an auto generated low-level Go binding around an user-defined struct.
type ElectionVoter struct {
	Addr common.Address
}

// ElectionMetaData contains all meta data concerning the Election contract.
var ElectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AddVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVoters\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structElection.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"RegisterCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"SeeCandidateVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structElection.Vote[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060065534801561001557600080fd5b5061123e806100256000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638807f2711161005b5780638807f2711461013b5780638f7a842d1461016b5780639763884d146101895780639e1fa02b146101a757610088565b80630ad2eca71461008d5780633e77583b146100bd57806343ad4498146100db57806371ff924f1461010b575b600080fd5b6100a760048036038101906100a29190610a72565b6101d7565b6040516100b49190610aba565b60405180910390f35b6100c56102d5565b6040516100d29190610aee565b60405180910390f35b6100f560048036038101906100f09190610a72565b6102e2565b6040516101029190610be3565b60405180910390f35b61012560048036038101906101209190610a72565b6103d1565b6040516101329190610aba565b60405180910390f35b61015560048036038101906101509190610c31565b61055d565b6040516101629190610c80565b60405180910390f35b6101736105b4565b6040516101809190610d66565b60405180910390f35b610191610665565b60405161019e9190610f20565b60405180910390f35b6101c160048036038101906101bc9190611077565b6107ac565b6040516101ce9190610aba565b60405180910390f35b6000600654600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600460405180602001604052808473ffffffffffffffffffffffffffffffffffffffff168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600660008154809291906102c790611102565b919050555060019050919050565b6000600480549050905090565b60606000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156103c6578382906000526020600020016040518060200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505081526020019060010190610342565b505050509050919050565b600080600090505b600180549050811015610552578273ffffffffffffffffffffffffffffffffffffffff16600182815481106104115761041061114a565b5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361053f57600060405180602001604052803373ffffffffffffffffffffffffffffffffffffffff1681525090506000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600192505050610558565b808061054a90611102565b9150506103d9565b50600090505b919050565b6000602052816000526040600020818154811061057957600080fd5b90600052602060002001600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081565b60606004805480602002602001604051908101604052809291908181526020016000905b8282101561065c578382906000526020600020016040518060200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050815260200190600101906105d8565b50505050905090565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156107a357838290600052602060002090600202016040518060400160405290816000820180546106bc906111a8565b80601f01602080910402602001604051908101604052809291908181526020018280546106e8906111a8565b80156107355780601f1061070a57610100808354040283529160200191610735565b820191906000526020600020905b81548152906001019060200180831161071857829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505081526020019060010190610689565b50505050905090565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001816001815401808255809150500390600052602060002050506000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480610855576108546111d9565b5b60019003818190600052602060002001600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550509055600060405180604001604052808481526020018573ffffffffffffffffffffffffffffffffffffffff1681525090506001819080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001908051906020019061090892919061095d565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600191505092915050565b828054610969906111a8565b90600052602060002090601f01602090048101928261098b57600085556109d2565b82601f106109a457805160ff19168380011785556109d2565b828001600101855582156109d2579182015b828111156109d15782518255916020019190600101906109b6565b5b5090506109df91906109e3565b5090565b5b808211156109fc5760008160009055506001016109e4565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a3f82610a14565b9050919050565b610a4f81610a34565b8114610a5a57600080fd5b50565b600081359050610a6c81610a46565b92915050565b600060208284031215610a8857610a87610a0a565b5b6000610a9684828501610a5d565b91505092915050565b60008115159050919050565b610ab481610a9f565b82525050565b6000602082019050610acf6000830184610aab565b92915050565b6000819050919050565b610ae881610ad5565b82525050565b6000602082019050610b036000830184610adf565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610b3e81610a34565b82525050565b602082016000820151610b5a6000850182610b35565b50505050565b6000610b6c8383610b44565b60208301905092915050565b6000602082019050919050565b6000610b9082610b09565b610b9a8185610b14565b9350610ba583610b25565b8060005b83811015610bd6578151610bbd8882610b60565b9750610bc883610b78565b925050600181019050610ba9565b5085935050505092915050565b60006020820190508181036000830152610bfd8184610b85565b905092915050565b610c0e81610ad5565b8114610c1957600080fd5b50565b600081359050610c2b81610c05565b92915050565b60008060408385031215610c4857610c47610a0a565b5b6000610c5685828601610a5d565b9250506020610c6785828601610c1c565b9150509250929050565b610c7a81610a34565b82525050565b6000602082019050610c956000830184610c71565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b602082016000820151610cdd6000850182610b35565b50505050565b6000610cef8383610cc7565b60208301905092915050565b6000602082019050919050565b6000610d1382610c9b565b610d1d8185610ca6565b9350610d2883610cb7565b8060005b83811015610d59578151610d408882610ce3565b9750610d4b83610cfb565b925050600181019050610d2c565b5085935050505092915050565b60006020820190508181036000830152610d808184610d08565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610dee578082015181840152602081019050610dd3565b83811115610dfd576000848401525b50505050565b6000601f19601f8301169050919050565b6000610e1f82610db4565b610e298185610dbf565b9350610e39818560208601610dd0565b610e4281610e03565b840191505092915050565b60006040830160008301518482036000860152610e6a8282610e14565b9150506020830151610e7f6020860182610b35565b508091505092915050565b6000610e968383610e4d565b905092915050565b6000602082019050919050565b6000610eb682610d88565b610ec08185610d93565b935083602082028501610ed285610da4565b8060005b85811015610f0e5784840389528151610eef8582610e8a565b9450610efa83610e9e565b925060208a01995050600181019050610ed6565b50829750879550505050505092915050565b60006020820190508181036000830152610f3a8184610eab565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f8482610e03565b810181811067ffffffffffffffff82111715610fa357610fa2610f4c565b5b80604052505050565b6000610fb6610a00565b9050610fc28282610f7b565b919050565b600067ffffffffffffffff821115610fe257610fe1610f4c565b5b610feb82610e03565b9050602081019050919050565b82818337600083830152505050565b600061101a61101584610fc7565b610fac565b90508281526020810184848401111561103657611035610f47565b5b611041848285610ff8565b509392505050565b600082601f83011261105e5761105d610f42565b5b813561106e848260208601611007565b91505092915050565b6000806040838503121561108e5761108d610a0a565b5b600061109c85828601610a5d565b925050602083013567ffffffffffffffff8111156110bd576110bc610a0f565b5b6110c985828601611049565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061110d82610ad5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361113f5761113e6110d3565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806111c057607f821691505b6020821081036111d3576111d2611179565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212204c2169dff2f9ee291b91b1fc1afbff5af53039f6f159dbb35a7ed34f4145019c64736f6c634300080e0033",
}

// ElectionABI is the input ABI used to generate the binding from.
// Deprecated: Use ElectionMetaData.ABI instead.
var ElectionABI = ElectionMetaData.ABI

// ElectionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ElectionMetaData.Bin instead.
var ElectionBin = ElectionMetaData.Bin

// DeployElection deploys a new Ethereum contract, binding an instance of Election to it.
func DeployElection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Election, error) {
	parsed, err := ElectionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ElectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Election{ElectionCaller: ElectionCaller{contract: contract}, ElectionTransactor: ElectionTransactor{contract: contract}, ElectionFilterer: ElectionFilterer{contract: contract}}, nil
}

// Election is an auto generated Go binding around an Ethereum contract.
type Election struct {
	ElectionCaller     // Read-only binding to the contract
	ElectionTransactor // Write-only binding to the contract
	ElectionFilterer   // Log filterer for contract events
}

// ElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionSession struct {
	Contract     *Election         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionCallerSession struct {
	Contract *ElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionTransactorSession struct {
	Contract     *ElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionRaw struct {
	Contract *Election // Generic contract binding to access the raw methods on
}

// ElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionCallerRaw struct {
	Contract *ElectionCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionTransactorRaw struct {
	Contract *ElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElection creates a new instance of Election, bound to a specific deployed contract.
func NewElection(address common.Address, backend bind.ContractBackend) (*Election, error) {
	contract, err := bindElection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Election{ElectionCaller: ElectionCaller{contract: contract}, ElectionTransactor: ElectionTransactor{contract: contract}, ElectionFilterer: ElectionFilterer{contract: contract}}, nil
}

// NewElectionCaller creates a new read-only instance of Election, bound to a specific deployed contract.
func NewElectionCaller(address common.Address, caller bind.ContractCaller) (*ElectionCaller, error) {
	contract, err := bindElection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionCaller{contract: contract}, nil
}

// NewElectionTransactor creates a new write-only instance of Election, bound to a specific deployed contract.
func NewElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionTransactor, error) {
	contract, err := bindElection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionTransactor{contract: contract}, nil
}

// NewElectionFilterer creates a new log filterer instance of Election, bound to a specific deployed contract.
func NewElectionFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectionFilterer, error) {
	contract, err := bindElection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectionFilterer{contract: contract}, nil
}

// bindElection binds a generic wrapper to an already deployed contract.
func bindElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Election *ElectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Election.Contract.ElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Election *ElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.Contract.ElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Election *ElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Election.Contract.ElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Election *ElectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Election.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Election *ElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Election *ElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Election.Contract.contract.Transact(opts, method, params...)
}

// GetCandidates is a free data retrieval call binding the contract method 0x9763884d.
//
// Solidity: function GetCandidates() view returns((string,address)[])
func (_Election *ElectionCaller) GetCandidates(opts *bind.CallOpts) ([]ElectionCandidate, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetCandidates")

	if err != nil {
		return *new([]ElectionCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]ElectionCandidate)).(*[]ElectionCandidate)

	return out0, err

}

// GetCandidates is a free data retrieval call binding the contract method 0x9763884d.
//
// Solidity: function GetCandidates() view returns((string,address)[])
func (_Election *ElectionSession) GetCandidates() ([]ElectionCandidate, error) {
	return _Election.Contract.GetCandidates(&_Election.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x9763884d.
//
// Solidity: function GetCandidates() view returns((string,address)[])
func (_Election *ElectionCallerSession) GetCandidates() ([]ElectionCandidate, error) {
	return _Election.Contract.GetCandidates(&_Election.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0x8f7a842d.
//
// Solidity: function GetVoters() view returns((address)[])
func (_Election *ElectionCaller) GetVoters(opts *bind.CallOpts) ([]ElectionVoter, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetVoters")

	if err != nil {
		return *new([]ElectionVoter), err
	}

	out0 := *abi.ConvertType(out[0], new([]ElectionVoter)).(*[]ElectionVoter)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0x8f7a842d.
//
// Solidity: function GetVoters() view returns((address)[])
func (_Election *ElectionSession) GetVoters() ([]ElectionVoter, error) {
	return _Election.Contract.GetVoters(&_Election.CallOpts)
}

// GetVoters is a free data retrieval call binding the contract method 0x8f7a842d.
//
// Solidity: function GetVoters() view returns((address)[])
func (_Election *ElectionCallerSession) GetVoters() ([]ElectionVoter, error) {
	return _Election.Contract.GetVoters(&_Election.CallOpts)
}

// SeeCandidateVotes is a free data retrieval call binding the contract method 0x43ad4498.
//
// Solidity: function SeeCandidateVotes(address _address) view returns((address)[])
func (_Election *ElectionCaller) SeeCandidateVotes(opts *bind.CallOpts, _address common.Address) ([]ElectionVote, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "SeeCandidateVotes", _address)

	if err != nil {
		return *new([]ElectionVote), err
	}

	out0 := *abi.ConvertType(out[0], new([]ElectionVote)).(*[]ElectionVote)

	return out0, err

}

// SeeCandidateVotes is a free data retrieval call binding the contract method 0x43ad4498.
//
// Solidity: function SeeCandidateVotes(address _address) view returns((address)[])
func (_Election *ElectionSession) SeeCandidateVotes(_address common.Address) ([]ElectionVote, error) {
	return _Election.Contract.SeeCandidateVotes(&_Election.CallOpts, _address)
}

// SeeCandidateVotes is a free data retrieval call binding the contract method 0x43ad4498.
//
// Solidity: function SeeCandidateVotes(address _address) view returns((address)[])
func (_Election *ElectionCallerSession) SeeCandidateVotes(_address common.Address) ([]ElectionVote, error) {
	return _Election.Contract.SeeCandidateVotes(&_Election.CallOpts, _address)
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x3e77583b.
//
// Solidity: function TotalNumberOfVoters() view returns(uint256)
func (_Election *ElectionCaller) TotalNumberOfVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "TotalNumberOfVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x3e77583b.
//
// Solidity: function TotalNumberOfVoters() view returns(uint256)
func (_Election *ElectionSession) TotalNumberOfVoters() (*big.Int, error) {
	return _Election.Contract.TotalNumberOfVoters(&_Election.CallOpts)
}

// TotalNumberOfVoters is a free data retrieval call binding the contract method 0x3e77583b.
//
// Solidity: function TotalNumberOfVoters() view returns(uint256)
func (_Election *ElectionCallerSession) TotalNumberOfVoters() (*big.Int, error) {
	return _Election.Contract.TotalNumberOfVoters(&_Election.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x8807f271.
//
// Solidity: function candidates(address , uint256 ) view returns(address addr)
func (_Election *ElectionCaller) Candidates(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "candidates", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidates is a free data retrieval call binding the contract method 0x8807f271.
//
// Solidity: function candidates(address , uint256 ) view returns(address addr)
func (_Election *ElectionSession) Candidates(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Election.Contract.Candidates(&_Election.CallOpts, arg0, arg1)
}

// Candidates is a free data retrieval call binding the contract method 0x8807f271.
//
// Solidity: function candidates(address , uint256 ) view returns(address addr)
func (_Election *ElectionCallerSession) Candidates(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Election.Contract.Candidates(&_Election.CallOpts, arg0, arg1)
}

// AddVoter is a paid mutator transaction binding the contract method 0x0ad2eca7.
//
// Solidity: function AddVoter(address _address) returns(bool)
func (_Election *ElectionTransactor) AddVoter(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "AddVoter", _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0x0ad2eca7.
//
// Solidity: function AddVoter(address _address) returns(bool)
func (_Election *ElectionSession) AddVoter(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.AddVoter(&_Election.TransactOpts, _address)
}

// AddVoter is a paid mutator transaction binding the contract method 0x0ad2eca7.
//
// Solidity: function AddVoter(address _address) returns(bool)
func (_Election *ElectionTransactorSession) AddVoter(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.AddVoter(&_Election.TransactOpts, _address)
}

// CastVote is a paid mutator transaction binding the contract method 0x71ff924f.
//
// Solidity: function CastVote(address _address) returns(bool)
func (_Election *ElectionTransactor) CastVote(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "CastVote", _address)
}

// CastVote is a paid mutator transaction binding the contract method 0x71ff924f.
//
// Solidity: function CastVote(address _address) returns(bool)
func (_Election *ElectionSession) CastVote(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.CastVote(&_Election.TransactOpts, _address)
}

// CastVote is a paid mutator transaction binding the contract method 0x71ff924f.
//
// Solidity: function CastVote(address _address) returns(bool)
func (_Election *ElectionTransactorSession) CastVote(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.CastVote(&_Election.TransactOpts, _address)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x9e1fa02b.
//
// Solidity: function RegisterCandidate(address _address, string _name) returns(bool)
func (_Election *ElectionTransactor) RegisterCandidate(opts *bind.TransactOpts, _address common.Address, _name string) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "RegisterCandidate", _address, _name)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x9e1fa02b.
//
// Solidity: function RegisterCandidate(address _address, string _name) returns(bool)
func (_Election *ElectionSession) RegisterCandidate(_address common.Address, _name string) (*types.Transaction, error) {
	return _Election.Contract.RegisterCandidate(&_Election.TransactOpts, _address, _name)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x9e1fa02b.
//
// Solidity: function RegisterCandidate(address _address, string _name) returns(bool)
func (_Election *ElectionTransactorSession) RegisterCandidate(_address common.Address, _name string) (*types.Transaction, error) {
	return _Election.Contract.RegisterCandidate(&_Election.TransactOpts, _address, _name)
}
