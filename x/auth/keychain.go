package auth

import (
	"sync"

	_const "github.com/andrewnguyen22/pocket-interview-test/const"
)

// NOTE: does not do any sanity checks on keypair, ignore this for now

// TODO make keychain thread safe, no test needed ✅

type KeyChain struct {
	Keys []string
	l    *sync.Mutex
}

// holds private Keys keychain
var (
	globalKeyChain *KeyChain
)

func New() *KeyChain {
	var lock = &sync.Mutex{}
	return &KeyChain{[]string{}, lock}
}

// retrieve the keychain
func GetKeyChain() *KeyChain {
	if globalKeyChain == nil {
		globalKeyChain = New()
		(*globalKeyChain).Keys = append((*globalKeyChain).Keys, _const.PrivateKey)
	}
	return globalKeyChain
}

// add key to keychain
func (kc *KeyChain) AddKey(key string) {
	kc.l.Lock()
	defer kc.l.Unlock()
	if !kc.Contains(key) {
		kc.Keys = append(kc.Keys, key)
	}
}

// delete key from keychain
func (kc *KeyChain) DeleteKey(key string) {
	kc.l.Lock()
	defer kc.l.Unlock()
	if i := kc.containsWhere(key); i != -1 {
		// can't delete coinbase
		if i != 0 {
			kc.Keys = append(kc.Keys[:i], kc.Keys[i+1:]...)
		}
	}
}

// returns the first `main` account
func (kc KeyChain) GetCoinbase() string {
	return kc.Keys[0]
}

// returns which index the slice contains
func (kc KeyChain) containsWhere(element string) int {
	for index, a := range kc.Keys {
		if a == element {
			return index
		}
	}
	return -1
}

// returns if the slice contains
func (kc KeyChain) Contains(element string) bool {
	for _, a := range kc.Keys {
		if a == element {
			return true
		}
	}
	return false
}
