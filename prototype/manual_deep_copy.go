package main

type ManualCopiedPerson struct {
	Person
}

func (a *Address) ManualDeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country}
}

func (p *ManualCopiedPerson) ManualDeepCopy() *ManualCopiedPerson {
	q := *p // copies Name
	q.Address = p.Address.ManualDeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}
