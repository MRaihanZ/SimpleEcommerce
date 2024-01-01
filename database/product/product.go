package product

import (
	"SimpleEcommerce/database"
	"SimpleEcommerce/database/entities"
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := context.Background()

	rows, err := database.DB.QueryContext(ctx, "SELECT * FROM products")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var responses []entities.ProductResponse

	for rows.Next() {
		var product entities.ProductResponse
		product.Code = "200"
		err := rows.Scan(&product.Data.ID, &product.Data.Title, &product.Data.Price, &product.Data.Description, &product.Data.Category, &product.Data.Image, &product.Data.Rating, &product.Data.TotalRating)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		responses = append(responses, product)
	}

	responseJSON, err := json.Marshal(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJSON)
}
