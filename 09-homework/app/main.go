package main

import (
	"fmt"
	"go_search/09-homework/pkg/users2"
	"io"
	"os"
)

func main() {
	customers := []users2.Customer{
		{22}, {99}, {33}, {44}, {55},
	}

	WriteAll(os.Stdout, users2.Elder(customers...))
}

func WriteAll(w io.Writer, args ...any) {
	w.Write([]byte(fmt.Sprint(args...)))
}
