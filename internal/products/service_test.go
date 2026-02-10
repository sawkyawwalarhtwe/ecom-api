package products

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/sawkyawwalarhtwe/ecom-api/internal/adapters/postgresql/sqlc"
)

// MockQuerier is a mock implementation of repo.Querier for testing
type MockQuerier struct {
	ListProductsFunc          func(ctx context.Context) ([]repo.Product, error)
	FindProductByIdFunc       func(ctx context.Context, id int64) (repo.Product, error)
	CreateOrderFunc           func(ctx context.Context, customerID int64) (repo.Order, error)
	CreateOrderItemFunc       func(ctx context.Context, arg repo.CreateOrderItemParams) (repo.OrderItem, error)
	UpdateProductQuantityFunc func(ctx context.Context, arg repo.UpdateProductQuantityParams) error
}

func (m *MockQuerier) ListProducts(ctx context.Context) ([]repo.Product, error) {
	if m.ListProductsFunc != nil {
		return m.ListProductsFunc(ctx)
	}
	return nil, nil
}

func (m *MockQuerier) FindProductById(ctx context.Context, id int64) (repo.Product, error) {
	if m.FindProductByIdFunc != nil {
		return m.FindProductByIdFunc(ctx, id)
	}
	return repo.Product{}, errors.New("not implemented")
}

func (m *MockQuerier) CreateOrder(ctx context.Context, customerID int64) (repo.Order, error) {
	if m.CreateOrderFunc != nil {
		return m.CreateOrderFunc(ctx, customerID)
	}
	return repo.Order{}, errors.New("not implemented")
}

func (m *MockQuerier) CreateOrderItem(ctx context.Context, arg repo.CreateOrderItemParams) (repo.OrderItem, error) {
	if m.CreateOrderItemFunc != nil {
		return m.CreateOrderItemFunc(ctx, arg)
	}
	return repo.OrderItem{}, errors.New("not implemented")
}

func (m *MockQuerier) UpdateProductQuantity(ctx context.Context, arg repo.UpdateProductQuantityParams) error {
	if m.UpdateProductQuantityFunc != nil {
		return m.UpdateProductQuantityFunc(ctx, arg)
	}
	return nil
}

func TestListProducts_Success(t *testing.T) {
	ts := pgtype.Timestamptz{}
	if err := ts.Scan("2026-02-11T10:00:00Z"); err != nil {
		t.Fatalf("failed to scan timestamp: %v", err)
	}

	expectedProducts := []repo.Product{
		{
			ID:             1,
			Name:           "Laptop",
			PriceInCenters: 99999,
			Quantity:       10,
			CreatedAt:      ts,
		},
		{
			ID:             2,
			Name:           "Mouse",
			PriceInCenters: 2999,
			Quantity:       50,
			CreatedAt:      ts,
		},
	}

	mock := &MockQuerier{
		ListProductsFunc: func(ctx context.Context) ([]repo.Product, error) {
			return expectedProducts, nil
		},
	}

	svc := NewService(mock)
	products, err := svc.ListProducts(context.Background())

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(products) != 2 {
		t.Errorf("expected 2 products, got %d", len(products))
	}

	if products[0].Name != "Laptop" {
		t.Errorf("expected first product name 'Laptop', got %s", products[0].Name)
	}

	if products[1].ID != 2 {
		t.Errorf("expected second product ID 2, got %d", products[1].ID)
	}
}

func TestListProducts_Empty(t *testing.T) {
	mock := &MockQuerier{
		ListProductsFunc: func(ctx context.Context) ([]repo.Product, error) {
			return []repo.Product{}, nil
		},
	}

	svc := NewService(mock)
	products, err := svc.ListProducts(context.Background())

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestListProducts_DatabaseError(t *testing.T) {
	expectedErr := errors.New("database connection failed")
	mock := &MockQuerier{
		ListProductsFunc: func(ctx context.Context) ([]repo.Product, error) {
			return nil, expectedErr
		},
	}

	svc := NewService(mock)
	products, err := svc.ListProducts(context.Background())

	if err == nil {
		t.Error("expected error, got nil")
	}

	if err.Error() != "database connection failed" {
		t.Errorf("expected error message 'database connection failed', got %s", err.Error())
	}

	if products != nil {
		t.Errorf("expected nil products on error, got %v", products)
	}
}

func TestListProducts_ContextCancellation(t *testing.T) {
	mock := &MockQuerier{
		ListProductsFunc: func(ctx context.Context) ([]repo.Product, error) {
			return nil, context.Canceled
		},
	}

	svc := NewService(mock)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	products, err := svc.ListProducts(ctx)

	if err != context.Canceled {
		t.Errorf("expected context.Canceled error, got %v", err)
	}

	if products != nil {
		t.Errorf("expected nil products on context cancellation, got %v", products)
	}
}
