package main

//Need to read in values from a CSV and make it a game.

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var userAnswer int

		fmt.Println(rec[0])
		fmt.Println("Enter you answer: ")

		timer2 := time.NewTimer(5 * time.Second)
		go func() {
			<-timer2.C
			fmt.Println("Times up!")
			os.Exit(0)
		}()

		fmt.Scanln(&userAnswer)
		solution, err := strconv.Atoi(rec[1])

		if err != nil {
			log.Fatal(err)
		}
		if userAnswer == solution {
			fmt.Println("success!")
			timer2.Stop()
		} else {
			fmt.Println("incorrect :(")
			os.Exit(0)
		}

		//Need to add a reading to get the users input
		///then compare that input ot rec[1] if it matches have points

	}

}
