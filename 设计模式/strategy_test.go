package main

import "testing"

func TestFunc ( t *testing.T) {
	printer := Printer{}
	printer.setStrategy(&PrintDoc{}).todo("abcdef")

	printer.setStrategy(&PrintPdf{}).todo("123456")

}
