package public

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"shared/dbpath"
)

const (
	REPO_URL = "https://github.com/JuniorBecari10/classicist-public"
	FOLDER = "public"
)

func FetchPublicFolder() {
	log.Println("Fetching 'public' folder..")

	log.Println("Checking if it exists..")

	path := dbpath.GetPath(FOLDER)

	_, err := os.Stat(path)
	if err == nil {
    	log.Println("It does. Let's update it by pulling any changes..")

		cmd := exec.Command("git", "-C", path, "pull", "origin", "main")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("ERROR: Failed to pull changes from repository: %v", err)
		}
	} else if errors.Is(err, os.ErrNotExist) {
		log.Println("It doesn't. Let's clone it.")

		cmd := exec.Command("git", "clone", REPO_URL, path)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("ERROR: Failed to clone repository: %v", err)
		}
	}

	log.Println("Fetched 'public' folder successfully.")
}
