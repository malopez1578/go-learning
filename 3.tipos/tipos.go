package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := "22"

	edad_Str, _ := strconv.Atoi(edad)

	fmt.Println(30 + edad_Str)

}
