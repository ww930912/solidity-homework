// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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

// CounterNFTMetaData contains all meta data concerning the CounterNFT contract.
var CounterNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"Incremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600b81526020017f436f756e746572204e46540000000000000000000000000000000000000000008152506040518060400160405280600481526020017f434e465400000000000000000000000000000000000000000000000000000000815250815f908161008a9190610b03565b50806001908161009a9190610b03565b5050506100ae3360016100b360201b60201c565b610c79565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610123575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161011a9190610c11565b60405180910390fd5b5f61013583835f6101ac60201b60201c565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146101a7575f6040517f73c6ac6e00000000000000000000000000000000000000000000000000000000815260040161019e9190610c11565b60405180910390fd5b505050565b5f5f6101bd846103c960201b60201c565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146102045761020381848661040260201b60201c565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610295576102495f855f5f6104cb60201b60201c565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461031457600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b61041383838361069660201b60201c565b6104c6575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361048757806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161047e9190610c39565b60405180910390fd5b81816040517f177e802f0000000000000000000000000000000000000000000000000000000081526004016104bd929190610c52565b60405180910390fd5b505050565b808061050357505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15610641575f6105188461076260201b60201c565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561058257508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561059b575061059981846107ee60201b60201c565b155b156105dd57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016105d49190610c11565b60405180910390fd5b811561063f57838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561075957508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610714575061071384846107ee60201b60201c565b5b8061075857508273ffffffffffffffffffffffffffffffffffffffff166107408361087c60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f5f610773836103c960201b60201c565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107e557826040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016107dc9190610c39565b60405180910390fd5b80915050919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061093057607f821691505b602082108103610943576109426108ec565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026109a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261096a565b6109af868361096a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6109f36109ee6109e9846109c7565b6109d0565b6109c7565b9050919050565b5f819050919050565b610a0c836109d9565b610a20610a18826109fa565b848454610976565b825550505050565b5f5f905090565b610a37610a28565b610a42818484610a03565b505050565b5f5b82811015610a6857610a5d5f828401610a2f565b600181019050610a49565b505050565b601f821115610abb5782821115610aba57610a8781610949565b610a908361095b565b610a998561095b565b6020861015610aa6575f90505b808301610ab582840382610a47565b505050505b5b505050565b5f82821c905092915050565b5f610adb5f1984600802610ac0565b1980831691505092915050565b5f610af38383610acc565b9150826002028217905092915050565b610b0c826108b5565b67ffffffffffffffff811115610b2557610b246108bf565b5b610b2f8254610919565b610b3a828285610a6d565b5f60209050601f831160018114610b6b575f8415610b59578287015190505b610b638582610ae8565b865550610bca565b601f198416610b7986610949565b5f5b82811015610ba057848901518255600182019150602085019450602081019050610b7b565b86831015610bbd5784890151610bb9601f891682610acc565b8355505b6001600288020188555050505b505050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610bfb82610bd2565b9050919050565b610c0b81610bf1565b82525050565b5f602082019050610c245f830184610c02565b92915050565b610c33816109c7565b82525050565b5f602082019050610c4c5f830184610c2a565b92915050565b5f604082019050610c655f830185610c02565b610c726020830184610c2a565b9392505050565b61227780610c865f395ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c806370a0823111610095578063c87b56dd11610064578063c87b56dd146102a8578063d09de08a146102d8578063e985e9c5146102e2578063fb7f21eb14610312576100fe565b806370a082311461022257806395d89b4114610252578063a22cb46514610270578063b88d4fde1461028c576100fe565b8063095ea7b3116100d1578063095ea7b31461019e57806323b872dd146101ba57806342842e0e146101d65780636352211e146101f2576100fe565b806301ffc9a71461010257806306661abd1461013257806306fdde0314610150578063081812fc1461016e575b5f5ffd5b61011c6004803603810190610117919061179d565b610330565b60405161012991906117e2565b60405180910390f35b61013a610411565b6040516101479190611813565b60405180910390f35b610158610417565b604051610165919061189c565b60405180910390f35b610188600480360381019061018391906118e6565b6104a6565b6040516101959190611950565b60405180910390f35b6101b860048036038101906101b39190611993565b6104c1565b005b6101d460048036038101906101cf91906119d1565b6104d7565b005b6101f060048036038101906101eb91906119d1565b6105d6565b005b61020c600480360381019061020791906118e6565b6105f5565b6040516102199190611950565b60405180910390f35b61023c60048036038101906102379190611a21565b610606565b6040516102499190611813565b60405180910390f35b61025a6106bc565b604051610267919061189c565b60405180910390f35b61028a60048036038101906102859190611a76565b61074c565b005b6102a660048036038101906102a19190611be0565b610762565b005b6102c260048036038101906102bd91906118e6565b610787565b6040516102cf919061189c565b60405180910390f35b6102e061081d565b005b6102fc60048036038101906102f79190611c60565b610888565b60405161030991906117e2565b60405180910390f35b61031a610916565b604051610327919061189c565b60405180910390f35b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103fa57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b8061040a575061040982610933565b5b9050919050565b60065481565b60605f805461042590611ccb565b80601f016020809104026020016040519081016040528092919081815260200182805461045190611ccb565b801561049c5780601f106104735761010080835404028352916020019161049c565b820191905f5260205f20905b81548152906001019060200180831161047f57829003601f168201915b5050505050905090565b5f6104b08261099c565b506104ba82610a22565b9050919050565b6104d382826104ce610a5b565b610a62565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610547575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161053e9190611950565b60405180910390fd5b5f61055a8383610555610a5b565b610a74565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146105d0578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016105c793929190611cfb565b60405180910390fd5b50505050565b6105f083838360405180602001604052805f815250610762565b505050565b5f6105ff8261099c565b9050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610677575f6040517f89c62b6400000000000000000000000000000000000000000000000000000000815260040161066e9190611950565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6060600180546106cb90611ccb565b80601f01602080910402602001604051908101604052809291908181526020018280546106f790611ccb565b80156107425780601f1061071957610100808354040283529160200191610742565b820191905f5260205f20905b81548152906001019060200180831161072557829003601f168201915b5050505050905090565b61075e610757610a5b565b8383610c7f565b5050565b61076d8484846104d7565b610781610778610a5b565b85858585610e58565b50505050565b60606107928261099c565b505f61079d83611004565b6107bf60405180610100016040528060d6815260200161216c60d691396110ce565b6107ca600654611004565b6040516020016107dc93929190611ebc565b60405160208183030381529060405290506107f6816110ce565b6040516020016108069190611f46565b604051602081830303815290604052915050919050565b600160065f82825461082f9190611f98565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa542060065460405161087e9190611813565b60405180910390a2565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60405180610100016040528060d6815260200161216c60d6913981565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f5f6109a7836110e1565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a1957826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610a109190611813565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610a6f838383600161111a565b505050565b5f5f610a7f846110e1565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610ac057610abf8184866112d9565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b4b57610aff5f855f5f61111a565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610bca57600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610cef575f6040517fa9fbf51f000000000000000000000000000000000000000000000000000000008152600401610ce69190611950565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d5f57816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401610d569190611950565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610e4b91906117e2565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115610ffd578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401610eb6949392919061201d565b6020604051808303815f875af1925050508015610ef157506040513d601f19601f82011682018060405250810190610eee919061207b565b60015b610f72573d805f8114610f1f576040519150601f19603f3d011682016040523d82523d5f602084013e610f24565b606091505b505f815103610f6a57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610f619190611950565b60405180910390fd5b805160208201fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610ffb57836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610ff29190611950565b60405180910390fd5b505b5050505050565b60605f60016110128461139c565b0190505f8167ffffffffffffffff8111156110305761102f611abc565b5b6040519080825280601f01601f1916602001820160405280156110625781602001600182028036833780820191505090505b5090505f82602083010190505b6001156110c3578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816110b8576110b76120a6565b5b0494505f850361106f575b819350505050919050565b60606110da825f6114ed565b9050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061115257505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611284575f6111618461099c565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141580156111cb57508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156111de57506111dc8184610888565b155b1561122057826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016112179190611950565b60405180910390fd5b811561128257838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6112e4838383611677565b611397575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361135857806040517f7e27328900000000000000000000000000000000000000000000000000000000815260040161134f9190611813565b60405180910390fd5b81816040517f177e802f00000000000000000000000000000000000000000000000000000000815260040161138e9291906120d3565b60405180910390fd5b505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106113f8577a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083816113ee576113ed6120a6565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310611435576d04ee2d6d415b85acef8100000000838161142b5761142a6120a6565b5b0492506020810190505b662386f26fc10000831061146457662386f26fc10000838161145a576114596120a6565b5b0492506010810190505b6305f5e100831061148d576305f5e1008381611483576114826120a6565b5b0492506008810190505b61271083106114b25761271083816114a8576114a76120a6565b5b0492506004810190505b606483106114d557606483816114cb576114ca6120a6565b5b0492506002810190505b600a83106114e4576001810190505b80915050919050565b60605f83510361150d5760405180602001604052805f8152509050611671565b5f8261153e576003600285516115239190611f98565b61152d91906120fa565b6004611539919061212a565b611565565b6003600285516004611550919061212a565b61155a9190611f98565b61156491906120fa565b5b905060405191507f4142434445464748494a4b4c4d4e4f505152535455565758595a616263646566601f5261067083027f6768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2f18603f526020820181810185865187016020810180515f82525b82841015611623576003840193508351603f8160121c16518753600187019650603f81600c1c16518753600187019650603f8160061c16518753600187019650603f8116518753600187019650506115d0565b808252886116625760038a510660018114611645576002811461165857611660565b603d6001880353603d6002880353611660565b603d60018803535b505b86885284604052505050505050505b92915050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561172e57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806116ef57506116ee8484610888565b5b8061172d57508273ffffffffffffffffffffffffffffffffffffffff1661171583610a22565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61177c81611748565b8114611786575f5ffd5b50565b5f8135905061179781611773565b92915050565b5f602082840312156117b2576117b1611740565b5b5f6117bf84828501611789565b91505092915050565b5f8115159050919050565b6117dc816117c8565b82525050565b5f6020820190506117f55f8301846117d3565b92915050565b5f819050919050565b61180d816117fb565b82525050565b5f6020820190506118265f830184611804565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61186e8261182c565b6118788185611836565b9350611888818560208601611846565b61189181611854565b840191505092915050565b5f6020820190508181035f8301526118b48184611864565b905092915050565b6118c5816117fb565b81146118cf575f5ffd5b50565b5f813590506118e0816118bc565b92915050565b5f602082840312156118fb576118fa611740565b5b5f611908848285016118d2565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61193a82611911565b9050919050565b61194a81611930565b82525050565b5f6020820190506119635f830184611941565b92915050565b61197281611930565b811461197c575f5ffd5b50565b5f8135905061198d81611969565b92915050565b5f5f604083850312156119a9576119a8611740565b5b5f6119b68582860161197f565b92505060206119c7858286016118d2565b9150509250929050565b5f5f5f606084860312156119e8576119e7611740565b5b5f6119f58682870161197f565b9350506020611a068682870161197f565b9250506040611a17868287016118d2565b9150509250925092565b5f60208284031215611a3657611a35611740565b5b5f611a438482850161197f565b91505092915050565b611a55816117c8565b8114611a5f575f5ffd5b50565b5f81359050611a7081611a4c565b92915050565b5f5f60408385031215611a8c57611a8b611740565b5b5f611a998582860161197f565b9250506020611aaa85828601611a62565b9150509250929050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611af282611854565b810181811067ffffffffffffffff82111715611b1157611b10611abc565b5b80604052505050565b5f611b23611737565b9050611b2f8282611ae9565b919050565b5f67ffffffffffffffff821115611b4e57611b4d611abc565b5b611b5782611854565b9050602081019050919050565b828183375f83830152505050565b5f611b84611b7f84611b34565b611b1a565b905082815260208101848484011115611ba057611b9f611ab8565b5b611bab848285611b64565b509392505050565b5f82601f830112611bc757611bc6611ab4565b5b8135611bd7848260208601611b72565b91505092915050565b5f5f5f5f60808587031215611bf857611bf7611740565b5b5f611c058782880161197f565b9450506020611c168782880161197f565b9350506040611c27878288016118d2565b925050606085013567ffffffffffffffff811115611c4857611c47611744565b5b611c5487828801611bb3565b91505092959194509250565b5f5f60408385031215611c7657611c75611740565b5b5f611c838582860161197f565b9250506020611c948582860161197f565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611ce257607f821691505b602082108103611cf557611cf4611c9e565b5b50919050565b5f606082019050611d0e5f830186611941565b611d1b6020830185611804565b611d286040830184611941565b949350505050565b7f7b226e616d65223a22436f756e74657220230000000000000000000000000000815250565b5f81905092915050565b5f611d6a8261182c565b611d748185611d56565b9350611d84818560208601611846565b80840191505092915050565b7f222c226465736372697074696f6e223a22436f756e74657220636f6e747261635f8201527f74206c6f676f204e4654222c22696d616765223a22646174613a696d6167652f60208201527f7376672b786d6c3b6261736536342c0000000000000000000000000000000000604082015250565b5f611e10604f83611d56565b9150611e1b82611d90565b604f82019050919050565b7f222c2261747472696275746573223a5b7b2274726169745f74797065223a22635f8201527f6f756e74222c2276616c7565223a000000000000000000000000000000000000602082015250565b5f611e80602e83611d56565b9150611e8b82611e26565b602e82019050919050565b7f7d5d7d0000000000000000000000000000000000000000000000000000000000815250565b5f611ec682611d30565b601282019150611ed68286611d60565b9150611ee182611e04565b9150611eed8285611d60565b9150611ef882611e74565b9150611f048284611d60565b9150611f0f82611e96565b600382019150819050949350505050565b7f646174613a6170706c69636174696f6e2f6a736f6e3b6261736536342c000000815250565b5f611f5082611f20565b601d82019150611f608284611d60565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611fa2826117fb565b9150611fad836117fb565b9250828201905080821115611fc557611fc4611f6b565b5b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f611fef82611fcb565b611ff98185611fd5565b9350612009818560208601611846565b61201281611854565b840191505092915050565b5f6080820190506120305f830187611941565b61203d6020830186611941565b61204a6040830185611804565b818103606083015261205c8184611fe5565b905095945050505050565b5f8151905061207581611773565b92915050565b5f602082840312156120905761208f611740565b5b5f61209d84828501612067565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6040820190506120e65f830185611941565b6120f36020830184611804565b9392505050565b5f612104826117fb565b915061210f836117fb565b92508261211f5761211e6120a6565b5b828204905092915050565b5f612134826117fb565b915061213f836117fb565b925082820261214d816117fb565b9150828204841483151761216457612163611f6b565b5b509291505056fe3c73766720786d6c6e733d27687474703a2f2f7777772e77332e6f72672f323030302f737667272076696577426f783d27302030203634203634273e3c636972636c652063783d273332272063793d2733322720723d273330272066696c6c3d2723363336366631272f3e3c7465787420783d2733322720793d2734322720666f6e742d73697a653d2732382720666f6e742d66616d696c793d276d6f6e6f73706163652720746578742d616e63686f723d276d6964646c65272066696c6c3d2723666666273e43363c2f746578743e3c2f7376673ea2646970667358221220d17c4699535e7da550533f1a0e3720cc89ff15737b1baa8ab239141c4094326b64736f6c63430008250033",
}

// CounterNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterNFTMetaData.ABI instead.
var CounterNFTABI = CounterNFTMetaData.ABI

// CounterNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterNFTMetaData.Bin instead.
var CounterNFTBin = CounterNFTMetaData.Bin

// DeployCounterNFT deploys a new Ethereum contract, binding an instance of CounterNFT to it.
func DeployCounterNFT(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CounterNFT, error) {
	parsed, err := CounterNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterNFTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CounterNFT{CounterNFTCaller: CounterNFTCaller{contract: contract}, CounterNFTTransactor: CounterNFTTransactor{contract: contract}, CounterNFTFilterer: CounterNFTFilterer{contract: contract}}, nil
}

// CounterNFT is an auto generated Go binding around an Ethereum contract.
type CounterNFT struct {
	CounterNFTCaller     // Read-only binding to the contract
	CounterNFTTransactor // Write-only binding to the contract
	CounterNFTFilterer   // Log filterer for contract events
}

// CounterNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterNFTSession struct {
	Contract     *CounterNFT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterNFTCallerSession struct {
	Contract *CounterNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CounterNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterNFTTransactorSession struct {
	Contract     *CounterNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CounterNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterNFTRaw struct {
	Contract *CounterNFT // Generic contract binding to access the raw methods on
}

// CounterNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterNFTCallerRaw struct {
	Contract *CounterNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CounterNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterNFTTransactorRaw struct {
	Contract *CounterNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounterNFT creates a new instance of CounterNFT, bound to a specific deployed contract.
func NewCounterNFT(address common.Address, backend bind.ContractBackend) (*CounterNFT, error) {
	contract, err := bindCounterNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CounterNFT{CounterNFTCaller: CounterNFTCaller{contract: contract}, CounterNFTTransactor: CounterNFTTransactor{contract: contract}, CounterNFTFilterer: CounterNFTFilterer{contract: contract}}, nil
}

// NewCounterNFTCaller creates a new read-only instance of CounterNFT, bound to a specific deployed contract.
func NewCounterNFTCaller(address common.Address, caller bind.ContractCaller) (*CounterNFTCaller, error) {
	contract, err := bindCounterNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterNFTCaller{contract: contract}, nil
}

// NewCounterNFTTransactor creates a new write-only instance of CounterNFT, bound to a specific deployed contract.
func NewCounterNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterNFTTransactor, error) {
	contract, err := bindCounterNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterNFTTransactor{contract: contract}, nil
}

// NewCounterNFTFilterer creates a new log filterer instance of CounterNFT, bound to a specific deployed contract.
func NewCounterNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterNFTFilterer, error) {
	contract, err := bindCounterNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterNFTFilterer{contract: contract}, nil
}

// bindCounterNFT binds a generic wrapper to an already deployed contract.
func bindCounterNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterNFT *CounterNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterNFT.Contract.CounterNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterNFT *CounterNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterNFT.Contract.CounterNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterNFT *CounterNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterNFT.Contract.CounterNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CounterNFT *CounterNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CounterNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CounterNFT *CounterNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CounterNFT *CounterNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CounterNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CounterNFT *CounterNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CounterNFT *CounterNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CounterNFT.Contract.BalanceOf(&_CounterNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CounterNFT *CounterNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CounterNFT.Contract.BalanceOf(&_CounterNFT.CallOpts, owner)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CounterNFT *CounterNFTCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CounterNFT *CounterNFTSession) Count() (*big.Int, error) {
	return _CounterNFT.Contract.Count(&_CounterNFT.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CounterNFT *CounterNFTCallerSession) Count() (*big.Int, error) {
	return _CounterNFT.Contract.Count(&_CounterNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CounterNFT *CounterNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CounterNFT *CounterNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CounterNFT.Contract.GetApproved(&_CounterNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CounterNFT *CounterNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CounterNFT.Contract.GetApproved(&_CounterNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CounterNFT *CounterNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CounterNFT *CounterNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CounterNFT.Contract.IsApprovedForAll(&_CounterNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CounterNFT *CounterNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CounterNFT.Contract.IsApprovedForAll(&_CounterNFT.CallOpts, owner, operator)
}

// Logo is a free data retrieval call binding the contract method 0xfb7f21eb.
//
// Solidity: function logo() view returns(string)
func (_CounterNFT *CounterNFTCaller) Logo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "logo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Logo is a free data retrieval call binding the contract method 0xfb7f21eb.
//
// Solidity: function logo() view returns(string)
func (_CounterNFT *CounterNFTSession) Logo() (string, error) {
	return _CounterNFT.Contract.Logo(&_CounterNFT.CallOpts)
}

// Logo is a free data retrieval call binding the contract method 0xfb7f21eb.
//
// Solidity: function logo() view returns(string)
func (_CounterNFT *CounterNFTCallerSession) Logo() (string, error) {
	return _CounterNFT.Contract.Logo(&_CounterNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CounterNFT *CounterNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CounterNFT *CounterNFTSession) Name() (string, error) {
	return _CounterNFT.Contract.Name(&_CounterNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CounterNFT *CounterNFTCallerSession) Name() (string, error) {
	return _CounterNFT.Contract.Name(&_CounterNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CounterNFT *CounterNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CounterNFT *CounterNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CounterNFT.Contract.OwnerOf(&_CounterNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CounterNFT *CounterNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CounterNFT.Contract.OwnerOf(&_CounterNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CounterNFT *CounterNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CounterNFT *CounterNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CounterNFT.Contract.SupportsInterface(&_CounterNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CounterNFT *CounterNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CounterNFT.Contract.SupportsInterface(&_CounterNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CounterNFT *CounterNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CounterNFT *CounterNFTSession) Symbol() (string, error) {
	return _CounterNFT.Contract.Symbol(&_CounterNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CounterNFT *CounterNFTCallerSession) Symbol() (string, error) {
	return _CounterNFT.Contract.Symbol(&_CounterNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CounterNFT *CounterNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CounterNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CounterNFT *CounterNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CounterNFT.Contract.TokenURI(&_CounterNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CounterNFT *CounterNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CounterNFT.Contract.TokenURI(&_CounterNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.Contract.Approve(&_CounterNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.Contract.Approve(&_CounterNFT.TransactOpts, to, tokenId)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterNFT *CounterNFTTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CounterNFT.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterNFT *CounterNFTSession) Increment() (*types.Transaction, error) {
	return _CounterNFT.Contract.Increment(&_CounterNFT.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_CounterNFT *CounterNFTTransactorSession) Increment() (*types.Transaction, error) {
	return _CounterNFT.Contract.Increment(&_CounterNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.Contract.SafeTransferFrom(&_CounterNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.Contract.SafeTransferFrom(&_CounterNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CounterNFT *CounterNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CounterNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CounterNFT *CounterNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CounterNFT.Contract.SafeTransferFrom0(&_CounterNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CounterNFT *CounterNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CounterNFT.Contract.SafeTransferFrom0(&_CounterNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CounterNFT *CounterNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CounterNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CounterNFT *CounterNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CounterNFT.Contract.SetApprovalForAll(&_CounterNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CounterNFT *CounterNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CounterNFT.Contract.SetApprovalForAll(&_CounterNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.Contract.TransferFrom(&_CounterNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CounterNFT *CounterNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CounterNFT.Contract.TransferFrom(&_CounterNFT.TransactOpts, from, to, tokenId)
}

// CounterNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CounterNFT contract.
type CounterNFTApprovalIterator struct {
	Event *CounterNFTApproval // Event containing the contract specifics and raw log

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
func (it *CounterNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterNFTApproval)
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
		it.Event = new(CounterNFTApproval)
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
func (it *CounterNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterNFTApproval represents a Approval event raised by the CounterNFT contract.
type CounterNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CounterNFT *CounterNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CounterNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CounterNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CounterNFTApprovalIterator{contract: _CounterNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CounterNFT *CounterNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CounterNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CounterNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterNFTApproval)
				if err := _CounterNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CounterNFT *CounterNFTFilterer) ParseApproval(log types.Log) (*CounterNFTApproval, error) {
	event := new(CounterNFTApproval)
	if err := _CounterNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CounterNFT contract.
type CounterNFTApprovalForAllIterator struct {
	Event *CounterNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CounterNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterNFTApprovalForAll)
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
		it.Event = new(CounterNFTApprovalForAll)
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
func (it *CounterNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterNFTApprovalForAll represents a ApprovalForAll event raised by the CounterNFT contract.
type CounterNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CounterNFT *CounterNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CounterNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CounterNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CounterNFTApprovalForAllIterator{contract: _CounterNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CounterNFT *CounterNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CounterNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CounterNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterNFTApprovalForAll)
				if err := _CounterNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CounterNFT *CounterNFTFilterer) ParseApprovalForAll(log types.Log) (*CounterNFTApprovalForAll, error) {
	event := new(CounterNFTApprovalForAll)
	if err := _CounterNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterNFTIncrementedIterator is returned from FilterIncremented and is used to iterate over the raw logs and unpacked data for Incremented events raised by the CounterNFT contract.
type CounterNFTIncrementedIterator struct {
	Event *CounterNFTIncremented // Event containing the contract specifics and raw log

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
func (it *CounterNFTIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterNFTIncremented)
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
		it.Event = new(CounterNFTIncremented)
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
func (it *CounterNFTIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterNFTIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterNFTIncremented represents a Incremented event raised by the CounterNFT contract.
type CounterNFTIncremented struct {
	By       common.Address
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIncremented is a free log retrieval operation binding the contract event 0x38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa5420.
//
// Solidity: event Incremented(address indexed by, uint256 newCount)
func (_CounterNFT *CounterNFTFilterer) FilterIncremented(opts *bind.FilterOpts, by []common.Address) (*CounterNFTIncrementedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _CounterNFT.contract.FilterLogs(opts, "Incremented", byRule)
	if err != nil {
		return nil, err
	}
	return &CounterNFTIncrementedIterator{contract: _CounterNFT.contract, event: "Incremented", logs: logs, sub: sub}, nil
}

// WatchIncremented is a free log subscription operation binding the contract event 0x38ac789ed44572701765277c4d0970f2db1c1a571ed39e84358095ae4eaa5420.
//
// Solidity: event Incremented(address indexed by, uint256 newCount)
func (_CounterNFT *CounterNFTFilterer) WatchIncremented(opts *bind.WatchOpts, sink chan<- *CounterNFTIncremented, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _CounterNFT.contract.WatchLogs(opts, "Incremented", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterNFTIncremented)
				if err := _CounterNFT.contract.UnpackLog(event, "Incremented", log); err != nil {
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
func (_CounterNFT *CounterNFTFilterer) ParseIncremented(log types.Log) (*CounterNFTIncremented, error) {
	event := new(CounterNFTIncremented)
	if err := _CounterNFT.contract.UnpackLog(event, "Incremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CounterNFT contract.
type CounterNFTTransferIterator struct {
	Event *CounterNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CounterNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterNFTTransfer)
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
		it.Event = new(CounterNFTTransfer)
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
func (it *CounterNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterNFTTransfer represents a Transfer event raised by the CounterNFT contract.
type CounterNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CounterNFT *CounterNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CounterNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CounterNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CounterNFTTransferIterator{contract: _CounterNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CounterNFT *CounterNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CounterNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CounterNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterNFTTransfer)
				if err := _CounterNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CounterNFT *CounterNFTFilterer) ParseTransfer(log types.Log) (*CounterNFTTransfer, error) {
	event := new(CounterNFTTransfer)
	if err := _CounterNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
