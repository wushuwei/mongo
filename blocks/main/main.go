package main

import (
	"fmt"

	"github.com/wushuwei/mongo/blocks/routines"
)

func main() {

	fmt.Print(utilities.Double(3))
	utilities.ConnetMongo()
}
