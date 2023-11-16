package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	// Iterate through the URLs sequentially
	for i, url := range urls {
		// Generate a unique filename for each image
		filename := fmt.Sprintf("image_%d.jpg", i)
		if err := downloadImage(url, filename); err != nil {

			// If an error occurs during download, print an error message
			fmt.Printf("Error downloading %s: %s\n", url, err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	// Create a wait group to manage concurrent goroutines
	var wg sync.WaitGroup

	// Channel to collect and handle errors from goroutines
	errors := make(chan error)

	// Loop through the URLs and start a goroutine for each download
	for i, url := range urls {
		// Increment the wait group counter for each goroutine
		wg.Add(1) // Increment the wait group counter for each goroutine

		go func(url string, i int) {
			// Signal the wait group when the goroutine finishes
			defer wg.Done()

			// Generate a unique filename for each image
			filename := fmt.Sprintf("image_%d.jpg", i)
			if err := downloadImage(url, filename); err != nil {
				// Send any download errors to the errors channel
				errors <- fmt.Errorf("Error downloading %s: %s", url, err)
			}
			// Pass the URL and index to the goroutine function
		}(url, i)
	}

	// Goroutine to close the errors channel once all downloads are completed
	go func() {
		// Wait for all goroutines to finish before closing the channel
		wg.Wait()

		// Close the errors channel to signal that no more errors will be sent
		close(errors)
	}()

	// Loop over the errors channel to print any encountered errors
	for err := range errors {
		// Print each error received from the errors channel
		fmt.Println(err)
	}
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// Fetch the image from the provided URL
	response, err := http.Get(url)
	if err != nil {
		// Return error if unable to fetch the image
		return err
	}
	// Close the response body after function execution
	defer response.Body.Close()

	// Check if the response status indicates successful retrieval (HTTP 200 OK)
	if response.StatusCode != http.StatusOK {
		// Return an error if unable to download
		return fmt.Errorf("unable to download: %s", response.Status)
	}

	// Create a file to save the downloaded image
	file, err := os.Create(filename)
	if err != nil {
		// Return error if unable to create the file
		return err
	}
	// Close the file after function execution
	defer file.Close()

	// Copy the content of the response body (image) to the created file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		// Return error if unable to copy the image content to the file
		return err
	}

	// Return nil if the image is successfully downloaded and saved
	return nil
}

func main() {
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",

		// My royalty free photos
		"https://images.pexels.com/photos/15545361/pexels-photo-15545361/free-photo-of-photo-of-a-pomegranate-fruit.jpeg",
		"https://images.pexels.com/photos/14264353/pexels-photo-14264353.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://images.pexels.com/photos/17528277/pexels-photo-17528277/free-photo-of-translucent-jellyfish-in-deep-ocean.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://images.pexels.com/photos/18895272/pexels-photo-18895272/free-photo-of-after-glow.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
		"https://images.pexels.com/photos/19081146/pexels-photo-19081146/free-photo-of-a-blurry-image-of-a-car-driving-down-a-street.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}
