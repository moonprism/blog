package model

type Gist struct {
	*BaseModel `xorm:"extends -"`

	Description string `xorm:"notnull default('')" json:"desc"`
	Lang        string `xorm:"notnull default('')" json:"lang"`
	Tags        string `xorm:"notnull default('')" json:"tags"`
	Content     string `xorm:"text notnull default('')" json:"content"`
	Deleted     *int   `xorm:"deleted" json:"deleted"`
}
