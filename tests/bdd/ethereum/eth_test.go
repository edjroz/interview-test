package ethereum

import (
	"github.com/andrewnguyen22/pocket-interview-test/x/ethereum"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ethereum", func() {
	Describe("Retriveing client", func() {
		Context("Valid API Key", func() {
			It("Should return nil error and non-nil client", func() {
				cli, err := ethereum.GetClient()
				Expect(cli).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
