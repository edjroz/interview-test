package block

import (
	"math/big"

	"github.com/andrewnguyen22/pocket-interview-test/x/block"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
	Describe("GettingBlockByHash", func() {
		Context("Invalid hash query", func() {
			It("should return an error", func() {
				_, err := block.GetBlockByNumber(big.NewInt(-1))
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring(block.BlockByNumberError.Error()))
			})
		})
		Context("Valid hash", func() {
			It("Should return the ethereum genesis block", func() {
				// NOTE: use a valid has hex address otherwise test will fail
				hx := "0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177"
				res, err := block.GetBlockByHash(hx)
				Expect(err).To(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(res.Height).To(BeZero())
				Expect(res.Time).ToNot(BeZero())
				Expect(res.Hash).To(ContainSubstring("0x"))
			})
		})
	})

})
