package day9

import (
	"testing"
)

func TestFindUnrulyNumber(t *testing.T) {
	res, err := findUnrulyNumber(5, []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != 127 {
		t.Fatalf("Wrong result: %d", res)
	}
}

func TestFindEncryptionWeakness(t *testing.T) {
	res, err := findEncryptionWeakness(5, []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	})

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res != 62 {
		t.Fatalf("Wrong result: %d", res)
	}
}
