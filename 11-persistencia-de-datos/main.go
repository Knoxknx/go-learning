package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// handlers.GetContactById(db, 6)

	newContact := models.Contact{
		Name:  "test new no correo",
		Email: "",
		Phone: "123456789",
	}

	updateContact := models.Contact{
		Id:    10,
		Name:  "Update",
		Email: "correo@correo.cl",
		Phone: "123456789",
	}

	handlers.CreateContact(db, newContact)
	handlers.UpdateContact(db, updateContact)
	handlers.DeleteContact(db, 6)
	handlers.ListContacts(db)
}
