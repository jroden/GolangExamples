package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createdeck()
	if len(d) != 6 {
		t.Errorf("Expected deck length of 6, but got %v", len(d))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := createdeck()
	d.savetofile("_decktesting")
	nd := newdeckfromfile("_decktesting")
	if len(nd) != 6 {
		t.Errorf("Expected deck length of 6, but got %v", len(d))
	}
	os.Remove("_decktesting")
}
