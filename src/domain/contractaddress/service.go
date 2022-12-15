package contractaddress

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/contracts"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/ethereumclient"
)

type ContractAddressService interface {
	GetContractAddress() (*string, error)
}

type contractAddressService struct {
}

func NewContractAddressService() ContractAddressService {
	return &contractAddressService{}
}

func (s contractAddressService) GetContractAddress() (*string, error) {
	client, err := ethereumclient.GetEthereumClient()
	if err != nil {
		return nil, err
	}

	priKey := "928f6acb4f4305c44d86b117d4a72d559cddeacdc801b6cde6a9c7f1a63f0215"
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address, _, _, err := contracts.DeployContracts(auth, client)
	if err != nil {
		return nil, err
	}

	res := address.String()

	return &res, nil
}
