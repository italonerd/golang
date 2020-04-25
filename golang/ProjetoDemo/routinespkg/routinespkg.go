package routinespkg

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}
