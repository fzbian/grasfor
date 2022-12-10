package utils

import (
	"fmt"
	"time"

	"github.com/fzbian/grasfor/database"
)

func AddPublication(author, message, ip string) (result string) {
	// Assigns the current date and time to the variable
	currentTime := time.Now()

	// Encrypts the ip for uploading to the database
	encrypted, err := Encrypt(ip)
	if err != nil {
		fmt.Println(err)
	}

	// The Connect function is called to obtain the SQL functions and upload the parameters passed to it.
	db := database.Connect()
	_, err = db.Exec("INSERT INTO `publications`(`id`, `author`, `message`, `ip`, `date`) VALUES (NULL, ?, ?, ?, ?)", author, message, encrypted, currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		result = fmt.Sprintf("database error: %s", err)
		return result
	}

	// Returns a success message (this will be removed in other deliveries)
	result = fmt.Sprintf("message created")
	return result
}
