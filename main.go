package main

import (
	gdb "gormdemo/db"
)

func main() {
	db, err := gdb.OpenConnection("test.db", &gdb.SQLiteConnectionFactory{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&gdb.Product{})

	// Create
	db.Create(&gdb.Product{Code: "D42", Price: 100})

	// Read
	var product gdb.Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(gdb.Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
