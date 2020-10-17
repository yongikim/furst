package repository

import (
	"furst/model"
)

type SubCategoryRepository struct {}

func (SubCategoryRepository) SetSubCategory(subCategory *model.SubCategory) error {
  result := db.Create(&subCategory)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (SubCategoryRepository) GetSubCategoryList() []model.SubCategory {
  subCategories := make([]model.SubCategory, 0)
  result := db.Limit(10).Find(&subCategories)
  if result.Error != nil {
    panic(result.Error)
  }
  return subCategories
}

func (SubCategoryRepository) UpdateSubCategory(newSubCategory *model.SubCategory) error {
  result := db.Save(&newSubCategory)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (SubCategoryRepository) DeleteSubCategory(id int) error {
  result := db.Delete(&model.SubCategory{}, id)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (SubCategoryRepository) FindByName(name string) (model.SubCategory, bool) {
  subCategory := model.SubCategory{}
  result := db.Where("name = ?", name).First(&subCategory)
  if result.Error != nil {
    return subCategory, false
  }
  return subCategory, true
}
