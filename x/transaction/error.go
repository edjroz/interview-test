package transaction

import "errors"

var (
	NonceError     = errors.New("unable to retrieve the pending nonce")
	GasError       = errors.New("error with raw tx gas")
	NetIDError     = errors.New("error with network id while creating raw tx")
	SignatureError = errors.New("error while signing the raw tx")
	SendError      = errors.New("error while trying to send transaction")
)

func NewNonceError(err error) error {
	return errors.New(NonceError.Error() + ": " + err.Error())
}

func NewGasError(err error) error {
	return errors.New(GasError.Error() + ": " + err.Error())
}

func NewNetIDError(err error) error {
	return errors.New(NetIDError.Error() + ": " + err.Error())
}

func NewTxSigError(err error) error {
	return errors.New(SignatureError.Error() + ": " + err.Error())
}
func NewSendTxError(err error) error {
	return errors.New(SendError.Error() + ": " + err.Error())
}
