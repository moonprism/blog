package model

type Attachment struct {
	BaseModel

	Link    string `gorm:"type:varchar(255);notnull;comment:访问地址" json:"link"`
	Summary string `gorm:"type:varchar(2000);notnull" json:"summary"`
}
