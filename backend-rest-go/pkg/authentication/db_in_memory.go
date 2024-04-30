package authentication

import "fmt"

type InMemoryAuthDB struct {
	usersBySub map[string]User
}

func NewInMemoryAuthDB() *InMemoryAuthDB {
	return &InMemoryAuthDB{
		usersBySub: make(map[string]User),
	}
}

func (db *InMemoryAuthDB) GetUserBySub(sub string) (User, bool) {
	task, ok := db.usersBySub[sub]

	return task, ok
}

func (db *InMemoryAuthDB) AddUser(user User) error {

	db.usersBySub[user.Sub] = user

	fmt.Println("db.usersBySub: ", db.usersBySub)
	return nil
}
