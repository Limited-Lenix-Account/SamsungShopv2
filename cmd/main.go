package main

import (
	"github.com/AlecAivazis/survey/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"samsungshop.go/internal/database"
	"samsungshop.go/internal/task"
	"samsungshop.go/internal/user/account"
	"samsungshop.go/internal/user/profiles"
)

var (
	mainOptions = []string{"manage accounts", "manage profiles", "run bot"}

	db  *database.DB
	log = logrus.New()
)

func init() {
	db, _ = database.GetDatabase()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false,
		ForceColors:   true,
	})
}

func main() {

	// loginJwt, err := jwt.MakeEcommJWT()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(loginJwt)
	// os.Exit(1)

	var mainSelect int

	prompt := &survey.Select{
		Message: "Select an Option",
		Options: mainOptions,
	}
	survey.AskOne(prompt, &mainSelect)

	switch mainSelect {
	case 0:
		// manage profiles
		account.AccountSurvey(db)
		break
	case 1:
		profiles.ProfileSurvey(db, log)
		break
	case 2:
		// run bot
		// fmt.Println("running bot")
		t := task.Task{
			DB: db,
		}

		t.Start(log)
		break
	}

}
