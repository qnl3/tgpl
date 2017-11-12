// Fetch prints the content found at a URL.(exercise 1.7)
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {
        if strings.HasPrefix(url, "http://") == false {
                url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        }
        _, err = io.Copy(os.Stdout, resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }
        //fmt.Printf("Retrived %d bytes.\n",c) // Don't contaminate stdout stream.
    }
}
