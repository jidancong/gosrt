package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jidancong/gosrt/config"

	"github.com/urfave/cli/v2"
)

func Cmd() error {
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		return fmt.Errorf("read config error: %s", err)
	}
	if cfg.Config(); err != nil {
		return fmt.Errorf("init config error: %s", err)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "srt",
				Usage: "字幕文件",
				Before: func(ctx *cli.Context) error {
					if ctx.Args().Len() < 1 {
						return fmt.Errorf("the number of parameters is less than 1")
					}
					fileName := ctx.Args().First()

					if filepath.Ext(fileName) != ".srt" {
						return fmt.Errorf("file name must end with .srt")
					}
					return cfg.Validate()
				},
				Action: func(cCtx *cli.Context) error {
					fileName := cCtx.Args().First()
					apikey := cfg.GetKey()
					apiUrl := cfg.GetUrl()

					err := translate(fileName, apikey, apiUrl)
					return err
				},
			},
			{
				Name:  "srtpath",
				Usage: "字幕目录",
				Before: func(ctx *cli.Context) error {
					if ctx.Args().Len() < 1 {
						return fmt.Errorf("the number of parameters is less than 1")
					}
					path := ctx.Args().First()
					pathInfo, err := os.Stat(path)
					if err != nil {
						return err
					}

					if !pathInfo.IsDir() {
						return fmt.Errorf("not a directory")
					}
					return cfg.Validate()
				},
				Action: func(cCtx *cli.Context) error {
					path := cCtx.Args().First()
					srtFiles, err := GetSrtFiles(path)
					if err != nil || len(srtFiles) == 0 {
						return err
					}

					apikey := cfg.GetKey()
					apiUrl := cfg.GetUrl()

					for _, file := range srtFiles {
						if err := translate(file, apikey, apiUrl); err != nil {
							return err
						}
					}
					return nil
				},
			},
			{
				Name:  "set",
				Usage: "配置",
				Action: func(cCtx *cli.Context) error {
					if cCtx.Args().Len() < 2 {
						return fmt.Errorf("the number of parameters is less than 2")
					}
					key := cCtx.Args().Get(0)
					value := cCtx.Args().Get(1)
					cfg.Set(key, value)
					return nil
				},
			},
		},
	}

	return app.Run(os.Args)
}

func GetSrtFiles(path string) ([]string, error) {
	var srtFiles []string
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".srt" {
			srtPath, _ := filepath.Abs(file.Name())
			srtFiles = append(srtFiles, srtPath)
		}
	}

	return srtFiles, nil
}
