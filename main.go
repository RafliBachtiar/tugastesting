package main

import (
	"errors"
	"fmt"
)

func PembayaranBarang(hargaTotal float64, metode string, dicicil bool) error {
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	switch metode {
	case "cod", "transfer", "debit", "gerai":
		if dicicil {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	case "credit":
		if dicicil {
			if hargaTotal <= 500000 {
				return errors.New("cicilan tidak memenuhi syarat")
			}
		} else {
			return errors.New("credit harus dicicil")
		}
	default:
		return errors.New("metode tidak dikenali")
	}

	return nil
}

func main() {
	harga := 600000.0
	metode := "cod"
	dicicil := false

	err := PembayaranBarang(harga, metode, dicicil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pembayaran berhasil")
	}
}
