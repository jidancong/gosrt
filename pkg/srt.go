package pkg

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Subtitle struct {
	Number int
	Start  time.Duration
	End    time.Duration
	Text   string
}

type WriteSrt struct {
	file *os.File
}

func NewWriteSrt(fileName string) (*WriteSrt, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	return &WriteSrt{file: file}, err
}

func (w *WriteSrt) Add(sub Subtitle) error {
	w.file.WriteString(strconv.Itoa(sub.Number))
	w.file.WriteString("\n")
	w.file.WriteString(fmt.Sprintf("%s --> %s", formatTime(sub.Start), formatTime(sub.End)))
	w.file.WriteString("\n")
	w.file.WriteString(sub.Text)
	_, err := w.file.WriteString("\n\n")
	return err
}

func (w *WriteSrt) Close() {
	w.file.Close()
}

func formatTime(duration time.Duration) string {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	milliseconds := duration.Milliseconds() % 1000
	return fmt.Sprintf("%02d:%02d:%02d,%03d", hours, minutes, seconds, milliseconds)
}
