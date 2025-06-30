package entity

import "github.com/ory/x/randx"

const (
	identifierLength = 128
)

type VerifyMfa struct {
	Identifier string
}

func NewVerifyMfa() *VerifyMfa {
	return &VerifyMfa{
		Identifier: generateIdentifier(),
	}
}

func generateIdentifier() string {
	result, err := randx.RuneSequence(identifierLength, randx.AlphaLower)
	if err != nil {
		panic(err)
	}
	return string(result)
}
