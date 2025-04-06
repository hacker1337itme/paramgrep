package main

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
)

func main() {
    // Check if the filename argument is provided
    if len(os.Args) < 2 {
        fmt.Println("Usage: paramgrep <filename>")
        return
    }

    // Get the filename from the command line arguments
    filename := os.Args[1]

    // Open the text file containing URLs
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a map to store unique parameter names
    paramNames := make(map[string]struct{})

    // Read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parsedURL, err := url.Parse(line)
        if err != nil {
            fmt.Println("Error parsing URL:", err)
            continue
        }

        // Extract query parameters
        queryParams := parsedURL.Query()
        for param := range queryParams {
            paramNames[param] = struct{}{} // Using an empty struct for uniqueness
        }
    }

    // Check for scanning errors
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Print unique parameter names
    fmt.Println("[!] UNIQUE parameter names:")
    for param := range paramNames {
        fmt.Println(param)
    }
}
