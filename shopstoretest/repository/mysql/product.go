package mysql

import (
	"fmt"
	"shopstoretest/entity"
	"shopstoretest/param"
)

func (mysql MySQLDB) AddProduct(product param.AddProductRequest) (entity.Product, error) {
	existCategory, eErr := mysql.CheckExistCategory(product.Category)
	if eErr != nil {

		return entity.Product{}, fmt.Errorf("unexpected error %w", eErr)
	}

	if !existCategory {
		newCategory := entity.Category{Name: product.Category}
		_, aErr := mysql.AddCategory(newCategory)
		if aErr != nil {

			return entity.Product{}, fmt.Errorf("unexpected error %w", aErr)
		}
	}

	category, gErr := mysql.GetCategoryByName(product.Category)
	if gErr != nil {
		fmt.Println()
		return entity.Product{}, fmt.Errorf("unexpected error %w", gErr)
	}

	res, exErr := mysql.DB.Exec("insert into products(name, price, description, category_id, count) values(?, ?, ?, ?, ?)",
		product.Name, product.Price, product.Description, category.ID, product.Count)
	if exErr != nil {

		return entity.Product{}, exErr
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.Product{}, fmt.Errorf("cant get last id after save new product in table %w", iErr)
	}

	createdProduct := entity.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  category.ID,
		Count:       product.Count,
	}

	createdProduct.ID = uint(id)

	return createdProduct, nil
}
