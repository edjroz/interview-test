package auth

import "errors"

var (
	ECDSAError = errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
)
