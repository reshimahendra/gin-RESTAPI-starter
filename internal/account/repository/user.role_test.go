package repository_test

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"github.com/reshimahendra/gin-starter/internal/account/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
    testColumns = []string{"id", "name", "description"}
    tr = []model.Role{
        model.Role{ID:1, Name:"role1", Description:"description1"},
        model.Role{ID:2, Name:"role2", Description:"description2"},
        model.Role{ID:3, Name:"role3", Description:"description3"},
    }
)
// Suite is struct for the suite
type Suite struct {
    suite.Suite
    db *gorm.DB
    mock sqlmock.Sqlmock

    repository repository.UserRoleRepository
    userRole *model.Role
}

// SetupSuite is to prepare the suite instance
func (s *Suite) SetupSuite() {
    var (
        db *sql.DB
        err error
    )

    db, s.mock, err = sqlmock.New()
    require.NoError(s.T(), err)

    s.db, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
    require.NoError(s.T(), err)

    s.repository = repository.NewUserRole(s.db)
}

// AfterTest is hook methods performed after each test
func (s *Suite) AfterTest(_,_ string) {
    require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// TestInit is for initializing the test instance and the suite
func TestInit(t *testing.T) {
    suite.Run(t, new(Suite))
}


// TestUserRoleRepositoryGet is for testing 'Get' on UserRoleRepository
func (s *Suite) TestUserRoleRepositoryGet() {
    roleHeader := []string{
        "id",
        "name",
        "description",
    }

    type userRole struct {
        id uint
        name string
        description string
    }

    cases := []struct {
        name string
        role userRole
        want *model.Role
    }{
        {
            name: "GET with full record",
            role: userRole{
                id:1,
                name:"test role",
                description:"test role description",
            },
            want: &model.Role{
                ID:1,
                Name:"test role",
                Description:"test role description",
            },
        },
        {
            name: "GET record not found",
            role: userRole{
                id:2,
                name:"",
                description:"",
            },
            want: &model.Role{
                ID:1,
                Name:"test role",
                Description:"test role description",
            },
        },
    }

    s.T().Run(cases[0].name, func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role" WHERE id = $1`)).
            WithArgs(cases[0].role.id).
            WillReturnRows(sqlmock.NewRows(roleHeader).
                AddRow(cases[0].role.id, cases[0].role.name, cases[0].role.description))

        got, err := s.repository.Get(cases[0].role.id)
        assert.NoError(s.T(), err)

        assert.Equal(s.T(), cases[0].want, got)
        require.Nil(s.T(), deep.Equal(cases[0].want, got))
    })

    s.T().Run(cases[1].name, func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role" WHERE id = $1`)).
            WithArgs(cases[1].role.id).
            WillReturnError(gorm.ErrRecordNotFound)

        got, err := s.repository.Get(cases[1].role.id)
        assert.Error(s.T(), err)

        assert.NotEqual(s.T(), cases[1].want, got)
        require.NotNil(s.T(), deep.Equal(cases[1].want, got))
    })
}

// TestUserRoleRepositoryGets is for testing 'Gets' on UserRoleRepository
func (s *Suite) TestUserRoleRepositoryGets() {
    // Get all record, expect no error
    s.T().Run("GETS all records", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role"`)).
                WillReturnRows(sqlmock.NewRows(testColumns).
                AddRow(tr[0].ID,tr[0].Name,tr[0].Description).
                AddRow(tr[1].ID,tr[1].Name,tr[1].Description).
                AddRow(tr[2].ID,tr[2].Name,tr[2].Description),
        )

        got, err := s.repository.Gets()
        assert.NoError(s.T(), err)

        require.Nil(s.T(), deep.Equal(&tr, got))
    })
    // Get all record, expect an error
    s.T().Run("GETS empty record error", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role"`)).
                WillReturnError(gorm.ErrRecordNotFound)

        got, err := s.repository.Gets()
        assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)

        require.NotNil(s.T(), deep.Equal(&tr, got))
    })

}


func (s *Suite) TestUserRoleRepositoryCreate() {
    s.T().Skip()
    s.T().Run("Create insert 1 record", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`INSERT into "Role" ("name", "description") VALUES (?, ?, ?) RETURNING "id"`)).
            WithArgs(tr[0].Name, tr[0].Description).
            WillReturnRows(
                // sqlmock.NewRows(testColumns).
                // AddRow(tr[0].ID,tr[0].Name,tr[0].Description),
                sqlmock.NewRows([]string{"id"}).
                AddRow(tr[0].ID),
            )
        // s.mock.ExpectBegin()
        // s.mock.ExpectExec(regexp.QuoteMeta(
        //     `INSERT into "Role"`)).
        //      WithArgs(tr[0].ID, tr[0].Name, tr[0].Description).
        //      WillReturnResult(sqlmock.NewResult(int64(tr[0].ID), 1))
        // s.mock.ExpectCommit()

        got, err := s.repository.Create(model.Role{
            ID: uint(1),
            Name: "Pletan",
            Description: "Deskripsi pletan",
        })
        // assert.NoError(s.T(), err)
        //
        t.Logf("got: %v,\nError: %v", got, err)
        // require.NoError(t, err)

        // assert.Equal(t, tr[0], got)
    })
}
