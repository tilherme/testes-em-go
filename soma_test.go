package main

import "testing"

func TestSumCorrect(t *testing.T) {
	teste := sum(3, 3, 3) //arrenge
	resultado := 9        //act

	if teste != resultado { //assert
		t.Errorf("Esperava %d, mas obteve %d", resultado, teste)
	}
}
func TestSumIncorrect(t *testing.T) {
	teste := sum(3, 3, 3) //arrenge
	resultado := 91       //act

	if teste != resultado { //assert
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}

func TestMultiplicatIonCorrect(t *testing.T) {
	teste := multiplication(10, 10)
	resultado := 100

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}

func TestMultiplicationIncorrect(t *testing.T) {
	teste := multiplication(10, 10)
	resultado := 10

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}
func TestSubtractionCorrect(t *testing.T) {
	teste := subtraction(10, 1)
	resultado := 9

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}

func TestSubtractionIncorrect(t *testing.T) {
	teste := subtraction(10, 1)
	resultado := 8

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}

func TestDivisionCorrect(t *testing.T) {
	teste := division(4, 2)
	resultado := 2

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}

func TestDivisionIncorrect(t *testing.T) {
	teste := division(4, 2)
	resultado := 3

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", teste, resultado)
	}
}
