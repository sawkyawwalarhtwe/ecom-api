package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/sawkyawwalarhtwe/ecom-api/internal/adapters/postgresql/sqlc"
	"github.com/sawkyawwalarhtwe/ecom-api/internal/orders"
	"github.com/sawkyawwalarhtwe/ecom-api/internal/products"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	// OpenAPI documentation endpoint
	r.Get("/docs/openapi.json", serveOpenAPI)
	r.Get("/docs", serveDocsHTML)

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	orderService := orders.NewService(repo.New(app.db), app.db)
	ordersHandler := orders.NewHandler(orderService)
	r.Post("/orders", ordersHandler.PlaceOrder)

	return r
}

// serveOpenAPI serves the OpenAPI specification
func serveOpenAPI(w http.ResponseWriter, r *http.Request) {
	openAPISpec := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "E-Commerce API",
			"description": "A simple e-commerce API for managing products and orders",
			"version":     "1.0.0",
		},
		"servers": []map[string]string{
			{
				"url":         "http://localhost:8080",
				"description": "Development server",
			},
		},
		"paths": map[string]interface{}{
			"/health": map[string]interface{}{
				"get": map[string]interface{}{
					"summary": "Health Check",
					"tags":    []string{"Health"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "Server is healthy",
						},
					},
				},
			},
			"/products": map[string]interface{}{
				"get": map[string]interface{}{
					"summary":     "List all products",
					"description": "Retrieve all available products",
					"tags":        []string{"Products"},
					"responses": map[string]interface{}{
						"200": map[string]interface{}{
							"description": "List of products",
						},
					},
				},
			},
			"/orders": map[string]interface{}{
				"post": map[string]interface{}{
					"summary":     "Create a new order",
					"description": "Place an order with products",
					"tags":        []string{"Orders"},
					"responses": map[string]interface{}{
						"201": map[string]interface{}{
							"description": "Order created successfully",
						},
						"400": map[string]interface{}{
							"description": "Invalid request",
						},
					},
				},
			},
		},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{
				"Product": map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"id":               map[string]interface{}{"type": "integer"},
						"name":             map[string]interface{}{"type": "string"},
						"price_in_centers": map[string]interface{}{"type": "integer"},
						"quantity":         map[string]interface{}{"type": "integer"},
						"created_at":       map[string]interface{}{"type": "string"},
					},
				},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(openAPISpec)
}

// serveDocsHTML serves a simple Swagger UI-like HTML page
func serveDocsHTML(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>E-Commerce API - OpenAPI Documentation</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        h1 { color: #333; }
        .endpoint { margin-top: 20px; padding: 10px; background: #f5f5f5; border-left: 4px solid #0066cc; }
        .get { border-left-color: #61affe; }
        .post { border-left-color: #49cc90; }
        .method { font-weight: bold; padding: 3px 8px; border-radius: 3px; display: inline-block; width: 50px; text-align: center; }
        .get .method { background: #61affe; color: white; }
        .post .method { background: #49cc90; color: white; }
        code { background: #f0f0f0; padding: 2px 6px; border-radius: 3px; }
        .json { background: white; border: 1px solid #ddd; padding: 10px; border-radius: 3px; overflow-x: auto; }
        pre { margin: 0; }
    </style>
</head>
<body>
    <h1>E-Commerce API Documentation</h1>
    
    <h2>Available Endpoints</h2>
    
    <div class="endpoint get">
        <span class="method get">GET</span> <code>/health</code>
        <p>Health check endpoint. Returns "all good" if the server is running.</p>
    </div>

    <div class="endpoint get">
        <span class="method get">GET</span> <code>/products</code>
        <p>List all available products with pricing and stock information.</p>
        <h4>Response (200):</h4>
        <div class="json">
            <pre>[
  {
    "id": 1,
    "name": "Laptop",
    "price_in_centers": 99999,
    "quantity": 10,
    "created_at": "2026-02-11T10:00:00Z"
  }
]</pre>
        </div>
    </div>

    <div class="endpoint post">
        <span class="method post">POST</span> <code>/orders</code>
        <p>Create a new order with items. Validates stock and creates order items within a transaction.</p>
        <h4>Request Body:</h4>
        <div class="json">
            <pre>{
  "customer_id": 1,
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    }
  ]
}</pre>
        </div>
        <h4>Response (201):</h4>
        <div class="json">
            <pre>{
  "id": 1,
  "customer_id": 1,
  "created_at": "2026-02-11T10:30:00Z"
}</pre>
        </div>
        <h4>Error Responses:</h4>
        <ul>
            <li><strong>400:</strong> Missing or invalid parameters (customer_id, items)</li>
            <li><strong>404:</strong> Product not found</li>
            <li><strong>409:</strong> Insufficient product stock</li>
            <li><strong>500:</strong> Internal server error</li>
        </ul>
    </div>

    <h2>API Specification</h2>
    <p>Full OpenAPI specification available at: <code><a href="/docs/openapi.json">/docs/openapi.json</a></code></p>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}

type application struct {
	config config
	// logger
	db *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
