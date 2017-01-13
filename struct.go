package main

import "fmt"

type realDev interface {
	setNamePseudo(name, pseudo string) Dev
}

type Dev struct {
	name   string
	pseudo string
}

type FakeDev struct {
	name   string
	pseudo string
}

func (d Dev) setNamePseudo(name, pseudo string) Dev {
	d.name = name
	d.pseudo = pseudo
	return d
}

func setAttributes(rD realDev) {
	rD = rD.setNamePseudo("Real", "Dev");
	fmt.Print(rD);
}

func main() {
	// d := Dev{name: "nom", pseudo: "pseudo"}
	// d = d.setNamePseudo("Test", "Bonjour")
	// fmt.Print(d)
	fd := FakeDev{name: "fake", pseudo: "dev"}
	fd = fd.setNamePseudo("Not", "Fake")
	fmt.Print(fd)
	setAttributes(fd)
}
