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

func (s *Service) Read(c context.Context, req *Read) ([]User, error) {
	offset := (req.Page - 1) * req.Size
	return s.Repo.Find(c, req.Name, req.Size, offset)
}

func (s *Service) Login(c context.Context, req *Login) (string, error) {
	storedHash, err := s.Repo.FindStoredHashByUsername(c, req.Username)
	if err != nil {
		return "", errInvalidCredentials
	}

	match, err := s.Hasher.Verify(req.Password, storedHash)
	if err != nil {
		fmt.Printf("[AUTH ERROR] User: %s, Error: %v\n", req.Username, err)
		return "", errInternalError
	}
	if !match {
		return "", errInvalidCredentials
	}
	//TODO: Need to change parameter for generate token
	ss, err := util.GenerateToken(1, "1", "1")
	if err != nil {
		return "", errInternalError
	}
	return ss, nil
}
