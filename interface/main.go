package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (p Persegi) HitungLuas() int {
	return p.Sisi * p.Sisi
}

type persegiPanjang struct {
	Panjang int
	Lebar   int
}

func (pp persegiPanjang) HitungLuas() int {
	return pp.Panjang * pp.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 1001
}

func main() {
	persegiPanjang := persegiPanjang{Panjang: 10, Lebar: 5}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas persegi panjang:", luas)

	persegi := Persegi{Sisi: 4}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas persegi:", luas)

	asal := Asal{Angka: 10}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas asal:", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
