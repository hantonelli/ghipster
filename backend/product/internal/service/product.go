package service

import (
	"context"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gogo/protobuf/types"
	"github.com/hantonelli/ghipster/middleware"
	"github.com/hantonelli/ghipster/product/internal/repository/entgen"
	"github.com/hantonelli/ghipster/product/internal/repository/entgen/product"
	pb "github.com/hantonelli/ghipster/proto/product"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductsService is the server API for ProductsService service.
type ProductsService struct {
	client *entgen.Client
	logger *zap.SugaredLogger
}

// NewProductsService creates a product service.
func NewProductsService(logger *zap.Logger, client *entgen.Client) *ProductsService {
	return &ProductsService{
		logger: logger.Sugar(),
		client: client,
	}
}

func toProtoProduct(prod *entgen.Product) *pb.Product {
	return &pb.Product{
		Id:   prod.ID,
		Name: prod.Name,
	}
}

func toProtoProducts(prods []*entgen.Product) []*pb.Product {
	list := make([]*pb.Product, 0, len(prods))
	for i, x := range prods {
		list[i] = &pb.Product{
			Id:   x.ID,
			Name: x.Name,
		}
	}
	return list
}

func (s *ProductsService) GetProduct(ctx context.Context, req *pb.IDRequest) (*pb.Product, error) {
	get, err := s.client.Product.Get(ctx, req.Id)
	switch {
	case err == nil:
		return toProtoProduct(get), nil
	case entgen.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}

func (s *ProductsService) GetProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	q := s.client.Product.Query()
	if req.Filter != nil {
		if req.Filter.NameLike != "" {
			q = q.Where(
				product.NameEQ(*&req.Filter.NameLike),
			)
		}
	}
	totalCount, err := q.Count(ctx)
	if err != nil {
		return nil, err
	}

	if req.OrderBy != nil {
		if req.OrderBy.Direction == pb.DESC {
			q = q.Order(entgen.Desc(string(req.OrderBy.Field)))
		} else {
			q = q.Order(entgen.Asc(string(req.OrderBy.Field)))
		}
	}
	if req.Offset != 0 {
		q = q.Offset(int(req.Offset))
	}
	nodes, err := q.Limit(int(req.Limit)).All(ctx)
	switch {
	case err == nil:
		return &pb.ListProductsResponse{
			TotalCount: int64(totalCount),
			Products:   toProtoProducts(nodes),
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}

func (s *ProductsService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.Product, error) {
	log := middleware.AddTraceToLog(ctx, s.logger)
	log.Infow("Calling CreateProduct")
	res, err := s.client.Product.
		Create().
		SetName(req.Name).
		Save(ctx)
	switch {
	case err == nil:
		return toProtoProduct(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case entgen.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

func (s *ProductsService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	res, err := s.client.Product.UpdateOneID(req.Id).SetName(req.Name).Save(ctx)
	switch {
	case err == nil:
		return toProtoProduct(res), nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case entgen.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal: %s", err)
	}
}

func (s *ProductsService) DeleteProduct(ctx context.Context, req *pb.IDRequest) (*types.Empty, error) {
	err := s.client.Product.
		DeleteOneID(req.Id).
		Exec(ctx)
	switch {
	case err == nil:
		return &types.Empty{}, nil
	case entgen.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}
