package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
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

type Meteo struct {
	Current map[string]interface{} `json:"current"`
	Time string `json:"code"`
	Temperature float64 `json:"places"`
	Vent float64 `json:"projecteur"`
}

func decodeMeteo(resp *http.Response) error {
	var m Meteo
	err := json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return fmt.Errorf("JSON illisible : %w", err)
	}
	fmt.Println("Météo du Campus à",m.Time, " :", m.Temperature, "°C, vent :", m.Vent, "km/h")
	return err
}

func recupererMeteo(baseURL string, lat, lon float64) (Meteo, error) {
	url := fmt.Sprintf("%s/v1/forecast?latitude=%g&longitude=%g&current=temperature_2m,wind_speed_10m", baseURL, lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête :", err)
		return Meteo{}, fmt.Errorf("Mauvais URL : %d", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture :", err)
		return Meteo{}, fmt.Errorf("Erreur lors du décodage : %d", err)
	}
	fmt.Println(string(body))
	
	
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur trouvée :", resp.StatusCode)
		return Meteo{}, fmt.Errorf("L'API a renvoyé : %d", resp.StatusCode)
	}
	return Meteo{}, nil
}