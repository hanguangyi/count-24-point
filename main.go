// Copyright 2018. All rights reserved.
// license Apache 2.0.
//Author：HanGuangyi
//Time：2018/04/23
//I did not consider using slices to reduce memory usage for calls between functions call.

//Package nextPermutation

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	flag := false
	num := []int{6, 3, 6, 3}
	fmt.Println("Here count 24 point program begin.")
	sort.Ints(num)
	//fmt.Printf("sort numbers:%v\n", nums)
	for nextPermutation(num, 0, len(num)-1) {
		//fmt.Printf("%v\n", num)
		expression := make([]string, 14)
		dfs(num[0], num[1], 1, flag, num, expression)
	}

}

func reserve(nums []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}

//Arrange arrays to return fully-aligned two-dimensional arrays.
func nextPermutation(array []int, first, last int) bool {
	if array == nil { //空区间
		fmt.Println("sort array is nil")
		return false
	}
	if last == 1 { //只有一个元素
		return false
	}
	//i := len(array) - 1
	//i -= i
	j := last

	for i := last; i > first; i-- {
		/* 以上锁定一组(两个)相邻元素 */
		if array[i-1] < array[i] { //如果前一个元素小于后一个元素
			for array[i-1] >= array[j] {
				j--
			}
			array[i-1], array[j] = array[j], array[i-1]
			reserve(array, i, last) //将i之后的元素全部逆序重排
			return true
		}
		if i-1 == first {
			reserve(array, first, last) // 全部逆序重排
			return false
		}

	}
	return false
}

