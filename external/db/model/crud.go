package model

import (
	"github.com/mochganjarn/go-template-project/external/db"
)

type Query interface {
	create(dbconn *db.Client) error
}

func CreateData(q Query, dbconn *db.Client) error {
	if err := q.create(dbconn); err != nil {
		return err
	} else {
		return nil
	}
}
