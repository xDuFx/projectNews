package repository

import (
	"database/sql"
	"fmt"
	"gin_news/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewNewsPostgres(db)

	type args struct {
		listId int
		item   models.News
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name    string
		mock    mockBehavior
		input   args
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			input: args{
				listId: 1,
				item: models.News{
					Title: "test title",
					Body:  "test body",
				},
			},
			want: 2,
			mock: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO "+NewsTable).
					WithArgs(args.item.Title, args.item.Body, args.item.Image, args.item.Mark, args.item.Reliz).WillReturnRows(rows)

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.input, tt.want)

			got, err := r.Create(tt.input.item)
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

func TestPostgres_GetAll(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewNewsPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		want    []models.News
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "body", "image", "mark", "reliz"}).
					AddRow(1, "title1", "description1", "true", "true", "true").
					AddRow(2, "title2", "description2", "false", "false", "false").
					AddRow(3, "title2", "description2", "false", "false", "false")
				query := fmt.Sprintf("SELECT (.+) FROM %s", NewsTable)
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []models.News{
				{1, "title1", "description1", "true", "true", "true"},
				{2, "title2", "description2", "false", "false", "false"},
				{3, "title2", "description2", "false", "false", "false"},
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "body", "image", "mark", "reliz"})
				query := fmt.Sprintf("SELECT (.+) FROM %s", NewsTable)
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAll()
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

func TestPostgres_GetByIdNews(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewNewsPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		input   int
		want    models.News
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "body", "image", "mark", "reliz"}).
					AddRow(1, "title1", "description1", "description2", "description3", "description4")

				mock.ExpectQuery("SELECT (.+) FROM " + NewsTable).
					WithArgs(1).WillReturnRows(rows)
			},
			input: 1,
			want:  models.News{1, "title1", "description1", "description2", "description3", "description4"},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "done"})

				mock.ExpectQuery("SELECT (.+) FROM " + NewsTable).
					WithArgs(404).WillReturnRows(rows)
			},
			input:   404,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetByIdNews(tt.input)
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

func TestItemPostgres_Delete(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewNewsPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		input   int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("DELETE FROM (.+) WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: 1,
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM (.+) WHERE (.+)").
					WithArgs(404).WillReturnError(sql.ErrNoRows)
			},
			input:   404,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.DeleteNews(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestTodoItemPostgres_Update(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewNewsPostgres(db)

	type args struct {
		id    int
		input models.UpdateNews
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "OK_AllFields",
			mock: func() {
				query := fmt.Sprintf("UPDATE %s SET (.+) WHERE id = (.+)", NewsTable)
				mock.ExpectExec(query).
					WithArgs("new title", "new Body", "new Image", "new Mark", "new Reliz", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
				input: models.UpdateNews{
					Title: stringPointer("new title"),
					Body:  stringPointer("new Body"),
					Image: stringPointer("new Image"),
					Mark:  stringPointer("new Mark"),
					Reliz: stringPointer("new Reliz"),
				},
			},
		},
		{
			name: "OK_WithoutReliz",
			mock: func() {
				mock.ExpectExec("UPDATE (.+) SET (.+) WHERE id = (.+)").
					WithArgs("new title", "new Body", "new Image", "new Mark", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
				input: models.UpdateNews{
					Title: stringPointer("new title"),
					Body:  stringPointer("new Body"),
					Image: stringPointer("new Image"),
					Mark:  stringPointer("new Mark"),
				},
			},
		},
		{
			name: "OK_WithoutImageAndMark",
			mock: func() {
				mock.ExpectExec("UPDATE (.+) SET (.+) WHERE (.+)").
					WithArgs( "new title", "new Body", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
				input: models.UpdateNews{
					Title: stringPointer("new title"),
					Body:  stringPointer("new Body"),
				},
			},
		},
		{
			name: "OK_NoInputFields",
			mock: func() {
				mock.ExpectExec("UPDATE (.+) SET WHERE (.+)").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				id: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := r.UpdateNews(tt.input.id, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func stringPointer(s string) *string {
	return &s
}
