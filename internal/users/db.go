package users

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Users struct {
	db *sql.DB
}

func NewUsers(dbPath string) (*Users, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, errors.Wrap(err, "Could not open database")
	}

	users := &Users{
		db: db,
	}

	users.seed()

	if err != nil {
		return nil, errors.Wrap(err, "Could not setup database structure")
	}

	return users, nil
}

func (u *Users) seed() error {
	_, err := u.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			session TEXT PRIMARY KEY,
			name TEXT
		);
	`)

	return err
}

func (u *Users) Put(session string, name string) (User, error) {
	user := User{
		Name: name,
	}

	statement, err := u.db.Prepare(`INSERT OR REPLACE INTO users (session, name) VALUES (?, ?)`)

	if err != nil {
		return user, errors.Wrap(err, "Could not prepare user put statement")
	}

	defer statement.Close()

	_, err = statement.Exec(session, name)

	if err != nil {
		return user, errors.Wrapf(err, "Could not update user %s -> %s", session, name)
	}

	return user, nil
}

func (u *Users) Get(session string) (User, error) {
	var user User

	statement, err := u.db.Prepare("SELECT name FROM users WHERE session = ?")

	if err != nil {
		return user, errors.Wrap(err, "Could not prepare user get statement statement")
	}

	defer statement.Close()

	row := statement.QueryRow(session)

	err = row.Scan(&user)

	if err != nil {
		return user, errors.Wrapf(err, "Could not query user %s", session)
	}

	return user, nil
}
