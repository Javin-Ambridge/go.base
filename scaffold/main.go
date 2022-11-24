package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"github.com/fatih/color"
	"os/exec"
)

var (
	acceptableInputs = map[string]bool{
		"yes": true,
		"y": true,
		"ya": true,
		"no": false,
		"n": false,
	}
	maxPrompts = 5
)

func promptWrapper(msg string, retryValidator func(input string) (interface{}, error)) (interface{}, error) {
	color.Green("\n"+msg)

	retry := func() (interface{}, error) {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		sentence, err := buf.ReadBytes('\n')
		if err != nil {
			return nil, err
		}

		return retryValidator(strings.TrimSpace(string(sentence)))
	}

	for i := 0; i < maxPrompts; i++ {
		if i > 0 {
			color.Red("\nInvalid input. Please try again.")
		}

		resp, err := retry()
		if err != nil {
			continue
		}

		return resp, nil
	}

	return "", errors.New(fmt.Sprint("invalid value for %d attempts", maxPrompts))
}

func responsePrompt(msg string) (interface{}, error) {
	resp, err := promptWrapper(
		msg,
		func(input string) (interface{}, error) {
			if input == "" {
				return "", errors.New("invalid value")
			}

			return input, nil
		},
	)
	if err != nil {
		return "", err
	}

	return resp, nil
}

func yesNoPrompt(msg string) (interface{}, error) {
	resp, err := promptWrapper(
		msg + " (Y/n):",
		func(input string) (interface{}, error) {
			val, ok := acceptableInputs[strings.ToLower(input)]
			if !ok {
				return false, errors.New("invalid value")
			}

			return val, nil
		},
	)
	if err != nil {
		return false, err
	}

	return resp, nil
}

type TotalContext struct {
	FailError error

	RepositoryName string
	OriginURL string
	Description string
	GitHubUsername string
}

type Prompt struct {
	Message string
	PromptFn func(msg string) (interface{}, error)
	Action func(ctx *TotalContext, response interface{})
}

type ExecutionCmd struct {
	Message string
	Command []string
}

func executeOsCommand(cmd string) {
	if err := exec.Command("bash", "-c", cmd).Run(); err != nil {
		exit(err)
	}
}

func exit(err error) {
	color.Red("Exiting due to: " + err.Error())
	os.Exit(1)
}

