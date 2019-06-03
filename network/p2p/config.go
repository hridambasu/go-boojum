package p2p

import (
	"io/ioutil"
	"strings"

	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
)

// Identity ...
func Identity(pkPath string) (libp2p.Option, error) {
	pkRaw, err := ioutil.ReadFile(pkPath)
	if err != nil {
		return nil, err
	}

	pkStr := string(pkRaw)
	pkStr = strings.TrimRight(pkStr, "\n")
	pkStr = strings.TrimRight(pkStr, "\r")
	
	pk, err := crypto.UnmarshalSecp256k1PrivateKey([]byte(pkStr))
	if err != nil {
		return nil, err
	}

	return libp2p.Identity(pk), nil
}

// RandomIdentity ...
func RandomIdentity() (libp2p.Option) {
	sk, _, _ := crypto.GenerateSecp256k1Key(nil)
	return libp2p.Identity(sk)
}

// ListenAddress ...
func ListenAddress(addr string) (libp2p.Option, error) {
	return libp2p.ListenAddrStrings(addr), nil
}