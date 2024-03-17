package main

import (
	"bufio"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

func main() {
	client := resty.New()

	endpoint := "http://localhost:8080/"
	fmt.Println("Please, enter URL: ")
	reader := bufio.NewReader(os.Stdin)
	long, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	long = strings.TrimSuffix(long, "\n")

	resp, err := client.R().
		SetHeader("Content-Type", "application/text").
		SetBody(long).
		Post(endpoint)

	if err != nil {
		panic(err)
	}
	fmt.Println("Status code: ", resp.StatusCode())
	fmt.Println(string(resp.Body()))
}
