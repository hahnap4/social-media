package database

import (
	"encoding/json"
	"fmt"
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

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	if _, ok := db.Users[email]; !ok {
		newUser := User{
			CreatedAt: time.Now().UTC(),
			Email:     email,
			Password:  password,
			Name:      name,
			Age:       age,
		}

		db.Users[email] = newUser

		c.updateDB(db)

		fmt.Println("new user added to db - success!")

		return db.Users[email], nil

	} else {
		fmt.Println(db.Users)
		return db.Users[email], fmt.Errorf("createUser: User already exists")
	}
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	if _, ok := db.Users[email]; !ok {
		return User{}, fmt.Errorf("updateUser: user doesn't exist")
	} else {
		db.Users[email] = User{
			CreatedAt: db.Users[email].CreatedAt,
			Email:     email,
			Password:  password,
			Name:      name,
			Age:       age,
		}

		c.updateDB(db)

		fmt.Println("user updated in db - success!")

		return db.Users[email], nil
	}
}

func (c Client) GetUser(email string) (User, error) {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	if _, ok := db.Users[email]; !ok {
		return User{}, fmt.Errorf("getUser: user doesn't exist")
	} else {
		fmt.Println("user info pulled from db - success!")
		return db.Users[email], nil
	}
}

func (c Client) DeleteUser(email string) error {
	latestData, _ := os.ReadFile(c.filePathToDB)
	db := databaseSchema{}
	_ = json.Unmarshal([]byte(latestData), &db)

	if _, ok := db.Users[email]; !ok {
		return fmt.Errorf("deleteUser: user doesn't exist")
	} else {
		delete(db.Users, email)
		if _, ok := db.Users[email]; !ok {
			fmt.Println("deleted user from map successfully")
			c.updateDB(db)

			fmt.Println("user deleted in db - success!")
		} else {
			return fmt.Errorf("delete failed")
		}
		fmt.Printf("delete %v is successful", db.Users[email])

		return nil
	}
}
