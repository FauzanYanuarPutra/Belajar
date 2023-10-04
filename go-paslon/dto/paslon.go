package dto

type CreatePaslonDTO struct {
	Name  string `json:"name" binding:"required"`
	Visi  string `json:"visi" binding:"required"`
	Image string `json:"image" binding:"omitempty"`
}

type UpdatePaslonDTO struct {
	Name  string `json:"name,omitempty"`
	Visi  string `json:"visi,omitempty"`
	Image string `json:"image,omitempty"`
}
