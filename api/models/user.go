package models

// User model
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

// JSONResult model
type JSONResult struct {
	Status bool        `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

// GetUser function
func GetUser(userID int) User {
	var u = User{
		ID:        userID,
		Email:     "zafercelenk@gmail.com",
		Password:  "demo",
		FirstName: "Zafer",
		LastName:  "Çelenk",
		Phone:     "+90 544 245 75 99",
	}
	return u
}
