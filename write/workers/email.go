package workers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"

	_ "embed"

	"git.kicoe.com/blog/write/config"
	db "git.kicoe.com/blog/write/database"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
)

type CommentMessage struct {
	Name      string
	ArtLink   string
	ArtTitle  string
	Text      string
	ReplyName string
}

func RunSendEmail() {
	os.MkdirAll("log", os.ModePerm)
	f, err := os.OpenFile("log/email.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	lo := log.New(f, "", log.Ldate|log.Ltime)
	for ;; {
		// todo log
		commentID, err := utils.NewRedisClient().BLPop(20*time.Second, "aqua:comment_message").Result()
		if err != nil {
			continue;
		}
		// comment
		comment := new(model.Comment)
		has, err := db.MysqlXEngine.Id(commentID).Get(comment)

		if err != nil || !has {
			lo.Printf("select comment id:%s error: %v\n", commentID, err)
		}
		messageComment := new(model.Comment)
		has, err = db.MysqlXEngine.Id(comment.ToID).Get(messageComment)
		if err != nil || !has {
			lo.Printf("select comment to id:%s error:%v\n", commentID, err)
		}
		// article
		article, has, err := model.FetchArticle(comment.ArtID)
		if err != nil || !has {
			lo.Printf("select article id:%d, %v\n", comment.ArtID, err)
		}
		data := CommentMessage {
			Name: messageComment.Name,
			ArtLink: "https://www.kicoe.com/article/id/"+strconv.FormatInt(article.ID, 10),
			ArtTitle: article.Title,
			Text: comment.Text,
			ReplyName: comment.Name,
		}
		sendMail(messageComment.Email, data)
		lo.Printf("send email success [%s ,%v]\n", messageComment.Email, data)
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
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (e *Email) Parse() []byte {
	to := "To: "+strings.Join(e.to, ",")+"\n"
	from := "From: kicoe<kicoe@qq.com>\n"
	subject := "Subject: "+e.subject+"\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	return []byte(to+from+subject+mime+e.body)
}

//go:embed email.html
var html string

func sendMail(email string, info CommentMessage) {
	t, err := template.New("comment").Parse(html)
	ww := bytes.NewBufferString("")
	err = t.Execute(ww, info)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	user := config.SMTP.User
	port := ":" + config.SMTP.Port
	passwd := config.SMTP.Pass
	host := config.SMTP.Host
	to      :=      []string{email}

	content := fmt.Sprint(ww)

	auth := smtp.PlainAuth("", user, passwd, host)
	mail := NewEmail(to, "评论回复提醒", content)
	err = SendMail(host+port, auth, user, to, mail.Parse())
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func SendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	c, err := smtp.Dial(addr)
	host, _, _ := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if ok, _ := c.Extension("STARTTLS"); ok {
		config := &tls.Config{ServerName:  host, InsecureSkipVerify: true}
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}
	if a != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(a); err != nil {
                fmt.Println("check auth with err:", err)
                return err
            }
        }
    }
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
