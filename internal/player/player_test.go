package player

import "testing"

func TestNewPlayer(t *testing.T) {
	p, _ := NewPlayer("Ajay", 100, 10, 15)
	if p.Name() != "Ajay" || p.Health() != 100 || p.Strength() != 10 || p.Attack() != 15 {
		t.Error("NewPlayer returned incorrect values")
	}

    _, err := NewPlayer("", 100, 10, 15)
    if err == nil || err.Error() != "name cannot be empty" {
		t.Error("NewPlayer did not return expected error for empty name")
	}

    _, err = NewPlayer("Ajay", -100, 10, 15)
    if err == nil || err.Error() != "atrributes cannot be negative" {
		t.Error("NewPlayer did not return expected error for empty name")
	}
}

func TestSetName(t *testing.T) {
	p, _ := NewPlayer("Ajay", 100, 10, 15)
	err := p.SetName("")
	if err == nil || err.Error() != "name cannot be empty" {
		t.Error("SetName did not return expected error for empty name")
	}
	err = p.SetName("Vijay")
	if err != nil {
		t.Error("SetName returned error for valid name")
	}
}

func TestSetHealth(t *testing.T) {
	p, _ := NewPlayer("Ajay", 100, 10, 15)
	err := p.SetHealth(-50)
	if err == nil || err.Error() != "health cannot be negative" {
		t.Error("SetHealth did not return expected error for negative health")
	}
	err = p.SetHealth(50)
	if err != nil {
		t.Error("SetHealth returned error for valid health")
	}
}

func TestSetAttack(t *testing.T) {
	p, _ := NewPlayer("Ajay", 100, 10, 15)
	err := p.SetAttack(-20)
	if err == nil || err.Error() != "attack cannot be negative" {
		t.Error("SetAttack did not return expected error for negative attack")
	}
	err = p.SetAttack(20)
	if err != nil {
		t.Error("SetAttack returned error for valid attack")
	}
}

func TestSetStrength(t *testing.T) {
	p, _ := NewPlayer("Ajay", 100, 10, 15)
	err := p.SetStrength(-10)
	if err == nil || err.Error() != "strength cannot be negative" {
		t.Error("SetStrength did not return expected error for negative strength")
	}
	err = p.SetStrength(10)
	if err != nil {
		t.Error("SetStrength returned error for valid strength")
	}
}
