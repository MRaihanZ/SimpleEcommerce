package product

import (
	"SimpleEcommerce/database"
	"SimpleEcommerce/database/entities"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetCarts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := context.Background()

	rows, err := database.DB.QueryContext(ctx, "SELECT * FROM cart")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var responses []entities.CartResponse

	for rows.Next() {
		var cart entities.CartResponse
		err := rows.Scan(&cart.Id, &cart.IdProduct, &cart.Title, &cart.Price, &cart.Image, &cart.Rating, &cart.Quantity)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		responses = append(responses, cart)
	}

	responseJSON, err := json.Marshal(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJSON)
}

func GetCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	ctx := context.Background()
	var responses []entities.CartResponse
	var cart entities.CartResponse
	row := database.DB.QueryRowContext(ctx, `SELECT * FROM cart WHERE id_product = ?`, id)
	err := row.Scan(&cart.Id, &cart.IdProduct, &cart.Title, &cart.Price, &cart.Image, &cart.Rating, &cart.Quantity)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	responses = append(responses, cart)

	responseJSON, err := json.Marshal(responses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the status code to 200 OK
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJSON)
}

func PostCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// Handle the main request logic
	var cart entities.CartResponse
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&cart); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	_, err := database.DB.Exec(`
        INSERT INTO cart(id, id_product, title, price, image, rating, quantity) 
        VALUES (?, ?, ?, ?, ?, ?, ?)`, "", cart.IdProduct, cart.Title, cart.Price, cart.Image, cart.Rating, cart.Quantity)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var response entities.Response
	response.Code = 200
	response.Messege = "Success"
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the status code to 200 OK
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJSON)
}

func PatchCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	quantityStr := params.ByName("quantity")

	// Parse quantity parameter to integer
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(writer, "Invalid quantity", http.StatusBadRequest)
		return
	}

	quantity++

	var patch entities.Response

	_, err = database.DB.Exec(`UPDATE cart SET quantity = ? WHERE id_product = ?`, quantity, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	patch.Code = 200
	patch.Messege = "Success"
	responseJSON, err := json.Marshal(patch)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the status code to 200 OK
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJSON)
}


func DeleteCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	var del entities.Response

	_, err := database.DB.Exec(`DELETE FROM cart WHERE id_product = ?`, id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	del.Code = 200
	del.Messege = "Success"
	responseJSON, err := json.Marshal(del)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the status code to 200 OK
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(responseJSON)
}