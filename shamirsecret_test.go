package polypasshash

import (
	"fmt"
	"testing"
	"math/rand"
)

func TestComputeShare(t *testing.T) {
  random = rand.New(rand.NewSource(10))
	ss := ShamirSecret(2, "abc")
	share := ss.ComputeShare(1)

	number := share.shareNumber
  bytes := share.shareBytes

	if number != 1 || bytes[0] != 75 || bytes[1] != 215 || bytes[2] != 238 {
		t.Errorf("Received wrong result")
	}
}

func TestIsValidShare(t *testing.T) {
  random = rand.New(rand.NewSource(10))
	ss := ShamirSecret(2, "abc")
	share := ss.ComputeShare(1)

	rst := ss.IsValidShare(share)

	if rst != true {
		t.Errorf("Received wrong result")
	}
}

/*
func TestRecoverSecretData(t *testing.T) {
  random = rand.New(rand.NewSource(10))
	ss := ShamirSecret(2, "abc")
	share1 := ss.ComputeShare(1)
	share2 := ss.ComputeShare(2)
	share3 := ss.ComputeShare(3)

	var shares []share = []share{share1, share2}

	ss = ShamirSecret(2, "")
	secret := ss.recoverSecretData(shares)

  if secret != "abc" {
		t.Errorf("Received wrong result")
	}
}
*/

func TestFullLagurange(t *testing.T){
	a := []int{2,4,5}
	b := []int{14,30,32}
	rst := fullLagurange(a, b)

	if rst[0] != 43 || rst[1] != 168 || rst[2] != 150 {
		t.Errorf("Received wrong result")
	}
}

func TestF(t *testing.T) {
	a := 2
	b := []byte{97,98,99} //"abc" as string
	expected := 50

	rst := f(a, b)

	if rst != expected {
		t.Errorf("Received wrong result")
	}

}

func TestMultiplyPolynomials(t *testing.T) {
	a := []int{1, 3, 4}
	b := []int{4, 5}
	rst := multiplyPolynomials(a, b)

	if rst[0] != 4 || rst[1] != 9 || rst[2] != 31 || rst[3] != 20 {
		t.Errorf("Received wrong result")
	}
}

func TestAddPolynomials1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5}

	rst := addPolynomials(a, b)

	if rst[0] != 5 || rst[1] != 7 || rst[2] != 3 {
		t.Errorf("Received wrong result")
	}

}

func TestGf256Add(t *testing.T) {
	a, b := 3, 5
	expected := 6
	rst := gf256Add(a, b)

	if rst != expected {
		t.Errorf("Received wrong result")
	}
}

func TestGf256Sub(t *testing.T) {
	a, b := 3, 5
	expected := 6
	rst := gf256Sub(a, b)

	if rst != expected {
		t.Errorf("received wrong result")
	}
}

func TestGf256Mul(t *testing.T) {
	a, b := 3, 5
	expected := 15

	rst := gf256Mul(a, b)

	fmt.Println("")

	if rst != expected {
		t.Errorf("received wrong result")
	}
}

func TestGf256Div(t *testing.T) {
	a, b := 5, 3
	expected := 3

	rst:= gf256Div(a, b)

	if rst != expected {
		t.Errorf("received wrong result")
	}
}
