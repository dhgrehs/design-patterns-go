// An example of Interface Segregation Principle
package main

// Document
type Document struct {}

// Machine interface, too generic. It may have too much data unnecessarily.
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}


// May cause problem because an OldFashionedPrinter may not offer valid implementations for the entire intf.

type OldFashionedPrinter struct {}

func (o OldFashionedPrinter) Print(d Document) {
	// ok
}

func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// Printer An approach that splits into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// MyPrinter printer only
type MyPrinter struct {
	// ...
}

func (m MyPrinter) Print(d Document) {
	// ...
}

// If required, you can always combine interfaces

type Photocopier struct {}

func (p Photocopier) Scan(d Document) {
	//
}

func (p Photocopier) Print(d Document) {
	//
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

// MultiFunctionMachine interface composition + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	//Not much here
}