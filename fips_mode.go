//go:build boringcrypto
// +build boringcrypto

package main

import _ "crypto/tls/fipsonly"

func init() {
	infoLog.Println("FIPS crypto enforced with boringcrypto")
}
