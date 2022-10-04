package main

import (
	"io/ioutil"
	"log"
)

func main() {
	lama := "file_lama.txt"
	baru := "file_baru.txt"
	copy(lama, baru)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func copy(lama string, baru string) {
	data, err := ioutil.ReadFile(lama)
	checkErr(err)
	err = ioutil.WriteFile(baru, data, 0644)
	checkErr(err)
}
