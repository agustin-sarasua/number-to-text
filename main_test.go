package main_test

import (
	"fmt"
	"testing"

	main "github.com/agustin-sarasua/number-to-text"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message + " " + a.(string) + " vs " + b.(string))
}

func TestParseNumber(t *testing.T) {

	assertEqual(t, main.ParseNumber(0), "zero", "ERROR")
	assertEqual(t, main.ParseNumber(63), "sesenta y tres", "ERROR")
	assertEqual(t, main.ParseNumber(21), "veinte y uno", "ERROR")
	assertEqual(t, main.ParseNumber(54), "cincuenta y cuatro", "ERROR")
	assertEqual(t, main.ParseNumber(45), "cuarenta y cinco", "ERROR")
	assertEqual(t, main.ParseNumber(99), "noventa y nueve", "ERROR")

	assertEqual(t, main.ParseNumber(144), "ciento cuarenta y cuatro", "ERROR")
	assertEqual(t, main.ParseNumber(131), "ciento treinta y uno", "ERROR")
	assertEqual(t, main.ParseNumber(654), "seiscientos cincuenta y cuatro", "ERROR")
	assertEqual(t, main.ParseNumber(999), "novecientos noventa y nueve", "ERROR")
	assertEqual(t, main.ParseNumber(301), "trescientos uno", "ERROR")

	assertEqual(t, main.ParseNumber(1000), "mil", "ERROR")
	assertEqual(t, main.ParseNumber(1404), "mil cuatrocientos cuatro", "ERROR")
	assertEqual(t, main.ParseNumber(1310), "mil trescientos diez", "ERROR")
	assertEqual(t, main.ParseNumber(6540), "seis mil quinientos cuarenta", "ERROR")
	assertEqual(t, main.ParseNumber(99900), "noventa y nueve mil novecientos", "ERROR")
	assertEqual(t, main.ParseNumber(309991), "trescientos nueve mil novecientos noventa y uno", "ERROR")
	assertEqual(t, main.ParseNumber(800001), "ochocientos mil uno", "ERROR")
	assertEqual(t, main.ParseNumber(909991), "novecientos nueve mil novecientos noventa y uno", "ERROR")
	assertEqual(t, main.ParseNumber(999999), "novecientos noventa y nueve mil novecientos noventa y nueve", "ERROR")

	assertEqual(t, main.ParseNumber(1000000), "un millon", "ERROR")
	assertEqual(t, main.ParseNumber(2000000), "dos millones", "ERROR")
	assertEqual(t, main.ParseNumber(1000001), "un millon uno", "ERROR")
	assertEqual(t, main.ParseNumber(4300000), "cuatro millones trescientos mil", "ERROR")
	assertEqual(t, main.ParseNumber(999999999), "novecientos noventa y nueve millones novecientos noventa y nueve mil novecientos noventa y nueve", "ERROR")

}
