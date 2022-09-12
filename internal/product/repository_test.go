package product

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllFail(t *testing.T) {
	testError := errors.New("Error")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM product;")).
		WillReturnError(testError)
	mock.ExpectCommit()
	mockedRepo := NewRepository(db)

	_, errResult := mockedRepo.GetAll()

	assert.Equal(t, testError, errResult)

}