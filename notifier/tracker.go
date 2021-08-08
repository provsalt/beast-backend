package notifier

import (
	config2 "github.com/provsalt/beast-backend/config"
	"github.com/provsalt/beast-backend/notifier/user"
)

type Notifier struct {
	Users chan user.User

	Config config2.Config
}

func New(config *config2.Config) error {
	n := Notifier{}
	n.Users = make(chan user.User)
	return nil
}
