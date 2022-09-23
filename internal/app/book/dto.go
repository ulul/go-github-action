package book

type BookFormatter struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func FormatBook(book Book) BookFormatter {
	FormattedBook := BookFormatter{
		ID:     int(book.ID),
		Title:  book.Title,
		Author: book.Author,
		Genre:  book.Genre,
	}

	return FormattedBook
}

func FormatBooks(books []Book) []BookFormatter {
	booksFormatter := []BookFormatter{}

	for _, book := range books {
		BookFormatter := FormatBook(book)
		booksFormatter = append(booksFormatter, BookFormatter)
	}

	return booksFormatter
}
