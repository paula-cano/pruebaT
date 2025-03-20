package application

import (
	"Backend/models"
	"math"
	"strconv"
	"strings"
)

var topBrokers = map[string]bool{
	"Morgan Stanley":          true,
	"The Goldman Sachs Group": true,
	"Wells Fargo & Company":   true,
}

// Puntajes asignados para cada cambio de calificación
var ratingChanges = map[string]int{
	"Sell->Neutral":            3,
	"Sell->Buy":                5,
	"Neutral->Buy":             3,
	"Underweight->Neutral":     2,
	"Underweight->Overweight":  4,
	"Equal Weight->Overweight": 3,
	"Buy->Neutral":             -3,
	"Buy->Sell":                -5,
	"Neutral->Sell":            -3,
	"Overweight->Underweight":  -4,
	"Overweight->Equal Weight": -3,
}

func parsePrice(priceStr string) float64 {
	priceStr = strings.Replace(priceStr, "$", "", -1)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return math.NaN() // para evitar errores en la lógica
	}
	return price
}

// calculateScore asigna un puntaje a cada acción en base a cambios de precio, calificación y broker
func CalculateScore(stock *models.Item) {
	stock.Score = 0

	priceFrom := parsePrice(stock.TargetFrom)
	priceTo := parsePrice(stock.TargetTo)

	if !math.IsNaN(priceFrom) && !math.IsNaN(priceTo) {
		if priceTo > priceFrom {
			stock.Score += 3
		} else if priceTo < priceFrom {
			stock.Score -= 3
		}
	}

	if stock.RatingFrom != "" && stock.RatingTo != "" {
		key := stock.RatingFrom + "->" + stock.RatingTo
		if val, exists := ratingChanges[key]; exists {
			stock.Score += float64(val)
		}
	}

	if _, exists := topBrokers[stock.Brokerage]; exists {
		stock.Score += 2
	}
}

func RecommendBestStock(stocks []models.Item) *models.Item {
	if len(stocks) == 0 {
		return nil
	}

	bestStock := &stocks[0]
	for i := 1; i < len(stocks); i++ {
		if stocks[i].Score > bestStock.Score {
			bestStock = &stocks[i]
		}
	}
	return bestStock
}
