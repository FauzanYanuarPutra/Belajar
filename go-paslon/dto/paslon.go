package dto

type CreatePaslonDTO struct {
	Name string `json:"name" binding:"required"`
	Visi string `json:"visi" binding:"required"`
}

type UpdatePaslonDTO struct {
	Name  string `json:"name,omitempty"`
	Visi  string `json:"visi,omitempty"`
	Image string `json:"image,omitempty"`
}
