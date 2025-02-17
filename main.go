package main

import "fmt"

func main() {
	fmt.Println("==========================================================================================")
	// var i int = 10;
	i := 20; //แบบนี้สั้นสุด เพราะคอมไปเลอร์สามารถเข้าใจว่าเป็น int ได้เอง
	//i = 20 อันนี้ก็เปลี่ยนค่าได้
	//i = false แบบนี้ไม่ได้

	// var j bool = true; default = flase
	j := true //อันนี้ก็เหมือนกัน
	
	// var x float64 = 10.5 
	x := 10.5
	y := int(x)

	fmt.Println("this var i:", i)
	fmt.Println("this bool j:", j)
	fmt.Println("thi is flot :", x)
	fmt.Println("convert from flot to int:", y)
	fmt.Println("==========================================================================================")
	useConst()
}

func useConst(){
	const pi = 3.14 // const ไม่สามารถเปลี่ยนค่าได้
	// pi = 2 ห้ามๆ

	const (
		sunday = iota //เริ่มต้นที่ 0 ที่เหลือก็จะเพิ่มขึ้นไปเรื่อยๆ ถ้าอยากได้ค่าเริ่มต้นเป็น 1 ก็แค่เอา iota+1
		monday 
		tuesday
		wednesday
		thursday
		_ //วิธีการข้าม
		saturday
	)

	fmt.Println(sunday, monday, tuesday, wednesday, thursday, saturday)
}
