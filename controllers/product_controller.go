package controllers

import (
	model "hacktiv8-chapter-three-session-two/models"
	"hacktiv8-chapter-three-session-two/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{
		productService: service,
	}
}

func (controller *ProductController) CreateProduct(c *gin.Context) {
	var newProduct model.ProductCreateRequest

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	userId, userIsExist := c.Get("user_id")
	if !userIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	response, err := controller.productService.CreateProduct(newProduct, userId.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.ProductCreateResponse{
		ProductID:   response.ProductID,
		Title:       response.Title,
		Description: response.Description,
		UserID:      response.UserID,
	})
}

func (controller *ProductController) GetProduct(c *gin.Context) {
	role, roleIsExist := c.Get("role")
	if !roleIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	if role.(bool) == false {
		userId, userIsExist := c.Get("user_id")
		if !userIsExist {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Err: "Unauthorized",
			})
			return
		}

		response, err := controller.productService.GetProductByUserID(userId.(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Err: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	} else {
		response, err := controller.productService.GetAllProduct()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Err: err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (controller *ProductController) DeleteProduct(c *gin.Context) {
	productID := c.Param("product_id")

	err := controller.productService.DeleteProduct(productID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted succcessfully",
	})
}

func (controller *ProductController) UpdateProduct(c *gin.Context) {
	userId, userIsExist := c.Get("user_id")
	if !userIsExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: "Unauthorized",
		})
		return
	}

	productID := c.Param("product_id")
	var updatedProductReq model.ProductUpdateRequest

	if err := c.ShouldBindJSON(&updatedProductReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	updatedProductRes, err := controller.productService.UpdatedProduct(productID, updatedProductReq, userId.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProductRes)
}
