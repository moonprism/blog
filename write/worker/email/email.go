package email

import (
	"fmt"
	"git.kicoe.com/blog/write/config"
	db "git.kicoe.com/blog/write/database"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
	"html/template"
	"bytes"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type CommentMessage struct {
	Name	string
	ArtLink	string
	ArtTitle        string
	Text	string
	ReplyName	string
}

func Run() {
	for ;; {
		// todo log
		commentID, err := utils.NewRedisClient().BLPop(5*time.Second, "comment_message").Result()
		if err != nil {
			continue;
		}
		// comment
		comment := new(model.Comment)
		has, err := db.MysqlXEngine.Id(commentID).Get(comment)
		
		if err != nil || !has {
		}
		messageComment := new(model.Comment)
		has, err = db.MysqlXEngine.Id(comment.ToID).Get(messageComment)
		if err != nil || !has {
		}
		// article
		article, has, err := model.FetchArticle(comment.ArtID)
		if err != nil || !has {
		}
		data := CommentMessage {
			Name: messageComment.Name,
			ArtLink: "https://www.kicoe.com/article/id/"+strconv.FormatInt(article.ID, 10),
			ArtTitle: article.Title,
			Text: comment.Text,
			ReplyName: comment.Name,
		}
		println(messageComment.Email)
		sendMail(messageComment.Email, data)
		fmt.Printf("%v\n", data)
	}
}

type Email struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewEmail(to []string, subject, body string) *Email {
	return &Email{
		to:     to,
		subject:        subject,
		body:   body,
	}
}

func (e *Email) Parse() []byte {
	to := "To: "+strings.Join(e.to, ",")+"\n"
	from := "From: kicoe<kicoe@qq.com>\n"
	subject := "Subject: "+e.subject+"\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	return []byte(to+from+subject+mime+e.body)
}

const html = `
<table align="center" style="border: 1px solid #ccc;padding: 0 17px 18px;border-radius: 5px;" cellpadding="0" cellspacing="0" width="600" style="border-collapse: collapse;">
		<tr>
				<td style="line-height: 1;">
						<div style="font-family:微软雅黑; font-size:15px;padding: 5px 10px 10px;margin: 0 10px; background-color: #fff;color: #233; position: relative; bottom: 10px;width: max-content;"><span style="color: #f54291;font-weight: 600;margin-right: 4px;">@</span><span style="font-family: Arial, 微软雅黑;margin-right: 1px;">{{.Name}}</span>酱</div>
				</td>
		</tr>
		<tr>
				<td><div style="font-family:微软雅黑; font-size:15px;color: #42426F;margin: 4px 3px 8px;">你在<a style="color: #12cad6;text-decoration: none;margin: 0 1px;" href="{{.ArtLink}}#com_list">「{{.ArtTitle}}」</a>下的留言有了新回复哦</div></td>
		</tr>
		<tr>
				<td><div style="font-family:微软雅黑; font-size: 15px;margin: 20px 0; padding: 22px 20px 14px; background-color: #f5f5f5;color: #233;border-radius: 5px;">{{.Text}}<div style="font-family: cursive;text-align: right;color: #383e56;margin-top: 5px">—— {{.ReplyName}}</div></div></td>
		</tr>
		<tr>
				<td align="center"><div style="font-family: Arial, 微软雅黑;width: max-content; background-color: #444;color: #fff;font-size: 12px;text-shadow: 0 1px 1px rgba(255,255,255,.2); padding: 0 4px;border-radius: 1px;margin: 10px 0 0;">系统邮件 by <a style="color: aqua;text-decoration: none;" href="https://www.kicoe.com/">kicoe.com</a></div></td>
		</tr>
</table>
`

func sendMail(email string, info CommentMessage) {
	t, err := template.New("comment").Parse(html)
	ww := bytes.NewBufferString("")
	err = t.Execute(ww, info)

	user := config.SMTP.User
	port := ":" + config.SMTP.Port
	passwd := config.SMTP.Pass
	host := config.SMTP.Host
	to      :=      []string{email}

	content := fmt.Sprint(ww)

	auth := smtp.PlainAuth("", user, passwd, host)
	mail := NewEmail(to, "评论回复提醒", content)
	err = smtp.SendMail(host+port, auth, user, to, mail.Parse())
	if err != nil {
	}
}