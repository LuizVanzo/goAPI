package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CreatedAt   string  `json:"created_at" gorm:"autoCreateTime"`
}

func initDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=products port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados")
	}
	log.Println("Conectado ao banco de dados")
	db.AutoMigrate(&Product{})
	DB = db
}

func getProducts(c *fiber.Ctx) error {
	var products []Product
	DB.Find(&products)
	return c.JSON(products)
}

func getProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produto não encontrado"})
	}
	return c.JSON(product)
}

func createProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida"})
	}
	DB.Create(&product)
	return c.Status(201).JSON(product)
}

func updateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produto não encontrado"})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Requisição inválida"})
	}
	DB.Save(&product)
	return c.JSON(product)
}

func deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := DB.Delete(&Product{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Erro ao deletar produto"})
	}
	return c.Status(204).Send(nil)
}

func setupRoutes(app *fiber.App) {
	app.Get("/products", getProducts)
	app.Get("/products/:id", getProduct)
	app.Post("/products", createProduct)
	app.Put("/products/:id", updateProduct)
	app.Delete("/products/:id", deleteProduct)
	app.Get("/swagger/*", swagger.HandlerDefault)
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
