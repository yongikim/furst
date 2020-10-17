package repository

import (
	"furst/model"
)

type MainCategoryRepository struct{}

func (MainCategoryRepository) SetMainCategory(mainCategory *model.MainCategory) error {
	result := db.Create(&mainCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (MainCategoryRepository) GetMainCategoryList() []model.MainCategory {
	mainCategories := make([]model.MainCategory, 0)
	result := db.Limit(10).Find(&mainCategories)
	if result.Error != nil {
		panic(result.Error)
	}
	return mainCategories
}

func (MainCategoryRepository) UpdateMainCategory(newMainCategory *model.MainCategory) error {
	result := db.Save(&newMainCategory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (MainCategoryRepository) DeleteMainCategory(id int) error {
	result := db.Delete(&model.MainCategory{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (MainCategoryRepository) FindByName(name string) (model.MainCategory, bool) {
	mainCategory := model.MainCategory{}
	result := db.Where("name = ?", name).First(&mainCategory)
	if result.Error != nil {
		return mainCategory, false
	}
	return mainCategory, true
}
