package models

type Lock struct {
    Model
    Brand       string `gorm:"column:brand"         json:"brand"`
    ModelName   string `gorm:"column:model_name"    json:"model_name"`
    CoreType    string `gorm:"column:core_type"     json:"core_type"`
}

type CreateLockInput struct {
    Brand       string `json:"brand"        binding:"required"`
    ModelName   string `json:"model_name"   binding:"required"`
    CoreType    string `json:"core_type"    binding:"required"`
}

type UpdateLockInput struct {
    Brand       string `json:"brand"`
    ModelName   string `json:"model_name"`
    CoreType    string `json:"core_type"`
}