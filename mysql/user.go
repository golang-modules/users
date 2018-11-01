package mysql

import (
	"log"
	"users/models"
	"users/mysql/db"

	"golang.org/x/crypto/bcrypt"
)

type (
	//UserController - Function related to Users
	UserController struct{}
)

//New user controller
func New() *UserController {
	return &UserController{}
}

//Create - Create new employee
func (uc UserController) Create(user models.User) (int64, error) {

	insQuery, err := db.DB.Prepare("INSERT INTO users(firstname, lastname, email, password, active, created, updated) VALUES(?,?,?,?,?,NOW(),NOW())")

	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := insQuery.Exec(user.FirstName, user.LastName, user.Email, hash, false)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, err
}

//GetAll - Fetch user list
func (uc UserController) GetAll() ([]models.User, error) {

	query, err := db.DB.Query("SELECT id, firstname, lastname, email, active FROM users ORDER BY id")

	defer db.DB.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	user := models.User{}
	res := []models.User{}

	for query.Next() {
		var id int64
		var fname, lname, email string
		var active bool
		err = query.Scan(&id, &fname, &lname, &email, &active)
		if err != nil {
			return res, err
		}
		user.ID = id
		user.FirstName = fname
		user.LastName = lname
		user.Email = email
		user.Active = active
		res = append(res, user)
	}

	return res, err
}

//Get - Fetch user details
func (uc UserController) Get(id int64) (models.User, error) {

	user := models.User{}
	query, err := db.DB.Query("SELECT id, firstname, lastname, email, active FROM users WHERE id=?", id)

	if err != nil {
		log.Println(err)
		return user, err
	}

	for query.Next() {
		var id int64
		var fname, lname, email string
		var active bool
		err = query.Scan(&id, &fname, &lname, &email, &active)
		if err != nil {
			return user, err
		}
		user.ID = id
		user.FirstName = fname
		user.LastName = lname
		user.Email = email
	}

	return user, err
}

//Update - Update
func (uc UserController) Update(user models.User) (bool, error) {

	query, err := db.DB.Prepare("UPDATE users SET firstname=?, lastname=?, email=?, updated=NOW() WHERE id=?")

	if err != nil {
		log.Println(err)
		return false, err
	}

	res, err := query.Exec(user.FirstName, user.LastName, user.Email, user.ID)

	log.Println(res)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, err
}

//UpdateStatus - Update status
func (uc UserController) UpdateStatus(id int64, status bool) (bool, error) {

	query, err := db.DB.Prepare("UPDATE users SET active=?, updated=NOW() WHERE id=?")

	if err != nil {
		log.Println(err)
		return false, err
	}

	res, err := query.Exec(status, id)

	log.Println(res)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, err
}

//UpdatePassword - Update password
func (uc UserController) UpdatePassword(id int64, pass string) (bool, error) {

	query, err := db.DB.Prepare("UPDATE users SET password=?, updated=NOW() WHERE id=?")

	if err != nil {
		log.Println(err)
		return false, err
	}

	// Generate "hash" to store from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false, err
	}

	res, err := query.Exec(hash, id)

	log.Println(res)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, err
}

//Delete - Update
func (uc UserController) Delete(id int64) (bool, error) {

	query, err := db.DB.Prepare("DELETE FROM users WHERE id=?")

	if err != nil {
		log.Println(err)
		return false, err
	}

	res, err := query.Exec(id)

	log.Println(res)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, err
}
