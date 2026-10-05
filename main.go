package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type User struct {
	Name  string
	Email string
}

func main() {
	file, err := os.Open("data/people-100.csv")
	if err != nil {
		fmt.Println("file open err")
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, _ := reader.ReadAll()

	for _, row := range record {
		user := User{
			Name:  row[1] + " " + row[2],
			Email: row[4],
		}
		fmt.Println(user)
	}

}
