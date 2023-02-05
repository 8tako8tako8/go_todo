package main

import (
	"fmt"

	"github.com/8tako8tako8/go_todo/app/controllers"
	"github.com/8tako8tako8/go_todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
