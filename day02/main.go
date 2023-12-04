package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
//    val, err := task1("./data/two.txt", MaxCubes{
//        RedCubes:   12,
//        GreenCubes: 13,
//        BlueCubes:  14,
//    })
    val, err := task2("./data/two.txt")
    if err != nil {
        fmt.Println("Error processing file:", err)
        return
    }

    fmt.Println(val)
}

type MaxCubes struct {
    RedCubes    int
    GreenCubes  int
    BlueCubes   int
}

func task1(filePath string, maxCubes MaxCubes) (int, error) {
    sumOfGameIds := 0

    lines, errChan := fileLines(filePath)
    for line := range lines {
        words := strings.Fields(line)

        gameId, err := strconv.Atoi(words[1][:len(words[1])-1])
        if err != nil {
            return 0, err
        }

        cubeCounts := map[string]int {
            "red":   0,
            "green": 0,
            "blue":  0,
        }

        for index, value := range words[2:] {
            for color := range cubeCounts {
                if strings.Contains(value, color) {
                    val, err := strconv.Atoi(words[index + 1])
                    if err != nil {
                        return 0, err
                    }
                    cubeCounts[color] = max(cubeCounts[color], val)
                }
            }
        }

        impossibleGame := cubeCounts["red"] > maxCubes.RedCubes || 
        cubeCounts["green"] > maxCubes.GreenCubes || 
        cubeCounts["blue"] > maxCubes.BlueCubes

        if !impossibleGame {
            sumOfGameIds += gameId
        }
    }
    if err := <-errChan; err != nil {
        return 0, err
    }

    return sumOfGameIds, nil
}

func task2(filePath string) (int, error) {
    sumOfPowerSet := 0

    lines, errChan := fileLines(filePath)
    for line := range lines {
        words := strings.Fields(line)

        cubeCounts := map[string]int {
            "red":   0,
            "green": 0,
            "blue":  0,
        }

        for index, value := range words[2:] {
            for color := range cubeCounts {
                if strings.Contains(value, color) {
                    val, err := strconv.Atoi(words[index + 1])
                    if err != nil {
                        return 0, err
                    }
                    cubeCounts[color] = max(cubeCounts[color], val)
                }
            }
        }

        sumOfPowerSet += cubeCounts["red"]*cubeCounts["green"]*cubeCounts["blue"]
    }
    if err := <-errChan; err != nil {
        return 0, err
    }

    return sumOfPowerSet, nil
}

func fileLines(filePath string) (<-chan string, <-chan error) {
    lines := make(chan string)
    errChan := make(chan error, 1)

    go func() {
        defer close(lines)
        defer close(errChan)

        file, err := os.Open(filePath)
        if err != nil {
            errChan <- err
            return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            lines <- scanner.Text()
        }

        if err := scanner.Err(); err != nil {
            errChan <- err
            return
        }
    }()

    return lines, errChan
}

