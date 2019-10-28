package transaction

import (
	"context"
	_const "github.com/andrewnguyen22/pocket-interview-test/const"
	"github.com/andrewnguyen22/pocket-interview-test/x/auth"
	"github.com/andrewnguyen22/pocket-interview-test/x/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

// creates new transaction object
func NewRawTransaction() ([]byte, error) {
	client, err := ethereum.GetClient()
	defer client.Close()
	if err != nil {
		return nil, err
	}
	privKey, pubKey, err := auth.GetCoinbaseKeypair()
	if err != nil {
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*pubKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, NewNonceError(err)
	}

	value := big.NewInt(50000000000000000) // in wei (.05 eth)
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, NewGasError(err)
	}

	toAddress := common.HexToAddress(_const.ToAddress)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, NewNetIDError(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		return nil, NewTxSigError(err)
	}

	ts := types.Transactions{signedTx}
	return ts.GetRlp(0), nil
}

// todo above creates a rawTx object -> use it to create SendTransaction, create bdd test, expose rpc endpoint to send
