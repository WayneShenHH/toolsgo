package repositoryimpl

import (
	"github.com/WayneShenHH/toolsgo/models/entities"
)

// CreateUser creates a new user account.
func (db *datastore) CreateUser(u *entities.User) error {
	return db.mysql.Create(&u).Error
}

// GetUser gets an user by the user identifier.
func (db *datastore) GetUser(username string) (*entities.User, error) {
	u := &entities.User{}
	d := db.mysql.Where(&entities.User{Username: username}).First(&u)
	return u, d.Error
}

// GetLastUser gets the last user.
func (db *datastore) GetLastUser() (*entities.User, error) {
	u := &entities.User{}
	d := db.mysql.Last(&u)
	return u, d.Error
}

// GetUserByID 以ID取得使用者
func (db *datastore) GetUserByID(ID uint) (entities.User, error) {
	var user entities.User
	err := db.mysql.First(&user, ID).Error
	return user, err
}

// GetUserWithProfileByID 以ID取得user及其profile
func (db *datastore) GetUserWithProfileByID(ID uint) *entities.User {
	var m entities.User
	db.mysql.Find(&m, ID)
	return &m
}

// GetUserByBankID 找ID的使用者下線
func (db *datastore) GetUserByBankID(userID uint) ([]entities.User, error) {
	var users []entities.User
	err := db.mysql.Where("bank_id = ?", userID).Find(&users).Error
	return users, err
}

// GetUserByToken get user data by token
func (db *datastore) GetUserByToken(token string) (*entities.User, error) {
	var user entities.User
	err := db.mysql.Model(entities.User{}).
		Where(entities.User{
			AccessToken: token,
		}).Find(&user).Error
	return &user, err
}

// GetUserAfter gets the user who is registered after the specified user.
func (db *datastore) GetUserAfter(id int) (*entities.User, error) {
	u := &entities.User{}
	d := db.mysql.Where("id > ?", id).First(&u)
	return u, d.Error
}

// DeleteUser deletes the user by the user identifier.
func (db *datastore) DeleteUser(id uint) error {
	deleteUser := &entities.User{}
	deleteUser.ID = id
	return db.mysql.Delete(deleteUser).Error
}

// UpdateUser updates an user account information.
func (db *datastore) UpdateUser(u *entities.User) error {
	return db.mysql.Save(&u).Error
}

func (db *datastore) UpdatesUser(id uint, fields *entities.User) error {
	return db.mysql.Model(&entities.User{}).Where(`id = ?`, id).Updates(fields).Error
}

// UpdateUserPassword 更新使用者密碼
func (db *datastore) UpdateUserPassword(user *entities.User, password string) error {
	err := db.mysql.Model(user).Update("encrypted_password", password).Error
	return err
}