//parttern:*A**B***C**D*
func dfs(sum, cur, m int, flag bool, num []int, expression []string) {
	if flag {
		return
	}
	//fmt.Printf("bill debug dfs %d\r\n", m)
	if m == 1 {
		//situation 1: A+B,C
		//expression[1] = sum
		expression[1] = strconv.Itoa(sum)
		expression[4] = strconv.Itoa(cur)
		expression[8] = strconv.Itoa(num[m+1])
		//expression1 := make([]string, len(expression))
		expression[2] = "+"
		expression1 := expression
		dfs(sum+cur, num[m+1], m+1, flag, num, expression1) //先计算前一部分

		expression[2] = "-"
		expression2 := expression
		dfs(sum-cur, num[m+1], m+1, flag, num, expression2)

		expression[2] = "*"
		expression5 := expression
		dfs(sum*cur, num[m+1], m+1, flag, num, expression5)

		if cur != 0 && sum%cur == 0 {
			expression[2] = "/"
			expression6 := expression
			dfs(sum/cur, num[m+1], m+1, flag, num, expression6)
		}

		//situation 2: A,B+C
		//First calculate the back following elements.
		expression[2] = ""

		expression[6] = "+"
		expression3 := expression
		dfs(sum, cur+num[m+1], m+1, flag, num, expression3) //先计算后一部分

		expression[6] = "-"
		expression4 := expression
		dfs(sum, cur-num[m+1], m+1, flag, num, expression4)

		expression[6] = "*"
		expression7 := expression
		dfs(sum, cur*num[m+1], m+1, flag, num, expression7)

		if num[m+1] != 0 && cur%num[m+1] == 0 {
			expression[6] = "/"
			expression8 := expression
			dfs(sum, cur/num[m+1], m+1, flag, num, expression8)
		}
	}
	//fmt.Printf("bill debug m=%d : %s\r\n", m, expression)
	if m == 2 {
		expression[11] = strconv.Itoa(num[m+1]) //now 4 elements complete

		expressionA := make([]string, 14)
		copy(expressionA, expression)

		expressionB := make([]string, 14)
		copy(expressionB, expression)

		expressionC := make([]string, 14)
		copy(expressionC, expression)

		expressionD := make([]string, 14)
		copy(expressionD, expression)

		//situation 1: A+B+C,D and A+B,C+D
		if expression[2] == "+" || expression[2] == "-" || expression[2] == "*" || expression[2] == "/" {
			expressionA[13] = "A"
			expressionA[6] = "+"
			//expression1 := expression
			dfs(sum+cur, num[m+1], m+1, flag, num, expressionA) //先计算前一部分

			expressionA[6] = "-"
			//expression2 := expression
			dfs(sum-cur, num[m+1], m+1, flag, num, expressionA) //先计算前一部分

			//expressionA = expression
			//expressionA[13] = "A"
			if expression[2] == "+" || expression[2] == "-" {
				expressionA[0] = "("
				expressionA[5] = ")"
			}
			expressionA[6] = "*"
			//expression5 := expression
			dfs(sum*cur, num[m+1], m+1, flag, num, expressionA)

			if cur != 0 && sum%cur == 0 {
				expressionA[6] = "/"
				//expression6 := expression
				dfs(sum/cur, num[m+1], m+1, flag, num, expressionA)
			}

			expressionB[13] = "B"
			expressionB[10] = "+"
			//expression3 := expression
			dfs(sum, cur+num[m+1], m+1, flag, num, expressionB) //先计算后一部分

			expressionB[10] = "-"
			//expression4 := expression
			dfs(sum, cur-num[m+1], m+1, flag, num, expressionB)

			expressionB[10] = "*" //should consider ()
			dfs(sum, cur*num[m+1], m+1, flag, num, expressionB)

			if num[m+1] != 0 && cur%num[m+1] == 0 {
				expressionB[10] = "/"
				dfs(sum, cur/num[m+1], m+1, flag, num, expressionB)
			}
		}

		//situation 2: A+(B+C),D and A,(B+C)+D
		if expression[6] == "+" || expression[6] == "-" || expression[6] == "*" || expression[6] == "/" {
			expressionC[13] = "C"
			expressionC[2] = "+"
			//expression1 := expression
			dfs(sum+cur, num[m+1], m+1, flag, num, expressionC) //先计算前一部分

			expressionC[2] = "-"
			//expression2 := expression
			dfs(sum-cur, num[m+1], m+1, flag, num, expressionC) //先计算前一部分

			if expression[6] == "+" || expression[6] == "-" {
				expressionC[3] = "("
				expressionC[9] = ")"
			}
			expressionC[2] = "*"
			dfs(sum*cur, num[m+1], m+1, flag, num, expressionC)

			if cur != 0 && sum%cur == 0 {
				expressionC[2] = "/"
				dfs(sum/cur, num[m+1], m+1, flag, num, expressionC)
			}
			//fmt.Printf("bill debug slice %s\r\n", expressionD)
			expressionD[13] = "D"
			expressionD[10] = "+"
			dfs(sum, cur+num[m+1], m+1, flag, num, expressionD) //先计算后一部分

			expressionD[10] = "-"
			//expression4 := expression
			dfs(sum, cur-num[m+1], m+1, flag, num, expressionD)

			if expression[6] == "+" || expression[6] == "-" {
				expressionD[3] = "("
				expressionD[9] = ")"
			}

			expressionD[10] = "*" //should consider ()
			//expression7 := expression
			dfs(sum, cur*num[m+1], m+1, flag, num, expressionD)

			if num[m+1] != 0 && cur%num[m+1] == 0 {
				expressionD[10] = "/"
				//expression8 := expression
				dfs(sum, cur/num[m+1], m+1, flag, num, expressionD)
			}
		}
	}

	if m == 3 {
		switch expression[13] {
		case "A":
			if sum+cur == 24 {
				expression[10] = "+"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum-cur == 24 {
				expression[10] = "-"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum*cur == 24 {
				expression[10] = "*"
				if expression[6] == "+" || expression[6] == "-" {
					expression[0] = "("
					expression[9] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true

			}
			if cur != 0 && sum%cur == 0 && sum/cur == 24 {
				expression[10] = "/"
				if expression[6] == "+" || expression[6] == "-" {
					expression[0] = "("
					expression[9] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
		case "B":
			if sum+cur == 24 {
				expression[6] = "+"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum-cur == 24 {
				expression[6] = "-"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum*cur == 24 {
				expression[6] = "*"
				if expression[2] == "+" || expression[2] == "-" {
					expression[0] = "("
					expression[5] = ")"
				}
				if expression[10] == "+" || expression[10] == "-" {
					expression[7] = "("
					expression[12] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true

			}
			if cur != 0 && sum%cur == 0 && sum/cur == 24 {
				expression[6] = "/"
				if expression[2] == "+" || expression[2] == "-" {
					expression[0] = "("
					expression[5] = ")"
				}
				if expression[10] == "+" || expression[10] == "-" {
					expression[7] = "("
					expression[12] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
		case "C":
			if sum+cur == 24 {
				expression[10] = "+"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum-cur == 24 {
				expression[10] = "-"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum*cur == 24 {
				expression[10] = "*"
				if expression[2] == "+" || expression[2] == "-" {
					expression[0] = "("
					expression[9] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true

			}
			if cur != 0 && sum%cur == 0 && sum/cur == 24 {
				expression[10] = "/"
				if expression[2] == "+" || expression[2] == "-" {
					expression[0] = "("
					expression[9] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
		case "D":
			if sum+cur == 24 {
				expression[2] = "+"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum-cur == 24 {
				expression[2] = "-"
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
			if sum*cur == 24 {
				expression[2] = "*"
				if expression[10] == "+" || expression[10] == "-" {
					expression[3] = "("
					expression[12] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true

			}
			if cur != 0 && sum%cur == 0 && sum/cur == 24 {
				expression[2] = "/"
				if expression[10] == "+" || expression[10] == "-" {
					expression[3] = "("
					expression[12] = ")"
				}
				fmt.Printf("%v\n", expression[0:13])
				flag = true
			}
		}
		return
	}

}
