package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//canal := make(chan int) //Se crea un canal
	//canal <- 15             //envia datos del canal
	//valor := <-canal        // recibe datos del canal

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkApi(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Wena, solo tomo %v segundos \n", elapsed.Seconds())
}

func checkApi(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		// fmt.Printf("Error: %s esta caido todo \n", api)
		ch <- fmt.Sprintf("Error: %s esta caido todo \n", api)
		return
	}
	ch <- fmt.Sprintf("Wena: %s ta funcionando \n", api)

}
