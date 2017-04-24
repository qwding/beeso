package main

// // No C code needed.
import "C"
import (
	"beeso/models"
	"fmt"
)

var Data models.Data

func GET() (string, error) {
	fmt.Println("hello world! ", Data.Ctx.Input.Method())
	return "hello world! "+ Data.Ctx.Input.Method(),nil
}

func POST() (string, error) {
	fmt.Println("hello world! ", Data.Ctx.Input.Method())
	return "hello world! "+ Data.Ctx.Input.Method(),nil
}
