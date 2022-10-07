package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func nomor(a []int32) int32 {
    b := make(map[int32]int32)
    
    for i := 0; i<len(a); i++ {
    	b[a[i]]++
    }
    var nomor int32
    for t, k := range b {
		if (b[t+1] + k > nomor)  {
			nomor = b[t+1] + k
		}
    }
    return nomor
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}
	result := nomor(a) //Print Hasil

	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
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