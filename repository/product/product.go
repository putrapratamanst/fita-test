package product

import (
	"context"
	"encoding/json"
	"fita-test/infrastructure"
	"fita-test/model"
	"fita-test/pkg"
)

type Repository struct {
	rc *infrastructure.RedisCache
}

func NewRepository(rc *infrastructure.RedisCache) *Repository {
	return &Repository{
		rc: rc,
	}
}

func (repo *Repository) Insert(input model.Add) {
	encodeData, _ := json.Marshal(input)
	context := context.Background()
	repo.rc.Client.Set(context, input.SKU, encodeData, 0)
}

func (repo *Repository) Checkout() {
}

func (repo *Repository) SaveCart(input model.Cart) {
	idCart := pkg.RandomString()
	encodeData, _ := json.Marshal(input)
	context := context.Background()

	repo.rc.Client.Set(context, idCart, encodeData, 0)

}
