package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
    for i, url := range urls {
        filename := fmt.Sprintf("image_%d.jpg", i)
        if err := downloadImage(url, filename); err != nil {
            fmt.Printf("Error downloading %s: %s\n", url, err)
        }
    }
}


// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
    var wg sync.WaitGroup
    errors := make(chan error)

    for i, url := range urls {
        wg.Add(1)
        go func(url string, i int) {
            defer wg.Done()
            filename := fmt.Sprintf("image_%d.jpg", i)
            if err := downloadImage(url, filename); err != nil {
                errors <- fmt.Errorf("Error downloading %s: %s", url, err)
            }
        }(url, i)
    }

    go func() {
        wg.Wait()
        close(errors)
    }()

    for err := range errors {
        fmt.Println(err)
    }
}


func main() {
    urls := []string{
        "https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
        "https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640"
        // Add more image URLs
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

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
    response, err := http.Get(url)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return fmt.Errorf("unable to download: %s", response.Status)
    }

    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.Copy(file, response.Body)
    if err != nil {
        return err
    }

    return nil
}