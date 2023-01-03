package database

import (
	"encoding/json"
	"os"
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

func (c Client) updateDB(db databaseSchema) error {
	jsonData, err := json.Marshal(db)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.filePathToDB, jsonData, 0600)
	return nil
}

func (c Client) readDB() (databaseSchema, error) {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)
	return db, nil
}
