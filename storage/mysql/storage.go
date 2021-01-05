package mysql

import (
	"database/sql"
	"errors"
	"exponet/expo"

	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = "user:pass@tcp(localhost)/dbname"
)

//Storage main struct
type Storage struct {
	db *sql.DB
}

//NewStorage create new storage instance
func NewStorage(dsn string) (*Storage, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return &Storage{}, err
	}

	return &Storage{db: db}, nil
}

//Insert store data
func (st *Storage) Insert(exhs []expo.Expo) error {
	var (
		err  error
		stmt *sql.Stmt
		tx   *sql.Tx
	)
	if len(exhs) == 0 {
		return errors.New("Zero len of exhs. Should be at least one item")
	}

	if tx, err = st.db.Begin(); err != nil {
		return err
	}
	if _, err = tx.Exec("TRUNCATE TABLE expos"); err != nil {
		return err
	}
	if stmt, err = tx.Prepare("INSERT INTO expos (theme,announce,description,city,start,expire) VALUES(?,?,?,?,?,?)"); err != nil {
		return err
	}

	for _, ex := range exhs {
		if _, err = stmt.Exec(ex.Name, ex.Announce, ex.Description, ex.City, ex.DateStart, ex.DateEnd); err != nil {
			break
		}
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
