package encrypt_test

import (
	"testing"

	"github.com/howood/moggiecollector/infrastructure/encrypt"
)

func Test_GetHash(t *testing.T) {
	t.Parallel()

	testdata := "aaaaaaaccccvvvvv"
	hashdata := encrypt.DataHash{}.GetHash(testdata)
	t.Log(hashdata)
	if hashdata == testdata {
		t.Fatal("failed GetHash ")
	}
	t.Log("success GetHash")
}
