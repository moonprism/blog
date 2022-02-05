package email

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"net/smtp"
	"strconv"

	"git.kicoe.com/blog/write/modules/redis"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/services/article"
	"git.kicoe.com/blog/write/services/comment"
)

//go:embed templates/reply.html
var replyHtml string

type CommentEmailData struct {
	Name      string
	ArtLink   string
	ArtTitle  string
	Text      string
	ReplyName string
}

func RunCommentReplyNotice() {
	queue := NewRedisQueue(redis.Client, "aqua:")
	for {
		commentIdStr, err := queue.Pop("comment_message")
		if err != nil {
			continue
		}
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
		replyDetail, _ := comment.GetDetail(commentId)
		sendComment, _ := comment.GetDetail(replyDetail.ToID)
		artDetail, _ := article.GetDetail(replyDetail.ArtID)

		link := "https://www.kicoe.com"
		if artDetail.ID == 1 {
			link += "/page/link/#com_list"
		} else {
			link += "/article/id/" + strconv.FormatInt(artDetail.ID, 10)
		}

		data := CommentEmailData{
			Name:      sendComment.Name,
			ArtLink:   link,
			ArtTitle:  artDetail.Title,
			Text:      replyDetail.Text,
			ReplyName: replyDetail.Name,
		}

		sendReplyEmail(replyDetail.Email, &data)
	}
}

func sendReplyEmail(email string, data *CommentEmailData) {
	tmp, _ := template.New("reply").Parse(replyHtml)
	tmpBuf := bytes.NewBufferString("")
	err := tmp.Execute(tmpBuf, data)
	if err != nil {
		logger.Errorf("%v\n", err)
	}
	content := fmt.Sprint(tmpBuf)

	user := setting.SMTP.User
	port := ":" + setting.SMTP.Port
	passwd := setting.SMTP.Pass
	host := setting.SMTP.Host
	auth := smtp.PlainAuth("", user, passwd, host)
	sendEmail := []string{email}

	mail := NewEmail(sendEmail, "评论回复提醒", content)
	err = SendMail(host+port, auth, user, sendEmail, mail.Parse())
	if err != nil {
		logger.Errorf("%v\n", err)
	}
}
