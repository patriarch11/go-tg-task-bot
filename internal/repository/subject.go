package repository

import (
	"context"
	gq "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/gofrs/uuid"
	"github.com/patriarch11/go-tg-task-bot/internal/entity"
	"github.com/patriarch11/go-tg-task-bot/internal/protocol"
	"github.com/patriarch11/go-tg-task-bot/pkg/datasource"
)

const subjectTable = "subjects"

var subjectRetCols = []interface{}{
	gq.C("id"),
	gq.C("name"),
	gq.C("description"),
}

type postgresSubjectRepository struct {
	datasource *datasource.Datasource
}

func NewPostgresSubjectRepository(datasource *datasource.Datasource) protocol.PostgresSubjectRepository {
	return &postgresSubjectRepository{datasource: datasource}
}

func (r postgresSubjectRepository) Create(subject *entity.Subject) (*entity.Subject, error) {
	ctx := context.Background()
	q, args, _ := gq.
		Insert(
			gq.T(subjectTable)).
		Rows(subject).
		Returning(subjectRetCols...).
		ToSQL()
	if err := pgxscan.Get(ctx, r.datasource, subject, q, args...); err != nil {
		return nil, err
	}
	return subject, nil
}

func (r postgresSubjectRepository) GetAll() (entity.ListSubjects, error) {
	ctx := context.Background()
	var subjects entity.ListSubjects
	q, args, _ := gq.
		Select(subjectRetCols...).
		From(subjectTable).
		ToSQL()
	if err := pgxscan.Select(ctx, r.datasource, &subjects, q, args...); err != nil {
		return nil, err
	}
	return subjects, nil
}

func (r postgresSubjectRepository) Update(id uuid.UUID, subject *entity.Subject) (*entity.Subject, error) {
	ctx := context.Background()
	q, args, _ := gq.
		Update(
			gq.T(subjectTable)).
		Set(subject).
		Where(gq.C("id").Eq(id)).
		Returning(subjectRetCols...).
		ToSQL()
	if err := pgxscan.Get(ctx, r.datasource, subject, q, args...); err != nil {
		return nil, err
	}
	return subject, nil
}

func (r postgresSubjectRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()
	q, args, _ := gq.
		Delete(
			gq.T(subjectTable)).
		Where(gq.C("id").Eq(id)).
		ToSQL()
	if _, err := r.datasource.Exec(ctx, q, args...); err != nil {
		return err
	}
	return nil
}
