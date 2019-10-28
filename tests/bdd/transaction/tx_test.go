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
})
