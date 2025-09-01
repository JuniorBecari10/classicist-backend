package public

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	repoURL = "https://github.com/JuniorBecari10/classicist-public"
	folder = "public"
)

func FetchPublicFolder() {
	log.Println("Fetching 'public' folder..")

	path := filepath.Join(".", folder)

	log.Println("Checking if it exists..")
	if _, err := os.Stat(path); os.IsExist(err) {
    	log.Println("It does. Let's update it by pulling any changes..")

		cmd := exec.Command("git", "-C", folder, "pull", "origin", "main")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("ERROR: Failed to pull changes from repository: %v", err)
		}
	} else {
		log.Println("It doesn't. Let's clone it.")

		cmd := exec.Command("git", "clone", repoURL, folder)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("ERROR: Failed to clone repository: %v", err)
		}
	}

	log.Println("Fetched 'public' folder successfully.")
}
