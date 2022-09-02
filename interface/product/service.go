package product

import "fita-test/model"

type IProductService interface {
	Checkout()
	AddToCart(model.Cart)
	Insert(model.Add)
}
