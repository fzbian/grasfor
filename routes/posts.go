package routes

import (
	"database/sql"
	"log"

	"github.com/fzbian/grasfor/database"
	"github.com/gofiber/fiber/v2"
)

type Posts struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Date    string `json:"date"`
}

func GetPosts(c *fiber.Ctx) error {

	// We connect to the database and make a request to call the data in the table publications
	db := database.Connect()

	rows, err := db.Query("SELECT id, author, message, date FROM publications")
	if err != nil {
		log.Panic(err.Error())
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic(err.Error())
		}
	}(rows)

	/* After saving the request in the rows variable, we scan the data from the
	Posts structure and add it to an array. */
	bks := make([]*Posts, 0)
	for rows.Next() {
		bk := new(Posts)
		err := rows.Scan(&bk.Id, &bk.Author, &bk.Message, &bk.Date)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Panic(err.Error())
		return nil
	}

	// we use a for to store each data and store it in another variable
	postsToShow := make([]Posts, len(bks))
	for i, bk := range bks {
		postsToShow[i] = Posts{
			Id:      bk.Id,
			Author:  bk.Author,
			Message: bk.Message,
			Date:    bk.Date,
		}
	}

	// we return the variable in a JSON
	err = c.JSON(postsToShow)
	if err != nil {
		return err
	}
	return nil
}
