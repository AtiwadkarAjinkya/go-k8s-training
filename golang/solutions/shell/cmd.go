package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func execCommand(cmd string) bool{
	cmdString := strings.Fields(cmd)
	if len(cmdString) == 0 {
		return false
	}
	switch cmdString[0] {
	case "ls":
		execListCommand(cmdString[1:])
	case "pwd":
		getCurrentDirCommand()
	case "cd":
		changeDirCommand(cmdString[1:])
	case "exit":
		return true
	}
	return false
}

func execListCommand(args []string) {
	if len(args) == 0 {
		currentWorkingDir, _ := os.Getwd()
		files, err := os.ReadDir(currentWorkingDir)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	} else {
		for _, path := range args {
			_, err := os.Stat(path)
			if err != nil {
				log.Fatal(err.Error())
			} else {
				files, _ := os.ReadDir(path)
				if err != nil {
					log.Fatal(err)
				}	
				for _, file := range files {
					fmt.Println(file.Name())
				}
			}
		}
	}
}

func getCurrentDirCommand() {
	currentDir, err := os.Getwd()
	if err != nil{
		log.Default()
	} else {
		fmt.Println(currentDir)
	}
}

func changeDirCommand(args []string) {
	if len(args) != 1 {
		log.Fatal("Invalid Arguments")
	} else {
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Println(err)
		}
		currentDir, _ := os.Getwd()
		fmt.Println("Current working directory is: ",currentDir)
	}
}