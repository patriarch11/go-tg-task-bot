package postgres

import (
	"context"
	gq "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/patriarch11/telegram-task-manager-bot/internal/domain/entity"
	"github.com/patriarch11/telegram-task-manager-bot/pkg/datasource/postgres"
)

var taskTable = gq.T("tasks")

var taskRetCols = []interface{}{
	taskTable.Col("id"),
	taskTable.Col("subject_id"),
	taskTable.Col("description"),
}

type TaskStorage struct {
	ds postgres.Datasource
}

func NewTaskStorage(ds postgres.Datasource) *TaskStorage {
	return &TaskStorage{ds}
}

func (r *TaskStorage) Create(ctx context.Context, task *entity.Task) (t *entity.Task, err error) {
	query, args, _ := gq.
		Insert(taskTable).
		Rows(task).
		Returning(taskRetCols...).
		ToSQL()
	err = pgxscan.Get(ctx, r.ds, t, query, args...)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TaskStorage) Get(ctx context.Context, id entity.ID) (*entity.Task, error) {
	return r.getOneByFilters(ctx, taskTable.Col("id").Eq(id))
}

func (r *TaskStorage) GetBySubjectId(ctx context.Context, subjectId entity.ID) (entity.TaskList, error) {
	return r.getManyByFilters(ctx, taskTable.Col("subject_id").Eq(subjectId))
}

func (r *TaskStorage) Update(ctx context.Context, task *entity.Task) (t *entity.Task, err error) {
	query, args, _ := gq.
		Update(taskTable).
		Set(task).
		Where(
			taskTable.Col("id").Eq(task.Id),
		).
		Returning(taskRetCols...).
		ToSQL()
	err = pgxscan.Get(ctx, r.ds, t, query, args...)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (r *TaskStorage) Delete(ctx context.Context, id entity.ID) (err error) {
	var cmdTag pgconn.CommandTag

	query, args, _ := gq.
		Delete(taskTable).
		Where(
			taskTable.Col("id").Eq(id),
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

func (r *TaskStorage) getOneByFilters(ctx context.Context, filters ...gq.Expression) (task *entity.Task, err error) {
	query, args, _ := gq.
		Select(taskRetCols...).
		From(taskTable).
		Where(filters...).
		Limit(1).
		ToSQL()
	err = pgxscan.Get(ctx, r.ds, task, query, args...)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskStorage) getManyByFilters(ctx context.Context, filters ...gq.Expression) (taskList entity.TaskList, err error) {
	query, args, _ := gq.
		Select(taskRetCols...).
		From(taskTable).
		Where(filters...).
		ToSQL()
	err = pgxscan.Get(ctx, r.ds, taskList, query, args...)
	if err != nil {
		return nil, err
	}
	return taskList, nil
}
