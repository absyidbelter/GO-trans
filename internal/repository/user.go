package repository

import (
	model "GO-Payment/internal/model/entity"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindById(id uint) (*model.User, error)
	FindByName(name string) ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(userID int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) checkDB() error {
	if r.db == nil {
		return fmt.Errorf("database connection is nil")
	}
	return nil
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	err := r.checkDB()
	if err != nil {
		return nil, err
	}

	var users []*model.User

	rows, err := r.db.Query(`SELECT id, name, email, password FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query FindAll: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows in FindAll: %v", err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to iterate rows in FindAll: %v", err)
	}

	return users, nil
}

func (r *userRepository) FindById(id uint) (*model.User, error) {
	err := r.checkDB()
	if err != nil {
		return nil, err
	}

	user := &model.User{}

	row := r.db.QueryRow(`SELECT id, name, email, password FROM users WHERE id = $1`, id)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query FindById: %v", err)
	}

	return user, nil
}

func (r *userRepository) FindByName(name string) ([]*model.User, error) {
	err := r.checkDB()
	if err != nil {
		return nil, err
	}

	var users []*model.User

	rows, err := r.db.Query(`SELECT id, name, email, password
	FROM users
	WHERE name = $1;`, name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query FindByName: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := &model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows in FindByName: %v", err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed to iterate rows in FindByName: %v", err)
	}

	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	err := r.checkDB()
	if err != nil {
		return nil, err
	}
	user := &model.User{}

	row := r.db.QueryRow(`SELECT id, name, email, password FROM users WHERE email = $1`, email)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query FindByEmail: %v", err)
	}

	return user, nil
}

func (r *userRepository) Save(user *model.User) (*model.User, error) {
	err := r.checkDB()
	if err != nil {
		return nil, err
	}
	stmt, err := r.db.Prepare(`INSERT INTO users (name, email, password) VALUES($1, $2, $3) RETURNING id`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement in Save: %v", err)
	}
	defer stmt.Close()

	var id uint
	err = stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query Save: %v", err)
	}

	user.ID = id

	return user, nil
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	err := r.checkDB()
	if err != nil {
		return nil, err
	}
	stmt, err := r.db.Prepare(`UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4 RETURNING id`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement in Update: %v", err)
	}
	defer stmt.Close()

	var id uint
	err = stmt.QueryRow(user.Name, user.Email, user.Password, user.ID).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query Update: %v", err)
	}

	return user, nil
}

func (r *userRepository) Delete(userID int) error {
	err := r.checkDB()
	if err != nil {
		return err
	}
	stmt, err := r.db.Prepare(`DELETE FROM users WHERE id = $1`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement in Delete: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return fmt.Errorf("failed to execute query Delete: %v", err)
	}

	return nil
}
