package main

import (
	"fmt"

	"github.com/mostafizurRahaman/student-api/internal/config"
)

func main() {

	var cfg = config.MustLoad()

	fmt.Println(cfg)
	fmt.Println("Waiting for configuration........")

}
