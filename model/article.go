package model

type Article struct {
	*BaseModel `xorm:"extends -"`

	Title   string `xorm:"notnull default('')" json:"title"`
	Status  int    `xorm:"tinyint notnull default(0)" json:"status"`
	Rune    int    `xorm:"notnull default(0)" json:"rune"`
	Image   string `xorm:"notnull default('')" json:"image"`
	Summary string `xorm:"varchar(1024) notnull default('')" json:"summary"`
	Content string `xorm:"text notnull default('')" json:"content"`
	Deleted *int   `xorm:"deleted" json:"deleted"`
}
