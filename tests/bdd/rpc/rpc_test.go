package rpc

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/andrewnguyen22/pocket-interview-test/rpc"
	"github.com/julienschmidt/httprouter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// error response for testing purposes only
type ErrorResponse struct {
	Code  int    `json:"code"`
	Title string `json:"title,omitempty"`
}

// response mock for testing purposes only
type response struct {
	Error      ErrorResponse `json:"error"`
	Height     uint64        `json:"height"`
	Time       uint64        `json:"time"`
	Difficulty uint64        `json:"difficulty"`
	Hash       string        `json:"hash"`
}

var _ = Describe("RPC", func() {
	Describe("Block By Hash", func() {
		Context("Invalid Hash", func() {
			var (
				req    *http.Request
				rec    *httptest.ResponseRecorder
				params httprouter.Params
				form   url.Values
				err    error
			)

			BeforeEach(func() {
				req, err = http.NewRequest("POST", "localhost:1050/block/ByHash", strings.NewReader(form.Encode()))
				if err != nil {
					log.Fatal("Could not create request")
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				rec = httptest.NewRecorder()
				params = httprouter.Params{}
			})
			It("should return Error without without form", func() {
				req, err = http.NewRequest("POST", "localhost:1050/block/ByHash", nil)
				if err != nil {
					log.Fatal("Could not create request")
				}
				rpc.BlockbyHash(rec, req, params)
				res := rec.Result()
				defer res.Body.Close()

				body := &response{}
				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					log.Fatal("could not read response")
				}

				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(body.Error.Title).To(ContainSubstring("the http form could not be parsed"))
			})
			It("should return Error with malformed hash", func() {
				rpc.BlockbyHash(rec, req, params)
				res := rec.Result()
				defer res.Body.Close()
				body := &response{}
				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					log.Fatal("could not read response")
				}
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(body.Error.Title).To(ContainSubstring("unable to be processed"))
			})
		})
		Context("Valid Hash", func() {
			params := httprouter.Params{}
			rec := httptest.NewRecorder()
			form := url.Values{}
			addr := "0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177"
			It("Should return Pokt Ethereum Block", func() {
				form.Add("hash", addr)
				req, err := http.NewRequest("POST", "localhost:1050/block/ByHash", strings.NewReader(form.Encode()))
				if err != nil {
					log.Fatal("Could not create request")
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				rpc.BlockbyHash(rec, req, params)
				res := rec.Result()
				defer res.Body.Close()

				body := &response{}
				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					log.Fatal(err)
					log.Fatal("could not read response")
				}

				Expect(body).ToNot(BeNil())
				Expect(body.Height).To(BeZero())
				Expect(body.Time).ToNot(BeZero())
				Expect(body.Hash).To(ContainSubstring("0x"))
			})
		})
	})
	Describe("Block By Number", func() {
		Context("Invalid Height ", func() {
			var (
				req    *http.Request
				rec    *httptest.ResponseRecorder
				params httprouter.Params
				form   url.Values
				err    error
			)

			BeforeEach(func() {
				req, err = http.NewRequest("POST", "localhost:1050/block/ByNumber", strings.NewReader(form.Encode()))
				if err != nil {
					log.Fatal("Could not create request")
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				rec = httptest.NewRecorder()
				params = httprouter.Params{}
			})
			It("should return Error without without form", func() {
				req, err = http.NewRequest("POST", "localhost:1050/block/ByNumber", nil)
				if err != nil {
					log.Fatal("Could not create request")
				}
				rpc.BlockbyNumber(rec, req, params)
				res := rec.Result()
				defer res.Body.Close()

				body := &response{}
				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					log.Fatal("could not read response")
				}

				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(body.Error.Title).To(ContainSubstring("the http form could not be parsed"))
			})
			It("should return Error with malformed height", func() {
				rpc.BlockbyNumber(rec, req, params)
				res := rec.Result()
				defer res.Body.Close()
				body := &response{}
				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					log.Fatal("could not read response")
				}
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(body.Error.Title).To(ContainSubstring("unable to be processed"))
			})
		})
		Context("Valid Height", func() {
			params := httprouter.Params{}
			rec := httptest.NewRecorder()
			form := url.Values{}
			It("Should return a Pokt ethereum block", func() {
				form.Add("height", "0")
				req, err := http.NewRequest("POST", "localhost:1050/block/ByNumber", strings.NewReader(form.Encode()))
				if err != nil {
					log.Fatal("Could not create request")
				}
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				rpc.BlockbyNumber(rec, req, params)
				res := rec.Result()
				defer res.Body.Close()

				body := &response{}
				err = json.NewDecoder(res.Body).Decode(&body)
				if err != nil {
					log.Fatal(err)
					log.Fatal("could not read response")
				}

				Expect(body).ToNot(BeNil())
				Expect(body.Height).To(BeZero())
				Expect(body.Time).ToNot(BeZero())
				Expect(body.Hash).To(ContainSubstring("0x"))
			})
		})
	})
	Describe(("SendTransaction"), func() {
		It("sends a transaction", func() {
			params := httprouter.Params{}
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "localhost:1050/send/Transaction", nil)
			if err != nil {
				log.Fatal("Could not create request")
			}
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			rpc.SendTrx(rec, req, params)
			res := rec.Result()
			defer res.Body.Close()

			body := &response{}
			err = json.NewDecoder(res.Body).Decode(&body)
			if err != nil {
				log.Fatal(err)
				log.Fatal("could not read response")
			}
			Expect(body).ToNot(BeNil())
			Expect(body.Hash).To(ContainSubstring("0x"))
		})
	})
})
