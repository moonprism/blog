package models

type Comment struct {
	BaseModel

	Name    string `gorm:"type:varchar(255);notnull" json:"name"`
	Email   string `gorm:"type:varchar(255);notnull" json:"email"`
	Link    string `gorm:"type:varchar(255);notnull" json:"link"`
	Content string `gorm:"type:varchar(1000);notnull" json:"content"`

	ArticleID      uint `gorm:"notnull" json:"article_id"`
	ReplyCommentID uint `gorm:"notnull" json:"reply_comment_id"`
	RootCommentID  uint `gorm:"notnull" json:"root_comment_id"`
}
