package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
	Stock int
}
type Order struct {
	ID       uint `gorm:"primaryKey"`
	Products []ProductOder
}
type ProductOder struct {
	ID        uint
	ProductID uint
	Product   Product
	OrderID   uint
	Order     Order
	Amount    int
}
type User struct {
	//gorm.Model =>have field create update delete auto
	ID       uint
	Username string
	Profile  StudentProfile
	Role     string
}
type StudentProfile struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	CompanyName string
	JobTitle    string
	Level       string
}
type Course struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type Class struct {
	ID        uint `gorm:"primaryKey"`
	CourseID  uint
	Course    Course
	TrainerID uint
	Trainer   User
	Start     time.Time
	End       time.Time
	Seats     int
	Students  []ClassStudent
}

type ClassStudent struct {
	ID        uint `gorm:"primaryKey"`
	ClassID   uint
	StudentID uint
	Student   User
}

// type Trainer struct {
// 	ID       uint `gorm:"primaryKey"`
// 	UserID   uint
// 	User     User
// 	courseId uint
// 	Course   Course
// }

func (p *Product) UpdateStock(stock int) { //ต้องมี *
	p.Stock = stock
}

func main() {
	url := "host=localhost user=peagolang password=supersecret dbname=peagolang port=54329 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.Migrator().DropTable(
		&Product{},
		&Order{},
		&User{},
		&StudentProfile{},
		&ProductOder{},
		// &Trainer{},
		&Class{},
		&Course{},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrator().AutoMigrate(
		&Product{},
		&Order{},
		&User{},
		&StudentProfile{},
		&ProductOder{},
		// &Trainer{},
		&Class{},
		&Course{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// shirt := Product{
	// 	Name:  "T-Shirt",
	// 	Price: 350,
	// 	Stock: 200,
	// }
	// fmt.Printf("\nshirt: %+v \n", shirt) //%+v มี key #v มtype
	// id := db.Create(&shirt)
	// fmt.Printf("\nshirt id : %+v \n", id)
	// shirt.UpdateStock(999)
	// fmt.Printf("\nshirt: %+v \n", shirt) //%+v มี key #v มtype
	// // fmt.Printf("shirt: %v \nshirt: %+v \nshirt: %#v\n", shirt, shirt, shirt)  //%+v มี key #v มtype
	// db.Create(&shirt)
	// found := []Product{}
	// db.Find(&found)
	// fmt.Print(found)
	// // shirt2 := Product{
	// // 	Name:  "T-Shirt",
	// // 	Price: 350,
	// // 	Stock: 200,
	// // }

	// var shirt2 Product

	// db.First(&shirt2, 2)
	// fmt.Printf("\n %v \n", shirt2)

	// shirt2.Name="T-Shirt V2"

	// // shirt2.Name="T-Shirt V2"
	// db.Save(&shirt2)

	// db.Where("id=?",3).Delete(&Product{})

	// db.Model(&Product{}).Where("id=?",2).Update("name","hello")
	// update:=Product{
	// 	Name: "Foobar",
	// 	Stock: 500,
	// }
	// db.Model(&Product{}).Where("id = ?",5).Updates(&update)

	// db.Model(&Product{}).Where("id = ?", 4).Updates(map[string]interface{}{
	// 	"name":  "hello",
	// 	"price": 999,
	// })

	// var x int    // default val =0
	// var y string // default val =0
	// var z float32
	// var m bool
	// var n time.Time
	// fmt.Printf("x=%v y=%q z=%f m=%T n=%v", x, y, z, m, n) //%T => Type %v =>value %q=>Show String
	tdd := Course{
		Name:        "TDD",
		Description: "TDD is fun!",
	}
	db.Create(&tdd)

	pong := User{Username: "pong"}
	gap := User{Username: "gap"}
	kane := User{Username: "kane"}
	jua := User{Username: "jua"} // Trainer

	db.Create(&pong)
	db.Create(&gap)
	db.Create(&kane)
	db.Create(&jua)

	class := Class{
		CourseID:  tdd.ID,
		TrainerID: jua.ID,
		Start:     time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local),
		End:       time.Date(2023, 5, 12, 17, 0, 0, 0, time.Local),
		Seats:     10,
		Students: []ClassStudent{
			{StudentID: pong.ID},
			{StudentID: gap.ID},
		},
	}

	db.Save(&class)

	var foundClass Class
	db.Preload("Course").Preload("Trainer").Preload("Students.Student").First(&foundClass, class.ID)

	fmt.Println("#ID: ", foundClass.ID)
	fmt.Println("Name: ", foundClass.Course.Name)
	fmt.Println("Description: ", foundClass.Course.Description)
	fmt.Println("\tBy: ", foundClass.Trainer.Username)
	fmt.Println("\tDate: ", foundClass.Start, foundClass.End)
	fmt.Println("Students: ")
	for _, student := range foundClass.Students {
		fmt.Println("\tName: ", student.Student.Username)
	}
	// a := 5
	// b := &a // & => address
	// *b += 1 // * => value in address
	// fmt.Printf("a=%v b=%v x=%v", a, *b, x)
	// course := Course{
	// 	Name:        "Golang",
	// 	Description: "Bla bah bla",
	// }
	// db.Create(&course)
	// // db.Find(&Product{})
	// mike := User{
	// 	Username: "Mike",
	// }
	// pong := User{
	// 	Username: "pong",
	// }
	// abbc := User{
	// 	Username: "abbc",
	// }
	// db.Create(&mike)
	// db.Create(&pong)
	// db.Create(&abbc)
	// class := Class{
	// 	CourseID: course.id,
	// 	Start:    time.Date(2023, 5, 10, 9, 0, 0, 0, time.Local),
	// 	End:      time.Date(2023, 5, 13, 16, 0, 0, 0, time.Local),
	// 	Seats:    10,
	// 	Students: []ClassStudent{
	// 		{StudentID: mike.ID},
	// 		{StudentID: abbc.ID}}}
	// db.Create(&class)
	// user2 := User{
	// 	Username: "awdwerfweg",
	// 	Profile: StudentProfile{CompanyName: "Odd",
	// 		JobTitle: "Golang",
	// 		Level:    "555"},
	// }
	// db.Create(&mike)
	// db.Create(&user2)

	// trainer := Trainer{
	// 	UserID: user2.ID, courseId: course.ID,
	// }
	// db.Create(&trainer)
	// shirt := Product{
	// 	Name:  "T-Shirt",
	// 	Price: 350,
	// 	Stock: 200,
	// }
	// short := Product{
	// 	Name:  "Short",
	// 	Price: 600,
	// 	Stock: 150,
	// }
	// toy := Product{
	// 	Name:  "Toy",
	// 	Price: 999,
	// 	Stock: 20,
	// }
	// db.Create(&shirt)
	// db.Create(&short)
	// db.Create(&toy)
	// order1 := Order{
	// 	Products: []ProductOder{
	// 		{ProductID: shirt.ID, Amount: 1},
	// 		{ProductID: short.ID, Amount: 1},
	// 	},
	// }
	// order2 := Order{
	// 	Products: []ProductOder{
	// 		{ProductID: shirt.ID, Amount: 1},
	// 		{ProductID: toy.ID, Amount: 1},
	// 	},
	// }
	// db.Create(&order1)
	// db.Create(&order2)

	// var foundOrder Order
	// db.Preload("Products.Product").First(&foundOrder, order1.ID)
	// fmt.Printf("foundOrder=%+v ", foundOrder)
	// printOrder(foundOrder)
	// var found Product
	// db.Preload("Orders").First(&found, 1)
	// fmt.Printf("found=%+v ", found)

	// var found2 Order
	// db.Preload("Product").First(&order2, 2)
	// fmt.Printf("found=%+v ", found2)

	// user := User{

	// 	Username: "paod",
	// 	Profile: StudentProfile{
	// 		CompanyName: "Odd",
	// 		JobTitle:    "Golang",
	// 		Level:       "teqwr",
	// 	},
	// }
	// fmt.Printf("user=%+v ", user)
	// db.Save(&user)
	// db.Save(&user)
	// // db.Create(&user)
	// var foundUser User
	// db.First(&foundUser, user.ID)
	// fmt.Printf("foundUser=%+v ", foundUser)

}
func printOrder(order Order) {
	fmt.Printf("Order ID:%v\n", order.ID)
	fmt.Printf("Products:%v\n", order.Products)
	for _, p := range order.Products {
		fmt.Printf("\t%v\t%v\t%v\n", p.Product.Name, p.Product.Price, p.Amount)
	}

}
