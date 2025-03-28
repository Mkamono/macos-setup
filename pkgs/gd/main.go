package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
)

func main() {
	branch, err := defaultBranch()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "checkout", branch)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("git", "fetch", "--prune")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("git", "pull", "--all")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	err = deleteAllLocalBranches(branch)
	if err != nil {
		panic(err)
	}
}

func defaultBranch() (string, error) {
	cmd := exec.Command("git", "remote", "show", "origin")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "HEAD branch") {
			fields := strings.Fields(line)
			return fields[len(fields)-1], nil
		}
	}

	return "", fmt.Errorf("HEAD branch not found in git remote output")
}

func deleteAllLocalBranches(defaultBranch string) error {
	cmd := exec.Command("git", "branch", "--list", "--format=%(refname:short)")
	out, err := cmd.Output()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if line == defaultBranch {
			continue
		}

		cmd := exec.Command("git", "branch", "-D", line)
		err := cmd.Run()
		if err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	return nil
}
