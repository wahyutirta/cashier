package product

import "fmt"

func FormatProduct(product Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.ProductCode = product.ProductCode
	productFormatter.ProductName = product.ProductName
	productFormatter.ProductUnit = product.ProductUnit
	// productFormatter.CreatedAt = product.CreatedAt
	// productFormatter.UpdatedAt = product.UpdatedAt
	productFormatter.Prices = FormatPrices(product.Price)

	return productFormatter
}

func FormatProducts(products []Product) []ProductFormatter {

	productsFormatter := []ProductFormatter{}

	fmt.Printf("1 %+v\n", products)

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}
	return productsFormatter
}

func FormatPrice(price Price) PriceFormatter {
	priceFormatter := PriceFormatter{}
	priceFormatter.ID = price.ID
	priceFormatter.ProductID = price.ProductID
	priceFormatter.BuyPrice = price.BuyPrice
	priceFormatter.SellPrice = price.SellPrice
	priceFormatter.IsPrimary = price.IsPrimary
	// priceFormatter.CreatedAt = price.CreatedAt
	// priceFormatter.UpdatedAt = price.UpdatedAt

	return priceFormatter
}

func FormatPrices(prices []Price) []PriceFormatter {

	pricesFormatter := []PriceFormatter{}

	for _, price := range prices {
		priceFormatter := FormatPrice(price)
		pricesFormatter = append(pricesFormatter, priceFormatter)
	}
	return pricesFormatter
}
