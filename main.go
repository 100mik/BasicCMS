package main

import (
	"github.com/100mik/BasicCMS/app"
	"os"
)

func main() {
mycms := app.App{}
mycms.Initialise()
mycms.Run(os.Args)
}