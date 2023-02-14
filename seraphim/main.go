package main

import (
	"fmt"
	"github.com/zcubbs/crucible/seraphim/pkg/gocotea"
	"log"
)

func main() {
	pbPath := "playbook.yaml"
	inv := "inventory"

	err := gocotea.InitPythonInterpretetor()
	if err != nil {
		log.Fatal(err)
	}

	var argMaker gocotea.ArgumentMaker

	err = argMaker.InitArgMaker()
	if err != nil {
		log.Fatal(err)
	}
	err = argMaker.AddArgument("-i", inv)
	if err != nil {
		log.Fatal(err)
	}

	var r gocotea.Runner

	err = r.InitRunner(&argMaker, pbPath)
	if err != nil {
		log.Fatal(err)
	}

	for r.HasNextPlay() {
		for r.HasNextTask() {
			fmt.Println("Next task name: ", r.GetNextTaskName())

			taskResults := r.RunNextTask()
			if len(taskResults) > 0 {
				fmt.Println("Task IsChanged:", taskResults[0].IsChanged)
			}
		}
	}

	r.FinishAnsibleWork()

	if r.WasError() {
		fmt.Printf("Ansible failed. Error:\n%s\n", r.GetErrorMsg())
	}

	err = gocotea.FinalizePythonInterpretetor()
	if err != nil {
		log.Fatal(err)
	}
}
