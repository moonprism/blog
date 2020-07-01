package email

import (
	db "git.kicoe.com/blog/write/database"
	"git.kicoe.com/blog/write/model"
	"git.kicoe.com/blog/write/utils"
	"html/template"
	"io/ioutil"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

type CommentMessage struct {
	Name	string
	ArtLink	string
	ArtTitle        string
	Text	string
	ReplyName	string
}

func run() {
	for ;; {
		// todo log
		commentID, err := utils.NewRedisClient().Do("brpop", "comment_message", 10).Result()
		if err != nil {
			break
		}
		// comment
		comment := new(model.Comment)
		has, err := db.MysqlXEngine.Id(commentID).Get(comment)
		if err != nil || !has {
			break
		}
		messageComment := new(model.Comment)
		has, err = db.MysqlXEngine.Id(comment.ToID).Get(comment)
		if err != nil || !has {
			break
		}
		// article
		article, has, err := model.FetchArticle(comment.ArtID)
		if err != nil || !has {
			break
		}
		data := CommentMessage {
			Name: messageComment.Name,
			ArtLink: "https://www.kicoe.com/article/id/"+strconv.FormatInt(article.ID, 10),
			ArtTitle: article.Title,
			Text: comment.Text,
			ReplyName: comment.Name,
		}

		t, err := template.New("comment").Parse()
		err = t.Execute(os.Stdout, data)
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

func sendMail() {
	user := ""
	port := ""
	passwd := ""
	host := ""
	to      :=      []string{"1370808234@qq.com", "ankicoe@gmail.com"}

	file, err := os.Open("./email.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)

	auth := smtp.PlainAuth("", user, passwd, host)
	mail := NewEmail(to, "测试", string(content))
	err = smtp.SendMail(host+port, auth, user, to, mail.Parse())
	if err != nil {

	}
}