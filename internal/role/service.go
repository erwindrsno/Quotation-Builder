package role

import "context"

type Service struct {
	Repo *Repository
}

func (s *Service) Create(c context.Context, req *CreateReq) error {
	if err := s.Repo.Save(c, req.Name); err != nil {
		return err
	}
	return nil
}

func (s *Service) Read(c context.Context) ([]Role, error) {
	return s.Repo.Find(c)
}
