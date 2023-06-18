package repository

import (
	"github.com/rendramakmur/freefolks-fc/model/entity"
	"github.com/rendramakmur/freefolks-fc/model/support"
	"gorm.io/gorm"
)

type GlobalParamRepository struct {
	db *gorm.DB
}

func NewGlobalParamRepository(db *gorm.DB) *GlobalParamRepository {
	return &GlobalParamRepository{db}
}

func (gpr *GlobalParamRepository) GetDefaultDataBySlug(slug string) ([]support.DefaultData, error) {
	var result []entity.GlobalParam
	var response []support.DefaultData

	if err := gpr.db.Where("gp_slug = ?", slug).Find(&result).Error; err != nil {
		return response, err
	}

	for i, r := range result {
		response[i] = support.DefaultData{Id: r.CodeId, Name: r.Description}
	}

	return response, nil
}

func (gpr *GlobalParamRepository) GetDefaultDataBySlugAndCodeId(slug string, codeId *int) (support.DefaultData, error) {
	var result entity.GlobalParam

	if codeId == nil {
		return support.DefaultData{Id: nil, Name: nil}, nil
	}

	if err := gpr.db.Where("gp_slug = ? AND gp_code_id = ?", slug, codeId).First(&result).Error; err != nil {
		return support.DefaultData{Id: nil, Name: nil}, err
	}

	return support.DefaultData{Id: result.CodeId, Name: result.Description}, nil
}
