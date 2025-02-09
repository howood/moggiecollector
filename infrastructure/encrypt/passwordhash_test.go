package encrypt_test

import (
	"context"
	"testing"

	"github.com/howood/moggiecollector/infrastructure/encrypt"
)

func Test_PasswordHash(t *testing.T) {
	t.Parallel()

	password := "ddddssssdvvb"
	passwordhash := encrypt.PasswordHash{
		Type:         "scrypt",
		ScryptN:      32768,
		ScryptR:      8,
		ScryptP:      1,
		ScryptKeylen: 32,
	}
	hashedpassword, salt, err := passwordhash.GetHashed(context.Background(), password)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if err = passwordhash.Compare(hashedpassword, password, salt); err != nil {
		t.Fatalf("failed test %#v", err)
	}
	t.Log("success PasswordHash")
}
