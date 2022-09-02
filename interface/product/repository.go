package product

import "fita-test/model"

type IProductRepository interface {
	Checkout()
	Insert(model.Add)
	SaveCart(model.Cart)
}
