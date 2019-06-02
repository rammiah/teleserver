package main

import "time"

type User struct {
	Uid     string `gorm:"column:uid;primary_key;not_null"`
	Name    string
	Pass    string
	Mid     int32
	Money   float32
	Overdue bool
}

func (User) TableName() string {
	return "user"
}

type Cashier struct {
	Uid  string
	Name string
	Pass string
	Oid  int32
}

func (Cashier) TableName() string {
	return "cashier"
}

type Charge struct {
	Id        int32   `json:"id"`
	UserId    string  `json:"-"`
	CashierId string  `json:"cashier_id"`
	Money     float32 `json:"money"`
	Year      int32   `json:"year"`
	Month     int32   `json:"month"`
	Day       int32   `json:"day"`
	Tm        []byte  `json:"tm"`
}

func (Charge) TableName() string {
	return "charge"
}

type Consume struct {
	Id     int32
	UserId string
	Year   int32
	Month  int32
	Day    int32
	Tm     time.Time
	Type   int32
	Cost   float32
}

func (Consume) TableName() string {
	return "consume"
}

type CustomerService struct {
	Uid  string
	Name string
	Pass string
}

func (CustomerService) TableName() string {
	return "cus_serv"
}

type Menu struct {
	Mid   int32   `json:"mid"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func (Menu) TableName() string {
	return "menu"
}

type Office struct {
	Oid  int32
	Name string
	Addr string
}

func (Office) TableName() string {
	return "office"
}

type Service struct {
	Id     int32
	UserId string
	SerId  string
	Year   int32
	Month  int32
	Day    int32
	Tm     time.Time
}

func (Service) TableName() string {
	return "services"
}

type Admin struct {
	Uid  string
	Name string
	Pass string
}

func (Admin) TableName() string {
	return "admin"
}
