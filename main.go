package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("==========================================================================================")
	// var i int = 10;
	i := 20 //แบบนี้สั้นสุด เพราะคอมไปเลอร์สามารถเข้าใจว่าเป็น int ได้เอง
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
	fmt.Println("==========================================================================================")
	ifAndElse()
	fmt.Println("==========================================================================================")
	loop()
	fmt.Println("==========================================================================================")
	switches()
	fmt.Println("==========================================================================================")
	fmt.Println(addNumber(10, 20))
	m, l := sum(10)
	fmt.Println(m, l)
	fmt.Println("==========================================================================================")
	pointer()
	fmt.Println("==========================================================================================")
	kk := 50
	incPointer(&kk)
	fmt.Println(kk)

	newnum := inc(kk)
	fmt.Println(newnum)
}

func useConst() {
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

func ifAndElse() {
	num := 20

	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	os := runtime.GOOS
	if os == "darwin" {
		fmt.Println("MacOS")
	} else if os == "linux" {
		fmt.Println("Linux")
	} else {
		fmt.Println("Window")
	}
}

func loop() {
	for i := 1; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println(i)
	}
}

func switches() {
	day := "Sunday"
	var days string = "day"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("weekends")
		break
	default:
		fmt.Println("Workday")
		break
	}

	fmt.Println(days)
}

func addNumber(a int, b int) int {
	return a + b
}

func sum(num int) (int, int) {
	return num + 1, num + 2
}

func pointer() {
	//0xc00000a1e0   21
	var p *int
	i := 21

	p = &i

	fmt.Println(p)
	fmt.Println(*p)
}

func inc(num int) int {
	return num + 5 //ทำสำเนาจาก kk
}

func incPointer(num *int) int {
	*num++
	return *num
}
