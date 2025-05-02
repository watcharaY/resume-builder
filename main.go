package main

import (
	"context"
	"log"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8"
	"github.com/starwalkn/gotenberg-go-client/v8/document"
	"github.com/watcharaY/resume-builder/config"
)

func main() {
	htmlPath, pdfPath, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	client, err := gotenberg.NewClient("http://localhost:3000", http.DefaultClient)
	if err != nil {
		panic(err)
	}

	index, err := document.FromPath("resume.html", htmlPath)
	if err != nil {
		log.Fatalln("get file error:", err)
	}

	req := gotenberg.NewHTMLRequest(index)

	req.PaperSize(gotenberg.A4)

	err = client.Store(context.Background(), req, pdfPath)
	if err != nil {
		log.Fatalln("store error:", err)
	}
	resp, err := client.Send(context.Background(), req)
	if err != nil {
		log.Fatalln("send error:", err)
	}
	log.Println(resp)
	log.Println("PDF generated successfully: resume.pdf")
}
