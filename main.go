package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/controller"
	_ "github.com/swaggo/swag/example/celler/docs"
	"github.com/swaggo/swag/example/celler/httputil"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)
func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/")
	{

		bottles := v1.Group("/cityCode")
		{
			bottles.POST(":stationName", c.cityCode)
		}
		
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			c.Abort()
		}
		c.Next()
	}
}
