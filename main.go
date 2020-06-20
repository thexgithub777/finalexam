package main

import (
	"github.com/thexgithub777/finalexam/customer"
)

func main() {
	r := customer.Router()
	r.Run(":2019")
}
