package main

import (
	"strings"
)

var (
	nums = []string{
		"uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve", "diez",
		"once", "doce", "trece", "catorce", "quince", "dieciseis", "diecisiete", "dieciocho", "diecinueve"}
	tNums = []string{"veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
	cNums = []string{"cien", "docientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos"}
)

const (
	oneHundred  = 100
	oneThousand = 1000
	oneMillion  = 1000000
	oneBillion  = 1000000000
)

func main() {

}

func ParseCents(n int) string {
	if n < 20 {
		if n == 0 {
			return ""
		}
		return nums[n-1]
	}

	if n < oneHundred {
		if n%10 == 0 {
			return tNums[n/10-2]
		}
		return tNums[n/10-2] + " y " + ParseCents(n%10)
	}

	if n < oneThousand {
		if n%oneHundred == 0 {
			return cNums[n/oneHundred-1]
		}
		if n/oneHundred == 1 {
			return cNums[n/oneHundred-1] + "to " + ParseCents(n%oneHundred)
		}
		return cNums[n/oneHundred-1] + " " + ParseCents(n%oneHundred)
	}
	panic("This function should not be called with values bigger than 999")
}

func ParseThousands(n int) string {
	if n == oneThousand {
		return "mil"
	}
	if n/oneThousand == 0 {
		return ParseCents(n % oneThousand)
	} else if n/oneThousand == 1 {
		return "mil " + ParseCents(n%oneThousand)
	} else {
		return ParseCents(n/oneThousand) + " mil " + ParseCents(n%oneThousand)
	}

}

func ParseMillions(n int) string {
	if n == oneMillion {
		return "un millon"
	}

	if n/oneMillion == 1 {
		return "un millon " + ParseThousands(n%oneMillion)
	} else if n%oneMillion > 0 {
		return ParseCents(n/oneMillion) + " millones " + ParseThousands(n%oneMillion)
	} else {
		return ParseCents(n/oneMillion) + " millones"
	}
}

func ParseNumber(n int) string {
	if n == 0 {
		return "zero"
	}
	if n < oneThousand {
		return strings.Trim(ParseCents(n), " ")
	}

	if n < oneMillion {
		return strings.Trim(ParseThousands(n), " ")
	}

	if n < oneBillion {
		return strings.Trim(ParseMillions(n), " ")
	}
	return "Not implemented"
}
