package contract

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"schema"
	"testing"
)

func Test_Cert(t *testing.T) {
	cert1 := schema.Cert{
		Owner: "p123",
		Content: "Cert-CS-p123-i123-it123",
		Issuer: "i123",
		Item: "it123",
		Point: 3.7,
		Date: "2021-06-17 19:22:30",
		Signature: "ecnu",
	}

	certByte1, _ := json.Marshal(cert1)
	h1 := sha256.New()
	h1.Write(certByte1)
	certHash1 := hex.EncodeToString(h1.Sum(nil))

	args1 := cert1.ToStringArray()
	args1 = append(args1, certHash1)

	if _, err := Send(args1, "addCert", Peer0); err != nil {
		t.Error("Test `addCert` not pass!")
	} else {
		t.Log("Test `addCert` pass!")
	}

	if _, err := Query([]string{"certHash1"}, "queryCert", Peer0); err != nil {
		t.Error("Test `queryCert` not pass!")
	} else {
		t.Log("Test `queryCert` pass!")
	}

	cert2 := schema.Cert{
		Owner: "p123",
		Content: "Cert-BC-p123-i234-it345",
		Issuer: "i234",
		Item: "it345",
		Point: 3.6,
		Date: "2021-06-17 19:22:40",
		Signature: "stju",
	}

	certByte2, _ := json.Marshal(cert2)
	h2 := sha256.New()
	h2.Write(certByte2)
	certHash2 := hex.EncodeToString(h2.Sum(nil))

	args2 := cert2.ToStringArray()
	args2 = append(args2, certHash2)

	if _, err := Send(args2, "addCert", Peer0); err != nil {
		t.Error("Test `addCert` not Pass!")
	}

	hashs := []string{certHash1, certHash2}
	if _, err := Query(hashs, "queryCerts", Peer0); err != nil {
		t.Error("Test `queryCerts` not pass!")
	} else {
		t.Log("Test `queryCerts` pass!")
	}

	cert3 := schema.Cert{
		Owner: "p234",
		Content: "Cert-DB-p234-i123-it234",
		Issuer: "i234",
		Item: "it234",
		Point: 3.5,
		Date: "2021-06-17 19:22:50",
		Signature: "ecnu",
	}

	certByte3, _ := json.Marshal(cert3)
	h3 := sha256.New()
	h3.Write(certByte3)
	certHash3 := hex.EncodeToString(h3.Sum(nil))

	args3 := cert3.ToStringArray()
	args3 = append(args3, certHash3)

	if _, err := Send(args3, "addCert", Peer0); err != nil {
		t.Error("Test `addCert` not Pass!")
	}

	cert4 := schema.Cert{
		Owner: "p234",
		Content: "Cert-DS-p234-i234-it456",
		Issuer: "i234",
		Item: "it456",
		Point: 3.7,
		Date: "2021-06-17 19:23:00",
		Signature: "stju",
	}

	certByte4, _ := json.Marshal(cert4)
	h4 := sha256.New()
	h4.Write(certByte4)
	certHash4 := hex.EncodeToString(h4.Sum(nil))

	args4 := cert4.ToStringArray()
	args4 = append(args4, certHash4)

	if _, err := Send(args4, "addCert", Peer0); err != nil {
		t.Error("Test `addCert` not Pass!")
	}
}

func Test_Item(t *testing.T) {
	item1 := schema.Item{
		Id: "it123",
		Name: "CS",
		Description: "computer science",
		Issuer: "i123",
		Point: 3.7,
	}
	args1 := item1.ToStringArray()
	if _, err := Send(args1, "addItem", Peer0); err != nil {
		t.Error("Test `addItem` not pass!")
	} else {
		t.Log("Test `assItem` pass!")
	}

	if _, err := Query([]string{item1.Id}, "queryItem", Peer0); err != nil {
		t.Error("Test `queryItem` not pass!")
	} else {
		t.Log("Test `queryItem` pass!")
	}

	item2 := schema.Item{
		Id: "it234",
		Name: "DB",
		Description: "Database",
		Issuer: "i123",
		Point: 3.5,
	}
	args2 := item2.ToStringArray()

	if _, err := Send(args2, "addItem", Peer0); err != nil {
		t.Error("Test `addItem` not pass!")
	}

	hashs := []string{item1.Id, item2.Id}
	if _, err := Query(hashs, "queryItems", Peer0); err != nil {
		t.Error("Test `queryItems` not pass!")
	} else {
		t.Log("Test `queryItems` pass!")
	}

	item3 := schema.Item{
		Id: "it345",
		Name: "BC",
		Description: "Blockchain",
		Issuer: "i234",
		Point: 3.6,
	}
	args3 := item3.ToStringArray()

	if _, err := Send(args3, "addItem", Peer0); err != nil {
		t.Error("Test `addItem` not pass!")
	}

	item4 := schema.Item{
		Id: "it456",
		Name: "DS",
		Description: "Data Struct",
		Issuer: "i234",
		Point: 3.7,
	}
	args4 := item4.ToStringArray()

	if _, err := Send(args4, "addItem", Peer0); err != nil {
		t.Error("Test `addItem` not pass!")
	}
}

func Test_Person(t *testing.T) {
	person := schema.Person{
		Id: "p123",
		Name: "Li",
		Password: "123",
		PublicKey: "abc",
	}
	args := person.ToStringArray()
	if _, err := Send(args, "addPerson", Peer0); err != nil {
		t.Error("Test `addPerson` not pass!")
	} else {
		t.Log("Test `addPerson` pass!")
	}

	if _, err := Query([]string{person.Id}, "queryPerson", Peer0); err != nil {
		t.Error("Test `queryPerson` not pass!")
	} else {
		t.Log("Test `queryPerson` pass!")
	}

	person1 := schema.Person{
		Id: "p234",
		Name: "Zhang",
		Password: "234",
		PublicKey: "bcd",
	}
	args1 := person1.ToStringArray()
	if _, err := Send(args1, "addPerson", Peer0); err != nil {
		t.Error("Test `addPerson` not pass!")
	}
}

func Test_Issuer(t *testing.T) {
	issuer := schema.Issuer{
		Id: "i123",
		Name: "ECNU",
		Password: "123",
		PublicKey: "cde",
	}
	args := issuer.ToStringArray()
	if _, err := Send(args, "addIssuer", Peer0); err != nil {
		t.Error("Test `addIssuer` not pass!")
	} else {
		t.Log("Test `addIssuer` pass!")
	}

	if _, err := Query([]string{issuer.Id}, "queryIssuer", Peer0); err != nil {
		t.Error("Test `queryIssuer` not pass!")
	} else {
		t.Log("Test `queryIssuer` pass!")
	}

	issuer1 := schema.Issuer{
		Id: "i234",
		Name: "STJU",
		Password: "234",
		PublicKey: "def",
	}
	args1 := issuer1.ToStringArray()
	if _, err := Send(args1, "addIssuer", Peer0); err != nil {
		t.Error("Test `addIssuer` not pass!")
	}
}

func Test_GetPublick(t *testing.T)  {
	if _, err := Query([]string{"p123"}, "getPublicKey", Peer0); err != nil {
		t.Error("Test `getPublicKey` not pass!")
	} else {
		t.Log("Test `getPublicKey` pass!")
	}

	if _, err := Query([]string{"i123"}, "getPublicKey", Peer0); err != nil {
		t.Error("Test `getPublicKey` not pass!")
	}
}
