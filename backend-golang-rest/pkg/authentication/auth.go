package authentication

import "github.com/google/uuid"

type User struct {
	ID    string
	Sub   string
	Email string
}

type AuthDataSource interface {
	GetUserBySub(sub string) (User, bool)
	AddUser(user User) error
}

type Auth struct {
	db AuthDataSource
}

func (a Auth) GetUserIdBySub(sub string) (string, bool) {
	user, ok := a.db.GetUserBySub(sub)

	if !ok {

		user = User{
			ID:  uuid.New().String(),
			Sub: sub,
		}

		err := a.db.AddUser(user)

		if err != nil {
			return "", false
		}
	}

	return user.ID, true
}

func NewAuthenticator(db AuthDataSource) Auth {
	return Auth{
		db: db,
	}
}
