package utils

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

func CloneRepo(url string, token string, dirname string) error {
	// Clone the repository into the current directory
	_, err := git.PlainClone("./"+dirname, false, &git.CloneOptions{
		URL:      url,
		Auth:     &http.BasicAuth{Username: "token", Password: token},
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	fmt.Println("Repository cloned successfully!")
	return nil
}

func ClonePublicRepo(url string, dirname string) error {
	// Clone the repository into the current directory
	_, err := git.PlainClone("./"+dirname, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		return err
	}

	fmt.Println("Repository cloned successfully!")
	return nil
}
