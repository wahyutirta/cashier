package handler

import (
	"cashier/helper"
	"cashier/product"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	productCode := c.Query("product_code")
	if productCode != "" {

		products, err := h.service.GetProductByCode(productCode)
		if err != nil {
			response := helper.APIResponse("Gets Products Failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List Of Products", http.StatusOK, "success", product.FormatProducts(products))
		c.JSON(http.StatusOK, response)
		return
	}

	productName := c.Query("product_name")
	if productName != "" {

		products, err := h.service.GetProductByName(productName)
		if err != nil {
			response := helper.APIResponse("Gets Products Failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List Of Products", http.StatusOK, "success", product.FormatProducts(products))
		c.JSON(http.StatusOK, response)
		return
	}

	products, err := h.service.GetProducts()
	if err != nil {
		response := helper.APIResponse("Gets Products Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Of Products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
	return

}

func (h *productHandler) GetProduct(c *gin.Context) {
	//api/v1/products/2
	// handler : mapping id yg di url ke struct input => service call formatter
	// service : input struct => tangkap param id di url, panggil rep
	// repository : get campaign by id
	var input product.IDParamsInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to bind ID from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.service.GetProductByID(input)
	fmt.Printf("%+v\n", productDetail)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Product detail", http.StatusOK, "success", product.FormatProduct(productDetail))
	c.JSON(http.StatusOK, response)
	return
}

// tangkap parameter dari user ke input struct
// ambil current user dari jwt
// panggil service, parameter input struct (dan juga buat slug)
// panggil repository untuk simpan data campaign baru

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.JsonProductInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errMeesage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Product Input", http.StatusBadRequest, "error", errMeesage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := h.service.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed To Create Product Input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to Create Product", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
	return

}

// user masukan input
// handler
// mapping dari input ke input struct
// input dari user, dan juga input yang ada di uri
// service
// repository update data camaign

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID product.IDParamsInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to bind ID from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData product.JsonProductInput

	err = c.ShouldBindJSON(&inputData)

	if err != nil {
		errMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Product Input", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newProduct, err := h.service.UpdateProduct(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed To Update Product Input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to Update Product", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
	return

}

func (h *productHandler) GetPrice(c *gin.Context) {
	var inputID product.IDParamsInput
	err := c.ShouldBindUri(&inputID)

	if err != nil {
		errMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Price Input", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productPrices, err := h.service.GetPriceByID(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of price", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Price detail", http.StatusOK, "success", product.FormatPrice(productPrices))
	c.JSON(http.StatusOK, response)
	return
}

func (h *productHandler) GetPrices(c *gin.Context) {
	//api/v1/products/2
	// handler : mapping id yg di url ke struct input => service call formatter
	// service : input struct => tangkap param id di url, panggil rep
	// repository : get campaign by id

	productID := c.Query("product_id")
	if productID != "" {
		prodID, _ := strconv.Atoi(productID)
		productPrices, err := h.service.GetPriceByProductID(prodID)
		if err != nil {
			response := helper.APIResponse("Gets Price Failed", http.StatusBadRequest, "error", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		response := helper.APIResponse("List Of Price", http.StatusOK, "success", product.FormatPrices(productPrices))
		c.JSON(http.StatusOK, response)
		return
	}

	productPrices, err := h.service.GetPrices()
	if err != nil {
		response := helper.APIResponse("Failed to get detail of Price", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Price detail", http.StatusOK, "success", product.FormatPrices(productPrices))
	c.JSON(http.StatusOK, response)
	return
}

func (h *productHandler) CreatePrice(c *gin.Context) {
	var inputPrice product.JsonPriceInput
	err := c.ShouldBindJSON(&inputPrice)

	if err != nil {
		errMeesage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Price Input", http.StatusBadRequest, "error", errMeesage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPrice, err := h.service.CreateProductPrice(inputPrice)
	if err != nil {
		response := helper.APIResponse("Failed To Create Price Input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to Create Price", http.StatusOK, "success", product.FormatPrice(newPrice))
	c.JSON(http.StatusOK, response)
	return

}

func (h *productHandler) UpdatePrice(c *gin.Context) {
	var inputPriceID product.IDParamsInput
	err := c.ShouldBindUri(&inputPriceID)

	if err != nil {
		errMessage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Price Input", http.StatusBadRequest, "error", errMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputPrice product.JsonPriceInput
	err = c.ShouldBindJSON(&inputPrice)

	if err != nil {
		errMeesage := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed To Bind Price Input", http.StatusBadRequest, "error", errMeesage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPrice, err := h.service.UpdateProductPrice(inputPriceID, inputPrice)
	if err != nil {
		response := helper.APIResponse("Failed To Update Price Input", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to Update Price", http.StatusOK, "success", product.FormatPrice(newPrice))
	c.JSON(http.StatusOK, response)
	return

}
