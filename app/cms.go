package app

import (
	"fmt"
	"net/http"
)

func (mycms *App) hello (w http.ResponseWriter, r *http.Request) {
	fmt.Print("hello")
}
