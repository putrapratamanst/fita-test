package pkg

import (
	"fita-test/model/response"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/id"
	UNIV_TRANS "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ID_TRANS "github.com/go-playground/validator/v10/translations/id"
)

func ValidateRequest(bindType int, funcName string, ctx *gin.Context, input interface{}) *response.Response {
	// check request body
	switch bindType {
	case BIND_TYPE_JSON:
		if errBind := ctx.ShouldBindJSON(&input); errBind != nil {
			fmt.Println(errBind.Error())
			return &response.Response{
				Code:    http.StatusBadRequest,
				Message: ErrFormatRequestBody.Error(),
			}
		}
	case BIND_TYPE_PARAM:
		if errBind := ctx.ShouldBindQuery(input); errBind != nil {
			return &response.Response{
				Code:    http.StatusBadRequest,
				Message: ErrFormatRequestBody.Error(),
			}
		}

	case BIND_TYPE_URI:
		if errBind := ctx.ShouldBindUri(input); errBind != nil {
			return &response.Response{
				Code:    http.StatusBadRequest,
				Message: ErrFormatRequestBody.Error(),
			}
		}

	default:
		return &response.Response{
			Code:    http.StatusBadRequest,
			Message: ErrFormatRequestBody.Error(),
		}
	}

	// validate request body
	validate := validator.New()

	uni := UNIV_TRANS.New(id.New())
	trans, _ := uni.GetTranslator("id")

	// Verifier registration translator
	errTranslation := ID_TRANS.RegisterDefaultTranslations(validate, trans)
	if errTranslation != nil {
		return &response.Response{
			Code:    http.StatusBadRequest,
			Message: errTranslation.Error(),
		}
	}

	errTranslation = validate.Struct(input)
	msgError := ""

	if errTranslation != nil {
		for _, e := range errTranslation.(validator.ValidationErrors) {
			translatedErr := fmt.Errorf(e.Translate(trans))
			msgError = msgError + translatedErr.Error() + ". "
		}

		return &response.Response{
			Code:    http.StatusBadRequest,
			Message: msgError,
		}
	}

	return nil
}

func RandomString() string {
	length:=10
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length)
    rand.Read(b)
    return fmt.Sprintf("%x", b)[:length]
}
