package main

type SimpleBird struct {
	Age int
}

func (b *SimpleBird) Fly() string {
	if b.Age >= 10 {
		return "Simple Bird Flying!"
	}
	return ""
}

type SimpleLizard struct {
	Age int
}

func (l *SimpleLizard) Crawl() string {
	if l.Age < 10 {
		return "Simple Lizard Crawling!"
	}
	return ""
}

type MultipleAggregationDragon struct {
	SimpleBird
	SimpleLizard
}

func (d *MultipleAggregationDragon) Age() int {
	return d.SimpleBird.Age
}

func (d *MultipleAggregationDragon) SetAge(age int) {
	d.SimpleBird.Age = age
	d.SimpleLizard.Age = age
}
