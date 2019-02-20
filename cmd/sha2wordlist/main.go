package main

import (
	"fmt"
	"os"

	pgpwords "github.com/jimmypw/golang-sha2wordlist"
)

func main() {
	pgpfn := pgpwords.File{
		Filename:   os.Args[1],
		Hashmethod: "sha256",
	}

	fmt.Printf("sha256 checksum: %s\n", pgpfn.Hexchecksum())
	fmt.Print("PGP wordlist: ")
	words := pgpfn.WordList()

	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[i])
	}

	fmt.Print("\n")
}
