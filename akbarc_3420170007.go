package main

import (
	"fmt"
)

type barang struct {
	nama  string
	harga int
}

func main() {
	var br1 = barang{"Sabun Cuci Piring", 50000}
	var br2 = barang{"Sabun Mandi", 25000}
	var br3 = barang{"Shampoo", 35000}
	var br4 = barang{"Pasta Gigi", 10000}
	var br5 = barang{"Tissue", 7500}
	var a, b, c, d, e int
	fmt.Println("+=========[ Daftar Barang ]==========+")
	fmt.Println("| 1.", br1.nama, "  : Rp", br1.harga, " |")
	fmt.Println("| 2.", br2.nama, "        : Rp", br2.harga, " |")
	fmt.Println("| 3.", br3.nama, "            : Rp", br3.harga, " |")
	fmt.Println("| 4.", br4.nama, "         : Rp", br4.harga, " |")
	fmt.Println("| 5.", br5.nama, "             : Rp ", br5.harga, " |")
	fmt.Println("+====================================+")
	fmt.Println("Masukan Jumlah Barang Yang Akan Dibeli : ")
	fmt.Print(br1.nama, " : ")
	fmt.Scan(&a)
	fmt.Print(br2.nama, " : ")
	fmt.Scan(&b)
	fmt.Print(br3.nama, " : ")
	fmt.Scan(&c)
	fmt.Print(br4.nama, " : ")
	fmt.Scan(&d)
	fmt.Print(br5.nama, " : ")
	fmt.Scan(&e)
	fmt.Println("+=================================+")
	fmt.Println("|          Total Belanja          |")
	fmt.Println("+=================================+")
	fmt.Println("|", br1.nama, "Rp", br1.harga, "     |")
	fmt.Println("| X", a, "item            : Rp", br1.harga*a, "|")
	fmt.Println("|", br2.nama, "Rp", br2.harga, "           |")
	fmt.Println("| X", b, "item            : Rp", br2.harga*b, "|")
	fmt.Println("|", br3.nama, "Rp", br3.harga, "               |")
	fmt.Println("| X", c, "               : Rp", br3.harga*c, "|")
	fmt.Println("|", br4.nama, "Rp  ", br4.harga, "          |")
	fmt.Println("| X", d, "               : Rp", br4.harga*d, "|")
	fmt.Println("|", br5.nama, "Rp ", br5.harga, "                |")
	fmt.Println("| X", e, "               : Rp", br5.harga*e, "|")
	fmt.Println("+=================================+")
	fmt.Println("| Sub Total         = Rp ", (br1.harga*a)+(br2.harga*b)+(br3.harga*c)+(br4.harga*e)+(br5.harga*e), "|")
	fmt.Println("+=================================+")
}
