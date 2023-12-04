package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    val, err := task1("./data/three.txt")
    if err != nil {
        fmt.Println("Error processing file:", err)
        return
    }

    fmt.Println(val)
}

func task1(filePath string) (int, error) {
    schema, err := readFileAs2DByteArray(filePath)
    if err != nil {
        return 0, err
    }

    sumOfParts := 0

    var num []byte
    lastRowIndex, lastColIndex := 0, 0
    for rowIndex, row := range schema {
        for colIndex, unit := range row {
            if len(num) > 0 && (!isDigit(unit) || colIndex == 0) {
                startX  := max(lastColIndex - len(num), 0)
                endX    := min(lastColIndex + 1, len(row) - 1)
                startY  := max(lastRowIndex - 1, 0)
                endY    := min(lastRowIndex + 1, len(schema) - 1)

                hasSymbol := isPart(schema, startX, endX, startY, endY)

                if hasSymbol {
                    val, err := strconv.Atoi(string(num[:]))
                    if err != nil {
                        return 0, err
                    }
                    sumOfParts += val
                }

                num = nil 
            }

            if isDigit(unit) {
                num = append(num, unit)
            } 

            lastRowIndex = rowIndex
            lastColIndex = colIndex
        }
    }

    return sumOfParts, nil
}

func isPart(schema [][]byte, startX int, endX int, startY int, endY int) (bool) {
    hasSymbol := false

    for x := startX; x <= endX; x++ {
        for y := startY; y <= endY; y++ {
            hasSymbol = hasSymbol || isSymbol(schema[y][x])
        }
    }

    return hasSymbol
}

func readFileAs2DByteArray(filePath string) ([][]byte, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines [][]byte
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Bytes()
        lineCopy := make([]byte, len(line))
        copy(lineCopy, line)

        lines = append(lines, lineCopy)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}

func isDigit(unit byte) (bool) {
    return unit >= '0' && unit <= '9'
}

func isSymbol(unit byte) (bool) {
    return (unit >= '!' && unit <= '-') || 
    (unit == '/') || (unit >= ':' && unit <= '@')

}
