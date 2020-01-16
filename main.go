package main

import (
	"fmt"
	"github.com/linonymous/goerr/log"
)

func CheckQuestion(i int) error {
	switch i{
	case 1:
		return fmt.Errorf("")
	default:
		return nil
	}
}

func main()  {
	Check()
	fmt.Println("My work is done LOL!")
}

func Check()  {
	defer log.Recover()
	err := CheckQuestion(1)
	log.Fatal(err)
}
