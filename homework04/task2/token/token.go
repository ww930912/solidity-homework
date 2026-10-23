// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = time.Tick
	_ = context.Background
)

// CounterTokenMetaData contains all meta data concerning the CounterToken contract.
var CounterTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"Incremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600d81526020017f436f756e74657220546f6b656e000000000000000000000000000000000000008152506040518060400160405280600381526020017f434e540000000000000000000000000000000000000000000000000000000000815250816003908161008b91906105cd565b50806004908161009b91906105cd565b5050506100d4336100b06100d960201b60201c565b600a6100bc9190610804565b6103e86100c9919061084e565b6100e160201b60201c565b610977565b5f6012905090565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610151575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161014891906108ce565b60405180910390fd5b6101625f838361016660201b60201c565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036101b6578060025f8282546101aa91906108e7565b92505081905550610284565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508181101561023f578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161023693929190610929565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036102cb578060025f8282540392505081905550610315565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610372919061095e565b60405180910390a3505050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103fa57607f821691505b60208210810361040d5761040c6103b6565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261046f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610434565b6104798683610434565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6104bd6104b86104b384610491565b61049a565b610491565b9050919050565b5f819050919050565b6104d6836104a3565b6104ea6104e2826104c4565b848454610440565b825550505050565b5f5f905090565b6105016104f2565b61050c8184846104cd565b505050565b5f5b82811015610532576105275f8284016104f9565b600181019050610513565b505050565b601f82111561058557828211156105845761055181610413565b61055a83610425565b61056385610425565b6020861015610570575f90505b80830161057f82840382610511565b505050505b5b505050565b5f82821c905092915050565b5f6105a55f198460080261058a565b1980831691505092915050565b5f6105bd8383610596565b9150826002028217905092915050565b6105d68261037f565b67ffffffffffffffff8111156105ef576105ee610389565b5b6105f982546103e3565b610604828285610537565b5f60209050601f831160018114610635575f8415610623578287015190505b61062d85826105b2565b865550610694565b601f19841661064386610413565b5f5b8281101561066a57848901518255600182019150602085019450602081019050610645565b868310156106875784890151610683601f891682610596565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b600185111561071e578086048111156106fa576106f961069c565b5b60018516156107095780820291505b8081029050610717856106c9565b94506106de565b94509492505050565b5f8261073657600190506107f1565b81610743575f90506107f1565b8160018114610759576002811461076357610792565b60019150506107f1565b60ff8411156107755761077461069c565b5b8360020a91508482111561078c5761078b61069c565b5b506107f1565b5060208310610133831016604e8410600b84101617156107c75782820a9050838111156107c2576107c161069c565b5b6107f1565b6107d484848460016106d5565b925090508184048111156107eb576107ea61069c565b5b81810290505b9392505050565b5f60ff82169050919050565b5f61080e82610491565b9150610819836107f8565b92506108467fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610727565b905092915050565b5f61085882610491565b915061086383610491565b925082820261087181610491565b915082820484148315176108885761088761069c565b5b5092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108b88261088f565b9050919050565b6108c8816108ae565b82525050565b5f6020820190506108e15f8301846108bf565b92915050565b5f6108f182610491565b91506108fc83610491565b92508282019050808211156109145761091361069c565b5b92915050565b61092381610491565b82525050565b5f60608201905061093c5f8301866108bf565b610949602083018561091a565b610956604083018461091a565b949350505050565b5f6020820190506109715f83018461091a565b92915050565b610e91806109845f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063313ce5671161006f578063313ce5671461016557806370a082311461018357806395d89b41146101b3578063a9059cbb146101d1578063d09de08a14610201578063dd62ed3e1461020b576100a7565b806306661abd146100ab57806306fdde03146100c9578063095ea7b3146100e757806318160ddd1461011757806323b872dd14610135575b5f5ffd5b6100b361023b565b6040516100c09190610ab2565b60405180910390f35b6100d1610241565b6040516100de9190610b3b565b60405180910390f35b61010160048036038101906100fc9190610be3565b6102d1565b60405161010e9190610c3b565b60405180910390f35b61011f6102f3565b60405161012c9190610ab2565b60405180910390f35b61014f600480360381019061014a9190610c54565b6102fc565b60405161015c9190610c3b565b60405180910390f35b61016d61032a565b60405161017a9190610cbf565b60405180910390f35b61019d60048036038101906101989190610cd8565b610332565b6040516101aa9190610ab2565b60405180910390f35b6101bb610377565b6040516101c89190610b3b565b60405180910390f35b6101eb60048036038101906101e69190610be3565b610407565b6040516101f89190610c3b565b60405180910390f35b610209610429565b005b61022560048036038101906102209190610d03565b610494565b6040516102329190610ab2565b60405180910390f35b60055481565b60606003805461025090610d6e565b80601f016020809104026020016040519081016040528092919081815260200182805461027c90610d6e565b80156102c75780601f1061029e576101008083540402835291602001916102c7565b820191905f5260205f20905b8154815290600101906020018083116102aa57829003601f168201915b5050505050905090565b5f5f6102db610516565b90506102e881858561051d565b600191505092915050565b5f600254905090565b5f5f610306610516565b905061031385828561052f565b61031e8585856105c2565b60019150509392505050565b5f6012905090565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60606004805461038690610d6e565b80601f01602080910402602001604051908101604052809291908181526020018280546103b290610d6e565b80156103fd5780601f106103d4576101008083540402835291602001916103fd565b820191905f5260205f20905b8154815290600101906020018083116103e057829003601f168201915b5050505050905090565b5f5f610411610516565b905061041e8185856105c2565b600191505092915050565b600160055f82825461043b9190610dcb565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa542060055460405161048a9190610ab2565b60405180910390a2565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b61052a83838360016106b2565b505050565b5f61053a8484610494565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156105bc57818110156105ad578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016105a493929190610e0d565b60405180910390fd5b6105bb84848484035f6106b2565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610632575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016106299190610e42565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106a2575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016106999190610e42565b60405180910390fd5b6106ad838383610881565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610722575f6040517fe602df050000000000000000000000000000000000000000000000000000000081526004016107199190610e42565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610792575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016107899190610e42565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550801561087b578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516108729190610ab2565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108d1578060025f8282546108c59190610dcb565b9250508190555061099f565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508181101561095a578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161095193929190610e0d565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109e6578060025f8282540392505081905550610a30565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610a8d9190610ab2565b60405180910390a3505050565b5f819050919050565b610aac81610a9a565b82525050565b5f602082019050610ac55f830184610aa3565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610b0d82610acb565b610b178185610ad5565b9350610b27818560208601610ae5565b610b3081610af3565b840191505092915050565b5f6020820190508181035f830152610b538184610b03565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b8882610b5f565b9050919050565b610b9881610b7e565b8114610ba2575f5ffd5b50565b5f81359050610bb381610b8f565b92915050565b610bc281610a9a565b8114610bcc575f5ffd5b50565b5f81359050610bdd81610bb9565b92915050565b5f5f60408385031215610bf957610bf8610b5b565b5b5f610c0685828601610ba5565b9250506020610c1785828601610bcf565b9150509250929050565b5f8115159050919050565b610c3581610c21565b82525050565b5f602082019050610c4e5f830184610c2c565b92915050565b5f5f5f60608486031215610c6b57610c6a610b5b565b5b5f610c7886828701610ba5565b9350506020610c8986828701610ba5565b9250506040610c9a86828701610bcf565b9150509250925092565b5f60ff82169050919050565b610cb981610ca4565b82525050565b5f602082019050610cd25f830184610cb0565b92915050565b5f60208284031215610ced57610cec610b5b565b5b5f610cfa84828501610ba5565b91505092915050565b5f5f60408385031215610d1957610d18610b5b565b5b5f610d2685828601610ba5565b9250506020610d3785828601610ba5565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610d8557607f821691505b602082108103610d9857610d97610d41565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610dd582610a9a565b9150610de083610a9a565b9250828201905080821115610df857610df7610d9e565b5b92915050565b610e0781610b7e565b82525050565b5f606082019050610e205f830186610dfe565b610e2d6020830185610aa3565b610e3a6040830184610aa3565b949350505050565b5f602082019050610e555f830184610dfe565b9291505056fea2646970667358221220c0edb8611e21d445d5f983668a73403e97827395d86e022ac1ea88ffff3b7f6764736f6c63430008250033",
}

// CounterTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterTokenMetaData.ABI instead.
var CounterTokenABI = CounterTokenMetaData.ABI

// CounterTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterTokenMetaData.Bin instead.
var CounterTokenBin = CounterTokenMetaData.Bin

// DeployCounterToken deploys a new Ethereum contract, binding an instance of CounterToken to it.
func DeployCounterToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CounterToken, error) {
	parsed, err := CounterTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterToken{CounterTokenCaller: CounterTokenCaller{contract: contract}, CounterTokenTransactor: CounterTokenTransactor{contract: contract}, CounterTokenFilterer: CounterTokenFilterer{contract: contract}}, nil
}

// CounterToken is an auto generated Go binding around an Ethereum contract.
type CounterToken struct {
	CounterTokenCaller     // Read-only binding to the contract
	CounterTokenTransactor // Write-only binding to the contract
	CounterTokenFilterer   // Log filterer for contract events
}

// CounterTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterTokenSession struct {
	Contract     *CounterToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterTokenCallerSession struct {
	Contract *CounterTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CounterTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTokenTransactorSession struct {
	Contract     *CounterTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CounterTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterTokenRaw struct {
	Contract *CounterToken // Generic contract binding to access the raw methods on
}

// CounterTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterTokenCallerRaw struct {
	Contract *CounterTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTokenTransactorRaw struct {
	Contract *CounterTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterToken creates a new instance of CounterToken, bound to a specific deployed contract.
func NewCounterToken(address common.Address, backend bind.ContractBackend) (*CounterToken, error) {
	contract, err := bindCounterToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterToken{CounterTokenCaller: CounterTokenCaller{contract: contract}, CounterTokenTransactor: CounterTokenTransactor{contract: contract}, CounterTokenFilterer: CounterTokenFilterer{contract: contract}}, nil
}

// NewCounterTokenCaller creates a new read-only instance of CounterToken, bound to a specific deployed contract.
func NewCounterTokenCaller(address common.Address, caller bind.ContractCaller) (*CounterTokenCaller, error) {
	contract, err := bindCounterToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTokenCaller{contract: contract}, nil
}

// NewCounterTokenTransactor creates a new write-only instance of CounterToken, bound to a specific deployed contract.
func NewCounterTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTokenTransactor, error) {
	contract, err := bindCounterToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTokenTransactor{contract: contract}, nil
}

// NewCounterTokenFilterer creates a new log filterer instance of CounterToken, bound to a specific deployed contract.
func NewCounterTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterTokenFilterer, error) {
	contract, err := bindCounterToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterTokenFilterer{contract: contract}, nil
}

// bindCounterToken binds a generic wrapper to an already deployed contract.
func bindCounterToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterToken *CounterTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterToken.Contract.CounterTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterToken *CounterTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterToken.Contract.CounterTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterToken *CounterTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterToken.Contract.CounterTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterToken *CounterTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterToken *CounterTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterToken *CounterTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CounterToken *CounterTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CounterToken *CounterTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CounterToken.Contract.Allowance(&_CounterToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CounterToken *CounterTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CounterToken.Contract.Allowance(&_CounterToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CounterToken *CounterTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CounterToken *CounterTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CounterToken.Contract.BalanceOf(&_CounterToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CounterToken *CounterTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CounterToken.Contract.BalanceOf(&_CounterToken.CallOpts, account)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CounterToken *CounterTokenCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CounterToken *CounterTokenSession) Count() (*big.Int, error) {
	return _CounterToken.Contract.Count(&_CounterToken.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CounterToken *CounterTokenCallerSession) Count() (*big.Int, error) {
	return _CounterToken.Contract.Count(&_CounterToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CounterToken *CounterTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CounterToken *CounterTokenSession) Decimals() (uint8, error) {
	return _CounterToken.Contract.Decimals(&_CounterToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CounterToken *CounterTokenCallerSession) Decimals() (uint8, error) {
	return _CounterToken.Contract.Decimals(&_CounterToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CounterToken *CounterTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CounterToken *CounterTokenSession) Name() (string, error) {
	return _CounterToken.Contract.Name(&_CounterToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CounterToken *CounterTokenCallerSession) Name() (string, error) {
	return _CounterToken.Contract.Name(&_CounterToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CounterToken *CounterTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CounterToken *CounterTokenSession) Symbol() (string, error) {
	return _CounterToken.Contract.Symbol(&_CounterToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CounterToken *CounterTokenCallerSession) Symbol() (string, error) {
	return _CounterToken.Contract.Symbol(&_CounterToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CounterToken *CounterTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CounterToken *CounterTokenSession) TotalSupply() (*big.Int, error) {
	return _CounterToken.Contract.TotalSupply(&_CounterToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CounterToken *CounterTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CounterToken.Contract.TotalSupply(&_CounterToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CounterToken *CounterTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CounterToken *CounterTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.Contract.Approve(&_CounterToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CounterToken *CounterTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.Contract.Approve(&_CounterToken.TransactOpts, spender, value)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterToken *CounterTokenTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterToken.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterToken *CounterTokenSession) Increment() (*types.Transaction, error) {
	return _CounterToken.Contract.Increment(&_CounterToken.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterToken *CounterTokenTransactorSession) Increment() (*types.Transaction, error) {
	return _CounterToken.Contract.Increment(&_CounterToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CounterToken *CounterTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CounterToken *CounterTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.Contract.Transfer(&_CounterToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CounterToken *CounterTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.Contract.Transfer(&_CounterToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CounterToken *CounterTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CounterToken *CounterTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.Contract.TransferFrom(&_CounterToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CounterToken *CounterTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CounterToken.Contract.TransferFrom(&_CounterToken.TransactOpts, from, to, value)
}

// CounterTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CounterToken contract.
type CounterTokenApprovalIterator struct {
	Event *CounterTokenApproval // Event containing the contract specifics and raw log

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
func (it *CounterTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterTokenApproval)
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
		it.Event = new(CounterTokenApproval)
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
func (it *CounterTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterTokenApproval represents a Approval event raised by the CounterToken contract.
type CounterTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CounterToken *CounterTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CounterTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CounterToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CounterTokenApprovalIterator{contract: _CounterToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CounterToken *CounterTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CounterTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CounterToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterTokenApproval)
				if err := _CounterToken.contract.UnpackLog(event, "Approval", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CounterToken *CounterTokenFilterer) ParseApproval(log types.Log) (*CounterTokenApproval, error) {
	event := new(CounterTokenApproval)
	if err := _CounterToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterTokenIncrementedIterator is returned from FilterIncremented and is used to iterate over the raw logs and unpacked data for Incremented events raised by the CounterToken contract.
type CounterTokenIncrementedIterator struct {
	Event *CounterTokenIncremented // Event containing the contract specifics and raw log

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
func (it *CounterTokenIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterTokenIncremented)
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
		it.Event = new(CounterTokenIncremented)
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
func (it *CounterTokenIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterTokenIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterTokenIncremented represents a Incremented event raised by the CounterToken contract.
type CounterTokenIncremented struct {
	By       common.Address
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncremented is a free log retrieval operation binding the contract event 0x38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa5420.
//
// Solidity: event Incremented(address indexed by, uint256 newCount)
func (_CounterToken *CounterTokenFilterer) FilterIncremented(opts *bind.FilterOpts, by []common.Address) (*CounterTokenIncrementedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _CounterToken.contract.FilterLogs(opts, "Incremented", byRule)
	if err != nil {
		return nil, err
	}
	return &CounterTokenIncrementedIterator{contract: _CounterToken.contract, event: "Incremented", logs: logs, sub: sub}, nil
}

// WatchIncremented is a free log subscription operation binding the contract event 0x38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa5420.
//
// Solidity: event Incremented(address indexed by, uint256 newCount)
func (_CounterToken *CounterTokenFilterer) WatchIncremented(opts *bind.WatchOpts, sink chan<- *CounterTokenIncremented, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _CounterToken.contract.WatchLogs(opts, "Incremented", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterTokenIncremented)
				if err := _CounterToken.contract.UnpackLog(event, "Incremented", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseIncremented is a log parse operation binding the contract event 0x38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa5420.
//
// Solidity: event Incremented(address indexed by, uint256 newCount)
func (_CounterToken *CounterTokenFilterer) ParseIncremented(log types.Log) (*CounterTokenIncremented, error) {
	event := new(CounterTokenIncremented)
	if err := _CounterToken.contract.UnpackLog(event, "Incremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CounterToken contract.
type CounterTokenTransferIterator struct {
	Event *CounterTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CounterTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterTokenTransfer)
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
		it.Event = new(CounterTokenTransfer)
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
func (it *CounterTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterTokenTransfer represents a Transfer event raised by the CounterToken contract.
type CounterTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CounterToken *CounterTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CounterTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CounterToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CounterTokenTransferIterator{contract: _CounterToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CounterToken *CounterTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CounterTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CounterToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterTokenTransfer)
				if err := _CounterToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					// If the signature doesn't match, skip this log.
					if errors.Is(err, bind.ErrEventSignatureMismatch) {
						continue
					}
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CounterToken *CounterTokenFilterer) ParseTransfer(log types.Log) (*CounterTokenTransfer, error) {
	event := new(CounterTokenTransfer)
	if err := _CounterToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
