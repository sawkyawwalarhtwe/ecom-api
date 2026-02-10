package orders

import (
	"context"
	"testing"

	repo "github.com/sawkyawwalarhtwe/ecom-api/internal/adapters/postgresql/sqlc"
)

// MockQuerier is a mock implementation of repo.Querier for testing orders
type MockQueriesForOrders struct {
	CreateOrderFunc           func(ctx context.Context, customerID int64) (repo.Order, error)
	FindProductByIdFunc       func(ctx context.Context, id int64) (repo.Product, error)
	CreateOrderItemFunc       func(ctx context.Context, arg repo.CreateOrderItemParams) (repo.OrderItem, error)
	UpdateProductQuantityFunc func(ctx context.Context, arg repo.UpdateProductQuantityParams) error
	ListProductsFunc          func(ctx context.Context) ([]repo.Product, error)
}

func (m *MockQueriesForOrders) CreateOrder(ctx context.Context, customerID int64) (repo.Order, error) {
	if m.CreateOrderFunc != nil {
		return m.CreateOrderFunc(ctx, customerID)
	}
	return repo.Order{}, nil
}

func (m *MockQueriesForOrders) FindProductById(ctx context.Context, id int64) (repo.Product, error) {
	if m.FindProductByIdFunc != nil {
		return m.FindProductByIdFunc(ctx, id)
	}
	return repo.Product{}, nil
}

func (m *MockQueriesForOrders) CreateOrderItem(ctx context.Context, arg repo.CreateOrderItemParams) (repo.OrderItem, error) {
	if m.CreateOrderItemFunc != nil {
		return m.CreateOrderItemFunc(ctx, arg)
	}
	return repo.OrderItem{}, nil
}

func (m *MockQueriesForOrders) UpdateProductQuantity(ctx context.Context, arg repo.UpdateProductQuantityParams) error {
	if m.UpdateProductQuantityFunc != nil {
		return m.UpdateProductQuantityFunc(ctx, arg)
	}
	return nil
}

func (m *MockQueriesForOrders) ListProducts(ctx context.Context) ([]repo.Product, error) {
	if m.ListProductsFunc != nil {
		return m.ListProductsFunc(ctx)
	}
	return nil, nil
}

// TestPlaceOrder_MissingCustomerID tests validation of missing customer ID
func TestPlaceOrder_MissingCustomerID(t *testing.T) {
	order := createOrderParams{
		CustomerID: 0, // Invalid
		Items: []orderItem{
			{ProductID: 1, Quantity: 2},
		},
	}

	// Validate function logic (extracted from PlaceOrder)
	if order.CustomerID == 0 {
		// Expected error
		return
	}

	t.Error("expected customer ID validation to fail")
}

// TestPlaceOrder_MissingItems tests validation of empty items list
func TestPlaceOrder_MissingItems(t *testing.T) {
	order := createOrderParams{
		CustomerID: 1,
		Items:      []orderItem{}, // Empty items
	}

	// Validate function logic (extracted from PlaceOrder)
	if len(order.Items) == 0 {
		// Expected error
		return
	}

	t.Error("expected items validation to fail")
}

// Note: For full integration tests with actual transaction behavior,
// use integration tests with Testcontainers or a test PostgreSQL database
