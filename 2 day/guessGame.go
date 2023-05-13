package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 官方版本
func main() {
	maxNum := 100
	// 定义随机种子为当前时间，如果没有设定随机种子，生成数一致
	rand.Seed(time.Now().UnixNano())
	// 设置随机数最高值n，最小值默认从零开始，即生成一个值在区间 [0, n) 的 Int 数
	secretNumber := rand.Intn(maxNum)
	fmt.Println("Please input your guess")
	// 读取文本
	reader := bufio.NewReader(os.Stdin)
	// 输入判断，猜数正确退出循环
	for {
		input, err := reader.ReadString('\n')
		// nil 即为 golang 的空值
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			// continue 返回循环开始处
			continue
		}
		// windows 需要修改换行符
		input = strings.Trim(input, "\r\n")
		// 利用 string 方法去除换行符，方便后续转化为数字，以及判断正确与否
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		// 判断数字大小，及其正确与否，不正确返回循环开始处，正确则结束循环
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			// 利用 break 结束循环
			break
		}
	}
}
