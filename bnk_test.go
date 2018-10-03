package main

import "testing"

func TestGetBNK48Cap(t *testing.T) {
	result := getBNK48Cap()
	if result != "Cherprang <3" {
		t.Errorf("No!!! , This is not Cherprang !!")
	}
}

func TestGetFullNameOfBnk48Member(t *testing.T) {
	ornResult := getFullNameOfBnk48Member("ORN")
	if ornResult != "PATCHANAN JIAJIRACHOTE" {
		t.Errorf("No!!! , This is not full name nong Orn *_* !!")
	}

	cherprangResult := getFullNameOfBnk48Member("CHERPRANG")
	if cherprangResult != "CHERPRANG AREEKUL" {
		t.Errorf("No!!! , This is not full name nong Cherang *_* !!")
	}
}
