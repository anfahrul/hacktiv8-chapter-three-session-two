package routers

import (
	"hacktiv8-chapter-three-session-two/controllers"
	"hacktiv8-chapter-three-session-two/middleware"
	"hacktiv8-chapter-three-session-two/repository"
	"hacktiv8-chapter-three-session-two/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoute(router *gin.Engine, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)
	userRepository := repository.NewUserRepository(db)
	productService := services.NewProductService(productRepository, userRepository)
	productController := controllers.NewProductController(*productService)

	productRouter := router.Group("/product", middleware.AuthMiddleware)
	{
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/", productController.GetProductByRole)
		productRouter.GET(":product_id", productController.GetOneProduct)
		productRouter.PUT(":product_id", productController.UpdateProduct)
		adminRouter := productRouter.Group("/", middleware.AdminMiddleware)
		{
			adminRouter.GET("/all", productController.GetAllProduct)
			adminRouter.DELETE(":product_id", productController.DeleteProduct)
		}
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "Berhasil")
}
