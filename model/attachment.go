package model

type Attachment struct {
	BaseModel

	Key     string `gorm:"type:varchar(255);notnull;comment:访问地址" json:"key"`
	Summary string `gorm:"type:varchar(2000);notnull" json:"summary"`

	// 保存文件大小，尺寸等信息
	// Info string `gorm:"type:varchar(1000);notnull" json:"info"`
}
