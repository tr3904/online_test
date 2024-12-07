package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "product-management-backend/models"
    "github.com/gorilla/mux"
)

var db *sql.DB 


func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, "Invalid JSON input", http.StatusBadRequest)
        return
    }

    
    if product.UserID == 0 || product.ProductName == "" || len(product.ProductImages) == 0 || product.ProductPrice <= 0 {
        http.Error(w, "Missing or invalid fields", http.StatusBadRequest)
        return
    }

    
    compressedImages := processImages(product.ProductImages)

    query := `
        INSERT INTO products (user_id, product_name, product_description, product_images, compressed_product_images, product_price)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id`
    err := db.QueryRow(query, product.UserID, product.ProductName, product.ProductDescription, 
        product.ProductImages, compressedImages, product.ProductPrice).Scan(&product.ProductID)
    if err != nil {
        http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully", "product_id": strconv.Itoa(product.ProductID)})
}


func GetProductByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID format", http.StatusBadRequest)
        return
    }

    var product models.Product
    query := `SELECT product_id, user_id, product_name, product_description, product_images, 
              compressed_product_images, product_price, created_at FROM products WHERE product_id = $1`
    row := db.QueryRow(query, id)
    if err := row.Scan(&product.ProductID, &product.UserID, &product.ProductName, &product.ProductDescription, 
        &product.ProductImages, &product.CompressedProductImages, &product.ProductPrice, &product.CreatedAt); err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Product not found", http.StatusNotFound)
        } else {
            http.Error(w, "Error fetching product: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(product)
}


func GetProducts(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("user_id")
    minPrice := r.URL.Query().Get("min_price")
    maxPrice := r.URL.Query().Get("max_price")
    productName := r.URL.Query().Get("product_name")

    query := `SELECT product_id, user_id, product_name, product_description, product_images, 
              compressed_product_images, product_price, created_at FROM products WHERE TRUE`
    args := []interface{}{}

    if userID != "" {
        query += ` AND user_id = $1`
        args = append(args, userID)
    }
    if minPrice != "" && maxPrice != "" {
        query += ` AND product_price BETWEEN $2 AND $3`
        args = append(args, minPrice, maxPrice)
    }
    if productName != "" {
        query += ` AND product_name ILIKE $4`
        args = append(args, "%"+productName+"%")
    }

    rows, err := db.Query(query, args...)
    if err != nil {
        http.Error(w, "Failed to retrieve products: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.ProductID, &product.UserID, &product.ProductName, &product.ProductDescription, 
            &product.ProductImages, &product.CompressedProductImages, &product.ProductPrice, &product.CreatedAt); err != nil {
            http.Error(w, "Failed to parse products: "+err.Error(), http.StatusInternalServerError)
            return
        }
        products = append(products, product)
    }

    json.NewEncoder(w).Encode(products)
}


func processImages(images []string) []string {
    var compressedImages []string
    for _, img := range images {
        compressedImages = append(compressedImages, img+"-compressed") // Simulate compression
    }
    return compressedImages
}


