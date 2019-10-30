// This package is rpc between the different RPC packages
package rpc

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/andrewnguyen22/pocket-interview-test/x/block"
	"github.com/ethereum/go-ethereum/common"

	"github.com/julienschmidt/httprouter"
)

// Populate the model from the parameters of the POST call.
func PopModel(_ http.ResponseWriter, r *http.Request, _ httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

// handles the localhost:<port>/block/byNumber call.
func BlockbyNumber(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		WriteErrorResponse(w, InternalError, InvalidHttpParsing)
		return
	}
	x := r.Form.Get("height")
	// convert to bigInteger
	i, ok := new(big.Int).SetString(x, 10)
	if !ok {
		WriteErrorResponse(w, InternalError, InvalidInputMessage)
		return
	}
	// call the underlying function
	blk, err := block.GetBlockByNumber(i)
	if err != nil {
		WriteErrorResponse(w, InternalError, err.Error())
		return
	}
	res, err := json.MarshalIndent(blk, "", "   ")
	WriteJSONResponse(w, res)
}

// handles the localhost:<port>/block/byHash call
func BlockbyHash(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		WriteErrorResponse(w, InternalError, InvalidHttpParsing)
		return
	}
	hash := r.Form.Get("hash")

	if hash == "" || common.IsHexAddress(hash) {
		WriteErrorResponse(w, InternalError, InvalidInputMessage)
		return
	}

	blk, err := block.GetBlockByHash(hash)
	if err != nil {
		WriteErrorResponse(w, InternalError, err.Error())
		return
	}
	res, err := json.MarshalIndent(blk, "", "   ")
	WriteJSONResponse(w, res)
}
