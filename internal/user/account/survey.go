package account

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"samsungshop.go/internal/database"
)

var (
	accountOptions = []string{ADD_ACCOUNT, VIEW_ACCOUNTS, EXIT}
)

func AccountSurvey(db *database.DB) {

	for {
		var accountSelect string
		prompt := &survey.Select{
			Message: "Select an Option",
			Options: accountOptions,
		}
		survey.AskOne(prompt, &accountSelect)

		switch accountSelect {
		case VIEW_ACCOUNTS:
			// view accounts

			rows, err := db.GetAllAccounts()
			if err != nil {
				fmt.Println("error executing command: %w", err)
			}

			for rows.Next() {
				var email string

				err = rows.Scan(&email)
				fmt.Printf("[*] %s\n", email)
			}

			break
		case ADD_ACCOUNT:
			// adds accounts in user:pass form
			var lines []string
			fmt.Println("paste account list in user:pass format")
			getLines(&lines)

			fmt.Println(lines)

			for _, v := range lines {
				s := strings.Split(v, ":")
				if len(s) != 2 {
					fmt.Printf("invalid format on input: %s\n", v)
					continue
				}

				u := s[0]
				p := s[1]

				err := db.InsertAccount(u, p)
				if err != nil {
					fmt.Printf("error adding account: %s\n", err)
				}

			}

			break
		case EXIT:
			// clear terminal maybe
			return
		}

	}

}

func getLines(lines *[]string) {
	scn := bufio.NewScanner(os.Stdin)
	// var lines []string
	for scn.Scan() {
		line := scn.Text()
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		(*lines) = append((*lines), line)
	}

}
