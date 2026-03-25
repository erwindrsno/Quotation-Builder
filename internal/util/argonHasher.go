package util

import (
	"github.com/alexedwards/argon2id"
)

type ArgonHasher struct {
}

func (ah ArgonHasher) Hash(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (ah ArgonHasher) Verify(password string, hashedContext string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashedContext)
	if err != nil {
		return false, err
	}
	return match, nil
}
