package book

type CreateBookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Genre  string `json:"genre" validate:"required"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Genre  string `json:"genre" validate:"required"`
}
