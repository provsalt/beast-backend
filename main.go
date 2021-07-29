package main

import (
	"fmt"
	"github.com/provsalt/beast-backend/notifier/config"
)

func main() {
	fmt.Println(config.LoadConfig().Email.Enabled)
}
