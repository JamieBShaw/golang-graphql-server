package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type PostgresDatabase struct {
	opts *pg.Options
}

type DBLogger struct{}

func NewDB(opts *pg.Options) *pg.DB {
	pdb := &PostgresDatabase{opts: opts}
	return pg.Connect(pdb.opts)
}

func (d DBLogger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d DBLogger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	qres, err := q.FormattedQuery()
	if err != nil {
		return err
	}

	fmt.Println("", string(qres))
	return nil
}
