package main

import (
	"fmt"
	"os"

	"github.com/peya/src/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
