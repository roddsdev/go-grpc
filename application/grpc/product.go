package grpc

import (
	"context"
	"example/grpc/application/grpc/pb"
	"example/grpc/application/usecase"
)

type ProductGrpcService struct {
	ProductUseCase usecase.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func (p *ProductGrpcService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product, err := p.ProductUseCase.Register(in.Name, in.Description, in.Price)
	if err != nil {
		return &pb.CreateProductResponse{}, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}, nil
}

func (p *ProductGrpcService) FindProducts(ctx context.Context, in *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	productsFromDb, err := p.ProductUseCase.FindAll()
	if err != nil {
		return &pb.FindProductsResponse{
			Products: []*pb.Product{},
		}, err
	}

	var products []*pb.Product

	for _, element := range productsFromDb {
		addProduct := pb.Product{
			Id:          element.ID,
			Name:        element.Name,
			Description: element.Description,
			Price:       element.Price,
		}

		products = append(products, &addProduct)
	}

	return &pb.FindProductsResponse{
		Products: products,
	}, nil
}

// func (p *ProductGrpcService) FindProduct(ctx context.Context, in *pb.FindProductRequest) (*pb.FindProductResponse, error) {
// 	product, err := p.ProductUseCase.FindByKey(in.Id)
// 	if err != nil {
// 		return &pb.FindProductResponse{
// 			Product: &pb.Product{},
// 		}, err
// 	}

// 	return &pb.FindProductResponse{
// 		Product: &pb.Product{
// 			Id:          product.ID,
// 			Name:        product.Name,
// 			Description: product.Description,
// 			Price:       product.Price,
// 		},
// 	}, nil
// }

func NewProductGrpcService(usecase usecase.ProductUseCase) *ProductGrpcService {
	return &ProductGrpcService{
		ProductUseCase: usecase,
	}
}
