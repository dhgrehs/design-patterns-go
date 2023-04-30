package main

type BadCopiedPerson struct {
	Person
}

func (p *BadCopiedPerson) DeepCopy() *BadCopiedPerson {
	// Buggy copy because its person has a ptr attribute
	newPerson := p
	return newPerson
}
