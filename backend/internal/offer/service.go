package offer

type Service interface {
    CreateOffer(offer *PriceOffer) error
    GetOfferByID(id uint) (*PriceOffer, error)
    GetAllOffers() ([]PriceOffer, error)
    GetOffersByUserID(userID uint) ([]PriceOffer, error)
    UpdateOffer(offer *PriceOffer) error
    DeleteOffer(id uint) error
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo}
}

func (s *service) CreateOffer(offer *PriceOffer) error {
    return s.repo.CreateOffer(offer)
}

func (s *service) GetOfferByID(id uint) (*PriceOffer, error) {
    return s.repo.GetOfferByID(id)
}

func (s *service) GetAllOffers() ([]PriceOffer, error) {
    return s.repo.GetAllOffers()
}

func (s *service) GetOffersByUserID(userID uint) ([]PriceOffer, error) {
    return s.repo.GetOffersByUserID(userID)
}

func (s *service) UpdateOffer(offer *PriceOffer) error {
    return s.repo.UpdateOffer(offer)
}

func (s *service) DeleteOffer(id uint) error {
    return s.repo.DeleteOffer(id)
}
