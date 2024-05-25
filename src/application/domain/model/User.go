package model

import "github.com/google/uuid"

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	UID   string `json:"uid"`
}

func CreateUser(email, uid string) *User {
	// uuid.NewRandom() はランダムなバージョン 4 の UUID を返す
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return &User{
		ID:    uuid.String(),
		Email: email,
		UID:   uid,
	}
}
