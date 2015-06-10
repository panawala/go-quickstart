package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
	Address  string

	Emails            []Email       // One-To-Many relationship (has many)
	BillingAddress    Address       // One-To-One relationship (has one)
	BillingAddressID  sql.NullInt64 // Foreign key of BillingAddress
	ShippingAddress   Address       // One-To-One relationship (has one)
	ShippingAddressID int           // Foreign key of ShippingAddress
	IgnoreMe          int           `sql:"-"`                          // Ignore this field
	Languages         []Language    `gorm:"many2many:user_languages;"` // Many-To-Many relationship, 'user_languages' is join table
}

type Email struct {
	ID         int
	UserID     int    `sql:"index"`                          // Foreign key (belongs to), tag `index` will create index for this field when using AutoMigrate
	Email      string `sql:"type:varchar(100);unique_index"` // Set field's sql type, tag `unique_index` will create unique index
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `sql:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `sql:"type:varchar(100);unique"`
	Post     sql.NullString `sql:"not null"`
}

type Language struct {
	ID   int
	Name string `sql:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
	Code string `sql:"index:idx_name_code"` // `unique_index` also works
}

func init() {
	var dbproxy DBProxy
	dbproxy.InitDB()
	dbproxy.AutoMigrate()
}

type DBProxy struct {
	db gorm.DB
}

func (this *DBProxy) InitDB() {
	var err error
	this.db, err = gorm.Open("mysql", "william:william@/gorm?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("err happend")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("init db success")
	this.db.DB().SetMaxIdleConns(10)
	this.db.DB().SetMaxOpenConns(100)
	this.db.SingularTable(true)
}

func (this *DBProxy) AutoMigrate() {
	this.db.AutoMigrate(&User{}, &Email{}, &Address{}, &Language{})
}
