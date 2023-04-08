// An example of Open Closed Principle
package main

import "fmt"



type Size int
const (
	large Size = iota
)

type Color int
const (
	green Color = iota
)

type Product struct {
	color Color
	size Size
	name string
}

// Filter type would be modified everytime a new requirement appears.
// Filter by color.. Then filter by size, then filter both together and so on..
type Filter struct {}

func (f *Filter) filterByColor(
	products []Product, color Color)[]*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color)[]*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Although, we can protect our already-implemented filters by creating them as Specification
// This way, we keep as Open for Expansion, but Closed for Modification.

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type OCPFilter struct {}

func (f *OCPFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	tree := Product{name:"Tree", color:green, size:large}

	products := []Product{tree}

	fmt.Print("Green products (not OCP):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Green products (OCP compliant):\n")
	greenSpec := ColorSpecification{green}
	f2 := OCPFilter{}
	for _, v := range f2.Filter(products, greenSpec) {
		fmt.Printf("%s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Print("Large green items (OCP compliant):\n")
	for _, v := range f2.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}






