package repository_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/internal/account/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Suite is struct for the suite
type Suite struct {
    suite.Suite
    db *gorm.DB
    mock sqlmock.Sqlmock

    repository repository.UserRoleRepository
    userRepo repository.UserRepository
    userRole *model.Role
}

// SetupSuite is to prepare the suite instance
func (s *Suite) SetupSuite() {
    var (
        db *sql.DB
        err error
    )

    db, s.mock, err = sqlmock.New()
    assert.NoError(s.T(), err)

    s.db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
    assert.NoError(s.T(), err)

    s.repository = repository.NewUserRole(s.db)
    s.userRepo = repository.NewUser(s.db)
}

// AfterTest is hook methods performed after each test
func (s *Suite) AfterTest(_,_ string) {
    assert.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// TestInit is for initializing the test instance and the suite
func TestInit(t *testing.T) {
    suite.Run(t, new(Suite))
}

