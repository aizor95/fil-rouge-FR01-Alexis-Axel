package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8081/v1/forecast?latitude=48.89&longitude=2.20&current=temperature_2m,wind_speed_10m"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture :", err)
		return
	}
	fmt.Println(string(body))
}
