package repository

import (
	"errors"
	"fmt"
	"gin_news/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	tests := []struct {
		name    string
		mock    func(user models.User)
		input   models.User
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func(user models.User) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO "+usersTable).
					WithArgs(user.Username, user.Password).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO "+personaldata).WithArgs(user.Email, user.Name).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			input: models.User{
				Email:    "testEmail",
				Name:     "testName",
				Username: "testUsername",
				Password: "testPass",
			},
			want: 1,
		},
		{
			name: "Empty Fields",
			input: models.User{
				Email:    "",
				Name:     "testName",
				Username: "testUsername",
				Password: "testPass",
			},
			mock: func(user models.User) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(1).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO "+usersTable).
					WithArgs(user.Username, user.Password).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input)

			got, err := r.CreateUser(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthPostgres_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
				AddRow(1) // return the user ID
			mock.ExpectQuery("SELECT (.+) FROM "+ usersTable).
				WithArgs("test", "password").WillReturnRows(rows)
		},
		input: args{"test", "password"},
		want: 1, // the user ID
		},
		{
			name: "Not Found",
			mock: func() {
			rows := sqlmock.NewRows([]string{"id"})
			mock.ExpectQuery("SELECT (.+) FROM "+ usersTable).
				WithArgs("not", "found").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetUser(tt.input.username, tt.input.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestAuthPostgres_GetStatusByID(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		input   int
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				query := fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", usersTable)
				rows := sqlmock.NewRows([]string{"status"}).AddRow(2)
				mock.ExpectQuery(query).
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want:  2,
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "status"})
				query := fmt.Sprintf("SELECT (.+) FROM %s WHERE (.+)", usersTable)
				mock.ExpectQuery(query).
					WithArgs(1).WillReturnRows(rows)
			},
			input:   1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetStatusByID(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
