package block

import (
	"context"
	"math/big"

	pkt "github.com/andrewnguyen22/pocket-interview-test/types"
	"github.com/andrewnguyen22/pocket-interview-test/x/ethereum"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

// retrieves the block by height
func GetBlockByNumber(height *big.Int) (*pkt.Block, error) {
	client, err := ethereum.GetClient()
	defer client.Close()
	if err != nil {
		return nil, err
	}
	block, err := client.BlockByNumber(context.Background(), height)
	if err != nil {
		return nil, NewBlockByNumberError(err)
	}
	return ethBlockToPocketBlock(block), nil
}

// converts to pocket type
func ethBlockToPocketBlock(block *eth.Block) *pkt.Block {
	return &pkt.Block{
		Height:     block.Number().Uint64(),
		Time:       block.Time(),
		Difficulty: block.Difficulty().Uint64(),
		Hash:       block.Hash().Hex(),
	}
}

// TODO getBlockByHash ✅
// TODO bdd test ✅
// TODO create rpc endpoint ✅
// retrieves the block by hash
func GetBlockByHash(hash string) (*pkt.Block, error) {
	client, err := ethereum.GetClient()
	defer client.Close()
	if err != nil {
		return nil, err
	}
	block, err := client.BlockByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return nil, NewBlockByHashError(err)
	}
	return ethBlockToPocketBlock(block), nil
}
