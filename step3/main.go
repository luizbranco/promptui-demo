package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Prompt{
		Label: "Search",
		Validate: func(input string) error {
			if len(input) < 3 {
				return errors.New("Search term must have at least 3 characters")
			}
			return nil
		},
	}

	keyword, err := prompt.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Search for %q\n", keyword)

	const apiURL = "http://www.recipepuppy.com/api/"

	query := apiURL + "?q=" + keyword

	client := &http.Client{Timeout: 10 * time.Second}

	r, err := client.Get(query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer r.Body.Close()
}
