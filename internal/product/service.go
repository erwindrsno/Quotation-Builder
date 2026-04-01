package product

import "context"

type Service struct {
	Repo *Repository
}

func (s *Service) Create(c context.Context, req *CreateReq) error {
	product := Product{
		Name:              req.Name,
		PartNumber:        req.PartNumber,
		Description:       req.Description,
		BaseUnit:          req.BaseUnit,
		ManufacturerBrand: req.ManufacturerBrand,
	}
	if err := s.Repo.Save(c, product); err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadList(c context.Context, req *ReadReq) ([]Product, error) {
	return s.Repo.FindList(c, req.Name)
}
