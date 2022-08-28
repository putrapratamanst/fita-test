package main

import (
	"fita-test/infrastructure"
	"fita-test/pkg"
	"fita-test/presenter"
	"fita-test/router"
	"fmt"
	"net/http"
	"os"

	repositoryProduct "fita-test/repository/product"
	serviceProduct "fita-test/service/product"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println("Listening LOAN API ...", serverString)

	r := gin.Default()
	redisConfig := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisCache, errRedis := infrastructure.NewDatabase(redisConfig, os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DB"))
	if errRedis != nil {
		panic(errRedis.Error())
	}

	repoProduct := repositoryProduct.NewRepository(redisCache)
	serviceProduct := serviceProduct.NewService(repoProduct)

	v1 := r.Group("/api/v1")
	{
		router.Route(v1, serviceProduct)
	}

	errorHandler(r)
	r.Run(serverString)
}


//handle error method and not found endpoint
func errorHandler(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, presenter.Response{
			Code:    http.StatusMethodNotAllowed,
			Message: pkg.ErrMethodNotAllow.Error(),
		})
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, presenter.Response{
			Code:    http.StatusNotFound,
			Message: pkg.ErrInvalidURL.Error(),
		})
		c.Abort()
	})
}
