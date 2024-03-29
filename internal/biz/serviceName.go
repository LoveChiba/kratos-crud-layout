package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrRecordNotFound = "数据不存在"
)

type ServiceName struct {
	Id int64
}

// 模块接口
type ServiceNameRepo interface {
	CreateServiceName(ctx context.Context, reqData *ServiceName) (*ServiceName, error)
	UpdateServiceName(ctx context.Context, reqData *ServiceName) (*ServiceName, error)
	GetServiceName(ctx context.Context, id int64) (*ServiceName, error)
	ListServiceName(ctx context.Context, pageNum, pageSize int64) ([]*ServiceName, int64, error)
	DeleteServiceName(ctx context.Context, id int64) error
}

type ServiceNameUseCase struct {
	repo ServiceNameRepo
	log  *log.Helper
}

func NewServiceNameUseCase(repo ServiceNameRepo, logger log.Logger) *ServiceNameUseCase {
	return &ServiceNameUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/beer"))}
}

func (uc *ServiceNameUseCase) Create(ctx context.Context, data *ServiceName) (*ServiceName, error) {
	return uc.repo.CreateServiceName(ctx, data)
}

func (uc *ServiceNameUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.DeleteServiceName(ctx, id)
}

func (uc *ServiceNameUseCase) Update(ctx context.Context, data *ServiceName) (*ServiceName, error) {
	return uc.repo.UpdateServiceName(ctx, data)
}

func (uc *ServiceNameUseCase) Get(ctx context.Context, id int64) (*ServiceName, error) {
	return uc.repo.GetServiceName(ctx, id)
}

func (uc *ServiceNameUseCase) List(ctx context.Context, pageNum, pageSize int64) ([]*ServiceName, int64, error) {
	return uc.repo.ListServiceName(ctx, pageNum, pageSize)
}
