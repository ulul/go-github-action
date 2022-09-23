package book

type Service interface {
	Get() ([]Book, error)
	FindByID(id string) (Book, error)
	Create(input CreateBookRequest) (Book, error)
	Delete(id string) error
	UpdateByID(id string, input UpdateBookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewBookService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Get() ([]Book, error) {
	books, err := s.repository.Get()
	if err != nil {
		return nil, err
	}

	return books, nil

}

func (s *service) FindByID(id string) (Book, error) {

	book, err := s.repository.FindByID(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) Create(input CreateBookRequest) (Book, error) {
	book := Book{}

	book.Title = input.Title
	book.Author = input.Author
	book.Genre = input.Genre

	newBook, err := s.repository.Create(book)

	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (s *service) Delete(id string) error {

	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateByID(id string, input UpdateBookRequest) (Book, error) {
	book := Book{}

	book.Title = input.Title
	book.Author = input.Author
	book.Genre = input.Genre

	newBook, err := s.repository.UpdateByID(id, book)

	if err != nil {
		return newBook, err
	}

	return newBook, nil

}
