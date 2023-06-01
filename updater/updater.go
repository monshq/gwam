package updater

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Download(url string, workingPath string) {
	// Send an HTTP GET request to the server
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Failed to download file:", err)
	}
	defer response.Body.Close()

	// Check if the response was successful
	if response.StatusCode != http.StatusOK {
		log.Fatalln("Failed to download file. Server returned:", response.Status)
	}

	file, err := os.Create(workingPath)
	if err != nil {
		log.Fatalln("error", "Error creating file:", err)
	}
	defer file.Close()

	// Copy the content from the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatalln("Failed to save content to file:", err)
	}

	log.Printf("Saved %s", workingPath)
}
