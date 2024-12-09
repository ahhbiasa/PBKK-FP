package productmodel

import (
	"pbkk-fp/config"
	"pbkk-fp/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id,
			products.name,
			categories.name AS category_name,
			products.stock,
			shops.name AS shop_name,
			products.description,
			products.created_At,
			products.updated_At
		FROM products
		JOIN categories ON products.category_id = categories.id
		JOIN shops ON products.shop_id = shops.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Shop.Name,
			&product.Description,
			&product.Created_At,
			&product.Updated_At,
		)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
	INSERT INTO products (name, category_id, stock, shop_id, description, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Shop.Id,
		product.Description,
		product.Created_At,
		product.Updated_At,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id,
			products.name,
			categories.name AS category_name,
			products.stock,
			shops.name AS shop_name,
			products.description,
			products.created_At,
			products.updated_At
		FROM products
		JOIN categories ON products.category_id = categories.id
		JOIN shops ON products.shop_id = shops.id
		where products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Shop.Name,
		&product.Description,
		&product.Created_At,
		&product.Updated_At,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		update products SET
			name = ?,
			category_id = ?,
			stock = ?,
			shop_id = ?,
			description = ?,
			updated_at = ?
		where id = ?
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Shop.Id,
		product.Description,
		product.Updated_At,
		id,
	)

	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
