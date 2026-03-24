package user

import (
	"context"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Create(c context.Context, u *Register) error {
	return s.Repo.Create(c, u)
}
