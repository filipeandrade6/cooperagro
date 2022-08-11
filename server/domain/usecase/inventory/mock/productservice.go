package mock

import (
	"time"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type MockProductService struct {
	validID entity.ID
}

func NewMockProductService(validID entity.ID) MockProductService {
	return MockProductService{validID: validID}
}

func (m MockProductService) GetProductByID(id entity.ID) (*entity.Product, error) {
	if id == m.validID {
		return &entity.Product{
			ID:            id,
			Name:          "laranja",
			BaseProductID: entity.NewID(),
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}, nil
	}

	return nil, entity.ErrNotFound
}

func (m MockProductService) SearchProduct(query string) ([]*entity.Product, error) {
	return nil, nil
}

func (m MockProductService) ListProduct() ([]*entity.Product, error) {
	return nil, nil
}

func (m MockProductService) CreateProduct(name string, baseProduct entity.ID) (entity.ID, error) {
	return entity.NewID(), nil
}

func (m MockProductService) UpdateProduct(e *entity.Product) error {
	return nil
}

func (m MockProductService) DeleteProduct(id entity.ID) error {
	return nil
}
