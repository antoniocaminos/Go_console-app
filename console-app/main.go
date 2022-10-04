package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Capuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Mocchiato"
	coffees[6] = "Expresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Capuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Mocciato")
	fmt.Println("6 - Expresso")
	fmt.Println("Q - Quit the program")
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}
		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i])) //anda y no porque

	}
	fmt.Println("program exiting")
}

/* reader := bufio.NewReader(os.Stdin)
for {
	fmt.Print("-> ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	if userInput == "quit" {
		break
	} else {
		fmt.Println(userInput)
	}
} */
