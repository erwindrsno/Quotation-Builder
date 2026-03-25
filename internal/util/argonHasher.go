package util

import (
	"github.com/alexedwards/argon2id"
)

type ArgonHasher struct {
}

func (ah ArgonHasher) Hash(password string) (string, error) {
	hash, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (ah ArgonHasher) Verify(password string, hashedContext string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, hashedContext)
	if err != nil && !match {
		return match, err
	}
	return match, nil
}
