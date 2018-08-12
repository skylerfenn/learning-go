package main

import (
    "testing"
)

func TestGetStarWarsCharacters(t *testing.T) {
    characters := getStarWarsCharacters(1)
    if len(characters) != 10 {
        t.Errorf("getStarWarsCharacters() == 10, but is len(%d)", len(characters))
    }

    getStarWarsCharacters(145)

    if len(characters) != 10 {
        t.Errorf("getStarWarsCharacters() == 10, but is len(%d)", len(characters))
    }   
}

func TestPrintPeople(t *testing.T) { 
    people := getStarWarsCharacters(1)
    printPeople(people)
}

func BenchmarkGetStarWarsCharacters(b *testing.B) {
    for i := 0; i < b.N; i++ {
        getStarWarsCharacters(1)
    }
}
