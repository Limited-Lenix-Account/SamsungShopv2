package database

import (
	"database/sql"
	"fmt"
	"sync"
)

type DB struct {
	Db *sql.DB
	Mu *sync.Mutex
}

func GetDatabase() (*DB, error) {
	fmt.Println("getting database...")
	db, err := sql.Open("sqlite3", "internal/database/data.db")
	if err != nil {
		return nil, fmt.Errorf("error opening databse: %w", err)
	}

	return &DB{
		Db: db,
	}, nil
}

func (db *DB) InsertAccount(username, password string) error {

	q := fmt.Sprintf("INSERT INTO accounts VALUES ('%s', '%s', '', '')", username, password)

	_, err := db.Db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetAllAccounts() (*sql.Rows, error) {
	r, err := db.Db.Query("SELECT email FROM accounts")
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db *DB) UpdateAuthToken(email, access_token, jwt string) error {

	q := fmt.Sprintf("UPDATE accounts SET access_token = '%s', jwt = '%s' WHERE email = '%s'", access_token, jwt, email)
	// fmt.Println(q)
	_, err := db.Db.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetUserJWT(email string) (string, error) {
	var jwt string
	q := fmt.Sprintf("SELECT jwt FROM accounts WHERE email = '%s'", email)
	r, err := db.Db.Query(q)
	if err != nil {
		return "", err
	}

	for r.Next() {
		r.Scan(&jwt)
	}

	r.Scan(&jwt)
	return jwt, nil
}

func (db *DB) InsertProfile(p []string) error {

	q := fmt.Sprintf("INSERT INTO profiles VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')", p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7], p[8], p[9], p[10], p[11], p[12])
	_, err := db.Db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetProfile(email string) ([]*string, error) {

	var p []*string
	q := fmt.Sprintf("SELECT * FROM profiles WHERE email = '%s'", email)
	row, err := db.Db.Query(q)
	if err != nil {
		return nil, err
	}

	var prof_email string
	var first_name string
	var last_name string
	var phone_number string
	var shipping_1 string
	var shipping_2 string
	var shipping_city string
	var shipping_state string
	var shipping_zip string
	var card_number string
	var card_exp_month string
	var card_exp_year string
	var card_cvv string

	// should only run once
	for row.Next() {
		row.Scan(&prof_email, &first_name, &last_name, &phone_number, &shipping_1, &shipping_2, &shipping_city, &shipping_state, &shipping_zip, &card_number, &card_exp_month, &card_exp_year, &card_cvv)
	}

	p = append(p, &prof_email, &first_name, &last_name, &phone_number, &shipping_1, &shipping_2, &shipping_city, &shipping_state, &shipping_zip, &card_number, &card_exp_month, &card_exp_year, &card_cvv)
	return p, nil
}
