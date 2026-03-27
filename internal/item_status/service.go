package item_status

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

func (s *Service) ReadList(c context.Context, req *ReadReq) ([]ItemStatus, error) {
	return s.Repo.FindList(c, req.Name)
}
