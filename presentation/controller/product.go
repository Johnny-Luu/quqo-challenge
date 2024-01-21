package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"quqo_challenge/domain/entity"
	"quqo_challenge/domain/network"
	"quqo_challenge/domain/service"
	"strconv"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{service: s}
}

// GetAllProductsController
// @Router /api/products [GET]
func (con *ProductController) GetAllProductsController(c *gin.Context) {
	data, err := con.service.GetAllProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, network.BuildErrorResponse(network.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, network.BuildSuccessResponse(network.Success, data))
}

// GetProductByIdController
// @Router /api/products/:id [GET]
func (con *ProductController) GetProductByIdController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	data, err := con.service.GetProductByIdService(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, network.BuildErrorResponse(network.DataNotFound))
			return
		}

		c.JSON(http.StatusInternalServerError, network.BuildErrorResponse(network.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, network.BuildSuccessResponse(network.Success, data))
}

// CreateProductsController
// @Router /api/products [POST]
func (con *ProductController) CreateProductsController(c *gin.Context) {
	var productList []entity.Product
	if err := c.ShouldBindJSON(&productList); err != nil {
		c.JSON(http.StatusBadRequest, network.BuildErrorResponse(network.InvalidRequest))
		return
	}

	data, err := con.service.CreateProductsService(productList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, network.BuildErrorResponse(network.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, network.BuildSuccessResponse(network.Success, data))
}

// UpdateProductController
// @Router /api/products/:id [PUT]
func (con *ProductController) UpdateProductController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, network.BuildErrorResponse(network.InvalidRequest))
		return
	}

	product.ID = id

	data, err := con.service.UpdateProductService(product)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, network.BuildErrorResponse(network.DataNotFound))
			return
		}

		c.JSON(http.StatusInternalServerError, network.BuildErrorResponse(network.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, network.BuildSuccessResponse(network.Success, data))
}

// DeleteProductController
// @Router /api/products/:id [DELETE]
func (con *ProductController) DeleteProductController(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	err := con.service.DeleteProductService(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, network.BuildErrorResponse(network.DataNotFound))
			return
		}

		c.JSON(http.StatusInternalServerError, network.BuildErrorResponse(network.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, network.BuildSuccessResponse(network.Success, ""))
}

// SearchProductController
// @Router /api/products/search?name= [GET]
func (con *ProductController) SearchProductController(c *gin.Context) {
	key := c.Query("name")

	data, err := con.service.SearchProductService(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, network.BuildErrorResponse(network.InternalServerError))
		return
	}

	c.JSON(http.StatusOK, network.BuildSuccessResponse(network.Success, data))
}
