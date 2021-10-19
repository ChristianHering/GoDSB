package main

import (
	"os/exec"
)

func audioEventCallback(position *int, button *int) (f callbackFunc, err error) {
	if *button == 0 {
		return nil, audioCallback(*position)
	}

	return nil, exec.Command("alacritty", "-e", "pulsemixer").Start()
}

func audioCallback(position int) error {
	b, err := exec.Command("pamixer", "--get-volume").Output()
	if err != nil {
		return err
	}

	sectionStatus[position] = string(b) + "%"

	return nil
}

func audioUpEventCallback(position *int, button *int) (f callbackFunc, err error) {
	*position = *position - 1
	*button = 0

	return audioEventCallback, exec.Command("pamixer", "-i", "2").Run()
}

func audioUpCallback(position int) error {
	sectionStatus[position] = " 🔊"

	return nil
}

func audioDownEventCallback(position *int, button *int) (f callbackFunc, err error) {
	*position = *position + 1
	*button = 0

	return audioEventCallback, exec.Command("pamixer", "-d", "2").Run()
}

func audioDownCallback(position int) error {
	sectionStatus[position] = "🔈 "

	return nil
}
