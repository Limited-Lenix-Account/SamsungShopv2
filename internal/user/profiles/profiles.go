package profiles

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"samsungshop.go/internal/database"
)

var (
	profileOptions = []string{VIEW_PROFILES, ADD_PROFILES, "edit profiles", LOGOUT_PROFILES, EXIT}
)

func ProfileSurvey(db *database.DB, log *logrus.Logger) {

	for {
		var profileSelect string
		prompt := &survey.Select{
			Message: "Select an Option",
			Options: profileOptions,
		}
		survey.AskOne(prompt, &profileSelect)

		switch profileSelect {
		case VIEW_PROFILES:
			// view profiles
			break
		case ADD_PROFILES:
			// read csv file and insert profiles

			profs := readProfiles()
			for i, p := range profs {
				// skip over headers
				if i == 0 {
					continue
				}
				err := db.InsertProfile(p)
				if err != nil {
					log.Error(err)
				}

				log.Infof("[%s] profile added!", p[0])
			}

			break
		case LOGOUT_PROFILES:

			// clear the JWT of all profiles in the database
			_, err := db.Db.Exec("UPDATE accounts SET jwt = '' WHERE jwt != ''")
			if err != nil {
				log.Error(err)
			}
			log.Info("removed all jwt's")

		case EXIT:
			return
		}
	}

}

func readProfiles() [][]string {

	var profs [][]string
	f, err := os.Open("profiles.csv")
	if err != nil {
		log.Fatal("cannot open profile csv")
	}
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
	}

	for _, r := range records {
		profs = append(profs, r)
	}

	return profs
}
