package usecase

import (
	"time"

	"github.com/pkg/errors"

	"github.com/LuigiAzevedo/public-library-v2/internal/domain/entity"
	"github.com/LuigiAzevedo/public-library-v2/internal/errs"
	r "github.com/LuigiAzevedo/public-library-v2/internal/ports/repository"
	u "github.com/LuigiAzevedo/public-library-v2/internal/ports/usecase"
)

type bookService struct {
	bookRepo r.BookRepository
}

// NewBookService creates a new instance of bookService
func NewBookService(repository r.BookRepository) u.BookUsecase {
	return &bookService{
		bookRepo: repository,
	}
}

func (s *bookService) GetBook(id int) (*entity.Book, error) {
	book, err := s.bookRepo.Get(id)
	if err != nil {
		return nil, errors.Wrap(err, errs.ErrGetBook)
	}

	return book, nil
}

func (s *bookService) SearchBook(query string) ([]*entity.Book, error) {
	books, err := s.bookRepo.Search(query)
	if err != nil {
		return nil, errors.Wrap(err, errs.ErrSearchBook)
	}

	return books, nil
}

func (s *bookService) CreateBook(b *entity.Book) (int, error) {
	book, err := entity.NewBook(b.Title, b.Author, b.Amount)
	if err != nil {
		return 0, errors.Wrap(err, errs.ErrCreateBook)
	}

	id, err := s.bookRepo.Create(book)
	if err != nil {
		return 0, errors.Wrap(err, errs.ErrCreateBook)
	}

	return id, nil
}

func (s *bookService) UpdateBook(b *entity.Book) error {
	b.UpdatedAt = time.Now()

	err := b.Validate()
	if err != nil {
		return errors.Wrap(err, errs.ErrUpdateBook)
	}

	err = s.bookRepo.Update(b)
	if err != nil {
		return errors.Wrap(err, errs.ErrUpdateBook)
	}

	return nil
}

func (s *bookService) DeleteBook(id int) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		return errors.Wrap(err, errs.ErrDeleteBook)
	}

	return nil
}