package util

type Hasher interface {
	Hash(password string) (string, error)
	Verify(password string, hashedContext string) (bool, error)
}
