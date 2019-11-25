package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"sync"
	"time"
)

func doLS(path string) ([]byte, error) {
	cmd := exec.Command("ls", "-l", path)
	return cmd.CombinedOutput()
}

func main() {

	outputs := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		out, _ := doLS("/usr")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		outputs <- string(out)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		out, _ := doLS("/")
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		outputs <- string(out)
		wg.Done()
	}()

	go func() {
		for m := range outputs {
			fmt.Println(m)
		}
	}()

	wg.Wait()
	close(outputs)

}
