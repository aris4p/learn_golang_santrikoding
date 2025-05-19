package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aris4p/database"
	"github.com/aris4p/helpers"
	"github.com/aris4p/models"
	"github.com/aris4p/structs"
	"github.com/gin-gonic/gin"
)

func FindProduct(c *gin.Context) {
	// inisialisasi slice untuk menampung data product
	var products []models.Product

	// Ambil data product dari database
	database.DB.Find(&products)

	// kirimkan response sukses dengan data product
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Lists Data Products",
		Data:    products,
	})
}

func CreateProduct(c *gin.Context) {
	//  struct product request
	var req = structs.ProductCreateRequest{}

	// Bind json request ke struct ProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// ambil file gambar dari form-data
	file, err := c.FormFile("image")
	var fileName string
	if err == nil {
		fileName = fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		if err := c.SaveUploadedFile(file, "uploads/"+fileName); err != nil {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to save image",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			return
		}
	}

	// skema nama file
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := c.Request.Host
	imageURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, host, fileName)

	// inisialisasi product baru
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Image:       imageURL,
	}

	// simpan product ke database
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create product",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}
	// kirimkan response sukses
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "Product Created",
		Data: structs.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Image:       product.Image,
			CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindProductById(c *gin.Context) {
	// ambil id dari parameter
	id := c.Param("id")

	// inisialisasi product
	var product models.Product

	// cari product berdasarkan id
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// kirimkan response sukses dengan data product
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Found",
		Data: structs.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateProduct(c *gin.Context) {
	// ambil id dari parameter
	id := c.Param("id")

	// inisialisasi product
	var product models.Product

	// cari product berdasarkan id
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// struct product request
	var req = structs.ProductUpdateRequest{}

	// Bind json request ke struct ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update product",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product Updated",
		Data: structs.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteProduct(c *gin.Context) {
	// ambil id dari parameter URL
	id := c.Param("id")

	// inisialisasi product
	var product models.Product

	// cari product berdasarkan ID
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Product not found",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// hapus product dari database
	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete product",
			Errors:  helpers.TranslateErrorMessage(err),
		})
		return
	}

	// kirimkan response sukses
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Product deleted successfully",
	})

}
