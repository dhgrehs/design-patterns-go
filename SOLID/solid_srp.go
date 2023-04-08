package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type Journal struct {
	entryCounter int
	entries []string
}

func (j *Journal) ToString() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	j.entryCounter++
	entry := fmt.Sprintf("%d: %s", j.entryCounter, text)
	j.entries = append(j.entries, entry)
	return j.entryCounter
}

// For any persistence logic, we should not add them into Journal type, as below.
// It would break the Single Responsibility Principe.

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.ToString()), 0644)
}

func (j *Journal) Load(filename string) {
	//nop
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	//nop
}

// Instead, we create a Persistence type to handle that.

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}


func main() {
	j := Journal{}
	j.AddEntry("first record")
	j.AddEntry("second record")
	fmt.Println(strings.Join(j.entries, "\n"))

	p := Persistence{"\n"}
	p.saveToFile(&j, "journal.txt")
}