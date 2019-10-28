package ethereum

import (
	_const "github.com/andrewnguyen22/pocket-interview-test/const"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient() (*ethclient.Client, error) {
	// Connect the client
	client, err := ethclient.Dial(_const.Api_Key)
	if err != nil {
		return nil, NewEthClientInitError(err)
	}
	return client, nil
}
