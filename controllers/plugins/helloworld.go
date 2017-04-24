package main

// // No C code needed.
import "C"
import (
	"beeso/models"
	"fmt"
)

var Data models.Data

func Get() (string, error) {
	fmt.Println("hello world! ", Data.Ctx.Input.Method())
}

func Post() (string, error) {
	fmt.Println("hello world! ", Data.Ctx.Input.Method())
}
