package auth

import (
	"crypto/ecdsa"
	_const "github.com/andrewnguyen22/pocket-interview-test/const"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetCoinbaseKeypair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := crypto.HexToECDSA(_const.PrivateKey)
	if err != nil {
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, ECDSAError
	}
	return privateKey, publicKeyECDSA, nil
}
