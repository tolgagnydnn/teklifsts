package models

// Offer model
type Offer struct {
	ID          int
	Title       string
	Description string
	Text        string
}

// Offers function
func Offers() []Offer {
	var teklif1 = Offer{ID: 1, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif2 = Offer{ID: 2, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif3 = Offer{ID: 3, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif4 = Offer{ID: 4, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif5 = Offer{ID: 5, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}

	var teklifler []Offer
	teklifler = append(teklifler, teklif1, teklif2, teklif3, teklif4, teklif5)

	return teklifler
}
