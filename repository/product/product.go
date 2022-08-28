package product

import "fita-test/infrastructure"

type Repository struct {
	rc *infrastructure.RedisCache
}

func NewRepository(rc *infrastructure.RedisCache) *Repository {
	return &Repository{
		rc: rc,
	}
}

func (repo *Repository) Checkout() {
}
