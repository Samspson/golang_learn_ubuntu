package main

import (
	"golang_learn_ubuntu/go_practice/ch7/runner"
	"log"
	"os"
	"time"
)

// timeout规定了必须在多少秒内处理完成
const timeout = 3 * time.Second

// main是程序的入口
func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

// createTask
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
