package database

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	if _, ok := db.Users[userEmail]; !ok {
		err := fmt.Errorf("createPost: user doesn't exist")
		return Post{}, err
	} else {
		id := uuid.New().String()

		newPost := Post{
			CreatedAt: time.Now().UTC(),
			ID:        id,
			UserEmail: db.Users[userEmail].Email,
			Text:      text,
		}

		db.Posts[id] = newPost

		c.updateDB(db)

		fmt.Println("new post added to db - success!")

		return db.Posts[id], nil
	}
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	postsArray := []Post{}

	for key := range db.Posts {
		if db.Posts[key].UserEmail == userEmail {
			postsArray = append(postsArray, db.Posts[key])
			fmt.Println("relevant post added to search list - success!")
		} else {
			fmt.Println("not relevant to search - skipping...")
		}

	}

	fmt.Println("here is the list of searched posts!")
	return postsArray, nil
}

func (c Client) DeletePost(id string) error {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	if _, ok := db.Posts[id]; !ok {
		return fmt.Errorf("deletePost: post doesn't exist")
	} else {
		delete(db.Posts, id)

		c.updateDB(db)

		fmt.Println("post deleted - success!")

		return nil
	}
}
