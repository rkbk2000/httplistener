package main

import (
	"os"
	"testing"
	"time"
)

var stop chan os.Signal

func TestMain(t *testing.T) {
	go func() {
		main()
	}()
	time.Sleep(time.Second * 2)
	stop <- os.Interrupt
	time.Sleep(time.Second * 5)
}
