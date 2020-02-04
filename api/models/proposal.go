package models

// Proposal model
type Proposal struct {
	ID          int
	Title       string
	Description string
	Text        string
}

// GetProposals function
func GetProposals() []Proposal {
	var teklif1 = Proposal{ID: 1, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif2 = Proposal{ID: 2, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif3 = Proposal{ID: 3, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif4 = Proposal{ID: 4, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}
	var teklif5 = Proposal{ID: 5, Title: "Baslik1", Description: "Aciklama1", Text: "Metin1"}

	var teklifler []Proposal
	teklifler = append(teklifler, teklif1, teklif2, teklif3, teklif4, teklif5)

	return teklifler
}
