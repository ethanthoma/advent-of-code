package main

import (
    "os"
    "fmt"
    "bufio"
    "errors"
)

func main() {
    sum, err := processFile("./data/one.txt")
    if err != nil {
        fmt.Println("Error processing file:", err)
        return
    }

    fmt.Println(sum)
}

func processFile(filePath string) (int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    sum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        str := scanner.Text()

        firstDigit,  err := getFirstDigit(str)
        if err != nil {
            return 0, err
        }

        lastDigit, err := getLastDigit(str)
        if err != nil {
            return 0, err
        }

        sum += firstDigit*10 + lastDigit
    }
    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return sum, nil
}

func getFirstDigit(str string) (int, error) {
    for i := 0; i < len(str); i++ {
        if isDigit(str[i]) {
            return int(str[i]), nil
        }
    }

    return 0, errors.New("getFirstDigit: no number in str")
}

func getLastDigit(str string) (int, error) {
    for i := len(str) - 1; i > -1; i-- {
        if isDigit(str[i]) {
            return int(str[i]), nil
        }
    }

    return 0, errors.New("getLastDigit: no number in str")
}

func isDigit(b byte) bool {
    return b >= '0' && b <= '9'
}

