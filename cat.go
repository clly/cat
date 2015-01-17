package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    args := os.Args
    for i := 1; i < len(args); i++ {
        file, err := os.Open(args[i])
        if err != nil {
            fmt.Printf("Unknown error trying to read file %v \n", args[i])
        } else {
            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                fmt.Println(scanner.Text())
            }
            if err := scanner.Err(); err != nil {
                fmt.Fprintln(os.Stderr, "Error reading file: ", err)
            }
        }
    }
}