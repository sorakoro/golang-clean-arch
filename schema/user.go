package schema

import "time"

// User User schema
//+migu table:"users"
type User struct {
	ID        int64     `migu:"pk,autoincrement"`
	Name      string    `migu:"type:varchar(255)"`
	Email     string    `migu:"type:varchar(255),unique"`
	Password  string    `migu:"type:varchar(255)"`
	CreatedAt time.Time `migu:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `migu:"default:CURRENT_TIMESTAMP,extra:ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt time.Time `migu:"null"`
}
