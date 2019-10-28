package block

import (
	"github.com/andrewnguyen22/pocket-interview-test/x/block"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
)

var _ = Describe("Block", func() {
	Describe("GettingBlockByNumber", func() {
		Context("Invalid height query", func() {
			It("should return an error", func() {
				_, err := block.GetBlockByNumber(big.NewInt(-1))
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring(block.BlockByNumberError.Error()))
			})
		})
		Context("Valid height", func() {
			It("Should return the ethereum genesis block", func() {
				res, err := block.GetBlockByNumber(big.NewInt(0))
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(res.Height).To(BeZero())
				Expect(res.Time).ToNot(BeZero())
			})
		})
	})
})
