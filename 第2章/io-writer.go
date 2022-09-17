package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	//createFile()
	//stdOut()
	//buffer()
	//buffer2()
	//connectNetwork()
	//decorator()
	//zip()
	//flush()
	format()
}

func format() {
	fmt.Fprintf(os.Stdout, "write sample %v", time.Now())
}

func flush() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("aaa\n")
	buffer.WriteString("+a\n")
	buffer.Flush()
	//buffer.WriteString("bbb\n")
	//buffer.Flush()
}

func zip() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		fmt.Println(err)
	}

	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	io.WriteString(writer, "example5")
	writer.Close()
}

func decorator() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}

	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "example4\n")
	io.WriteString(writer, "example5")
}

func createFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}

func stdOut() {
	os.Stdout.Write([]byte("sample"))
}

func buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("sample2"))
	fmt.Println(buffer.String())
}

func buffer2() {
	var builder strings.Builder
	builder.Write([]byte("sample3"))
	fmt.Println(builder.String())
}

func connectNetwork() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		panic(err)
	}

	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: example.com\r\n\r\n")
	io.Copy(os.Stdout, conn)
}
