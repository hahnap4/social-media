package database

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// User
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
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

		err := c.updateDB(db)

		fmt.Println("new user added to db - success!")

		return newUser, err

	}
	fmt.Println(db.Users)
	return User{}, fmt.Errorf("createUser: User already exists")
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
