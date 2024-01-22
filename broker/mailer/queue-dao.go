package mailer

import (
	"context"

	"github.com/pydio/cells/v4/common/dao"
	"github.com/pydio/cells/v4/common/dao/boltdb"
	"github.com/pydio/cells/v4/common/dao/mongodb"
	"github.com/pydio/cells/v4/common/proto/mailer"
)

type Queue interface {
	dao.DAO
	Push(email *mailer.Mail) error
	Consume(func(email *mailer.Mail) error) error
	Close(ctx context.Context) error
}

func NewQueueDAO(ctx context.Context, o dao.DAO) (dao.DAO, error) {
	switch v := o.(type) {
	case boltdb.DAO:
		return &BoltQueue{DAO: v}, nil
	case mongodb.DAO:
		return &mongoQueue{DAO: v, runtime: ctx}, nil
	}
	return nil, dao.UnsupportedDriver(o)
}

// MigrateQueue is a MigratorFunc to move queued emails from one Queue to another.
func MigrateQueue(from dao.DAO, to dao.DAO, dryRun bool, status chan dao.MigratorStatus) (map[string]int, error) {
	out := map[string]int{
		"Emails": 0,
	}
	ctx := context.Background()
	var queueFrom, queueTo Queue
	if qf, e := NewQueueDAO(ctx, from); e == nil {
		queueFrom = qf.(Queue)
	} else {
		return nil, e
	}
	if qt, e := NewQueueDAO(ctx, to); e == nil {
		queueTo = qt.(Queue)
	} else {
		return nil, e
	}
	er := queueFrom.Consume(func(email *mailer.Mail) error {
		out["Emails"]++
		if dryRun {
			return nil
		}
		return queueTo.Push(email)
	})
	return out, er
}
