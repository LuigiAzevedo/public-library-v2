package ports

import "github.com/LuigiAzevedo/public-library-v2/internal/domain/entity"

type LoanUsecase interface {
	BorrowBook(userID, bookID int) error
	ReturnBook(userID, bookID int) error
	SearchUserLoans(userID int) ([]*entity.Loan, error)
}