package services

import (
	"log"
)

func ProcessImages(urls []string) []string {
	
	if len(urls) == 0 {
		log.Println("No URLs provided for processing.")
		return nil
	}

	var compressedImages []string
	for _, url := range urls {
		if url == "" {
			log.Println("Skipping empty URL.")
			continue
		}

		log.Printf("Compressing image: %s\n", url)
		compressedURL := url + "_compressed"
		compressedImages = append(compressedImages, compressedURL)
	}

	log.Printf("Compression completed. Total images processed: %d\n", len(compressedImages))
	return compressedImages
}

