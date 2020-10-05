package models

type UserLock struct {
    Model
    LockID          int     `gorm:"column:lock_id"           json:"lock_id"`
    OwnerID         int     `gorm:"column:owner_id"          json:"owner_id"`
    TransferredTo   *int    `gorm:"column:transferred_to"    json:"transferred_to"`
    Picked          bool    `gorm:"column"picked"            json:"picked"`
    Notes           *string `gorm:"column:notes"             json:"notes"`
}

type CreateUserLockInput struct {
    LockID      int     `json:"lock_id"     binding:"required"`
    OwnerID     int     `json:"owner_id"    binding:"required"`
}

type UpdateUserLockInput struct {
    TransferredTo   int     `json:"transferred_to"`
    Picked          bool    `json:"picked"`
    Notes           string  `json:"notes"`
}