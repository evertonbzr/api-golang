package repository

import (
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/evertonbzr/api-golang/pkg/infra/db"
)

type BorrowingRepository struct {
}

func NewBorrowingRepository() *BorrowingRepository {
	return &BorrowingRepository{}
}

func (s *BorrowingRepository) Create(borrowing *model.Borrowing) error {
	return db.GetDB().Create(borrowing).Error
}

func (s *BorrowingRepository) GetByID(id uint) (model.Borrowing, error) {
	borrowing := model.Borrowing{}

	if err := db.GetDB().First(&borrowing, id).Error; err != nil {
		return model.Borrowing{
			ID: 0,
		}, err
	}

	return borrowing, nil
}

func (s *BorrowingRepository) Update(borrowing model.Borrowing) error {
	var borrowingDB model.Borrowing

	if err := db.GetDB().First(&borrowingDB, borrowing.ID).Error; err != nil {
		return err
	}

	return db.GetDB().Model(&borrowingDB).Updates(borrowing).Error
}

func (s *BorrowingRepository) List() (borrowings []model.Borrowing, err error) {
	if err := db.GetDB().Find(&borrowings).Error; err != nil {
		return nil, err
	}

	return borrowings, nil
}

func (s *BorrowingRepository) ListPending() (borrowings []model.Borrowing, err error) {
	if err := db.GetDB().Where("status = ?", "borrowed").Find(&borrowings).Error; err != nil {
		return nil, err
	}

	return borrowings, nil
}
