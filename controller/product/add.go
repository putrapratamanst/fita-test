package product

import (
	"fita-test/model"
	"fita-test/pkg"
	"fita-test/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Controller) Add(ctx *gin.Context) {
	var input model.Add
	err := pkg.ValidateRequest(pkg.BIND_TYPE_JSON, "Add Data", ctx, &input)
	if err != nil {
		pkg.Response(ctx, &presenter.Response{
			Code:    err.Code,
			Message: err.Message,
		})
		return
	}

	handler.prd.Create(input)
	pkg.Response(ctx, &presenter.Response{
		Code: http.StatusOK,
		Message: "Successfully Create Product",
	})

}