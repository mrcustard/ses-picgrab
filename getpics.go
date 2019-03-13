package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	newPartsListFile = "NewPartsList.txt" // file where parts list is stored
)

type partFile struct { // I don't really need this at this point
	PartNo      string // but it could be intersting to pull
	Description string // things into this struct
	Date        string // Could use this for reading in the CSV directly
}

func main() {
	readPartsListFile()
}

func readPartsListFile() {
	//Open the file and send the contents over to download_images
	file, err := ioutil.ReadFile(newPartsListFile)
	if err != nil {
		fmt.Printf("Error reading file: %v", newPartsListFile)
		os.Exit(1)
	}
	downloadImages(file)
}

func downloadImages(b []byte) {
	//here we want to download the images

	/* Seudo code:
	for part_number := range b {
		resp, err := http.GET(URL)
		if err != {
			fmt.Printf("Error downloading: %v", resp.Body)
		}
		copy to /location/resp.Body
	}

	*/
}

// We should also do something related to a flag system where we
// want to specify the parts file, and where we want to download
// the images to.
