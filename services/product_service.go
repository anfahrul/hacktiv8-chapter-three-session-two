package services

import (
	"errors"
	model "hacktiv8-chapter-three-session-two/models"
	"hacktiv8-chapter-three-session-two/repository"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepository repository.ProductRepository
	userRepository    repository.UserRepository
}

func NewProductService(productRepository repository.ProductRepository, userRepository repository.UserRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		userRepository:    userRepository,
	}
}

func (service *ProductService) CreateProduct(request model.ProductCreateRequest, userId string) (model.ProductCreateResponse, error) {
	productID := uuid.New()

	userRes, err := service.userRepository.UserCheck(userId)
	if err != nil {
		return model.ProductCreateResponse{}, err
	}

	product := model.Product{
		ProductID:   productID,
		Title:       request.Title,
		Description: request.Description,
		UserID:      userRes.UserID,
	}

	result, err := service.productRepository.CreateProduct(product)
	if err != nil {
		return model.ProductCreateResponse{}, err
	}

	response := model.ProductCreateResponse{
		ProductID:   result.ProductID,
		Title:       result.Title,
		Description: result.Description,
		UserID:      result.UserID,
	}

	return response, nil
}

func (service *ProductService) GetProductByUserID(userId string) ([]model.ProductResponse, error) {
	userRes, err := service.userRepository.UserCheck(userId)
	if err != nil {
		return []model.ProductResponse{}, err
	}

	response, err := service.productRepository.GetByUserID(userRes.UserID)
	if err != nil {
		return []model.ProductResponse{}, err
	}

	products := []model.ProductResponse{}
	for _, product := range response {
		product := model.ProductResponse{
			ProductID:   product.ProductID,
			Title:       product.Title,
			Description: product.Description,
			UserID:      product.UserID,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		}
		products = append(products, product)
	}

	return products, nil
}

func (service *ProductService) GetAllProduct() ([]model.ProductResponse, error) {
	response, err := service.productRepository.GetAllProduct()
	if err != nil {
		return []model.ProductResponse{}, err
	}

	products := []model.ProductResponse{}
	for _, product := range response {
		product := model.ProductResponse{
			ProductID:   product.ProductID,
			Title:       product.Title,
			Description: product.Description,
			UserID:      product.UserID,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		}
		products = append(products, product)
	}

	return products, nil
}

func (service *ProductService) DeleteProduct(productID string) error {
	result, err := service.productRepository.FindProduct(productID)
	if err != nil {
		return err
	}

	err = service.productRepository.DeleteProduct(*result)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProductService) UpdatedProduct(productID string, request model.ProductUpdateRequest, userId string) (model.ProductResponse, error) {
	resultProduct, err := service.productRepository.FindProduct(productID)
	if err != nil {
		return model.ProductResponse{}, err
	}

	resultUser, err := service.userRepository.UserCheck(userId)
	if err != nil {
		return model.ProductResponse{}, err
	}

	updatedProductReq := &model.Product{
		ProductID:   resultProduct.ProductID,
		Title:       request.Title,
		Description: request.Description,
	}

	var updateResult model.Product
	if resultUser.Role == true {
		updateResult, err = service.productRepository.UpdateProduct(*updatedProductReq)
		if err != nil {
			return model.ProductResponse{}, err
		}
	} else {
		if resultUser.UserID == resultProduct.UserID {
			updateResult, err = service.productRepository.UpdateProduct(*updatedProductReq)
			if err != nil {
				return model.ProductResponse{}, err
			}
		} else {
			return model.ProductResponse{}, errors.New("Unauthorized")
		}
	}

	response := model.ProductResponse{
		ProductID:   updateResult.ProductID,
		Title:       updateResult.Title,
		Description: updateResult.Description,
		UserID:      resultUser.UserID,
		CreatedAt:   updateResult.CreatedAt.String(),
		UpdatedAt:   updateResult.UpdatedAt.String(),
	}

	return response, nil
}
