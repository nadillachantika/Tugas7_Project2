package main

import "fmt"

type Data struct {
	Nama  string
	Kelas string
}

type Kelas struct {
	Kelas string
	Wali  string
}

func main() {

	var req Data
	var res Kelas

	req.Nama = "Yunus"
	req.Kelas = "Kelas 2"

	res.Kelas = "Kelas 1"

	fmt.Println(inspectResponseCode(res))

}

func inspectResponseCode(req interface{}) string {

	//var msg Data
	var ok bool
	//var Class string
	
	fmt.Println(req)
    
	if req, ok = req. ; ok {
		fmt.Println("masuk dia")
	} else {
		fmt.Println("keluar dia")
	}





	// if req, ok = req{}; ok {
	// 	fmt.Println("masuk dia")
	// } else {
	// 	fmt.Println("keluar dia")
	// }

	// if msg, ok = req.(Data); ok {
	// 	Class = msg.Kelas
	// } else if xml, ok := req.(Kelas); ok {
	// 	msg = Data{
	// 		Kelas: xml.Kelas,
	// 	}
	// 	Class = msg.Kelas
	// } else {
	// 	Class = "NONE"
	// }

	return "Test"

}