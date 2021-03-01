package schema

import "time"

// Article Article schema
//+migu table:"articles"
type Article struct {
	ID        int64 `migu:"pk,autoincrement"`
	UserID    int64
	Title     string    `migu:"type:varchar(255)"`
	Text      string    `migu:"type:text"`
	CreatedAt time.Time `migu:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `migu:"default:CURRENT_TIMESTAMP,extra:ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt time.Time `migu:"null"`
}
