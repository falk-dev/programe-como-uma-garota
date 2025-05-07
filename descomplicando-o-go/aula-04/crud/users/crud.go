package users

import "gorm.io/gorm"

func Create(db *gorm.DB, name string) uint {
	user := new(User)
	user.Name = name
	db.Create(user)
	return user.ID
}

func Update(db *gorm.DB, id uint, name string) {
	var user User
	db.First(&user, id)
	db.Model(&user).Update("Name", name)
}

func List(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func Delete(db *gorm.DB, id uint) {
	db.Delete(&User{}, id)
}
