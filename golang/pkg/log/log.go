package main

import (
    "bufio"
    "fmt"
    "io"
    "math"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'log' function below.
 *
 * The function is expected to return a DOUBLE.
 * The function accepts following parameters:
 *  1. DOUBLE z
 *  2. INTEGER numberOfIterations
 */

 func log(z float64, numberOfIterations int32) float64 {
    // Write your code here
    if z <= 0 {
        return -1.0
    }
    sum := 0.0
    var i int32
    for i = 0; i < numberOfIterations; i++ {
        sum += iteration(z, i)
    }
    return 2 * sum
}

func iteration(z float64, k int32) float64 {
    val := (z - 1)/(z + 1)
    exp := 2 * float64(k) + 1
    temp := math.Pow(val, exp)
    return (1/(2 * float64(k) + 1)) * temp
}




func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    z, err := strconv.ParseFloat(strings.TrimSpace(readLine(reader)), 64)
    checkError(err)

    numberOfIterationsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    numberOfIterations := int32(numberOfIterationsTemp)

    result := log(z, numberOfIterations)

    fmt.Fprintf(writer, "%s\n", formatResult(result))

    writer.Flush()
}

func formatResult(result float64) string {
	strResult := strconv.FormatFloat(result, 'f', 16, 64)
	if (strings.Contains(strResult, ".")) {
		strResult = strings.TrimRight(string(strResult), "0")
	}
	if (strings.HasSuffix(strResult, ".")) {
		strResult = strResult + "0"
	}
	return strResult
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

