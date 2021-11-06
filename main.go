package main

import (
	"ManGo/utils"
	"ManGo/websites"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"net/http"
)

func main() {
	var opt gofpdf.ImageOptions
	fmt.Println(websites.Manganato)
	pdf := gofpdf.New("P", "mm", "letter", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 11)
	pdf.SetX(60)
	opt.ImageType = "png"
	pdf.ImageOptions("images/1.jpg", 0, 0, 216, 279, false, opt, 0, "")
	opt.AllowNegativePosition = true
	pdf.AddPage()
	opt.ImageType = "jpg"
	pdf.ImageOptions("images/5.jpg", 0, 0, 216, 279, false, opt, 0, "")
	fileStr := "Mangas/img.pdf"
	err := pdf.OutputFileAndClose(fileStr)
	utils.Check(err)

}
