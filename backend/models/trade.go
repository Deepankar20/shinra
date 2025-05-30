package models

type Trade struct {
    BaseModel
    UserID    uint
    EventID   uint
    Amount    float64
    Prediction string  // e.g., 'yes' or 'no'
    Status     string  `gorm:"default:'pending'"` // e.g., 'pending', 'won', 'lost'
    Result     string  // e.g., 'yes', 'no', 'pending'
} 
