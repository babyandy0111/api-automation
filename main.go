package main

import "go-micro-project/utils"

func main() {
	passwords := utils.PasswordEncrypt("", "123456")
	print(passwords)
}
