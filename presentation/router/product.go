package router

import (
	"github.com/gin-gonic/gin"
	"quqo_challenge/domain/db"
	"quqo_challenge/infrastructure/repository_impl"
	"quqo_challenge/infrastructure/service_impl"
	"quqo_challenge/presentation/controller"
)

func ProductRoutes(r *gin.RouterGroup, p *db.Persistence) {
	redisRepo := repository_impl.NewRedisRepository(p)
	productRepo := repository_impl.NewProductRepository(p)

	service := service_impl.NewProductService(productRepo, redisRepo)
	con := controller.NewProductController(service)

	pR := r.Group("/products")
	pR.GET("/", con.GetAllProductsController)
	pR.GET("/:id", con.GetProductByIdController)
	pR.GET("/search", con.SearchProductController)
	pR.POST("/", con.CreateProductsController)
	pR.PUT("/:id", con.UpdateProductController)
	pR.DELETE("/:id", con.DeleteProductController)
}
