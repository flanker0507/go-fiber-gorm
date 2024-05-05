package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"log"
)

func ProductHandlerGetAll(ctx *fiber.Ctx) error {
	var products []entity.Product
	result := database.DB.Debug().Find(&products)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(products)
}

func ProductHandlerCreate(ctx *fiber.Ctx) error {
	product := new(request.ProductCreateRequest)
	if err := ctx.BodyParser(product); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(product)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newProduct := entity.Product{
		Name:      product.Name,
		Deskripsi: product.Deskripsi,
		Harga:     product.Harga,
		Stok:      product.Stok,
	}

	errCreateUser := database.DB.Create(&newProduct).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newProduct,
	})
}

func ProductHandlerGetById(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")

	var product entity.Product
	err := database.DB.First(&product, "id = ?", productId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    product,
	})
}

func ProducthandlerUpdate(ctx *fiber.Ctx) error {
	productRequest := new(request.ProductUpdateRequest)
	if err := ctx.BodyParser(productRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var product entity.Product

	productId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&product, "id = ?", productId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	// UPDATE USER DATA
	if productRequest.Name != "" {
		product.Name = productRequest.Name
	}
	product.Deskripsi = productRequest.Deskripsi
	product.Harga = productRequest.Harga
	product.Stok = productRequest.Stok
	errUpdate := database.DB.Save(&product).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    product,
	})

}

func ProductHandlerDelete(ctx *fiber.Ctx) error {
	productId := ctx.Params("id")
	var product entity.Product

	// CHECK AVAILABLE USER
	err := database.DB.Debug().First(&product, "id = ?", productId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}
	errDelete := database.DB.Debug().Delete(&product).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "product was deleted",
	})
}
