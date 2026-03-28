package client

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

func (s *Service) ReadList(c context.Context, req *ReadReq) ([]Client, error) {
	return s.Repo.FindList(c, req.Name)
}

func (s *Service) ReadPaginated(c context.Context, req *ReadReq) ([]Client, error) {
	offset := (req.Page - 1) * req.Size
	return s.Repo.FindPaginated(c, req.Name, req.Size, offset)
}
