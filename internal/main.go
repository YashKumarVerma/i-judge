package main

import (
	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/YashKumarVerma/i-judge/internal/config"
)

func main() {
	config.Initialize()
	ui.ContextPrint("balance_scale", config.Load.Name)
}
