package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("COUNT : %d\n", i)
			if i == 5 {
				goto GOTO_FINISH
			}
			i++
		}
	}

GOTO_FINISH:
	fmt.Println("GOTO_FINISH")
}
