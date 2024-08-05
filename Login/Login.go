package Login

import (
	"crypto/sha256"
	"fmt"
	"os/exec"

	"github.com/Carsen/PaxxNest/DB"
)

func Login(running bool) bool {
	var checker bool = false
	var i int = -2
	for running {
		cls()
		fmt.Println("Hello, and welcome to PaxxNest!")
		fmt.Println("Want to take a ride? y/n + Enter")
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" || answer == "Y" {
			i = 5
			cls()
		} else if answer == "n" || answer == "N" {
			cls()
			fmt.Println("I'm sorry to see you go so soon. We hope to see you back!")
			i = -1
			checker = false
			running = false
			break
		} else {
			i = -2
		}
		if i == -2 {
			continue
		}
		for i > 0 {
			fmt.Print("Please enter username: ")
			var inUsern string
			var hashUsern []byte
			fmt.Scanln(&inUsern)
			hashUsern = hashInput(inUsern)
			
			keycheck, _ := DB.CheckForKey(hashUsern)
			switch keycheck {
			case true:
				fmt.Print("Please enter password: ")
				var inPassw string
				var hashPassw []byte
				fmt.Scanln(&inPassw)
				hashPassw = hashInput(inPassw)
				keymatch, _ := DB.ValueMatchesKey(hashUsern, hashPassw)
				switch keymatch {
				case true:
					cls()
					checker = true
					return checker
				case false:
					cls()
					fmt.Println("Try again!")
					i--
				}

			case false:
				var mkAcc string
				var runChk int
				runChk = 5
				cls()
				fmt.Printf("It looks like you're new here!(Username: '%s')\n", inUsern)
				for runChk > 0 {
					fmt.Println("Would you like to create an account? y/n")
					fmt.Scanln(&mkAcc)
					if mkAcc == "y" || mkAcc == "Y" {
						fmt.Print("Please enter your new password: ")
						var inPassw string
						fmt.Scanln(&inPassw)
						fmt.Print("Please confirm password: ")
						var matchPassw string
						fmt.Scanln(&matchPassw)
						if inPassw == matchPassw {
							var hashPassw []byte = hashInput(inPassw)
							DB.NewKeyValue(hashUsern, hashPassw)
							checker = true
							return checker
						} else if inPassw != matchPassw {
							runChk--
							cls()
							fmt.Println("Try again!")
						}
					} else if mkAcc == "n" || mkAcc == "N" {
						cls()
						runChk--
						i--
						break
					} else {
						cls()
						fmt.Println("Bad Format! Try again.")
					}
				}
				if runChk == 0 {
					i = 0
				}
			}
		}
		if i == 0 || i == -1 {
			cls()
			fmt.Println("Too many tries!")
			checker = false
			break
		}
	}
	for !running {
		checker = false
		return checker
	}
	return checker
}

func hashInput(u string) []byte {
	hash := sha256.New()
	defer hash.Reset()
	hash.Write([]byte(u))
	b := hash.Sum(nil)
	return b
}

func cls() {
	exec.Command("clear\n")
}
