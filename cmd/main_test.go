package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestValidateConfiguration(t *testing.T) {
	valid := []Machine{
		{"Red", GreenOnly},
		{"Red and Green", RedOnly},
		{"Green", Mixed},
	}

	if !validateConfiguration(valid) {
		t.Error("Expected valid configuration to be accepted")
	}

	invalid := []Machine{
		{"Red", RedOnly},
		{"Red and Green", Mixed},
		{"Green", GreenOnly},
	}

	if validateConfiguration(invalid) {
		t.Error("Expected invalid configuration to be rejected")
	}
}

func TestBuyGum_Mixed(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := Machine{"Red and Green", Mixed}

	seenRed, seenGreen := false, false

	// Пробуем купить 100 раз, чтобы оба цвета точно появилис
	for i := 0; i < 100; i++ {
		gum := buyGum(m, r)
		if gum == RedOnly {
			seenRed = true
		} else if gum == GreenOnly {
			seenGreen = true
		}
		if seenRed && seenGreen {
			break
		}
	}

	if !seenRed || !seenGreen {
		t.Error("Expected mixed machine to produce both red and green gum over time")
	}
}

func TestRemoveItem(t *testing.T) {
	input := []GumColor{RedOnly, GreenOnly, Mixed}
	expected := []GumColor{RedOnly, Mixed}
	result := removeItem(input, GreenOnly)

	if len(result) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
	}
}

func TestGeneratePermutations(t *testing.T) {
	items := []GumColor{RedOnly, GreenOnly, Mixed}
	perms := generatePermutations(items)

	if len(perms) != 6 {
		t.Errorf("Expected 6 permutations, got %d", len(perms))
	}

	// Проверим уникальность
	seen := make(map[string]bool)
	for _, p := range perms {
		key := string(p[0]) + string(p[1]) + string(p[2])
		if seen[key] {
			t.Errorf("Duplicate permutation found: %v", p)
		}
		seen[key] = true
	}
}
