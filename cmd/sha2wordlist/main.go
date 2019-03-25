package main

import (
	"fmt"
	"os"

	sha2wordlist "github.com/jimmypw/sha2wordlist"
)

func usage() {
	fmt.Printf("Error: No input file specified.\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("%s /bin/bash\n", os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	pgpfn := sha2wordlist.File{
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
