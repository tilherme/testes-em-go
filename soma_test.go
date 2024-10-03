package main

import "testing"

func TestSumCorrect(t *testing.T) {
	teste := soma(3, 3, 3) //arrenge
	resultado := 9         //act

	if teste != resultado { //assert
		t.Errorf("Esperava %d, mas obteve %d", resultado, teste)
	}
}
func TestSumIncorrect(t *testing.T) {
	teste := soma(3, 3, 3) //arrenge
	resultado := 91        //act

	if teste != resultado { //assert
		t.Errorf("Esperava %d, mas obteve %d", resultado, teste)
	}
}

func TestMultiplicationCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", resultado, teste)
	}
}

func TestMultiplicationIncorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 10

	if teste != resultado {
		t.Errorf("Esperava %d, mas obteve %d", resultado, teste)
	}
}
