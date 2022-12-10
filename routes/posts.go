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
	id      int    `json:"id"`
	author  string `json:"author"`
	message string `json:"message"`
	ip      string `json:"ip"`
	date    string `json:"date"`
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

	for rows.Next() {
		bk := new(Posts)
		err := rows.Scan(&bk.id, &bk.author, &bk.message, &bk.ip, &bk.date)
		if err != nil {
			log.Panic(err.Error())
			return nil
		}
		fmt.Println(bk.id, bk.author, bk.message, bk.ip, bk.date)
		post := Posts{
			id:      bk.id,
			author:  bk.author,
			message: bk.message,
			ip:      bk.ip,
			date:    bk.date,
		}
		marshal, err := json.Marshal(post)
		if err != nil {
			return err
		}
		return c.JSON(marshal)
	}
	return nil
}
