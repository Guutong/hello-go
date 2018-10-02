package main

import "testing"

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := getBNK48Cap()
	if result != "Cherprang <3" {
		t.Errorf("No!!! , This is not Cherprang !!")
	}
}
