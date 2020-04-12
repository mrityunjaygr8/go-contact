package models

import (
	u "go-contact/utils"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	UserId uint `json:"user_id"`
}


func (contact *Contact) Validate() (map[string] interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Contact name should not be empty"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should not be empty"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (contact *Contact) Create() (map[string] interface{}) {
	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func UpdateContact(id uint, user uint, new_contact *Contact) (map[string] interface{}) {
	contact := GetContact(id, user)
	if contact == nil {
		return u.Message(false, "Either this contact does not exist or you do not have access to this contact")
	}
	
	if data, ok := new_contact.Validate(); !ok {
		return data
	}

	contact.Name = new_contact.Name
	contact.Phone = new_contact.Phone

	GetDB().Save(contact)

	resp := u.Message(true, "contact successfully updated")
	resp["contact"] = contact
	return resp
}

func DeleteContact(id uint, user uint) (map[string] interface{}) {
	contact := GetContact(id, user)
	if contact == nil {
		return u.Message(false, "Either this contact does not exist or you do not have access to this contact")
	}
	GetDB().Delete(contact)
	return u.Message(true, "Deleted Successfully")
}

func GetContact(id uint, user uint) (*Contact) {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).Where("user_id = ?", user).First(contact).Error
	if err != nil {
		return nil
	}

	return contact
}

func GetContacts(user uint) ([]*Contact) {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}