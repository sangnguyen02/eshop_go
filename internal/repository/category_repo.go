package repository

import (
	"go_ecommerce/internal/model"
	"go_ecommerce/pkg/database"
	"strconv"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{db: database.GetDB()}
}

func (r *CategoryRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *CategoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) FindByID(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) FindAll(page, pageSize int, search string, status bool) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	query := r.db.Model(&model.Category{}).Where("deleted_at IS NULL")

	if !status {
		query = query.Where("status = ?", true)
	}

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize

	err = query.Offset(offset).Limit(pageSize).Find(&categories).Error
	if err != nil {
		return nil, 0, err
	}

	parentIDs := make([]uint, 0)
	for _, category := range categories {
		if category.ParentID != nil {
			parentIDs = append(parentIDs, *category.ParentID)
		}
	}

	var parents []model.Category
	if len(parentIDs) > 0 {
		err = r.db.Where("id IN ? AND deleted_at IS NULL", parentIDs).Find(&parents).Error
		if err != nil {
			return nil, 0, err
		}
	}

	parentMap := make(map[string]model.Category)
	for _, parent := range parents {
		parentMap[strconv.FormatUint(uint64(parent.ID), 10)] = parent
	}

	for i, category := range categories {
		if category.ParentID != nil {
			if parent, exists := parentMap[strconv.FormatUint(uint64(*category.ParentID), 10)]; exists {
				categories[i].Parent = &model.Category{
					BaseModel: model.BaseModel{ID: parent.ID},
					Name:      parent.Name,
				}
			}
		}
	}

	return categories, total, err
}

func (r *CategoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}
