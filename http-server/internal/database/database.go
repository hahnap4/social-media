package database

import (
	"encoding/json"
	"os"
	"time"
)

type Client struct {
	filePathToDB string
}

func NewClient(filePath string) Client {
	c := Client{
		filePathToDB: filePath,
	}

	return c
}

// Post
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

// User
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.filePathToDB)
	if err != nil {
		return c.createDB()
	}

	return err
}

func (c Client) createDB() error {
	jsonData, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})

	if err != nil {
		return err
	}

	err = os.WriteFile(c.filePathToDB, jsonData, 0600)
	if err != nil {
		return err
	}

	return nil
}
