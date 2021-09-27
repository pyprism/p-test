package models

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func CreateUser(firstName string, lastName string, password string) uint{
	conn := Db()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Fatalln(err)
	}

	user := User{
		FirstName:  firstName,
		LastName:   lastName,
		Password: string(bytes),
	}

	conn.Create(&user)

	return user.ID
}

func GetUserById(uid int) User{
	conn := Db()

	var u User
	conn.Find(&u, "id = ?", uid)
	return u
}

func UpdateUser(uId int, tagName string) {
	var user User
	var tag Tag
	var tagR TagRelation

	userErr := db.First(&user, "id = ?", uId).Error
	if !errors.Is(userErr, gorm.ErrRecordNotFound) {
		tagErr := db.First(&tag, "name = ?", tagName).Error
		if errors.Is(tagErr, gorm.ErrRecordNotFound) {
			fmt.Println("1")
			t := Tag{
				Name: tagName,
			}
			db.Create(&t)
			tagRErr := db.First(&tagR, "user_id = ?", user.ID).Error
			if errors.Is(tagRErr, gorm.ErrRecordNotFound) {
				fmt.Println("2")
				tagR := TagRelation{
					UserID: user.ID,
				}

				db.Create(&tagR)
				db.Create(&tagR).Association("UserID").Append(&user)
				db.Create(&tagR).Association("Tags").Append([]Tag{t})
			}
		} else {
			fmt.Println("3")
			tagRErr := db.First(&tagR, "user_id = ?", user.ID).Error
			if !errors.Is(tagRErr, gorm.ErrRecordNotFound) {
				fmt.Println("4")
				tagR := TagRelation{
					UserID: user.ID,
				}

				db.Create(&tagR)
				db.Create(&tagR).Association("UserID").Append(&user)
				db.Create(&tagR).Association("Tags").Append([]Tag{tag})
			}

		}
	} else {
		fmt.Println("u found")
	}

}
