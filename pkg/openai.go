package pkg

import (
	"context"
	"errors"
	"fmt"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

type Openai struct {
	client *openai.Client
}

func NewOpenai(token string, proxyUrl string) *Openai {
	config := openai.DefaultConfig(token)
	if proxyUrl != "" {
		config.BaseURL = proxyUrl
	}
	client := openai.NewClientWithConfig(config)
	return &Openai{client: client}
}

func (o *Openai) VideoToText(path string) {
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		Format:   openai.AudioResponseFormatSRT,
		FilePath: path,
	}
	resp, err := o.client.CreateTranslation(context.TODO(), req)
	fmt.Println(resp.Text, err)
}

func (o *Openai) TextToText(content string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo0125,
		MaxTokens: 100,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Translate into Chinese without any content other than the translated content:" + content,
			},
		},
		Stream: true,
	}

	stream, err := o.client.CreateChatCompletionStream(context.TODO(), req)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	var result string
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return result, nil
		}

		if err != nil {
			return "", err
		}

		result += response.Choices[0].Delta.Content
	}
}
