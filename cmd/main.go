package main

import (
	"sync"
	"weeklytickits/services"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		services.FetchAndSaveActor()
	}()

	go func() {
		defer wg.Done()
		services.FetchAndSaveDirector()
	}()

	go func() {
		defer wg.Done()
		services.FetchAndSaveGenres()
	}()

	go func() {
		defer wg.Done()
		services.FetchMovie()
	}()

	wg.Wait()
}
