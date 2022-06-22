package main

import (
	"os/exec"
	"strings"
)

func musicEventCallback(position *int, button *int) (f callbackFunc, err error) {
	if *button == 0 {
		return nil, musicCallback(*position)
	}

	return nil, exec.Command("alacritty", "-e", "ncmpcpp").Start()
}

func musicCallback(position int) error {
	b, err := exec.Command("mpc", "volume").Output()
	if err != nil {
		return err
	}

	sectionStatus[position] = strings.Split(string(b), " ")[1]

	return nil
}

func musicUpEventCallback(position *int, button *int) (f callbackFunc, err error) {
	*position = *position - 1
	*button = 0

	return musicEventCallback, exec.Command("mpc", "volume", "+1").Run()
}

func musicUpCallback(position int) error {
	sectionStatus[position] = " 🎶"

	return nil
}

func musicDownEventCallback(position *int, button *int) (f callbackFunc, err error) {
	*position = *position + 1
	*button = 0

	return musicEventCallback, exec.Command("mpc", "volume", "-1").Run()
}

func musicDownCallback(position int) error {
	sectionStatus[position] = "🎵 "

	return nil
}
