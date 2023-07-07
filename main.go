package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	mapToRegistredPeople := make(map[string]string)

	file, err := os.ReadFile("users")
	if err != nil {
		fmt.Println("Not found conf file, please create it")
	}
	var (
		idline      = 0
		idAfterline int
		idToName    = 0
	)
	for i := 0; i < len(string(file)); i++ {
		if string(string(file)[i]) == "|" {
			idline = i
		} else if string(string(file)[i]) == "\n" {
			name := string(file)[idToName:idline]
			idAfterline = i
			path := string(file)[idline+1 : idAfterline]
			mapToRegistredPeople[name] = path
			idAfterline = i
		} else if i == len(string(file))-1 {
			name := string(file)[idAfterline:idline]
			path := string(file)[idline+1 : i]
			mapToRegistredPeople[name] = path
		}
	}

	for {
		for _, path := range mapToRegistredPeople {
			temp := path
			cmd := exec.Command("rsync", "-avz", "/home/adminusertest/test/*.txt", temp)
			out := cmd.Run()
			fmt.Println(out)
		}
		time.Sleep(time.Hour)
	}
	//path admintest@192.168.0.23:/root/test
}
