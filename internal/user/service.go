package user

import (
	"context"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Create(c context.Context, u *Register) error {
	return s.Repo.Save(c, u)
}

func (s *Service) Read(c context.Context, data *Read) ([]User, error) {
	offset := (data.Page - 1) * data.Size
	return s.Repo.Find(c, data.Name, data.Size, offset)
}
