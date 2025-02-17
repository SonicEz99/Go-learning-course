package main

import (
	"fmt"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e", "f"}

	result := arr[2:6]

	result[0] = "z"

	fmt.Println(result)

	arr = append(arr, "g") //เพิ่มข้อมูลใน array

	fmt.Println(arr)

	arr = append(arr[:2], arr[3:]...) //ลบข้อมูลใน array

	fmt.Println(arr)
}
