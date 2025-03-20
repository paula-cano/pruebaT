package connection

import (
	"Backend/models"
	"context"
	"log"
)

func SaveItems(item []models.Item) error {
	ctx := context.Background()

	query := `
		INSERT INTO item (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (ticker, time) DO NOTHING;
	`

	for _, item := range item {
		_, err := DB.Exec(ctx, query, item.Ticker, item.Company, item.Brokerage, item.Action, item.RatingFrom, item.RatingTo, item.TargetFrom, item.TargetTo, item.Time)
		if err != nil {
			log.Printf("Error al guardar el item %s: %v", item.Ticker, err)
		}
	}

	return nil
}
