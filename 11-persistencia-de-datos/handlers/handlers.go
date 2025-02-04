package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	query := "SELECT * FROM contact"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("\n LISTA DE CONTACTOS")
	fmt.Println("**********************************************************************")
	for rows.Next() {
		contact := models.Contact{}

		var valueEmail sql.NullString

		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Celular: %s \n", contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("**********************************************************************")
	}

}

func GetContactById(db *sql.DB, contactID int) {
	query := "SELECT * FROM contact where id = ?"

	row := db.QueryRow(query, contactID)
	contact := models.Contact{}
	var valueEmail sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encuentra contacto con el ID: %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo"
	}
	fmt.Println("\n LISTA DE CONTACTO")
	fmt.Println("**********************************************************************")

	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Celular: %s \n", contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("**********************************************************************")
}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Contacto creado con exito")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	result, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		log.Printf("No se encontró contacto con ID: %d", contact.Id)
		return
	}

	log.Printf("Contacto actualizado con éxito")
}

func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id = ?"

	result, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rowsAffected == 0 {
		log.Printf("No se encontró contacto con ID: %d", contactID)
		return
	}

	log.Printf("Contacto eliminado con éxito")
}
