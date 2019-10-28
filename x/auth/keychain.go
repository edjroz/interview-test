package auth

import _const "github.com/andrewnguyen22/pocket-interview-test/const"

// NOTE: does not do any sanity checks on keypair, ignore this for now

// todo make keychain thread safe, no test needed

type KeyChain []string

// holds private keys keychain
var (
	globalKeyChain *KeyChain
)

// retrieve the keychain
func GetKeyChain() *KeyChain {
	if globalKeyChain == nil {
		globalKeyChain = new(KeyChain)
		*globalKeyChain = []string{_const.PrivateKey}
	}
	return globalKeyChain
}

// add key to keychain
func (kc *KeyChain) AddKey(key string) {
	if !kc.Contains(key) {
		*kc = append(*kc, key)
	}
}

// delete key from keychain
func (kc *KeyChain) DeleteKey(key string) {
	if i := kc.containsWhere(key); i != -1 {
		// can't delete coinbase
		if i != 0 {
			*kc = append((*kc)[:i], (*kc)[i+1:]...)
		}
	}
}

// returns the first `main` account
func (kc KeyChain) GetCoinbase() string {
	return kc[0]
}

// returns which index the slice contains
func (kc KeyChain) containsWhere(element string) int {
	for index, a := range kc {
		if a == element {
			return index
		}
	}
	return -1
}

// returns if the slice contains
func (kc KeyChain) Contains(element string) bool {
	for _, a := range kc {
		if a == element {
			return true
		}
	}
	return false
}
