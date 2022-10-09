package app

import (
	"log"
)

type Ids64 struct {
	SU []int64
	US []int64
}

var Users = Ids64{
	SU: make([]int64, 0),
	US: make([]int64, 0),
}

//chats
type chat struct {
	Ms     int
	Status string
}

var Chat = chat{
	Ms:     0,
	Status: "none",
}

func GetSuperUsers(superuser int64) *Ids64 {
	if len(Users.SU) == 0 {
		Users.SU = append(Users.SU, superuser)
	}
	for _, us := range Users.SU {
		if superuser == us {
			return &Users
		} else {
			Users.SU = append(Users.SU, superuser)
			log.Println("len of Susers:", len(Users.SU))
		}
	}
	if len(Users.US) == 0 {
		Users.US = append(Users.US, superuser)
	}
	for _, us := range Users.US {
		if superuser == us {
			return &Users
		} else {
			Users.US = append(Users.US, superuser)
			log.Println("len of users:", len(Users.US))
		}
	}
	return &Users
}
