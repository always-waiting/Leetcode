package main

import (
	"testing"
)

func Test_isNumber(t *testing.T) {
	//t.Log("+100", isNumber("+100"))
	//t.Log("5e2", isNumber("5e2"))
	//t.Log("-123", isNumber("-123"))
	//t.Log("3.1415", isNumber("3.1415"))
	//t.Log("-1E-16", isNumber("-1E-16"))
	//t.Log("0123", isNumber("0123"))
	t.Log("1", isNumber("1"))
	t.Log("3.", isNumber("3."))
	t.Log("12e", isNumber("12e"))
	t.Log("1a3.14", isNumber("1a3.14"))
	t.Log("1.2.3", isNumber("1.2.3"))
	t.Log("+-5", isNumber("+-5"))
	t.Log("12e5.4", isNumber("12e+5.4"))
}
