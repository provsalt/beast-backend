package main

import (
	"beast-backend/notifier/config"
	"fmt"
)

func main() {
	fmt.Println(config.LoadConfig().Email.Enabled)
}
