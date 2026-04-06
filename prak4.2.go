package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int 
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			(*soal)++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int

	maxSoal := -1
	minSkor := 1<<31 - 1
	for {
		_, err := fmt.Scan(&nama)
		if err != nil || nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}