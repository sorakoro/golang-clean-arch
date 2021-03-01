package schema

import "time"

// Article Article schema
//+migu table:"articles"
type Article struct {
	ID        int `migu:"pk,autoincrement"`
	UserID    int
	Text      string    `migu:"type:text"`
	CreatedAt time.Time `migu:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `migu:"default:CURRENT_TIMESTAMP,extra:ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt time.Time `migu:"null"`
}
