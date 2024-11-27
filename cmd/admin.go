package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"syscall"

	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"gorm.io/gorm"
)

func NewAdminCommand(app *core.App) *cli.Command {
	command := &cli.Command{
		Name: "admin",
		Subcommands: []*cli.Command{
			{
				Name:  "passwd",
				Usage: "set account password",
				Action: func(_ *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitORM(); err != nil {
						return err
					}

					fmt.Print("Enter Username: ")
					reader := bufio.NewReader(os.Stdin)
					username, err := reader.ReadString('\n')
					if err != nil {
						return err
					}

					fmt.Print("Enter Password: ")
					bytePass, _ := term.ReadPassword(int(syscall.Stdin))
					pass, err := bcrypt.GenerateFromPassword(bytePass, 14)
					if err != nil {
						return err
					}

					user := models.User{
						Name: strings.TrimSpace(username),
						Pass: strings.TrimSpace(string(pass)),
					}
					// 单用户
					app.O.Where("is_del = ?", 0).Unscoped().Delete(&models.User{})
					return app.O.Create(&user).Error
				},
			},
			{
				Name:  "save-attachments",
				Usage: "download attachments to /attachments dir",
				Action: func(_ *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitORM(); err != nil {
						return err
					}
					err := app.O.Model(&models.Attachment{}).
						FindInBatches(&models.Attachment{}, 3, func(tx *gorm.DB, batch int) error {
							var wg sync.WaitGroup
							var attachments []*models.Attachment
							if err := tx.Find(&attachments).Error; err != nil {
								return err
							}
							for _, attachment := range attachments {
								wg.Add(1)
								go func(v *models.Attachment) {
									defer wg.Done()

									// 发送 GET 请求
									resp, err := http.Get(app.Settings.System.AttachmentCDN + v.Key)
									if err != nil {
										fmt.Println("Error fetching URL:", err)
										return
									}
									defer resp.Body.Close()

									os.Mkdir("attachments", 0755)
									// 创建目标文件
									file, err := os.Create("attachments/" + v.Key)
									if err != nil {
										fmt.Println("Error creating file:", err)
										return
									}
									// 将响应体复制到文件
									_, err = io.Copy(file, resp.Body)
									if err != nil {
										fmt.Println("Error copying response body to file:", err)
										return
									}
									fmt.Println("Downloaded:", v.Key)
									defer file.Close()
								}(attachment)
							}
							wg.Wait()
							return nil
						}).Error
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
	return command
}
