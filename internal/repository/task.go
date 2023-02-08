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

const taskTable = "tasks"

var taskRetCols = []interface{}{
	gq.C("id"),
	gq.C("subject_id"),
	gq.C("description"),
}

type postgresTaskRepository struct {
	datasource *datasource.Datasource
}

func NewPostgresTaskRepository(datasource *datasource.Datasource) protocol.PostgresTaskRepository {
	return &postgresTaskRepository{datasource: datasource}
}

func (r postgresTaskRepository) Create(task *entity.Task) (*entity.Task, error) {
	ctx := context.Background()
	q, args, _ := gq.
		Insert(
			gq.T(taskTable)).
		Rows(task).
		Returning(taskRetCols...).
		ToSQL()
	if err := pgxscan.Get(ctx, r.datasource, task, q, args...); err != nil {
		return nil, err
	}
	return task, nil
}

func (r postgresTaskRepository) GetAllBySubjectID(subjectID uuid.UUID) (entity.ListTasks, error) {
	ctx := context.Background()
	var tasks entity.ListTasks
	q, args, _ := gq.
		Select(taskRetCols...).
		From(gq.T(taskTable)).
		Where(gq.C("subject_id").Eq(subjectID)).
		ToSQL()
	if err := pgxscan.Select(ctx, r.datasource, &tasks, q, args...); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r postgresTaskRepository) Update(id uuid.UUID, task *entity.Task) (*entity.Task, error) {
	ctx := context.Background()
	q, args, _ := gq.
		Update(
			gq.T(taskTable)).
		Set(task).
		Where(gq.C("id").Eq(id)).
		Returning(taskRetCols...).
		ToSQL()
	if err := pgxscan.Get(ctx, r.datasource, task, q, args...); err != nil {
		return nil, err
	}
	return task, nil
}

func (r postgresTaskRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()
	q, args, _ := gq.
		Delete(
			gq.T(taskTable)).
		Where(gq.C("id").Eq(id)).
		ToSQL()
	_, err := r.datasource.Exec(ctx, q, args...)
	return err
}
