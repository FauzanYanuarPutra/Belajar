package entity

import (
	"strconv"
	"strings"
)

func TestAja(nama string, umur int) interface{} {
		namaKosong := strings.TrimSpace(nama)
	
		if namaKosong != "" {
			return "Hallo nama saya " + namaKosong + ", Umur saya " + strconv.Itoa(umur) + " Tahun"
		} else {
			return "Make sure field name and age have been added"
		}
}