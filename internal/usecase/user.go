package usecase

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/hareta0109/graphql_sandbox/internal/domain/model"
	"github.com/hareta0109/graphql_sandbox/internal/domain/repository"
	"github.com/hareta0109/graphql_sandbox/internal/lib/graph/loader"
)

type User interface {
	Get(id uint64) (*model.User, error)
	BulkGet() ([]*model.User, error)
	Create(name string, departmentID uint64) (*model.User, error)

	BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result
	LoadUser(ctx context.Context, userID uint64) (*model.User, error)
}

type UserUsecase struct {
	userRepo repository.User
}

func NewUserUsecase(userRepo repository.User) User {
	userUsecase := UserUsecase{userRepo: userRepo}
	return &userUsecase
}

func (u *UserUsecase) Get(id uint64) (*model.User, error) {
	user, err := u.userRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) BulkGet() ([]*model.User, error) {
	users, err := u.userRepo.BulkGet()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUsecase) Create(name string, departmentID uint64) (*model.User, error) {
	user := model.NewUser(
		model.UserDescription{
			Name:         name,
			DepartmentID: departmentID,
		},
	)

	created, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (u *UserUsecase) BatchGetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	userIDs := make([]uint64, len(keys))
	for i, key := range keys {
		userIDs[i], _ = strconv.ParseUint(key.String(), 10, 64)
	}

	userByID, err := u.userRepo.GetMapInIDs(ctx, userIDs)
	if err != nil {
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		userID, _ := strconv.ParseUint(key.String(), 10, 64)
		user, ok := userByID[userID]
		if ok {
			output[i] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %d", userID)
			output[i] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	
	return output
}

func (u *UserUsecase) LoadUser(ctx context.Context, userID uint64) (*model.User, error) {
	loaders := loader.GetLoaders(ctx)

	// 遅延読み込み
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(fmt.Sprintf("%d", userID)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	user := result.(*model.User)
	
	return user, nil
}
