package crypto

import (
	"crypto/sha256"
)

func HashMessage(message []byte) (msgHashSum []byte , err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(message)
	if err != nil {
		return
	}
	msgHashSum = msgHash.Sum(nil)
	return
}