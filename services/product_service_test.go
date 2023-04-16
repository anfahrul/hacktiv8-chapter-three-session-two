package services

import (
	"errors"
	model "hacktiv8-chapter-three-session-two/models"
	"hacktiv8-chapter-three-session-two/repository/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestProductService_GetProductByID(t *testing.T) {

	productRepo := mocks.NewProductRepository(t)

	type args struct {
		productID string
		userID    string
		role      bool
	}
	tests := []struct {
		name      string
		service   *ProductService
		args      args
		want      model.ProductResponse
		wantErrIs string
		mockFunc  func()
		wantErr   bool
	}{
		{
			name: "Case #1 get one product found",
			service: &ProductService{
				productRepository: productRepo,
			},
			args: args{
				productID: "dbf5db3a-6f4a-4a11-8d6b-24c62bfe23a1",
				userID:    "2a80b842-0aa1-4e20-a998-32654131997a",
				role:      true,
			},
			want: model.ProductResponse{
				ProductID:   "dbf5db3a-6f4a-4a11-8d6b-24c62bfe23a1",
				Title:       "Product B",
				Description: "Product B Description by fahrul2",
				UserID:      "2a80b842-0aa1-4e20-a998-32654131997a",
				CreatedAt:   "2023-04-15 17:12:27",
				UpdatedAt:   "2023-04-15 17:12:27",
			},
			mockFunc: func() {
				createdAtFormat, _ := time.Parse("2006-01-02 15:04:05", "2023-04-15 17:12:27")
				updatedAtFormat, _ := time.Parse("2006-01-02 15:04:05", "2023-04-15 17:12:27")

				productRepo.On("FindProduct", "dbf5db3a-6f4a-4a11-8d6b-24c62bfe23a1").Return(&model.Product{
					ProductID:   uuid.MustParse("dbf5db3a-6f4a-4a11-8d6b-24c62bfe23a1"),
					Title:       "Product B",
					Description: "Product B Description by fahrul2",
					UserID:      uuid.MustParse("2a80b842-0aa1-4e20-a998-32654131997a"),
					CreatedAt:   createdAtFormat,
					UpdatedAt:   updatedAtFormat,
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "Case #2 get one product not found",
			service: &ProductService{
				productRepository: productRepo,
			},
			args: args{
				productID: "dbf5db3a-6f4a-4a11-8d6b-000000000000",
				userID:    "2a80b842-0aa1-4e20-a998-32654131997a",
				role:      true,
			},
			want:      model.ProductResponse{},
			wantErrIs: "record not found",
			mockFunc: func() {
				productRepo.On("FindProduct", "dbf5db3a-6f4a-4a11-8d6b-000000000000").Return(&model.Product{}, errors.New("record not found"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := tt.service.GetProductByID(tt.args.productID, tt.args.userID, tt.args.role)
			if err != nil && err.Error() != tt.wantErrIs {
				t.Errorf("ProductService.GetProductByID() error = %v, wantErrIs %v", err.Error(), tt.wantErrIs)
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.GetProductByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_GetProductByUserID(t *testing.T) {

	productRepo := mocks.NewProductRepository(t)

	type args struct {
		userId string
	}
	tests := []struct {
		name      string
		service   *ProductService
		args      args
		want      []model.ProductResponse
		wantErrIs string
		mockFunc  func()
		wantErr   bool
	}{
		{
			name: "Case #3 get all product found",
			service: &ProductService{
				productRepository: productRepo,
			},
			want: []model.ProductResponse{
				{
					ProductID:   "25eb23dd-b196-4e38-99ad-87e186009715",
					Title:       "Product A3",
					Description: "Product A3 Description by fahrul",
					UserID:      "e43d4d44-54e4-41f0-91cc-f618ad3f5cb5",
					CreatedAt:   "2023-04-16 12:22:29",
					UpdatedAt:   "2023-04-16 12:22:29",
				},
			},
			args: args{
				userId: "e43d4d44-54e4-41f0-91cc-f618ad3f5cb5",
			},
			mockFunc: func() {
				createdAtFormat, _ := time.Parse("2006-01-02 15:04:05", "2023-04-16 12:22:29")
				updatedAtFormat, _ := time.Parse("2006-01-02 15:04:05", "2023-04-16 12:22:29")

				productRepo.On("GetByUserID", "e43d4d44-54e4-41f0-91cc-f618ad3f5cb5").Return([]model.Product{
					{
						ProductID:   uuid.MustParse("25eb23dd-b196-4e38-99ad-87e186009715"),
						Title:       "Product A3",
						Description: "Product A3 Description by fahrul",
						UserID:      uuid.MustParse("e43d4d44-54e4-41f0-91cc-f618ad3f5cb5"),
						CreatedAt:   createdAtFormat,
						UpdatedAt:   updatedAtFormat,
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "Case #4 get all product not found",
			service: &ProductService{
				productRepository: productRepo,
			},
			want:      []model.ProductResponse{},
			wantErrIs: "Not Found",
			args: args{
				userId: "d3e53f12-1ce3-4119-a46b-6658c5e9655f",
			},
			mockFunc: func() {
				productRepo.On("GetByUserID", "d3e53f12-1ce3-4119-a46b-6658c5e9655f").Return([]model.Product{}, errors.New("Not Found"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := tt.service.GetProductByUserID(tt.args.userId)
			if err != nil && err.Error() != tt.wantErrIs {
				t.Errorf("ProductService.GetProductByUserID() error = %v, wantErrIs %v", err.Error(), tt.wantErrIs)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetProductByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.GetProductByUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
