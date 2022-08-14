package repositories

import (
	"back_end_waysbucks/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

type repositoryProduct struct {
	db *gorm.DB
}

func RepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db}
}

func (r *repositoryProduct) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error

	return products, err
}

func (r *repositoryProduct) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repositoryProduct) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repositoryProduct) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Debug().Save(&product).Error

	return product, err
}

func (r *repositoryProduct) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}
