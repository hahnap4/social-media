package database

func (c Client) updateDB(db databaseSchema) error {
	//save data and overwrite the data in filePathToDB
}

func (c Client) readDB() (databaseSchema, error) {
	return new databaseSchema (with latest data from os)
}

func (c Client) CreateUser(email, password, name string, age int) (User, error){
	if emailExists, ok := User[email]; !ok {
		read current state of db
		create new user struct
		set CreatedAt to time.Now().UTC()
		add to Users map in schema
		update data on disk
	}

}