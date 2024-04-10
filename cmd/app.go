package cmd

import (
	"path/filepath"
	"strings"

	"github.com/jidancong/gosrt/pkg"

	gosrt "github.com/konifar/go-srt"
)

func translate(fileName string, apiKey, apiUrl string) error {
	oa := pkg.NewOpenai(apiKey, apiUrl)

	s, err := gosrt.ReadFile(fileName)
	if err != nil {
		return err
	}

	ext := filepath.Ext(fileName)
	newFileName := strings.Split(fileName, ext)[0] + ".output" + ext
	newFile, err := pkg.NewWriteSrt(newFileName)
	if err != nil {
		return err
	}

	defer newFile.Close()

	for _, v := range s {
		text, err := oa.TextToText(v.Text)
		if err != nil {
			return err
		}
		newFile.Add(pkg.Subtitle{Number: v.Number, Start: v.Start, End: v.End, Text: text})
	}
	return nil
}
