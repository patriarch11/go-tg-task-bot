package postgres

import (
	"context"
	gq "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/patriarch11/go-tg-task-bot/internal/domain/entity"
	"github.com/patriarch11/go-tg-task-bot/pkg/datasource/postgres"
)

var subjectTable = gq.T("subjects")

var subjectRetCols = []interface{}{
	subjectTable.Col("id"),
	subjectTable.Col("name"),
	subjectTable.Col("description"),
}

type SubjectStorage struct {
	ds postgres.Datasource
}

func NewSubjectStorage(ds postgres.Datasource) *SubjectStorage {
	return &SubjectStorage{ds}
}

func (r *SubjectStorage) Create(ctx context.Context, subject *entity.Subject) (*entity.Subject, error) {
	query, args, _ := gq.
		Insert(subjectTable).
		Rows(subject).
		Returning(subjectRetCols...).
		ToSQL()

	err := pgxscan.Get(ctx, r.ds, subject, query, args...)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

func (r *SubjectStorage) Get(ctx context.Context, id entity.ID) (*entity.Subject, error) {
	return r.getOneByFilters(ctx, subjectTable.Col("id").Eq(id))
}

func (r *SubjectStorage) GetAll(ctx context.Context) (entity.SubjectList, error) {
	var subList entity.SubjectList
	query, args, _ := gq.
		Select(subjectRetCols...).
		From(subjectTable).
		ToSQL()

	err := pgxscan.Select(ctx, r.ds, &subList, query, args...)
	if err != nil {
		return nil, err
	}
	return subList, nil
}

func (r *SubjectStorage) Update(ctx context.Context, subject *entity.Subject) (*entity.Subject, error) {
	query, args, _ := gq.
		Update(subjectTable).
		Set(subject).
		Where(
			subjectTable.Col("id").Eq(subject.Id),
		).
		Returning(subjectRetCols...).
		ToSQL()

	err := pgxscan.Get(ctx, r.ds, subject, query, args...)
	if err != nil {
		return nil, err
	}
	return subject, nil

}

func (r *SubjectStorage) Delete(ctx context.Context, id entity.ID) (err error) {
	var cmdTag pgconn.CommandTag
	query, args, _ := gq.
		Delete(subjectTable).
		Where(
			subjectTable.Col("id").Eq(id),
		).
		ToSQL()

	cmdTag, err = r.ds.Exec(ctx, query, args...)
	if err != nil {
		return
	}
	if cmdTag.RowsAffected() == 0 {
		err = entity.ErrNotFound
		return
	}
	return nil
}

func (r *SubjectStorage) getOneByFilters(ctx context.Context, filters ...gq.Expression) (*entity.Subject, error) {
	subject := new(entity.Subject)
	query, args, _ := gq.
		Select(subjectRetCols...).
		From(subjectTable).
		Where(filters...).
		Limit(1).
		ToSQL()
	err := pgxscan.Get(ctx, r.ds, subject, query, args...)
	if err != nil {
		return nil, err
	}
	return subject, nil
}
