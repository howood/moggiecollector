package entity

import "github.com/ory/x/randx"

type VerifyMfa struct {
	Identifier string
}

func NewVerifyMfa() *VerifyMfa {
	return &VerifyMfa{
		Identifier: generateIdentifier(),
	}
}

func generateIdentifier() string {
	result, err := randx.RuneSequence(128, randx.AlphaLower)
	if err != nil {
		panic(err)
	}
	return string(result)
}
