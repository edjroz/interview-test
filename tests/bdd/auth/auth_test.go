package auth

import (
	_const "github.com/andrewnguyen22/pocket-interview-test/const"
	"github.com/andrewnguyen22/pocket-interview-test/x/auth"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	fakeKey = "1"
)

var _ = Describe("Auth", func() {
	Describe("Getting the global keychain", func() {
		Context("Single process retrieval of variable", func() {
			It("Should return the global keychain variable", func() {
				Expect(auth.GetKeyChain()).ToNot(BeNil())
				Expect(len(*auth.GetKeyChain())).ToNot(BeZero())
			})
		})
	})
	Describe("Contains function", func() {
		Context("Initialized with the coinbase privkey", func() {
			It("Should return true", func() {
				Expect(auth.GetKeyChain().Contains(_const.PrivateKey)).To(BeTrue())
			})
		})
	})
	Describe("Adding a key to the keychain", func() {
		Context("No error checking on key validity", func() {
			It("Should sucessfully add a new key to the global list", func() {
				auth.GetKeyChain().AddKey(fakeKey)
				Expect(auth.GetKeyChain().Contains(fakeKey)).To(BeTrue())
			})
		})
	})
	Describe("Deleting a key to the keychain", func() {
		Context("Contains the key", func() {
			It("Should sucessfully delete a new key to the global list", func() {
				auth.GetKeyChain().AddKey(fakeKey)
				auth.GetKeyChain().DeleteKey(fakeKey)
				Expect(auth.GetKeyChain().Contains(fakeKey)).To(BeFalse())
			})
		})
		Context("Doesn't contain the key", func() {
			It("Should not affect the state of the list", func() {
				auth.GetKeyChain().DeleteKey(fakeKey)
				Expect(auth.GetKeyChain().Contains(fakeKey)).To(BeFalse())
				Expect(len(*auth.GetKeyChain())).ToNot(BeZero())
			})
		})
	})
	Describe("Getting the coinbase", func() {
		Context("Always has a coinbase", func() {
			It("Should get the coinbase", func() {
				Expect(auth.GetKeyChain().GetCoinbase()).ToNot(BeEmpty())
			})
		})
	})
})
