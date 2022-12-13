package ethereumclient

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hiroki-Fukumoto/geth-dapp-demo/env"
)

func GetEthereumClient() (*ethclient.Client, error) {
	client, err := ethclient.Dial(fmt.Sprintf("%s:%s", env.GanacheHost(), env.GanachePort()))
	if err != nil {
		return nil, err
	}
	return client, nil
}
