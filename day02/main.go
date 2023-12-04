package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    maxCubes := MaxCubes{
        RedCubes:   12,
        GreenCubes: 13,
        BlueCubes:  14,
    }

    gameId, err := processFile("./data/two.txt", maxCubes)
    if err != nil {
        fmt.Println("Error processing file:", err)
        return
    }

    fmt.Println(gameId)
}

type MaxCubes struct {
    RedCubes    int
    GreenCubes  int
    BlueCubes   int
}

func processFile(filePath string, maxCubes MaxCubes) (int, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    possibleGames := 0 
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        gameId, err := processLine(line, maxCubes)
        if err != nil {
            return 0, err
        }
        
        possibleGames += gameId
    }
    if err := scanner.Err(); err != nil {
        return 0, err
    }

    return possibleGames, nil
}

func processLine(line string, maxCubes MaxCubes) (int, error) {
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

    if impossibleGame {
        return 0, nil
    } else {
        return gameId, nil
    }
}

