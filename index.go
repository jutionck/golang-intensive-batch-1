package main

import (
	"fmt"
)

// Hanya 1 tidak boleh di duakan.
func main() {
	// Kalo mau buat sebuah variable
	// penamaan itu harus menggunakan kata benda (noun)
	// kemudian gaya penulisan itu menggunakan format camelCase
	// pembuatan variabel itu harus sesuai dengan tujuan nya

	// var
	// :=
	// age := 17
	// address := "Jakarta"
	// isStatus := true

	// fmt.Printf("%d %s %t\n", age, address, isStatus)
	// untuk config ke db atau API
	// result := fmt.Sprintf("%d %s %t", age, address, isStatus)
	// fmt.Println(result)

	// Condition -> memilih antara 2 atau lebih pilihan
	// (condition) ? 'true' : 'false'
	// if else if && switch

	// price = 150,000,-
	// discount ? => 10%, 15%

	// point = 8000
	// persentase => 8000 / 100 = 80%
	// point := 8000
	// if percentase := point / 100; percentase > 80 {
	// 	fmt.Println("Good")
	// } else {
	// 	fmt.Println("Bad")
	// }

	// // manggil si function getAll
	// // _, err := getAll()
	// // if err != nil {

	// // }

	// var (
	// 	grade  string
	// 	tugas  float64
	// 	uts    float64
	// 	uas    float64
	// 	result float64
	// )

	// tugas = 85
	// uts = 90
	// uas = 100
	// result = tugas*0.25 + uts*0.3 + uas*0.45

	// if result >= 80 && result <= 100 {
	// 	grade = "A"
	// } else if result >= 70 {
	// 	grade = "B"
	// } else if result >= 55 {
	// 	grade = "C"
	// } else if result >= 40 {
	// 	grade = "D"

	// } else {
	// 	grade = "E"
	// }

	// fmt.Printf("Nilai %.2f dengan hasil akhir adalah %s\n", result, grade)

	// // Switch -> Bulan 1 s.d 12
	// month := 3
	// switch month {
	// case 1:
	// 	fmt.Println("Januari")
	// case 2:
	// 	fmt.Println("Februari")

	// default:
	// 	fmt.Println("Bulan tidak ditemukan")
	// }

	// // Penggunaan operator logika di switch (&&, ||)
	// score := 85
	// switch {
	// case score > 81:
	// 	fmt.Println("Perfecto")
	// case (score <= 80) && (score >= 60):
	// 	fmt.Println("Good")
	// default:
	// 	fmt.Println("Your score is bad!")
	// }

	// // _ mengabaikan used
	// // golang sangat posesif atau strict terhadap variabel yang di buat

	// data, _ := "Data Product", errors.New("Error")
	// fmt.Println(data)

	// // Array -> kumpulan data yang bertipe data sama e.g string, int, bool, dst
	// // Array -> [20] -> jumlah elemen yang di tampung ditulis, e.g 20
	// // Array -> [...] -> variadic
	// // Array dibuat bisa juga menggunakan keyword make
	// // Array ada sebuah index yang dimulai dari angka 0
	// // [5] -> 0,1,2,3,4

	// var students [5]string
	// students[0] = "John"
	// students[1] = "Max"
	// students[2] = "Nindya"
	// students[3] = "Maia"
	// students[4] = "Felix"
	// // students[5] = "Jack" // error

	// fmt.Println(students)
	// for i := 0; i < len(students); i++ {
	// 	fmt.Println(students[i])
	// }

	// fmt.Println("===================")
	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	// //
	// var employees = [...]string{"John", "Felix", "Maia", "Fadli", "Pablo", "Andry"}
	// fmt.Println("===================")
	// for _, employee := range employees {
	// 	fmt.Println(employee)
	// }

	// Keyword make
	var cars = make([]string, 5) // 0,1,2
	cars[0] = "Avanza"
	cars[1] = "Ferrary"
	cars[2] = "Lamborghini"
	// cars[3] = "Brio" // error -> index out of range
	// fmt.Println(cars[1]) // Ferary
	// cars[1] = "Bemo"
	// fmt.Println(cars[1])

	// Slice
	// -> reference dari sebuah array
	// kalo ada kata reference -> berarti 1 alamat / address memori yang sama
	// meng-reference array cars dari index ke-0 hingga sebelum index ke-2
	var sliceOfCars = cars[0:2] // Avanza, Ferary
	// fmt.Println(sliceOfCars)

	sliceOfCars = cars[0:4] // mengambil semua elemen si array of cars
	sliceOfCars[0] = "Sigra"
	fmt.Println("len(): ", len(sliceOfCars)) // 4
	fmt.Println("cap(): ", cap(sliceOfCars)) // 5
	// fmt.Println(sliceOfCars)

	// Menambah data di slice -> append
	var appendCar = append(sliceOfCars, "Jimmy")
	fmt.Println("appendCar: ", appendCar)
	// appendCar = append(sliceOfCars, "Daihatzu")
	appendCar = append(sliceOfCars, "Bugati", "Civic", "Terrios", "Jimny", "Alphard", "Mazda", "Wuling", "Datsun")
	fmt.Println("sliceOfCars:", sliceOfCars)
	fmt.Println("appendCar: ", appendCar)
	// fmt.Println(appendCar)

	// for _, car := range appendCar {
	// 	fmt.Println(car)
	// }

	fmt.Println("len(): ", len(appendCar)) // 2
	fmt.Println("cap(): ", cap(appendCar)) // 4

	// Map
	// Map itu bisa disebut sebagai array bertipe data asosiative,
	// Karena map memiliki sebuah key-value
	// map di buat dengan keyword map
	var months = map[int]string{
		1: "Januari",
		2: "Februari",
		3: "Maret",
		4: "April",
	}

	// fmt.Println(months[3])
	for _, month := range months {
		fmt.Println(month)
	}

	// Kombinasi antara slice dan map
	// var students = []map[string]string{
	// 	{"name": "John", "major": "Engineering"},
	// 	{"name": "Maia", "major": "Tax"},
	// }

	// for _, student := range students {
	// 	fmt.Println(student["name"], student["major"])
	// }

	var students map[string]map[string]map[string]string
	students = map[string]map[string]map[string]string{} // nil
	students["name"] = map[string]map[string]string{}    // student[name] = nil
	students["name"]["class"] = map[string]string{}      // student[name][class] = nil
	students["name"]["class"]["grade"] = "A"             // student[name][class][grade] = A
	fmt.Println(students)

	students02 := map[string]map[string]map[string]string{
		"name": {"class": {"grade": "A"}},
	}
	fmt.Println(students02)
}
