/*
INFO: 4.8-GORM | 4.9-Configurando e criando operaçoes | 4.10-Primeiras consultas
INFO: 4.11-Consultas com where | 4.12-Alterando e removendo | 4.13-Soft delete
*/
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
    gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Name: "Notebook", Price: 1000})

	// Create batch
	//products := []Product{
	//    {Name: "PlayStation 5", Price: 3299.00},
	//    {Name: "Nintedo Switch", Price: 1299.00},
	//    {Name: "Xbox Series X", Price: 2899.00},
	//}
	//db.Create(&products)

	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Nintedo Switch")
	// fmt.Println(product)

	//Select All
	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	//     fmt.Println(p)
	// }

	//Select All with options
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, p := range products {
	//     fmt.Println(p)
	// }

	/*
	   ===Where===
	   var products []Product
	   db.Where("price > ?", 1001).Find(&products)
	   for _, p := range products {
	       fmt.Println(p)
	   }

	   fmt.Println("====")

	   db.Where("name LIKE ?", "%Switch").Find(&products)
	   for _, p := range products {
	       fmt.Println(p)
	   }
	*/

	//Update
	    var p Product
	    db.First(&p, 1)
	    p.Name = "Nintendo Switch"
	    db.Save(&p)
	//
	//    var p2 Product
	//    db.First(&p2, 3)
	//    fmt.Println(p2.Name)

	//Delete => if using gorm.Model at the Struct, the gorm does a soft delete
	//db.Delete(&p)

}
