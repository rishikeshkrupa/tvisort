package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rishikeshkrupa/tvisort/algos"
)

func main() {

	algoPtr := flag.String("algo", "", "name of the sorting algorithm")
	flag.Parse()

	algo := *algoPtr

	if algo == "" {
		fmt.Println("Sorting algorithm not specified")
		fmt.Println("We will be running all availble ones...")
	}

	//if no algo is selected we run all by default

	/*

		fmt.Print("\x1b7")
		fmt.Print("Input: ")
		fmt.Println("[1 2 5]")
		fmt.Println("outside!!")
		fmt.Print("\x1b8")
		fmt.Println("Input:", "[1 2 5 -1]")

	*/

	var userInput []int

	scan := bufio.NewScanner(os.Stdin)

	fmt.Println("Add numbers to input [empty to stop reading]: ")

	fmt.Print("\x1b7") //remember cursor position
	fmt.Println("Input: ")
	fmt.Print("\n⏫ : ")

	for scan.Scan() {
		errIt(scan.Err())

		t := scan.Text()
		//fmt.Printf("%q\n", t)

		if t == "" {
			break
		}

		v, err := strconv.Atoi(t)

		if err != nil {
			fmt.Printf(" Cannot Parse %q as an Integer\n", t)
			continue
		}

		userInput = append(userInput, v)

		fmt.Print("\x1b8") //go back to to position
		fmt.Print("Input: ")
		fmt.Println(userInput)

		fmt.Print("\n⏫ : ")
	}

	fmt.Print("\x1b8") //go back to to position
	fmt.Print("Input: ")
	fmt.Println(userInput)
	//fmt.Print("\x1bc")

	fmt.Println("let's start cookin~")

	algos.Bubble(userInput)

}

func errIt(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
