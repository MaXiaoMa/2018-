package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answer := []string{
		1:  "B",
		2:  "C",
		3:  "A",
		4:  "C",
		5:  "A",
		6:  "C",
		7:  "D",
		8:  "A",
		9:  "B",
		10: "A",
	}
	file1, err := os.OpenFile("d:/001.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("文件打开失败")
		os.Exit(1)
	}
	defer file1.Close()
	numArr := [4]int{1, 2, 3, 4}
	for _, p1 := range numArr {
		for _, p2 := range numArr {
			for _, p3 := range numArr {
				for _, p4 := range numArr {
					for _, p5 := range numArr {
						for _, p6 := range numArr {
							for _, p7 := range numArr {
								for _, p8 := range numArr {
									for _, p9 := range numArr {
										for _, p10 := range numArr {
											strr := strconv.Itoa(p1) + strconv.Itoa(p2) + strconv.Itoa(p3) + strconv.Itoa(p4) + strconv.Itoa(p5) + strconv.Itoa(p6) + strconv.Itoa(p7) + strconv.Itoa(p8) + strconv.Itoa(p9) + strconv.Itoa(p10)
											for j := 0; j < 10; j++ {
												answer[j+1] = intToString(string(strr[j]))
											}
											fmt.Println("计算匹配中...", strr)
											if TwoTest(answer) {
												if ThreeTest(answer) {
													if FourTest(answer) {
														if FiveTest(answer) {
															if SixTest(answer) {
																if SevenTest(answer) {
																	if NineTest(answer) {
																		if TenTest(answer) {
																			fmt.Println("序号为: ", strr)
																			fmt.Println("找到一个匹配的答案排列: ", answer)
																			file1.WriteString("找到一个匹配的答案排列: ")
																			file1.Write([]byte("\n"))
																			file1.WriteString(strings.Join(answer, "-"))
																			file1.Write([]byte("\n"))
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	/* fmt.Println("2: ", TwoTest(answer))
	fmt.Println("3: ", ThreeTest(answer))
	fmt.Println("4: ", FourTest(answer))
	fmt.Println("5: ", FiveTest(answer))
	fmt.Println("6: ", SixTest(answer))
	fmt.Println("7: ", SevenTest(answer))
	fmt.Println("8: ", NineTest(answer))
	fmt.Println("9: ", TenTest(answer)) */
}
func generate() string {

	return generate()
}
func intToString(i string) string {
	switch i {
	case "1":
		return "A"
	case "2":
		return "B"
	case "3":
		return "C"
	case "4":
		return "D"
	default:
		return "O"
	}
}
func TwoTest(answer []string) bool {
	if answer[2] == "A" {
		if answer[5] != "C" {
			return false
		}
	}
	if answer[2] == "B" {
		if answer[5] != "D" {
			return false
		}
	}
	if answer[2] == "C" {
		if answer[5] != "A" {
			return false
		}
	}
	if answer[2] == "D" {
		if answer[5] != "B" {
			return false
		}
	}
	return true
}
func ThreeTest(answer []string) bool {
	if answer[3] == "A" {
		if answer[2] == answer[4] && answer[4] == answer[6] && answer[6] != answer[3] {
			return true
		}
	}
	if answer[3] == "B" {
		if answer[2] == answer[3] && answer[3] == answer[4] && answer[2] != answer[6] {
			return true
		}
	}
	if answer[3] == "C" {
		if answer[3] == answer[4] && answer[4] == answer[6] && answer[6] != answer[2] {
			return true
		}
	}
	if answer[3] == "D" {
		if answer[2] == answer[3] && answer[3] == answer[6] && answer[6] != answer[4] {
			return true
		}
	}
	return false
}
func FourTest(answer []string) bool {
	a := answer[1] == answer[5]
	b := answer[2] == answer[7]
	c := answer[1] == answer[9]
	d := answer[6] == answer[10]
	if answer[4] == "A" {
		if a && !b && !c && !d {
			return true
		}
	}
	if answer[4] == "B" {
		if !a && b && !c && !d {
			return true
		}
	}
	if answer[4] == "C" {
		if !a && !b && c && !d {
			return true
		}
	}
	if answer[4] == "D" {
		if !a && !b && !c && d {
			return true
		}
	}

	return false
}
func FiveTest(answer []string) bool {
	if answer[5] == "A" {
		if answer[8] == answer[5] && answer[8] != answer[4] && answer[8] != answer[9] && answer[8] != answer[7] {
			return true
		}
	}
	if answer[5] == "B" {
		if answer[8] != answer[5] && answer[4] == answer[5] && answer[9] != answer[5] && answer[7] != answer[5] {
			return true
		}
	}
	if answer[5] == "C" {
		if answer[8] != answer[5] && answer[4] != answer[5] && answer[9] == answer[5] && answer[7] != answer[5] {
			return true
		}
	}
	if answer[5] == "D" {
		if answer[8] != answer[5] && answer[4] != answer[5] && answer[9] != answer[5] && answer[7] == answer[5] {
			return true
		}
	}
	return false
}
func SixTest(answer []string) bool {
	a := answer[2] == answer[4] && answer[2] == answer[8]
	b := answer[1] == answer[6] && answer[6] == answer[8]
	c := answer[3] == answer[10] && answer[10] == answer[8]
	d := answer[5] == answer[9] && answer[9] == answer[8]
	if answer[6] == "A" {
		if a && !b && !c && !d {
			return true
		}
	}
	if answer[6] == "B" {
		if !a && b && !c && !d {
			return true
		}
	}
	if answer[6] == "C" {
		if !a && !b && c && !d {
			return true
		}
	}
	if answer[6] == "D" {
		if !a && !b && !c && d {
			return true
		}
	}
	return false
}
func SevenTest(answer []string) bool {
	_, str := Count(answer, "min")
	if str == "A" {
		if answer[7] != "C" {
			return false
		}
	}
	if str == "B" {
		if answer[7] != "B" {
			return false
		}
	}
	if str == "C" {
		if answer[7] != "A" {
			return false
		}
	}
	if str == "D" {
		if answer[7] != "D" {
			return false
		}
	}
	return true
}
func Count(answer []string, flag string) (count int, str string) {
	var a, b, c, d int
	hm := make(map[string]int)
	for _, value := range answer {
		if value == "A" {
			a++
			hm["A"] = a
		}
		if value == "B" {
			b++
			hm["B"] = b
		}
		if value == "C" {
			c++
			hm["C"] = c
		}
		if value == "D" {
			d++
			hm["D"] = d
		}
	}
	if a == 0 {
		hm["A"] = 0
	}
	if b == 0 {
		hm["B"] = 0
	}
	if c == 0 {
		hm["C"] = 0
	}
	if d == 0 {
		hm["D"] = 0
	}
	if flag == "max" {
		for key, value := range hm {
			if value > count {
				count = value
				str = key
			}
		}
	}
	if flag == "min" {
		count = 10
		for key, value := range hm {
			if value < count {
				count = value
				str = key
			}
		}
	}
	return
}
func NineTest(answer []string) bool {
	a := answer[1] == answer[6]
	b := answer[5] == answer[6]
	c := answer[5] == answer[10]
	d := answer[5] == answer[2]
	e := answer[5] == answer[9]
	if answer[9] == "A" {
		if a != b && a == c && a == d && a == e {
			return true
		}
	}
	if answer[9] == "B" {
		if a == b && a != c && a == d && a == e {
			return true
		}
	}
	if answer[9] == "C" {
		if a == b && a == c && a != d && a == e {
			return true
		}
	}
	if answer[9] == "D" {
		if a == b && a == c && a == d && a != e {
			return true
		}
	}

	return false
}

func TenTest(answer []string) bool {
	max, _ := Count(answer, "max")
	min, _ := Count(answer, "min")
	if max-min > 4 || max-min < 1 {
		return false
	}
	if max-min == 1 {
		if answer[10] != "D" {
			return false
		}
	}
	if max-min == 2 {
		if answer[10] != "B" {
			return false
		}
	}
	if max-min == 3 {
		if answer[10] != "A" {
			return false
		}
	}
	if max-min == 4 {
		if answer[10] != "C" {
			return false
		}
	}
	return true
}
