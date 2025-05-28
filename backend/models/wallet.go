package models

type Wallet struct {
    BaseModel
    UserID  uint    `gorm:"uniqueIndex"`
    Balance float64 `gorm:"default:0"`
}
