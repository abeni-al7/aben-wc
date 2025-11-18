package main

import (
	"github.com/abeni-al7/aben-wc/controllers"
	"github.com/abeni-al7/aben-wc/services"
)

func main() {
	fs := services.FileService{}
	fio := controllers.FileIO{Fs: fs}

	fio.AcceptInput()
}
