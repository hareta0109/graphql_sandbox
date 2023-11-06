package infra

import (
	"context"

	"github.com/hareta0109/graphql_sandbox/internal/domain/model/graph"
	"github.com/hareta0109/graphql_sandbox/internal/domain/repository"
	"github.com/hareta0109/graphql_sandbox/internal/infra/dao"
	"gorm.io/gorm"
)

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) repository.Department {
	return &departmentRepository{db}
}

func (r *departmentRepository) Get(id uint64) (*graph.Department, error) {
	var row dao.Department
	if err := r.db.Where("id = ?", id).First(row).Error; err != nil {
		return nil, err
	}

	department := row.ToModel()

	return department, nil
}

func (r *departmentRepository) BulkGet() ([]*graph.Department, error) {
	var rows []dao.Department
	if err := r.db.Find(rows).Error; err != nil {
		return nil, err
	}

	var departments []*graph.Department
	for _, row := range rows {
		department := row.ToModel()
		departments = append(departments, department)
	}

	return departments, nil
}

func (r *departmentRepository) Create(department *graph.Department) (*graph.Department, error) {
	dbModel := dao.NewDepartment(department)
	if err := r.db.Create(dbModel).Error; err != nil {
		return nil, err
	}

	return dbModel.ToModel(), nil
}

func (r *departmentRepository) GetMapInIDs(ctx context.Context, ids []uint64) (map[uint64]*graph.Department, error) {
	var rows []dao.Department
	if err := r.db.Find(&rows, ids).Error; err != nil {
		return nil, err
	}

	results := make(map[uint64]*graph.Department)
	for _, row := range rows {
		department := row.ToModel()
		results[department.ID] = department
	}

	return results, nil
}
