package product

import (
	"fita-test/model"
	"fita-test/pkg"
	"fita-test/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Controller) AddToCart(ctx *gin.Context) {
	var input model.Cart
	err := pkg.ValidateRequest(pkg.BIND_TYPE_JSON, "Add To Cart", ctx, &input)
	if err != nil {
		pkg.Response(ctx, &presenter.Response{
			Code:    err.Code,
			Message: err.Message,
		})
		return
	}

	handler.prd.AddToCart(input)
	pkg.Response(ctx, &presenter.Response{
		Code: http.StatusOK,
		Message: "Successfully Add To Cart",
	})
}
