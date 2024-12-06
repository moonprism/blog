package cmd

import (
	"github.com/moonprism/blog/core"
	"github.com/moonprism/blog/models"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"

	_ "embed"
)

//go:embed about_tmpl.md
var aboutTmpl string

//go:embed links_tmpl.md
var linksTmpl string

func NewDbCommand(app *core.App) *cli.Command {
	command := &cli.Command{
		Name: "db",
		Subcommands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "migrate database and construct fts data",
				Action: func(_ *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitORM(); err != nil {
						return err
					}

					err := app.O.OrmClient.AutoMigrate(
						&models.Article{},
						&models.ArticleContent{},
						&models.Tag{},
						&models.ArticleTags{},
						&models.Attachment{},
						&models.Gist{},
						&models.GistOutput{},
						&models.Comment{},
						&models.LoginRecord{},
						&models.User{},
						&models.Settings{},
					)

					var a1 models.Article
					// 载入默认模板 1: about 2: links
					if err := app.O.Where("id = ?", 1).Take(&a1).Error; err != nil {
						if app.O.IsRecordNotFoundErr(err) {
							a1.ID = 1
							a1.Title = "About"
							a1.ArticleContent = &models.ArticleContent{ArticleID: 1, Text: aboutTmpl}
							app.O.Create(&a1)
						} else {
							return err
						}
					}
					var a2 models.Article
					if err := app.O.Where("id = ?", 2).Take(&a2).Error; err != nil {
						if app.O.IsRecordNotFoundErr(err) {
							a2.ID = 2
							a2.Title = "Links"
							a2.ArticleContent = &models.ArticleContent{ArticleID: 2, Text: linksTmpl}
							app.O.Create(&a2)
						} else {
							return err
						}
					}

					var settings models.Settings
					if err := app.O.First(&settings).Error; err != nil {
						if app.O.IsRecordNotFoundErr(err) {
							settings.Title = "kicoe's Blog"
							settings.Background = "background: linear-gradient( 180deg, rgba(238, 174, 202, 1) 0%, rgba(148, 187, 233, 1) 100%);"
							settings.MarginBottom = 140
							app.O.Create(&settings)
						} else {
							return err
						}
					}

					if err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("gist"); err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("article"); err != nil {
						return err
					}
					return nil
					/* 还是别用触发器，自己用程序处理吧
					if err != nil {
						return err
					}
					_, err = app.O.SqliteFtsDB.Exec(`CREATE TRIGGER after_insert_gist
						AFTER INSERT ON gists
						BEGIN
						    INSERT INTO gists_fts (rowid, title, lang, content) VALUES (NEW.id, NEW.title, NEW.lang, NEW.content);
						END;
					`)
					*/
				},
			},
			{
				Name:  "sync-fts",
				Usage: "sync sqlite fts",
				Action: func(_ *cli.Context) error {
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitORM(); err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("gist"); err != nil {
						return err
					}
					if err := app.O.FtsCreateTable("article"); err != nil {
						return err
					}
					// 初始化gists查询
					err := app.O.Model(&models.Gist{}).
						FindInBatches(&models.Gist{}, 20, func(tx *gorm.DB, batch int) error {
							var gists []*models.Gist
							if err := tx.Find(&gists).Error; err != nil {
								return err
							}
							for _, g := range gists {
								if err := app.O.FtsInsert("gist", g.ID, models.Gist2TextPoint(g)); err != nil {
									return err
								}
							}
							return nil
						}).Error
					if err != nil {
						return err
					}
					// 初始化文章查询
					err = app.O.Model(&models.Article{}).
						FindInBatches(&models.Article{}, 10, func(tx *gorm.DB, batch int) error {
							var articles []*models.Article
							err := tx.Preload("ArticleContent", func(db *gorm.DB) *gorm.DB {
								return db.Omit("html")
							}).Find(&articles).Error
							if err != nil {
								return err
							}
							for _, art := range articles {
								if art.Status != models.ArticleStatusPublished {
									continue
								}
								if err := app.O.FtsInsert("article", art.ID, models.Art2TextPoint(art)); err != nil {
									return err
								}
							}
							return nil
						}).Error
					return err
				},
			},
			{
				Name:  "fix-cmnt-time",
				Usage: "comment updated time sync from created",
				Action: func(_ *cli.Context) error {
					// 修复旧版本数据
					if err := app.InitSetting(); err != nil {
						return err
					}
					if err := app.InitORM(); err != nil {
						return err
					}

					var cmnts []*models.Comment
					if err := app.O.Find(&cmnts).Error; err != nil {
						return err
					}
					for _, v := range cmnts {
						err := app.O.Model(&v).Updates(map[string]interface{}{"updated": v.Created}).Error
						if err != nil {
							return err
						}
					}
					for _, v := range cmnts {
						if v.RootCommentID == 0 {
							for _, v2 := range cmnts {
								if v2.RootCommentID == v.ID {
									v.Updated = max(v.Updated, v2.Created)
								}
							}
							err := app.O.Model(&v).Updates(map[string]interface{}{"updated": v.Updated}).Error
							if err != nil {
								return err
							}
						}
					}
					return nil
				},
			},
		},
	}
	return command
}
