package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    val, err := task2("./data/three.txt")
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
    prevY, prevX := 0, 0
    for y, row := range schema {
        for x, unit := range row {
            if len(num) > 0 && (!isDigit(unit) || x == 0) {
                if hasSymbolAroundNumber(schema, prevY, prevX, len(num)) {
                    number, err := strconv.Atoi(string(num[:]))
                    if err != nil {
                        return 0, err
                    }
                    sumOfParts += number
                }

                num = nil 
            }

            if isDigit(unit) {
                num = append(num, unit)
            } 

            prevX = x
            prevY = y
        }
    }

    return sumOfParts, nil
}

func hasSymbolAroundNumber(schema [][]byte, y int, x int, lengthOfNumber int) (bool) {
    startX  := max(x - lengthOfNumber, 0)
    endX    := min(x + 1, len(schema[0]) - 1)
    startY  := max(y - 1, 0)
    endY    := min(y + 1, len(schema) - 1)

    hasSymbol := false

    for y := startY; y <= endY; y++ {
        for x := startX; x <= endX; x++ {
            hasSymbol = hasSymbol || isSymbol(schema[y][x])
        }
    }

    return hasSymbol
}

func isSymbol(unit byte) (bool) {
    return (unit >= '!' && unit <= '-') || 
    (unit == '/') || (unit >= ':' && unit <= '@')

}

func task2(filePath string) (int, error) {
    schema, err := readFileAs2DByteArray(filePath)
    if err != nil {
        return 0, err
    }

    gearRatioSum := 0
    for y, row := range schema {
        for x, unit := range row {
            if unit == '*' {
                gearRatio, err := findGearRatio(schema, y, x)
                if err != nil {
                    return 0, err
                }

                gearRatioSum += gearRatio
            }
        }
    }

    return gearRatioSum, nil
}

func findGearRatio(schema [][]byte, yOfSymbol int, xOfSymbol int) (int, error) {
    var nums []int

    xLeftOfSymbol := max(xOfSymbol - 1, 0)
    xRightOfSymbol := min(xOfSymbol + 1, len(schema[0]) - 1)

    yAboveSymbol := max(yOfSymbol - 1, 0)
    yBelowSymbol := min(yOfSymbol + 1, len(schema) - 1)

    for y := yAboveSymbol; y <= yBelowSymbol; y++ {
        if isDigit(schema[y][xOfSymbol]) {
            number, err := getNumberAround(schema, y, xOfSymbol)
            if err != nil {
                return 0, err
            }
            nums = append(nums, number)
        } else {
            if isDigit(schema[y][xLeftOfSymbol]) {
                number, err := getNumberLeft(schema, y, xLeftOfSymbol)
                if err != nil {
                    return 0, err
                }
                nums = append(nums, number)
            }

            if isDigit(schema[y][xRightOfSymbol]) {
                number, err := getNumberRight(schema, y, xRightOfSymbol)
                if err != nil {
                    return 0, err
                }
                nums = append(nums, number)
            }
        }
    }

    if len(nums) != 2 {
        return 0, nil
    }

    return nums[0] * nums[1], nil
}

func getNumberRight(schema [][]byte, y int, x int) (int, error) {
    var numberAsBytes []byte
    for i := x; i < len(schema[0]) && isDigit(schema[y][i]); i++ {
        numberAsBytes = append(numberAsBytes, schema[y][i])
    }

    number, err := strconv.Atoi(string(numberAsBytes[:]))
    if err != nil {
        return 0, err
    }

    return number, nil
}

func getNumberLeft(schema [][]byte, y int, x int) (int, error) {
    var numberAsBytes []byte

    for i := x; i >= 0 && isDigit(schema[y][i]); i-- {
        numberAsBytes = append([]byte{schema[y][i]}, numberAsBytes...)
    }

    number, err := strconv.Atoi(string(numberAsBytes[:]))
    if err != nil {
        return 0, err
    }

    return number, nil
}

func getNumberAround(schema [][]byte, y int, x int) (int, error) {
    var numberAsBytes []byte
    for i := x - 1; i >= 0 && isDigit(schema[y][i]); i-- {
        numberAsBytes = append([]byte{schema[y][i]}, numberAsBytes...)
    }

    numberAsBytes = append(numberAsBytes, schema[y][x])

    for i := x + 1; i < len(schema[0]) && isDigit(schema[y][i]); i++ {
        numberAsBytes = append(numberAsBytes, schema[y][i])
    }

    number, err := strconv.Atoi(string(numberAsBytes[:]))
    if err != nil {
        return 0, err
    }

    return number, nil
}

func isDigit(unit byte) (bool) {
    return unit >= '0' && unit <= '9'
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

