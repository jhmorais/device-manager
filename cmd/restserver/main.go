package main

import (
	"fmt"

	"github.com/jhmorais/device-manager/internal/infra/di"
)

func main() {
	dependencies := di.NewBuild()
	fmt.Printf("%v", dependencies.DB.Statement.Vars...)
}
