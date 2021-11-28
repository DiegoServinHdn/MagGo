package utils

import (
	"github.com/jung-kurt/gofpdf"
)

func PDF(){
	var opt gofpdf.ImageOptions
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
	Check(err)
}

