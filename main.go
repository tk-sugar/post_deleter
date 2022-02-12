package main

import (
	"fmt"
	"log"
	"time"

	"post_deleter/database"
	"post_deleter/lib"
	"post_deleter/models"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnv()

	db := database.GormConnect()
	defer db.Close()

	wp_posts := []models.WpPost{}

	status := "publish"
	to := time.Now()
	from := to.Add(-1 * time.Hour)

	db.Where("post_status = ? AND post_date BETWEEN ? AND ?", status, from.String(), to.String()).Find(&wp_posts)
	for _, wp_post := range wp_posts {
		if !lib.IsHiragana(wp_post.PostTitle) {
			fmt.Println(wp_post.PostName)
			db.Where("id = ?", wp_post.ID).Delete(wp_post)
		}
	}
}
