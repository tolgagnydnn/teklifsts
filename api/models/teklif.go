package models

// Teklif model
type Teklif struct {
	ID       int
	Baslik   string
	Aciklama string
	Metin    string
}

// TeklifListe function
func TeklifListe() []Teklif {
	var teklif1 = Teklif{ID: 1, Baslik: "Baslik1", Aciklama: "Aciklama1", Metin: "Metin1"}
	var teklif2 = Teklif{ID: 2, Baslik: "Baslik1", Aciklama: "Aciklama1", Metin: "Metin1"}
	var teklif3 = Teklif{ID: 3, Baslik: "Baslik1", Aciklama: "Aciklama1", Metin: "Metin1"}
	var teklif4 = Teklif{ID: 4, Baslik: "Baslik1", Aciklama: "Aciklama1", Metin: "Metin1"}
	var teklif5 = Teklif{ID: 5, Baslik: "Baslik1", Aciklama: "Aciklama1", Metin: "Metin1"}

	var teklifler []Teklif
	teklifler = append(teklifler, teklif1, teklif2, teklif3, teklif4, teklif5)

	return teklifler
}
