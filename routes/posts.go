package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fzbian/grasfor/database"
	"github.com/gofiber/fiber/v2"
)

type Posts struct {
	Id      int    `json:"id"`
	Author  string `json:"author"`
	Message string `json:"message"`
	Ip      string `json:"ip"`
	Date    string `json:"date"`
}

func GetPosts(c *fiber.Ctx) error {
	db := database.Connect()
	rows, err := db.Query("SELECT * FROM publications")
	if err != nil {
		log.Panic(err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic(err.Error())
		}
	}(rows)

	bks := make([]*Posts, 0)
	for rows.Next() {
		bk := new(Posts)
		err := rows.Scan(&bk.Id, &bk.Author, &bk.Message, &bk.Ip, &bk.Date)
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

	postsToShow := make([]Posts, len(bks))
	for i, bk := range bks {
		postsToShow[i] = Posts{
			Id:      bk.Id,
			Author:  bk.Author,
			Message: bk.Message,
			Ip:      bk.Ip,
			Date:    bk.Date,
		}
	}
	b, _ := json.MarshalIndent(postsToShow, "", "\t")
	fmt.Printf("\n\n%v\n\n", string(b))
	err = c.JSON(postsToShow)
	if err != nil {
		return err
	}
	return nil
}
