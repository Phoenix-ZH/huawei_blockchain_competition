package schema

import (
	"strconv"
)

type ServerError struct {
	Error error
	Message string
	Code int
}

type Cert struct {
	Owner string
	Content string
	Issuer string
	Item string
	Point float64
	Date string
	Signature string
	Hash string
}

type Item struct {
	Id string
	Name string
	Description string
	Issuer string
	Point float64
}

type Person struct {
	Id string
	Name string
	Password string
	//PrivateKey string
	PublicKey string
}

type Issuer struct {
	Id string
	Name string
	Password string
	//PrivateKey string
	PublicKey string
}

type Res struct {
	Info string
	Status string
	Node string
}

func (c Cert) ToStringArray() []string {
	return []string{
		c.Owner,
		c.Content,
		c.Issuer,
		c.Item,
		strconv.FormatFloat(c.Point, 'E', -1, 32),
		c.Date,
		c.Signature,
	}
}

func (i Item) ToStringArray() []string {
	return []string{
		i.Id,
		i.Name,
		i.Description,
		i.Issuer,
		strconv.FormatFloat(i.Point, 'E', -1, 32),
	}
}

func (p Person) ToStringArray() []string {
	return []string{
		p.Id,
		p.Name,
		p.Password,
		p.PublicKey,
	}
}

func (i Issuer) ToStringArray() []string {
	return []string{
		i.Id,
		i.Name,
		i.Password,
		i.PublicKey,
	}
}
