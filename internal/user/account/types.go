package account

import "github.com/AlecAivazis/survey/v2"

const (
	VIEW_ACCOUNTS = "view accounts"
	ADD_ACCOUNT   = "add account"
	EXIT          = "exit"
)

var (
	accountQuestions = []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "Email"},
			Validate: survey.Required,
		},
		{
			Name:   "password",
			Prompt: &survey.Input{Message: "Password"},
		},
	}
)

type Account struct {
	Email    string `survey:"email"`
	Password string `survey:"password"`
}
