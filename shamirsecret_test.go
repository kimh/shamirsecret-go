package polypasshash

import (
	//"fmt"
	"testing"
	"math/rand"
)

func TestComputeShare(t *testing.T) {
  random = rand.New(rand.NewSource(10))
	ss := ShamirSecret(2, "abc")
	share := ss.ComputeShare(1)

	number := share.shareNumber
  bytes := share.shareBytes

	if number != 1 || bytes[0] != 15 || bytes[1] != 146 || bytes[2] != 88 {
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

func TestRecoverSecretData(t *testing.T) {
  random = rand.New(rand.NewSource(10))
	ss := ShamirSecret(2, "Adi Shamir")
	share1 := ss.ComputeShare(1)
	share2 := ss.ComputeShare(2)

	var shares []share = []share{share1, share2}

	ss = ShamirSecret(2, "")
	secret := ss.RecoverSecretData(shares)

  if secret != "Adi Shamir" {
		t.Errorf("Received wrong result")
	}
}

func TestFullLagurange(t *testing.T){
	a := []byte{1,2}
	b := []byte{19,133}
	rst := fullLagurange(a, b)

	if rst[0] != 97 || rst[1] != 114 {
		t.Errorf("Received wrong result")
	}
}

func TestF(t *testing.T) {
  var a byte = 2
	b := []byte{97,98,99} //"abc" as string
	var expected byte = 50

	rst := f(a, b)

	if rst != expected {
		t.Errorf("Received wrong result")
	}

}

func TestMultiplyPolynomials(t *testing.T) {
	a := []byte{1, 3, 4}
	b := []byte{4, 5}
	rst := multiplyPolynomials(a, b)

	if rst[0] != 4 || rst[1] != 9 || rst[2] != 31 || rst[3] != 20 {
		t.Errorf("Received wrong result")
	}
}

func TestAddPolynomials1(t *testing.T) {
	a := []byte{1, 2, 3}
	b := []byte{4, 5}

	rst := addPolynomials(a, b)

	if rst[0] != 5 || rst[1] != 7 || rst[2] != 3 {
		t.Errorf("Received wrong result")
	}

}

func TestGf256Add(t *testing.T) {
	var a, b, expected byte = 3, 5, 6
	rst := gf256Add(a, b)

	if rst != expected {
		t.Errorf("Received wrong result")
	}
}

func TestGf256Sub(t *testing.T) {
	var a, b, expected byte = 3, 5, 6
	rst := gf256Sub(a, b)

	if rst != expected {
		t.Errorf("received wrong result")
	}
}

func TestGf256Mul(t *testing.T) {
	var a, b byte = 3, 5
	var expected byte = 15

	rst := gf256Mul(a, b)

	if rst != expected {
		t.Errorf("received wrong result")
	}
}

func TestGf256Div(t *testing.T) {
	var a, b byte = 5, 3
	var expected byte = 3

	rst:= gf256Div(a, b)

	if rst != expected {
		t.Errorf("received wrong result")
	}
}
