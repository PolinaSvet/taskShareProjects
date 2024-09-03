package templTask1

import (
	"fmt"
	"sync"
	"time"
)

const nameTask string = "18.2.1> "

func StartTask(channelMess chan string, maxValue int, numGoroutines int) {
	var wg sync.WaitGroup

	channelMess <- fmt.Sprintf(nameTask + " === Задание <" + nameTask + " (HW-04) === ")

	wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		go func(id int) {
			for ii := 1; ii <= maxValue; ii++ {
				channelMess <- fmt.Sprintf(nameTask+"%v: goroutine_%03d print name: %03d\n", time.Now().Format("15:04:05.000"), id, ii)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	channelMess <- fmt.Sprintf(nameTask + " === end === ")
}
