package main

import "testing"

func TestPembayaranBarang(t *testing.T) {
	tests := []struct {
		name     string
		harga    float64
		metode   string
		dicicil  bool
		wantErr  bool
		errorMsg string
	}{
		{"Harga tidak boleh nol", 0, "cod", false, true, "harga tidak bisa nol"},
		{"Metode tidak dikenali", 100000, "paypal", false, true, "metode tidak dikenali"},
		{"Credit harus dicicil", 600000, "credit", false, true, "credit harus dicicil"},
		{"Cicilan tidak memenuhi syarat", 400000, "credit", true, true, "cicilan tidak memenuhi syarat"},
		{"Pembayaran berhasil", 600000, "credit", true, false, ""},
		{"Pembayaran berhasil", 400000, "debit", false, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PembayaranBarang(tt.harga, tt.metode, tt.dicicil)

			if tt.wantErr && err == nil {
				t.Errorf("PembayaranBarang() error = nil, want error")
				return
			}

			if !tt.wantErr && err != nil {
				t.Errorf("PembayaranBarang() error = %v, want nil", err)
				return
			}

			if err != nil && err.Error() != tt.errorMsg {
				t.Errorf("PembayaranBarang() error message = %v, want %v", err.Error(), tt.errorMsg)
			}
		})
	}
}
