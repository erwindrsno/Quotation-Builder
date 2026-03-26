package user

import (
	"context"
	"fmt"
	"github.com/erwindrsno/Quotation-Builder/internal/util"
)

type Service struct {
	Repo   *Repository
	Hasher util.Hasher
}

func (s *Service) Create(c context.Context, u *Register) error {
	hashed, err := s.Hasher.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed
	return s.Repo.Save(c, u)
}

func (s *Service) Read(c context.Context, data *Read) ([]User, error) {
	offset := (data.Page - 1) * data.Size
	return s.Repo.Find(c, data.Name, data.Size, offset)
}

func (s *Service) Login(c context.Context, data *Login) error {
	storedHash, err := s.Repo.FindStoredHashByUsername(c, data.Username)
	if err != nil {
		return errInvalidCredentials
	}

	match, err := s.Hasher.Verify(data.Password, storedHash)
	if err != nil {
		fmt.Printf("[AUTH ERROR] User: %s, Error: %v\n", data.Username, err)
		return errInternalError
	}
	if !match {
		return errInvalidCredentials
	}
	return nil
}
