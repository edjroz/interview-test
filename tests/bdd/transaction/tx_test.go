package transaction

import (
	"github.com/andrewnguyen22/pocket-interview-test/x/transaction"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction", func() {
	Describe("Creating a raw Ethereum transaction", func() {
		Context("Dummy data already filled", func() {
			It("Should return a rpl raw tx", func() {
				Expect(transaction.NewRawTransaction()).ToNot(BeEmpty())
			})
		})
	})
	Describe("Sending a raw Ethereum transaction", func() {
		Context("Invalid rawTx", func() {
			It("should return an error", func() {
				_, err := transaction.SendTransaction([]byte("malicious tx"))
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring(transaction.SendError.Error()))
			})
		})
		Context("Valid rawTx", func() {
			It("should return a hash", func() {
				raw, _ := transaction.NewRawTransaction()
				hash, _ := transaction.SendTransaction(raw)
				Expect(hash.Hex()).ToNot(BeNil())
			})
		})
	})
})
