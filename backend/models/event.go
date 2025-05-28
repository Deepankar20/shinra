package models

type Event struct {
    BaseModel
    Title       string  `gorm:"not null"`
    Description string
    OddsYes     float64
    OddsNo      float64
    Status      string  `gorm:"default:'open'"` // e.g., 'open', 'closed', 'settled'
    Outcome     string  // e.g., 'yes', 'no', 'pending'
}
