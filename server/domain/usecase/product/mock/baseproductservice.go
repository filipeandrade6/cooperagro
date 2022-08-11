package mock

import (
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type MockBaseProductService struct {
	validID entity.ID
}

func NewMockBaseProductService(validID entity.ID) MockBaseProductService {
	return MockBaseProductService{validID: validID}
}

func (m MockBaseProductService) GetBaseProductByID(id entity.ID) (*entity.BaseProduct, error) {
	if id == m.validID {
		return &entity.BaseProduct{
			ID:        id,
			Name:      "laranja",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil
	}

	return nil, entity.ErrNotFound
}

func (m MockBaseProductService) SearchBaseProduct(query string) ([]*entity.BaseProduct, error) {
	return nil, nil
}

func (m MockBaseProductService) ListBaseProduct() ([]*entity.BaseProduct, error) {
	return nil, nil
}

func (m MockBaseProductService) CreateBaseProduct(name string) (entity.ID, error) {
	return entity.NewID(), nil
}

func (m MockBaseProductService) UpdateBaseProduct(e *entity.BaseProduct) error {
	return nil
}

func (m MockBaseProductService) DeleteBaseProduct(id entity.ID) error {
	return nil
}
