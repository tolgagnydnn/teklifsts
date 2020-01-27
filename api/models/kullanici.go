package models

// Kullanici model
type Kullanici struct {
	ID     int
	Eposta string
	Parola string
	Adi    string
	Soyadi string
}

// KullaniciBilgi function
func KullaniciBilgi() Kullanici {
	var k = Kullanici{
		ID:     1,
		Eposta: "zafercelenk@gmail.com",
		Parola: "demo",
		Adi:    "Zafer",
		Soyadi: "Çelenk",
	}

	return k
}
