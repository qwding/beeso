package plugins

// // No C code needed.
import "C"
import (
	"beeso/models"
	"fmt"
)

var Data models.Data

func Get() {
	fmt.Println("hello world! ", Data.Ctx.Input.Method())
}

func Post() {
	fmt.Println("hello world! ",, Data.Ctx.Input.Method())
}
