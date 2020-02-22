package models

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

// User model
type User struct {
	ID        int64     `orm:"auto;column(id)" json:"id"`
	Email     string    `orm:"size(100);unique" json:"email"`
	Password  string    `orm:"size(100)" json:"password"`
	FirstName string    `orm:"size(150)" json:"firstName"`
	LastName  string    `orm:"size(150)" json:"lastName"`
	Phone     string    `orm:"size(150)" json:"phone"`
	Created   time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
}

// JSONResult model
type JSONResult struct {
	Status bool        `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func init() {
	orm.RegisterModel(new(User))
}

// TableName 'users' function
func (u *User) TableName() string {
	return "users"
}

// AddUser function
func AddUser(user User) (int64, error) {
	o := orm.NewOrm()

	id, err := o.Insert(&user)
	if err != nil {
		log.Println("Kullanici ekleme fonksiyonunda hata olustu. Hata: ", err)
	}
	return id, err
}

// UpdateUser function
func UpdateUser(user User) error {
	o := orm.NewOrm()

	_, err := o.Update(&user, "FirstName", "LastName", "Phone")
	if err != nil {
		log.Println("UpdateUser fonksiyonunda hata olustu. Hata: ", err)
	}

	return err
}

// GetUser function
func GetUser(userID int64) User {
	o := orm.NewOrm()

	var user = User{ID: userID}
	o.Read(&user)

	return user
}

// CheckUser function
func CheckUser(email string, password string) (User, bool) {
	o := orm.NewOrm()

	var user = User{Email: email, Password: password}
	o.Read(&user, "Email", "Password")

	return user, user.ID > 0
}
