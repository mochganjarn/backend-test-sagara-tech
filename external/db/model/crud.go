package model

import (
	"github.com/mochganjarn/go-template-project/external/db"
)

type Crud interface {
	create(dbconn *db.Client) error
	show(dbconn *db.Client) error
	read()
	update()
	delete()
}

func CreateData(crud Crud, dbconn *db.Client) error {
	if err := crud.create(dbconn); err != nil {
		return err
	} else {
		return nil
	}
}

func ShowData(crud Crud, dbconn *db.Client) error {
	if err := crud.show(dbconn); err != nil {
		return err
	} else {
		return nil
	}
}
