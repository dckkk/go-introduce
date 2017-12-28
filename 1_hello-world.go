// setiap project harus memiliki package main
// ex: package <nama-package>
package main

// import digunakan untuk meng-include atau memasukkan package
// ex: import "<nama-package>"
import "fmt"
/* import "fmt" adalah salah satu package yang disediakan Golang, memiliki banyak
fungsi untuk keperluan I/O yang berhubungan dengan text*/

func main() {
  // func main() adalah function pertama yang dipanggil saat meng-eksekusi program

  fmt.Println("hello world")
  // fungsi ini digunakan untuk memunculkan text ke layar dan menghasilkan baris baru diakhir text
  fmt.Printf("hello world")
  // fungsi ini digunakan untuk memunculkan text ke layar tetapi tidak menghasilkan baris baru diakhir text
  fmt.Print("hello", "world")
  // fungsi ini digunakan untuk memunculkan text ke layar tetapi tidak menghasilkan baris baru diakhir text
  // dan menggunakan , untuk memisahkan setiap kata
  
}