func main() {
	ctx := &TotalContext{}
	totalPrompts := []Prompt{
		{
			Message: "Are you sure you want to complete this action? This will seriously alter this repository",
			PromptFn: yesNoPrompt,
			Action: func(ctx *TotalContext, response interface{}) {
				responseVal := response.(bool)

				if !responseVal {
					ctx.FailError = errors.New("not continuing")
				}
			},
		},
		{
			Message: "Please enter your new repository name: ",
			PromptFn: responsePrompt,
			Action: func(ctx *TotalContext, response interface{}) {
				ctx.RepositoryName = response.(string)
			},
		},
		{
			Message: "Please enter your new GitHub Repository origin URL (should look something like: \"https://github.com/Javin-Ambridge/example.git\"): ",
			PromptFn: responsePrompt,
			Action: func(ctx *TotalContext, response interface{}) {
				responseVal := response.(string)

				if !strings.HasPrefix(responseVal, "https://github.com/") {
					ctx.FailError = errors.New("the GitHub Repository must begin with \"https://github.com/\"")
				} else if !strings.HasSuffix(responseVal, ".git") {
					ctx.FailError = errors.New("the GitHub Repository must begin with \".git\"")
				} else {
					ctx.OriginURL = responseVal
				}
			},
		},
		{
			Message: "Please enter a small description for your Readme:",
			PromptFn: responsePrompt,
			Action: func(ctx *TotalContext, response interface{}) {
				ctx.Description = response.(string)
			},
		},
		{
			Message: "What is your GitHub name?:",
			PromptFn: responsePrompt,
			Action: func(ctx *TotalContext, response interface{}) {
				ctx.GitHubUsername = response.(string)
			},
		},
	}

	pwd, err := os.Getwd()
	if err != nil {
		exit(err)
	}
	if strings.HasSuffix(pwd, "/scaffold") {
		exit(errors.New("please run this from the top level, not inside /scaffold"))
	}

	for _, p := range totalPrompts {
		resp, err := p.PromptFn(p.Message)
		if err != nil {
			exit(err)
		}

		p.Action(ctx, resp)
		if ctx.FailError != nil {
			exit(ctx.FailError)
		}
	}

	color.Yellow("\nNew Settings. Please Confirm.")
	fmt.Println("Repository Name: " + ctx.RepositoryName)
	fmt.Println("GitHub Origin URL: " + ctx.OriginURL)
	fmt.Println("Readme Description: " + ctx.Description)
	fmt.Println("GitHub User Name: " + ctx.GitHubUsername)

	proceed, err := yesNoPrompt("Do these settings look correct to you?")
	if err != nil {
		exit(err)
	} else if !(proceed.(bool)) {
		exit(errors.New("user mistyped settings"))
	}

	executionCommands := []ExecutionCmd{
		{
			Message: "Renaming all instances of 'go.base' with \"" + ctx.RepositoryName + "\" in .go files",
			Command: []string{
				"find . -type f -name \"*.go\" -print0 | xargs -0 sed -i '' -e 's/go.base/" + ctx.RepositoryName + "/g'",
			},
		},
		{
			Message: fmt.Sprintf("Renaming GitHub UserName, \"Javin-Ambridge\" -> %q.", ctx.GitHubUsername),
			Command: []string{
				"find . -type f -name \"*.go\" -print0 | xargs -0 sed -i '' -e 's/Javin-Ambridge/" + ctx.GitHubUsername + "/g'",
				"find . -type f -name \"Makefile\" -print0 | xargs -0 sed -i '' -e 's/Javin-Ambridge/" + ctx.GitHubUsername + "/g'",
			},
		},
		{
			Message: fmt.Sprintf("Renaming all instances of 'go.base' with \"" + ctx.RepositoryName + "\" in Makefile"),
			Command: []string{
				"find . -type f -name \"Makefile\" -print0 | xargs -0 sed -i '' -e 's/go.base/" + ctx.RepositoryName + "/g'",
			},
		},
		{
			Message: "Updating README.md",
			Command: []string{
				"echo '# " + ctx.RepositoryName + " \n\n" + ctx.Description + "' > README.md",
			},
		},
		{
			Message: fmt.Sprintf("Setting new github origin URL (%q)", ctx.GitHubUsername),
			Command: []string{
				"git remote set-url origin " + ctx.OriginURL,
			},
		},
		{
			Message: fmt.Sprintf("Renaming directory ../go.base/ -> ../" + ctx.RepositoryName + "/"),
			Command: []string{
				"mv ../go.base/ ../" + ctx.RepositoryName + "/",
			},
		},
		{
			Message: "Deleting the scaffold directory",
			Command: []string{
				"rm -r scaffold/",
			},
		},
		{
			Message: "Tracking all of the new files on GitHub",
			Command: []string{
				"git add -A",
			},
		},
		{
			Message: "Adding a new commit",
			Command: []string{
				"git commit -m \"First Commit from Javin-Ambridge/go.base Scaffold.\"",
			},
		},
	}

	fmt.Print("\n")
	for index, cmd := range executionCommands {
		color.Yellow("[%d] %s", index + 1, cmd.Message)
	}

	fmt.Println(fmt.Sprintf("\nDone [%d/%d].", len(executionCommands), len(executionCommands)))
	fmt.Println("Please go back one directory so your linux shell updates the directory structure (cd ..)")
	fmt.Println("Please add (git add -A), commit (git commit -m ...), and push (git push) to your new repository now.")
	fmt.Println("Thank you!")
}