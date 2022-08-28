package product

import "fita-test/service/product"

type Controller struct {
	prd *product.Service
}

func ControllerHandler(prd *product.Service) *Controller {
	handler := &Controller{
		prd: prd,
	}
	return handler
}
