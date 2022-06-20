// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package election

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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"AddVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVoters\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structElection.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"RegisterCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"SeeCandidateVotes\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structElection.Vote[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalNumberOfVoter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060075534801561001557600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611403806100656000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638807f2711161005b5780638807f2711461013b5780638f7a842d1461016b5780639763884d146101895780639e1fa02b146101a757610088565b80630ad2eca71461008d57806343ad4498146100bd578063509dd40f146100ed57806371ff924f1461010b575b600080fd5b6100a760048036038101906100a29190610b94565b6101d7565b6040516100b49190610bdc565b60405180910390f35b6100d760048036038101906100d29190610b94565b610364565b6040516100e49190610cd1565b60405180910390f35b6100f5610454565b6040516101029190610d0c565b60405180910390f35b61012560048036038101906101209190610b94565b610461565b6040516101329190610bdc565b60405180910390f35b61015560048036038101906101509190610d53565b6105ee565b6040516101629190610da2565b60405180910390f35b610173610645565b6040516101809190610e88565b60405180910390f35b6101916106f6565b60405161019e9190611042565b60405180910390f35b6101c160048036038101906101bc9190611199565b61083d565b6040516101ce9190610bdc565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025f90611278565b60405180910390fd5b600754600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600560405180602001604052808473ffffffffffffffffffffffffffffffffffffffff168152509080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505060076000815480929190610356906112c7565b919050555060019050919050565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610449578382906000526020600020016040518060200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681525050815260200190600101906103c5565b505050509050919050565b6000600580549050905090565b600080600090505b6002805490508110156105e3578273ffffffffffffffffffffffffffffffffffffffff16600282815481106104a1576104a061130f565b5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036105d057600060405180602001604052803373ffffffffffffffffffffffffffffffffffffffff168152509050600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506001925050506105e9565b80806105db906112c7565b915050610469565b50600090505b919050565b6001602052816000526040600020818154811061060a57600080fd5b90600052602060002001600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081565b60606005805480602002602001604051908101604052809291908181526020016000905b828210156106ed578382906000526020600020016040518060200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505081526020019060010190610669565b50505050905090565b60606002805480602002602001604051908101604052809291908181526020016000905b82821015610834578382906000526020600020906002020160405180604001604052908160008201805461074d9061136d565b80601f01602080910402602001604051908101604052809291908181526020018280546107799061136d565b80156107c65780601f1061079b576101008083540402835291602001916107c6565b820191906000526020600020905b8154815290600101906020018083116107a957829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250508152602001906001019061071a565b50505050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108c590611278565b60405180910390fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600181600181540180825580915050039060005260206000205050600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806109775761097661139e565b5b60019003818190600052602060002001600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550509055600060405180604001604052808481526020018573ffffffffffffffffffffffffffffffffffffffff16815250905060028190806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000019080519060200190610a2a929190610a7f565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600191505092915050565b828054610a8b9061136d565b90600052602060002090601f016020900481019282610aad5760008555610af4565b82601f10610ac657805160ff1916838001178555610af4565b82800160010185558215610af4579182015b82811115610af3578251825591602001919060010190610ad8565b5b509050610b019190610b05565b5090565b5b80821115610b1e576000816000905550600101610b06565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b6182610b36565b9050919050565b610b7181610b56565b8114610b7c57600080fd5b50565b600081359050610b8e81610b68565b92915050565b600060208284031215610baa57610ba9610b2c565b5b6000610bb884828501610b7f565b91505092915050565b60008115159050919050565b610bd681610bc1565b82525050565b6000602082019050610bf16000830184610bcd565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610c2c81610b56565b82525050565b602082016000820151610c486000850182610c23565b50505050565b6000610c5a8383610c32565b60208301905092915050565b6000602082019050919050565b6000610c7e82610bf7565b610c888185610c02565b9350610c9383610c13565b8060005b83811015610cc4578151610cab8882610c4e565b9750610cb683610c66565b925050600181019050610c97565b5085935050505092915050565b60006020820190508181036000830152610ceb8184610c73565b905092915050565b6000819050919050565b610d0681610cf3565b82525050565b6000602082019050610d216000830184610cfd565b92915050565b610d3081610cf3565b8114610d3b57600080fd5b50565b600081359050610d4d81610d27565b92915050565b60008060408385031215610d6a57610d69610b2c565b5b6000610d7885828601610b7f565b9250506020610d8985828601610d3e565b9150509250929050565b610d9c81610b56565b82525050565b6000602082019050610db76000830184610d93565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b602082016000820151610dff6000850182610c23565b50505050565b6000610e118383610de9565b60208301905092915050565b6000602082019050919050565b6000610e3582610dbd565b610e3f8185610dc8565b9350610e4a83610dd9565b8060005b83811015610e7b578151610e628882610e05565b9750610e6d83610e1d565b925050600181019050610e4e565b5085935050505092915050565b60006020820190508181036000830152610ea28184610e2a565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f10578082015181840152602081019050610ef5565b83811115610f1f576000848401525b50505050565b6000601f19601f8301169050919050565b6000610f4182610ed6565b610f4b8185610ee1565b9350610f5b818560208601610ef2565b610f6481610f25565b840191505092915050565b60006040830160008301518482036000860152610f8c8282610f36565b9150506020830151610fa16020860182610c23565b508091505092915050565b6000610fb88383610f6f565b905092915050565b6000602082019050919050565b6000610fd882610eaa565b610fe28185610eb5565b935083602082028501610ff485610ec6565b8060005b8581101561103057848403895281516110118582610fac565b945061101c83610fc0565b925060208a01995050600181019050610ff8565b50829750879550505050505092915050565b6000602082019050818103600083015261105c8184610fcd565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110a682610f25565b810181811067ffffffffffffffff821117156110c5576110c461106e565b5b80604052505050565b60006110d8610b22565b90506110e4828261109d565b919050565b600067ffffffffffffffff8211156111045761110361106e565b5b61110d82610f25565b9050602081019050919050565b82818337600083830152505050565b600061113c611137846110e9565b6110ce565b90508281526020810184848401111561115857611157611069565b5b61116384828561111a565b509392505050565b600082601f8301126111805761117f611064565b5b8135611190848260208601611129565b91505092915050565b600080604083850312156111b0576111af610b2c565b5b60006111be85828601610b7f565b925050602083013567ffffffffffffffff8111156111df576111de610b31565b5b6111eb8582860161116b565b9150509250929050565b600082825260208201905092915050565b7f4f6e6c792061646d696e7320617265207065726d697474656420746f20646f2060008201527f7468697300000000000000000000000000000000000000000000000000000000602082015250565b60006112626024836111f5565b915061126d82611206565b604082019050919050565b6000602082019050818103600083015261129181611255565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006112d282610cf3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361130457611303611298565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061138557607f821691505b6020821081036113985761139761133e565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122073f4db4010d1a2298ad68e3869f8a4d92941bb1ee07dd526ee3605ec7bef8cc264736f6c634300080e0033",
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

// TotalNumberOfVoter is a free data retrieval call binding the contract method 0x509dd40f.
//
// Solidity: function TotalNumberOfVoter() view returns(uint256)
func (_Election *ElectionCaller) TotalNumberOfVoter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "TotalNumberOfVoter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalNumberOfVoter is a free data retrieval call binding the contract method 0x509dd40f.
//
// Solidity: function TotalNumberOfVoter() view returns(uint256)
func (_Election *ElectionSession) TotalNumberOfVoter() (*big.Int, error) {
	return _Election.Contract.TotalNumberOfVoter(&_Election.CallOpts)
}

// TotalNumberOfVoter is a free data retrieval call binding the contract method 0x509dd40f.
//
// Solidity: function TotalNumberOfVoter() view returns(uint256)
func (_Election *ElectionCallerSession) TotalNumberOfVoter() (*big.Int, error) {
	return _Election.Contract.TotalNumberOfVoter(&_Election.CallOpts)
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
