package models

// Kullanici model
type Kullanici struct {
	ID     int `json:"id"`
	Eposta string `json:"eposta"`
	Parola string `json:"parola"`
	Adi    string `json:"adi"`
	Soyadi string `json:"soyadi"`
}

// JSONResult model
type JSONResult struct {
	Status bool `json:"status"`
	Error string `json:"error"`
	Data interface{} `json:"data"`
}

// KullaniciBilgi function
func KullaniciBilgi(kullaniciID int) Kullanici {
	var k = Kullanici{
		ID:     kullaniciID,
		Eposta: "zafercelenk@gmail.com",
		Parola: "demo",
		Adi:    "Zafer",
		Soyadi: "Çelenk",
	}

	return k
}
