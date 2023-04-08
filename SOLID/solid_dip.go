// An example of Dependency Inversion Principle
// High Level Modules should not depend on Low Level Modules
// They must depend on abstractions
package main

import "fmt"

type Person struct {
	name string
}

type Relationship int

const (
	Parent Relationship = iota
	Child
)

type Information struct {
	from *Person
	relationship Relationship
	to *Person
}

type Relationships struct {
	relations []Information
}

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations,
		Information{parent, Parent, child})
	rs.relations = append(rs.relations,
		Information{child, Child, parent})
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type BadResearch struct {
	relationships *Relationships// LLM
}

type GoodResearch struct {
	browser RelationshipBrowser // LLM
}

func (r *BadResearch) Investigate() {
	// This high level module depends on internal implementations of low level module to do its investigation!
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "Daniel" &&
			rel.relationship == Parent {
			fmt.Println("Daniel has a child called", rel.to.name)
		}
	}
}

func (r *GoodResearch) Investigate() {
	// This high level module relies on the interface/abstraction on how low level modules behaves.
	for _, p := range r.browser.FindAllChildrenOf("Daniel") {
		fmt.Println("JoDanielhn has a child called", p.name)
	}
}

func main() {
	parent := Person{"Daniel" }
	child1 := Person{ "Bruno" }
	child2 := Person{ "Arthur" }

	// Low Level Module
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := BadResearch{&relationships}
	research.Investigate()

	goodResearch := GoodResearch{&relationships}
	goodResearch.Investigate()
}