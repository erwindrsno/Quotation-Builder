package quotation

import "context"

type Service struct {
	Repo *Repository
}

// TODO: Quote number generator
func (s *Service) Create(c context.Context, req *CreateReq) error {
	quotation := Quotation{
		Subject:             req.Subject,
		QuoteNumber:         req.QuoteNumber,
		Validity:            req.Validity,
		DeliveryTime:        req.DeliveryTime,
		Deadline:            req.Deadline,
		DeliveryDestination: req.DeliveryDestination,
		TermsOfPayment:      req.TermsOfPayment,
		Notes:               req.Notes,
		Discount:            req.Discount,
		ClientId:            req.ClientId,
	}

	if err := s.Repo.Save(c, quotation); err != nil {
		return err
	}
	return nil
}

func (s *Service) ReadPaginated(c context.Context, req *ReadAllReq) ([]Quotation, error) {
	offset := (req.Page - 1) * req.Size
	return s.Repo.FindPaginated(c, req.Search, req.Size, offset)
}

func (s *Service) ReadOne(c context.Context, req *ReadOneReq) (Quotation, error) {
	return s.Repo.FindById(c, req.Id)
}
