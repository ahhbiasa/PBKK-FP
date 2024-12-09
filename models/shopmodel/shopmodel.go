package shopmodel

import (
	"pbkk-fp/config"
	"pbkk-fp/entities"
)

func GetAll() []entities.Shop {
	rows, err := config.DB.Query("SELECT * FROM shops")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var Shops []entities.Shop

	for rows.Next() {
		var Shop entities.Shop
		if err := rows.Scan(&Shop.Id, &Shop.Name, &Shop.Address, &Shop.CreatedAt, &Shop.UpdatedAt); err != nil {
			panic(err)
		}

		Shops = append(Shops, Shop)
	}

	return Shops
}

func Create(Shop entities.Shop) bool {
	result, err := config.DB.Exec(`
	INSERT INTO Shops (name, address, created_at, updated_at)
	VALUES (?, ?, ?, ?)`,
		Shop.Name, Shop.Address, Shop.CreatedAt, Shop.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

// func Detail(id int) entities.Shop {
// 	row := config.DB.QueryRow(`
// 		SELECT
// 			shops.id,
// 			shops.name,
// 			shops.address,
// 			products.name AS product_name,
// 			categories.name AS category_name,
// 			products.stock AS product_stock
// 		FROM Shops
// 		JOIN products ON product.shop_id = shops.id
// 		JOIN categories ON products.category_id = categories.id
// 		WHERE shops.id = ?
// 	`, id)
// 	// data := config.DB.QueryRow("Select ")

// 	var Shop entities.Shop
// 	var Product entities.Product

// 	if err := row.Scan(
// 		&Shop.Id,
// 		&Shop.Name,
// 		&Shop.Address,
// 		&Product.Name,
// 		&Product.Category.Name,
// 		&Product.Stock,
// 	)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return Shop
// }

func Detail(id int) entities.Shop {
	row := config.DB.QueryRow(`
		SELECT 
			shops.id, 
			shops.name, 
			shops.address
		FROM Shops 
		WHERE shops.id = ? 
	`, id)

	var Shop entities.Shop
	if err := row.Scan(
		&Shop.Id,
		&Shop.Name,
		&Shop.Address,
	); err != nil {
		panic(err.Error())
	}

	// Now fetch the products for the shop
	rows, err := config.DB.Query(`
		SELECT 
			products.name AS product_name,
			categories.name AS category_name,
			products.stock AS product_stock 
		FROM products 
		JOIN categories ON products.category_id = categories.id
		WHERE products.shop_id = ?
	`, id)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Collect all products for this shop
	for rows.Next() {
		var Product entities.Product
		if err := rows.Scan(
			&Product.Name,
			&Product.Category.Name,
			&Product.Stock,
		); err != nil {
			panic(err.Error())
		}
		Shop.Products = append(Shop.Products, Product)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	return Shop
}

func Update(id int, Shop entities.Shop) bool {
	query, err := config.DB.Exec("UPDATE Shops SET name = ?, updated_at = ? WHERE id = ?",
		Shop.Name, Shop.UpdatedAt, id)

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
	_, err := config.DB.Exec("DELETE FROM Shops WHERE id = ?", id)
	return err
}
