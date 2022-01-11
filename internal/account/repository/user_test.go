package repository_test

import (
	"time"

	"github.com/google/uuid"
	m "github.com/reshimahendra/gin-starter/internal/account/model"
	base "github.com/reshimahendra/gin-starter/internal/database/model"
)

var (
    datetime time.Time
    column = []string{
        "id",
        "username",
        "first_name",
        "last_name",
        "email",
        "password",
        "role_id",
    }

    users = []m.User{
        {
            BaseUUID : base.BaseUUID{
                ID : uuid.New(),
                CreatedAt : datetime,
                UpdatedAt : datetime,
            },
            Username : "User1",
            Firstname : "User",
            Lastname : "One",
            Email : "user@email.com",
            Password : "password", 
            Active : true,
            RoleID : uint(1),
            Role : &m.Role{
                ID : uint(1),
                Name : "Role Test",
                Description : "Role Test descriptions",
            },
        },
    }
)

func (s *Suite)Test_USER_REPOSITORY_C(){
    // hash the pass
    // pass, err := helper.HashPassword(users[0].Password)
    // assert.NoError(s.T(), err)
    // users[0].Password = pass

    s.T().Skip()

    // prepare mock
    // s.mock.ExpectBegin()
    // q := `INSERT INTO "User" ("created_at","updated_at","username","first_name","last_name","email","active","role_id","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`
    // s.mock.ExpectQuery(regexp.QuoteMeta(q)).
    //     WithArgs(
    //         users[0].CreatedAt,
    //         users[0].UpdatedAt,
    //         users[0].Username,
    //         users[0].Firstname,
    //         users[0].Lastname,
    //         users[0].Email,
    //         users[0].Password,
    //         users[0].Active,
    //         users[0].RoleID,
    //         users[0].ID,
    //     ).
    //     WillReturnRows(
    //         sqlmock.NewRows([]string{"id"}).
    //         AddRow(users[0].ID),
    //     )
    // // q = `SELECT * FROM "Role"`
    // s.mock.ExpectQuery(regexp.QuoteMeta(q)).
    //     WillReturnRows(
    //         sqlmock.NewRows(testColumns).
    //         AddRow(
    //             users[0].Role.ID,
    //             users[0].Role.Name,
    //             users[0].Role.Description,
    //         ),
    //     )
    // s.mock.ExpectCommit()

    // because our 
    // s.mock.ExpectQuery
    
    // got, err := s.userRepo.Create(users[0])
    // assert.NoError(s.T(), err)
    // assert.Equal(s.T(), &users[0], got)

}
