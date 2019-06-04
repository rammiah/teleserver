package main

type User struct {
	Uid     string  `gorm:"column:uid;primary_key;not_null"`
	Name    string  `gorm:"column:name;not_null"`
	Pass    string  `gorm:"column:pass;not_null"`
	Mid     int32   `gorm:"column:mid;not_null"`
	Money   float32 `gorm:"column:money;not_null"`
	Overdue bool    `gorm:"column:overdue;not_null"`
}

func (User) TableName() string {
	return "user"
}

type Cashier struct {
	Uid  string `gorm:"column:uid;primary_key;not_null"`
	Name string `gorm:"column:name;not_null"`
	Pass string `gorm:"column:pass;not_null"`
	Oid  int32  `gorm:"column:oid;not_null"`
}

func (Cashier) TableName() string {
	return "cashier"
}

type Charge struct {
	Id        int32   `json:"id"`
	UserId    string  `json:"user_id"`
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
	Id     int32  `gorm:"column:id;auto_increment"`
	UserId string `gorm:"column:uid"`
	Year   int32
	Month  int32
	Day    int32
	Tm     []byte
	Cost   float32
}

func (Consume) TableName() string {
	return "consume"
}

type CustomerService struct {
	Uid  string `gorm:"column:uid;primary_key;not_null"`
	Name string `gorm:"column:name;not_null"`
	Pass string `gorm:"column:pass;not_null"`
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

type ServiceRecord struct {
	Id     int32
	UserId string
	SerId  string
	Year   int32
	Month  int32
	Day    int32
	Tm     []byte
}

func (ServiceRecord) TableName() string {
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
