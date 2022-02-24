package comment

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"strconv"
	"sync/atomic"
	"time"

	"git.kicoe.com/blog/write/models"
	"git.kicoe.com/blog/write/modules/email"
	"git.kicoe.com/blog/write/modules/redis"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/services/article"
	"git.kicoe.com/blog/write/services/link"
)

const prefix = "aqua:"

//go:embed notice_template.html
var templateHTML string

type templateData struct {
	Name   string
	Link   string
	Title  string
	Text   string
	ToName string
	ToID   int64
}

var noticeFlag atomic.Value

func init() {
	noticeFlag.Store(false)
}

func RunNoticServer() {
	for {
		re, err := redis.Client.BLPop(time.Hour, prefix+"comment_message").Result()
		if err != nil {
			continue
		}
		if !noticeFlag.Load().(bool) {
			continue
		}

		commentId, _ := strconv.ParseInt(re[1], 10, 64)
		comment, _ := GetDetail(commentId)
		var toComment *models.Comment
		if comment.ToID == 0 {
			toComment = &models.Comment{
				Name:  "moonprism",
				Email: setting.SMTP.User,
			}
		} else {
			toComment, _ = GetDetail(comment.ToID)
			// 只发送申请通过的邮箱
			_, has, err := link.FetchDetailByEmail(toComment.Email)
			if err != nil || !has {
				continue
			}
		}

		article, _ := article.GetDetail(comment.ArtID)
		templateLink := "https://www.kicoe.com/article/id/" + strconv.FormatInt(article.ID, 10)
		data := templateData{
			Name:   toComment.Name,
			Link:   templateLink,
			Title:  article.Title,
			Text:   comment.Text,
			ToName: toComment.Name,
			ToID:   comment.ToID,
		}

		tmp, _ := template.New("notice").Parse(templateHTML)
		tmpBuf := bytes.NewBufferString("")
		tmp.Execute(tmpBuf, data)
		content := fmt.Sprint(tmpBuf)

		email := &email.Email{
			To:      []string{toComment.Email},
			Subject: "评论回复提醒",
			Body:    content,
		}
		email.Send()
	}
}

func GetNoticeFlag() bool {
	return noticeFlag.Load().(bool)
}

func SetNoticeFlag(flag bool) {
	noticeFlag.Store(flag)
}
