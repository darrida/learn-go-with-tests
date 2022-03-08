package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!\n"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigureableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigureableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigureableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
