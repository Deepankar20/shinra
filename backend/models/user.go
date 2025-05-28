package models

type User struct {
    BaseModel
    Username string  `gorm:"unique;not null"`
    Email    string  `gorm:"unique;not null"`
    Password string  `gorm:"not null"`
    Role     string  `gorm:"default:'user'"` // e.g., 'user' or 'admin'
    Wallet   Wallet
}