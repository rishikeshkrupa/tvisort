package algos

import (
	"fmt"
	"time"
)

func Bubble(a []int) {
	//blubble slort
	//show original input
	fmt.Printf("Input  : %v\n", a)
	fmt.Printf("Output : %v ", a)

	for i, _ := range a {

		totalIndex := len(a)
		lastIndex := totalIndex - i

		b := a[:lastIndex]

		rem := float64(len(a) - len(b))
		per := rem / float64(len(a))
		per = per * 100

		fmt.Printf(" %-3d%%", int(per))

		//hide the cursor
		fmt.Printf("\u001b[?25l")

		for j, v := range b {

			var (
				bigNum   int
				smallNum int
			)

			if j == lastIndex-1 {
				break
			}

			if b[j] > b[j+1] {
				b[j] = b[j+1]
				b[j+1] = v

				bigNum = j
				smallNum = j + 1

			} else {

				bigNum = j + 1
				smallNum = j

			}

			fmt.Print("\rOutput : [")

			for n, _ := range a {

				if n == bigNum {
					fmt.Print("\u001b[1m")
					fmt.Print("\u001b[31m")
				}

				if n == smallNum {
					fmt.Print("\u001b[32m")
				}

				if n > len(b)-1 {
					fmt.Print("\u001b[90m")
				}

				fmt.Printf("%v ", a[n])
				fmt.Print("\u001b[0m")
			}

			fmt.Print("\b] ")
			time.Sleep(time.Millisecond * 500)

		}

	}

	fmt.Printf("\r\u001b[2KOutput : %v %d %%\n", a, 100)

	//unhide the cursor
	fmt.Printf("\u001b[?25h")

}
