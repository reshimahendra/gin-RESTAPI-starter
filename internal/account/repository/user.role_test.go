package repository_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/reshimahendra/gin-starter/internal/account/model"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)
var (
    testColumns = []string{"id", "name", "description"}
    tr = []model.Role{
        {ID:1, Name:"role 1", Description:"description1"},
        {ID:2, Name:"roleku", Description:"description2"},
        {ID:3, Name:"roleaaaa", Description:"description3"},
    }
)
// Test_USER_ROLE_REPOSITORY_C is CREATE user.role test 
func (s *Suite) Test_USER_ROLE_REPOSITORY_C() {
    // CREATE record Success
    s.mock.ExpectBegin()
    q := `INSERT INTO "Role" ("name","description","id") VALUES ($1,$2,$3) RETURNING "id"`
    s.mock.ExpectQuery(
        regexp.QuoteMeta(q)).
        WithArgs(tr[0].Name, tr[0].Description, tr[0].ID).
        WillReturnRows(
            // sqlmock.NewRows(testColumns).
            // AddRow(tr[0].ID,tr[0].Name,tr[0].Description),
            sqlmock.NewRows([]string{"id"}).
            AddRow(tr[0].ID),
        )
    s.mock.ExpectCommit()
    s.mock.ExpectQuery(regexp.QuoteMeta(
        `SELECT * FROM "Role" WHERE "Role"."id" = $1`)).
         WithArgs(tr[0].ID).
         WillReturnRows(
             sqlmock.NewRows(testColumns).
             AddRow(tr[0].ID, tr[0].Name, tr[0].Description),
         )

    got, err := s.repository.Create(tr[0])
    assert.NoError(s.T(), err)
    assert.Equal(s.T(), &tr[0], got)
}


// TestUserRoleRepositoryGet is for testing 'Get' on UserRoleRepository
func (s *Suite) Test_USER_ROLE_REPOSITORY_R() {
    s.T().Run("READ", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role" WHERE id = $1`)).
            WithArgs(tr[0].ID).
            WillReturnRows(sqlmock.NewRows(testColumns).
                AddRow(tr[0].ID, tr[0].Name, tr[0].Description))

        got, err := s.repository.Get(tr[0].ID)
        assert.NoError(s.T(), err)

        assert.Equal(s.T(), &tr[0], got)
        assert.Nil(s.T(), deep.Equal(&tr[0], got))
    })

    s.T().Run("GET FAIL record not found", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role" WHERE id = $1`)).
            WithArgs(uint(4)).
            WillReturnError(gorm.ErrRecordNotFound)

        got, err := s.repository.Get(uint(4))
        assert.Error(s.T(), err)
        assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)
        assert.Nil(s.T(), got)
    })

    // GETs
    // Get all record, expect no error
    s.T().Run("READ(s)", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role"`)).
                WillReturnRows(sqlmock.NewRows(testColumns).
                AddRow(tr[0].ID,tr[0].Name,tr[0].Description).
                AddRow(tr[1].ID,tr[1].Name,tr[1].Description).
                AddRow(tr[2].ID,tr[2].Name,tr[2].Description),
        )

        got, err := s.repository.Gets()
        assert.NoError(s.T(), err)
        assert.Equal(s.T(), &tr, got)
        assert.Nil(s.T(), deep.Equal(&tr, got))
    })

    // Get all record, expect an error
    s.T().Run("GETS FAIL empty record", func(t *testing.T){
        s.mock.ExpectQuery(
            regexp.QuoteMeta(`SELECT * FROM "Role"`)).
                WillReturnError(gorm.ErrRecordNotFound)

        got, err := s.repository.Gets()
        assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)
        assert.Nil(s.T(), got)
    })
}


// TestUserRoleRepositoryUpdate is to test Update operation 
func (s *Suite) Test_USER_ROLE_REPOSITORY_U() {
    // update operation mock
    s.mock.ExpectBegin()
    q := `UPDATE "Role" SET "name"=$1,"description"=$2 WHERE "id" = $3 AND "id" = $4`
    s.mock.ExpectExec(regexp.QuoteMeta(q)).
        WithArgs(tr[0].Name, tr[0].Description, tr[0].ID, tr[0].ID).
        WillReturnResult(sqlmock.NewResult(1,1))
    s.mock.ExpectCommit()

    // get updated data
    s.mock.ExpectQuery(regexp.QuoteMeta(
        `SELECT * FROM "Role" WHERE  "id" = $1 AND "id" = $2 AND "Role"."id" = $3`)).
         WithArgs(tr[0].ID, tr[0].ID, tr[0].ID).
         WillReturnRows(
             sqlmock.NewRows(testColumns).
             AddRow(tr[0].ID, tr[0].Name, tr[0].Description),
         )

    got, err := s.repository.Update(tr[0].ID, tr[0])
    assert.NoError(s.T(), err)
    assert.Equal(s.T(), &tr[0], got)
}
