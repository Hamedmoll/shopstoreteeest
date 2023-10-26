package productservice

import (
	"fmt"
	"shopstoretest/param"
)

func (s Service) AddProduct(req param.AddProductRequest) (param.AddProductResponse, error) {
	createdProduct, aErr := s.Repository.AddProduct(req)
	if aErr != nil {

		return param.AddProductResponse{}, fmt.Errorf("unexpected error %w", aErr)
	}

	productInfo := param.ProductInfo{
		Name:        createdProduct.Name,
		Count:       createdProduct.Count,
		Description: createdProduct.Description,
		Category:    req.Category,
		Price:       createdProduct.Price,
	}

	res := param.AddProductResponse{ProductInfo: productInfo}

	return res, nil
}
