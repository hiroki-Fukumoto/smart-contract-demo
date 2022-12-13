package contractaddress

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/contracts"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/env"
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
	client, err := ethclient.Dial(fmt.Sprintf("%s:%s", env.GanacheHost(), env.GanachePort()))
	if err != nil {
		return nil, err
	}

	priKey := "ed19d0e3fc1e8d3bb92389bf993943949c6c96f17f4bf506bb0b5c5194ee780b"
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
