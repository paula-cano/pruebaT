package request

import (
	"Backend/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Request() (*models.Production, error) {
	url := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	apiKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MSwiZW1haWwiOiJwYXVsYS4xNGNhbm9iQGdtYWlsLmNvbSIsImV4cCI6MTc0MTkyNDc1OSwiaWQiOiIwIiwicGFzc3dvcmQiOiInIE9SICcxJz0nMScgQU5EICdhJz0nYSJ9.Mb7PsGite_ClVT7J8v5tOq4c1NB5LqzlEuFryx5nGZM"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error en la solicitud: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Error en la solicitud HTTP:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body) // Leer respuesta
	if err != nil {
		return nil, fmt.Errorf("error leyendo la respuesta: %v", err)
	}

	var production models.Production // se mapea el JSON a la estructura del modelo
	if err := json.Unmarshal(body, &production); err != nil {
		return nil, fmt.Errorf("error al decodificar JSON: %v", err)
	}
	
	return &production, nil
}
