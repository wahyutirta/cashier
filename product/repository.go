package product

import (
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAllProduct() ([]Product, error)
	FindByProductCode(productCode string) ([]Product, error)
	FindByProductName(productName string) ([]Product, error)
	FindByID(ID int) (Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(product Product) (Product, error)

	CreatePrice(price Price) (Price, error)
	UpdatePrice(price Price) (Price, error)
	FindAllPrice() ([]Price, error)
	FindPriceByID(ID int) (Price, error)
	FindPriceByProductID(productID int) ([]Price, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllProduct() ([]Product, error) {
	var products []Product
	err := r.db.Preload("Price").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindByProductCode(productCode string) ([]Product, error) {
	var product []Product
	codeBuilder := fmt.Sprintf("%%%s%%", productCode)
	err := r.db.Preload("Price").Where("product_code LIKE ?", codeBuilder).Order("id desc").Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByProductName(productName string) ([]Product, error) {
	var product []Product
	nameBuilder := fmt.Sprintf("%%%s%%", productName)
	err := r.db.Preload("Price").Where("product_name LIKE ?", nameBuilder).Order("id desc").Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product
	// chain load campaign -> campaign images
	err := r.db.Preload("Price").Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) CreateProduct(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) UpdateProduct(product Product) (Product, error) {
	fmt.Printf("%+v\n", product)
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) CreatePrice(price Price) (Price, error) {
	err := r.db.Create(&price).Error
	if err != nil {
		return price, err
	}
	return price, nil
}
func (r *repository) UpdatePrice(price Price) (Price, error) {
	err := r.db.Save(&price).Error
	if err != nil {
		return price, err
	}
	return price, nil
}
func (r *repository) FindAllPrice() ([]Price, error) {
	var prices []Price
	err := r.db.Find(&prices).Error
	if err != nil {
		return prices, err
	}
	return prices, nil
}
func (r *repository) FindPriceByID(ID int) (Price, error) {
	var price Price
	// chain load campaign -> campaign images
	err := r.db.Where("id = ?", ID).Find(&price).Error
	if err != nil {
		return price, err
	}
	return price, nil
}
func (r *repository) FindPriceByProductID(productID int) ([]Price, error) {
	var prices []Price
	err := r.db.Where("product_id = ?", productID).Find(&prices).Error
	if err != nil {
		return prices, err
	}
	return prices, nil
}
