package main

import "fmt"

type Printer interface {
	Print() string
}

type Scanner interface {
	Scan() string
}

type Copier interface {
	Copy() string
}

type PrinterScannerCopier interface {
	Printer
	Scanner
	Copier
}

type PrinterScanner interface {
	Printer
	Scanner
}

type FastPrinter struct{}

func (fp FastPrinter) Print() string {
	return "fast printer is printing..."
}

func (fp FastPrinter) Scan() string {
	return "fast printer is scanning..."
}

func (fp FastPrinter) Copy() string {
	return "fast printer is copying..."
}

type SlowPrinter struct{}

func (sp SlowPrinter) Print() string {
	return "slow printer is printing..."
}

func (sp SlowPrinter) Scan() string {
	return "slow printer is scanning.."
}

func (sp SlowPrinter) Copy() string {
	return "slow printer is copying..."
}

func process(printer PrinterScanner) {
	fmt.Println("Print: ", printer.Print())
	fmt.Println("Scan: ", printer.Scan())
	//fmt.Println("Copy: ", printer.Copy())
}

func main() {
	fp := FastPrinter{}
	sp := SlowPrinter{}
	process(fp)
	process(sp)
}
