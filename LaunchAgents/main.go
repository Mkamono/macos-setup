package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type launchctl struct{}

func (l launchctl) load(file string) error {
	return exec.Command("launchctl", "load", file).Run()
}

func (l launchctl) unload(file string) error {
	return exec.Command("launchctl", "unload", file).Run()
}

func main() {
	fmt.Println("Start")
	const repoDirName = "macos-setup"
	const modDirName = "LaunchAgents"
	const tmplFileName = "kamono.script.tmpl.plist"

	launchctl := launchctl{}
	if err := launchctl.unload("kamono.script.plist"); err != nil {
		panic(err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	tmplPath := filepath.Join(homeDir, repoDirName, modDirName, tmplFileName)
	if !isFile(tmplPath) {
		panic(fmt.Errorf("Template file not found: %s", tmplPath))
	}

	generateFileName := strings.Replace(tmplFileName, ".tmpl", "", 1)

	err = generatePlist(PlistInput{
		TEMPLATE_PATH: tmplPath,
		GENERATE_PATH: filepath.Dir(tmplPath) + "/" + generateFileName,
		TASKFILE_PATH: filepath.Join(homeDir, repoDirName, "Taskfile.yml"),
		LOG_DIR:       filepath.Join(homeDir, repoDirName, modDirName, "log"),
	})
	if err != nil {
		panic(err)
	}

	if err := launchctl.load(generateFileName); err != nil {
		panic(err)
	}
	fmt.Println("Loaded: " + generateFileName)
}

type PlistInput struct {
	TEMPLATE_PATH string
	GENERATE_PATH string
	TASKFILE_PATH string
	LOG_DIR       string
}

func generatePlist(input PlistInput) error {
	// テンプレートファイルが存在するか確認
	if !isFile(input.TEMPLATE_PATH) {
		return fmt.Errorf("Template file not found: %s", input.TEMPLATE_PATH)
	}

	// 既存のファイルがあれば削除
	if err := os.Remove(input.GENERATE_PATH); err != nil && !os.IsNotExist(err) {
		return err
	}

	// テンプレートファイルを読み込んで、生成する
	f, err := os.Create(input.GENERATE_PATH)
	defer f.Close()
	if err != nil {
		return err
	}
	tmpl, err := template.New(filepath.Base(input.TEMPLATE_PATH)).ParseFiles(input.TEMPLATE_PATH)
	if err != nil {
		return err
	}
	err = tmpl.Execute(f, input)
	if err != nil {
		return err
	}

	// 生成したファイルが正しいか確認
	if err := exec.Command("plutil", "-lint", input.GENERATE_PATH).Run(); err != nil {
		return err
	}
	fmt.Println("Generated: " + input.GENERATE_PATH)
	return nil
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}
