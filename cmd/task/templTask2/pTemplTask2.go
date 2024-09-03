package templTask2

import (
	"fmt"
	"time"
)

const nameTask string = "18.4.1> "

type Semaphore struct {
	sem     chan int
	timeout time.Duration
}

// Acquire - метод захвата семафора
func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- 0:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("Не удалось захватить семафор")
	}
}

// Release - метод освобождения семафора
func (s *Semaphore) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("Не удалось освободить семафор")
	}
}

// NewSemaphore - функция создания семафора
func NewSemaphore(counter int, timeout time.Duration) *Semaphore {
	return &Semaphore{
		sem:     make(chan int, counter),
		timeout: timeout,
	}
}

func doWork(channelMess chan string, item int, sem *Semaphore) {
	//semaphoreChan <- struct{}{}
	if err := sem.Acquire(); err != nil {
		channelMess <- fmt.Sprintf(nameTask+"%v: Код, в случае если кто-то уже работает с разделяемыми данными\n", time.Now().Format("15:04:05.000"))
	}
	go func() {
		defer func() {
			//<-semaphoreChan
			if err := sem.Release(); err != nil {
				channelMess <- fmt.Sprintf(nameTask+"%v: Код, в случае если по ошибке освобождаем свободный семафор\n", time.Now().Format("15:04:05.000"))
			}
		}()
		channelMess <- fmt.Sprintf(nameTask+"%v: item %04d\n", time.Now().Format("15:04:05.000"), item)
		time.Sleep(2 * time.Second)
	}()
}

func StartTask(channelMess chan string, maxValue int, numGoroutines int) {
	sem := NewSemaphore(numGoroutines, 3*time.Second)
	channelMess <- fmt.Sprintf(nameTask + " === Задание <" + nameTask + " (HW-04) === ")

	for i := 1; i <= maxValue; i++ {
		doWork(channelMess, i, sem)
	}

}
