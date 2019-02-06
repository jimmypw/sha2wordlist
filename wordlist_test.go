package pgpwords

import (
	"testing"
)

var binChecksumReference = [32]byte{
	46, 24, 42, 253, 33,
	51, 41, 68, 2, 45,
	223, 111, 22, 94, 233,
	177, 179, 124, 203, 113,
	164, 133, 125, 36, 76,
	170, 128, 213, 134, 85,
	47, 204,
}

var hexChecksumReference = "2e182afd21332944022ddf6f165ee9b1b37ccb71a4857d244caa80d586552fcc"

var wordListReference = [32]string{"buzzard", "borderline", "brickyard", "Wyoming", "blackjack",
	"concurrent", "breakup", "designing", "accrue", "clergyman",
	"talon", "hemisphere", "backward", "finicky", "treadmill",
	"photograph", "scallion", "informant", "spheroid", "hideaway",
	"regain", "leprosy", "klaxon", "Capricorn", "drainage",
	"pedigree", "merit", "specialist", "necklace", "equipment",
	"cement", "revolver"}

func TestBinCheckSum(t *testing.T) {
	testfile := File{
		Filename:   "testdata/testfile",
		Hashmethod: "sha256",
	}

	binChecksumResult := testfile.Binchecksum()

	for i := 0; i < len(binChecksumResult); i++ {
		if binChecksumResult[i] != binChecksumReference[i] {
			t.Error("Binary checksum does not match")
		}
	}
}

func TestHexCheckSum(t *testing.T) {
	testfile := File{
		Filename:   "testdata/testfile",
		Hashmethod: "sha256",
	}

	if testfile.Hexchecksum() != hexChecksumReference {
		t.Errorf("Hex checksum does not match. got %s, expected %s\n", testfile.Hexchecksum(), hexChecksumReference)
	}
}

func TestWordList(t *testing.T) {
	testfile := File{
		Filename:   "testdata/testfile",
		Hashmethod: "sha256",
	}

	resultWordList := testfile.WordList()

	for i := 0; i < len(resultWordList); i++ {
		if resultWordList[i] != wordListReference[i] {
			t.Errorf("Word list does not match. got %s, expected %s\n", resultWordList[i], wordListReference[i])
		}
	}
}
