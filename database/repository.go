package database

import (
  "database/sql"
  "errors"
  "github.com/mattn/go-sqlite3"
)

var (
  ErrDuplicate = errors.New("zaznam jiz existuje")
  ErrNotExists = errors.New("radek neexistuje")
  ErrUpdateFailed = errors.New("nelze updatovat")
  ErrDeleteFailed = errors.New("nelze smazat")
)

// db object
type SQLiteRepository struct {
  db *sql.DB
}

// create new database
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

// create table 
func (r *SQLiteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS persons(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL,
		rank INTEGER NOT NULL
	);
	`

	_, err := r.db.Exec(query)
	return err
}


// CRUD create entry
func (r *SQLiteRepository) Create(person Person) (*Person, error) {
	res, err := r.db.Exec("INSERT INTO persons(given_name, family_name, affiliation, gender, foreigner, labels) values(?,?,?,?,?,?)", person.given_name, person.family_name, person.affiliation, person.foreigner, person.labels)
        
        if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}
	
        person.ID = id

	return &person, nil
}


// return all entries
func (r *SQLiteRepository) All() ([]Person, error) {
	rows, err := r.db.Query("SELECT * FROM persons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Person
	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.given_name, &person.family_name, &person.affiliation, &person.foreigner, &person.labels); err != nil {
			return nil, err
		}
		all = append(all, person)
	}
	return all, nil
}

// get entry by family_name
func (r *SQLiteRepository) GetByFamilyName(family_name string) (*Person, error) {
	row := r.db.QueryRow("SELECT * FROM persons WHERE family_name = ?", family_name)

	var person Person
	if err := row.Scan(&person.id, &person.given_name, &person.family_name, &person.affiliation, &person.foreigner, &person.labels); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &person, nil
}


// UPDATE entry
func (r *SQLiteRepository) Update(id int64, updated Person) (*Person, error) {
	if id == 0 {
		return nil, errors.New("nespravne update ID")
	}
	res, err := r.db.Exec("UPDATE persons SET given_name = ?, family_name = ?, affiliation = ?, gender = ?, labels = ? WHERE id = ?", updated.given_name, updated.family_name, updated.affiliation, updated.gender, updated.labels, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}


// DELETE entry 
func (r *SQLiteRepository) Delete(id int64) error {
	res, err := r.db.Exec("DELETE FROM persons WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
