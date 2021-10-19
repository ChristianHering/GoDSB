package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func updateEventCallback(position *int, button *int) (f callbackFunc, err error) {
	if *button == 1 {
		return nil, exec.Command("alacritty", "-e", "paru", "-Syyuu").Run() //TODO: Remove hardcoded terminal
	} else if *button == 3 {
		return nil, exec.Command("alacritty", "-e", "sudo", "pacman", "-Syyuu").Run()
	}

	return nil, nil
}

func updateCallback(position int) error {
	b, err := exec.Command("paru", "-Qu").Output()
	if err != nil {
		return err
	}

	updates := strconv.Itoa(len(strings.Split(string(b), "\n")))

	if updates == "0" {
		updates = "No"
	}

	sectionStatus[position] = updates + " Updates 📦"

	return nil
}
