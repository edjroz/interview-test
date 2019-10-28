package block

import "errors"

var (
	BlockByNumberError = errors.New("unable to retrieve the eth block by height")
)

func NewBlockByNumberError(err error) error {
	return errors.New(BlockByNumberError.Error() + ": " + err.Error())
}
