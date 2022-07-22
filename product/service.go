package product

type Service interface {
	GetProducts() ([]Product, error)
	GetProductByID(input IDParamsInput) (Product, error)
	GetProductByCode(productCode string) ([]Product, error)
	GetProductByName(productName string) ([]Product, error)
	CreateProduct(input JsonProductInput) (Product, error)
	UpdateProduct(inputID IDParamsInput, inputData JsonProductInput) (Product, error)

	GetPrices() ([]Price, error)
	GetPriceByProductID(inputProductID int) ([]Price, error)
	GetPriceByID(inputID IDParamsInput) (Price, error)
	CreateProductPrice(inputPrice JsonPriceInput) (Price, error)
	UpdateProductPrice(inputProductID IDParamsInput, inputPrice JsonPriceInput) (Price, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts() ([]Product, error) {

	products, err := s.repository.FindAllProduct()
	if err != nil {
		return products, err
	}
	return products, nil

}

func (s *service) GetProductByCode(productCode string) ([]Product, error) {
	product, err := s.repository.FindByProductCode(productCode)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) GetProductByName(productName string) ([]Product, error) {
	product, err := s.repository.FindByProductName(productName)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) GetProductByID(input IDParamsInput) (Product, error) {
	product, err := s.repository.FindByID(input.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) CreateProduct(input JsonProductInput) (Product, error) {
	product := Product{}
	product.ProductName = input.ProductName
	product.ProductCode = input.ProductCode
	product.ProductUnit = input.ProductUnit

	newProduct, err := s.repository.CreateProduct(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) UpdateProduct(inputID IDParamsInput, inputData JsonProductInput) (Product, error) {

	// product := Product{}
	product, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return product, err
	}
	product.ProductName = inputData.ProductName
	product.ProductCode = inputData.ProductCode
	product.ProductUnit = inputData.ProductUnit

	updatedProduct, err := s.repository.UpdateProduct(product)
	if err != nil {
		return updatedProduct, err
	}
	return updatedProduct, nil
}

func (s *service) CreateProductPrice(inputPrice JsonPriceInput) (Price, error) {
	price := Price{}
	price.ProductID = inputPrice.ProductID
	price.BuyPrice = inputPrice.BuyPrice
	price.SellPrice = inputPrice.SellPrice
	price.IsPrimary = inputPrice.IsPrimary

	newPrice, err := s.repository.CreatePrice(price)
	if err != nil {
		return newPrice, err
	}

	return newPrice, nil
}

func (s *service) UpdateProductPrice(inputID IDParamsInput, inputPrice JsonPriceInput) (Price, error) {
	price, err := s.repository.FindPriceByID(inputID.ID)
	if err != nil {
		return price, err
	}
	price.ProductID = inputPrice.ProductID
	price.BuyPrice = inputPrice.BuyPrice
	price.SellPrice = inputPrice.SellPrice
	price.IsPrimary = inputPrice.IsPrimary

	newPrice, err := s.repository.UpdatePrice(price)
	if err != nil {
		return newPrice, err
	}

	return newPrice, nil
}

func (s *service) GetPrices() ([]Price, error) {
	prices, err := s.repository.FindAllPrice()
	if err != nil {
		return prices, err
	}
	return prices, nil
}

func (s *service) GetPriceByProductID(inputProductID int) ([]Price, error) {
	prices, err := s.repository.FindPriceByProductID(inputProductID)
	if err != nil {
		return prices, err
	}
	return prices, nil
}

func (s *service) GetPriceByID(inputID IDParamsInput) (Price, error) {
	price, err := s.repository.FindPriceByID(inputID.ID)
	if err != nil {
		return price, err
	}
	return price, nil
}
