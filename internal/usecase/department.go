package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
	"github.com/hareta0109/graphql_sandbox/internal/domain/model/graph"
	"github.com/hareta0109/graphql_sandbox/internal/domain/repository"
	"github.com/hareta0109/graphql_sandbox/internal/lib/graph/loader"
)

type Department interface {
	Get(id uint64) (*graph.Department, error)
	BulkGet() ([]*graph.Department, error)
	Create(name string) (*graph.Department, error)

	BatchGetDepartments(ctx context.Context, keys dataloader.Keys) []*dataloader.Result
	LoadDepartment(ctx context.Context, departmentID uint64) (*graph.Department, error)
}

type DepartmentUsecase struct {
	departmentRepo repository.Department
}

func NewDepartmentRepository(departmentRepo repository.Department) Department {
	departmentUsecase := DepartmentUsecase{departmentRepo: departmentRepo}
	return &departmentUsecase
}

func (u *DepartmentUsecase) Get(id uint64) (*graph.Department, error) {
	department, err := u.departmentRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (u *DepartmentUsecase) BulkGet() ([]*graph.Department, error) {
	departments, err := u.departmentRepo.BulkGet()
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (u *DepartmentUsecase) Create(name string) (*graph.Department, error) {
	department := model.NewDepartment(
		model.DepartmentDescription{
			Name: name,
		},
	)

	created, err := u.departmentRepo.Create(department)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (u *DepartmentUsecase) BatchGetDepartments(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	departmentIDs := make([]uint64, len(keys))
	for i, key := range keys {
		departmentIDs[i], _ = strconv.ParseUint(key.String(), 10, 64)
	}

	departmentByID, err := u.departmentRepo.GetMapInIDs(ctx, departmentIDs)
	if err != nil {
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		departmentID, _ := strconv.ParseUint(key.String(), 10, 64)
		department, ok := departmentByID[departmentID]
		if ok {
			output[i] = &dataloader.Result{Data: department, Error: nil}
		} else {
			err := fmt.Errorf("department not found %d", departmentID)
			output[i] = &dataloader.Result{Data: nil, Error: err}
		}
	}

	return output
}

func (u *DepartmentUsecase) LoadDepartment(ctx context.Context, departmentID uint64) (*graph.Department, error) {
	loaders := loader.GetLoaders(ctx)

	// 遅延読み込み
	thunk := loaders.DepartmentLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%d", departmentID)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	department := result.(*graph.Department)

	return department, nil
}
