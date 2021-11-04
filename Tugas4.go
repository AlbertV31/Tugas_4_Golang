package main

import "fmt"

func main()  {
  var Aldo,Yosep = tampil_mahasiswa()
  fmt.Println("Aldo : ", Aldo, "cm")
  fmt.Println("Yosep : ", Yosep, "cm")
}

func tampil_mahasiswa()(int,int)  {
  var tinggi = map[string]int {"Aldo":182, "Yosep":178}
  var Aldo = tinggi["Aldo"]
  var Yosep = tinggi["Yosep"]
  return Aldo,Yosep
}
