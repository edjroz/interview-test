package block

import "errors"

var (
	BlockByNumberError = errors.New("unable to retrieve the eth block by height")
	BlockByHashError   = errors.New("unable to retrieve the eth block by hash")
)

// return error for getting block by number error
func NewBlockByNumberError(err error) error {
	return errors.New(BlockByNumberError.Error() + ": " + err.Error())
}

// return error for getting block by hash error
func NewBlockByHashError(err error) error {
	return errors.New(BlockByHashError.Error() + ": " + err.Error())
}
