package tasks

import "time"

func Schedulerepeatingtask(what func()) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()
	return stop
}

func Scheduletask(what func()) {
	go func() {
		what()
	}()
}
