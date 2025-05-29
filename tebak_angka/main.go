package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var menu int
var tebakan int
var nyawa int

func teks() {
	fmt.Println("Selamat datang di Permainan Tebak Angka.")
	fmt.Println("Menu:")
	fmt.Println("1. Main")
	fmt.Println("2. Keluar")
	fmt.Println("Pilih menu:")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	angka := rand.Intn(1000) + 1 
	nyawa := 10

	teks()
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Println("Nyawamu masih:", nyawa)
		fmt.Println("Masukkan angka untuk menebak (1–1000):")
		for {

			fmt.Scan(&tebakan)
			selisih := int(math.Abs(float64(tebakan - angka)))
			nyawa -= 1
			if nyawa != 0 {
				fmt.Println("Nyawamu masih:", nyawa)
				if tebakan < angka {
					if selisih > 25 {
						fmt.Println("Tebakanmu terlalu kecil, sangat jauh. Coba lagi:")
					} else if selisih >= 10 {
						fmt.Println("Tebakanmu terlalu kecil, masih di bawah. Coba lagi:")
					} else {
						fmt.Println("Tebakanmu terlalu kecil, dikit lagi! Coba lagi:")
					}
				} else if tebakan > angka {
					if selisih > 25 {
						fmt.Println("Tebakanmu terlalu besar, sangat jauh. Coba lagi:")
					} else if selisih >= 10 {
						fmt.Println("Tebakanmu terlalu besar, masih di atas. Coba lagi:")
					} else {
						fmt.Println("Tebakanmu terlalu besar, dikit lagi! Coba lagi:")
					}
				} else {
					fmt.Println("Selamat! Tebakanmu benar!")
					break
				}
			} else {
				fmt.Println("Kesempatanmu Habis.")
				break
			}
		}
	} else if menu == 2 {
		fmt.Println("Sampai Jumpa!!!")
	} else {
		fmt.Println("Hanya menerima angka 1 & 2")
	}
}
