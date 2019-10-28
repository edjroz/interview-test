package ethereum

import "errors"

var (
	EthClientInitError = errors.New("unable to initialize the Ethereum client")
)

func NewEthClientInitError(err error) error {
	return errors.New(EthClientInitError.Error() + ": " + err.Error())
}
