package main

import (
	"time"
)

func dateEventCallback(position *int, button *int) (f callbackFunc, err error) {
	return nil, nil
}

func dateCallback(position int) (err error) {
	sectionStatus[position] = time.Now().Format("3:04pm 5s on Monday Jan 02 2006") + " 📅"

	return nil
}
