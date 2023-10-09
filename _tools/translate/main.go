package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Model    string     `json:"model"`
	Messages []*Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message *Message
}

func main() {
	flag.Parse()
	src := flag.Arg(0)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, src, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	for _, g := range node.Comments {
		if strings.HasPrefix(g.Text(), "Copyright ") {
			continue
		}
		if strings.HasPrefix(g.Text(), "Output:") {
			continue
		}

		input := ""
		for _, c := range g.List {
			input += c.Text + "\n"
		}
		log.Println("input:", input)
		output, err := Translate(context.Background(), input)
		if err != nil {
			log.Println(err)
			output = input
		}
		log.Println("output:", output)

		list := []*ast.Comment{}
		pos := g.End()
		for output != "" {
			var c *ast.Comment
			if strings.HasPrefix(output, "//") {
				idx := strings.Index(output, "\n")
				if idx < 0 {
					c = &ast.Comment{
						Slash: pos,
						Text:  output,
					}
					output = ""
				} else {
					c = &ast.Comment{
						Slash: pos,
						Text:  output[:idx],
					}
					output = output[idx+1:]
				}
			} else if strings.HasPrefix(output, "/*") {
				idx := strings.Index(output, "*/")
				if idx < 0 {
					c = &ast.Comment{
						Slash: pos,
						Text:  output + "*/",
					}
					output = ""
				} else {
					c = &ast.Comment{
						Slash: pos,
						Text:  output[:idx+2],
					}
					output = output[idx+2:]
				}
			} else {
				idx := strings.Index(output, "\n")
				if idx < 0 {
					output = ""
				} else {
					output = output[idx+1:]
				}
			}
			if c != nil {
				c.Slash = pos
				c.Text = strings.TrimSpace(c.Text)
				list = append(list, c)
			}
		}
		g.List = list
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		log.Println(err)
	}

	if err := os.WriteFile(src, buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

func Translate(ctx context.Context, input string) (string, error) {
	// ChatGPT APIのエンドポイントURLとAPIキーを設定します
	endpoint := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("OPENAI_API_KEY")

	// リクエストのヘッダーを設定します
	headers := map[string]string{
		"Authorization": "Bearer " + apiKey,
		"Content-Type":  "application/json",
	}

	// リクエストボディを準備します
	requestBody, err := json.Marshal(&Request{
		Model: "gpt-3.5-turbo",
		Messages: []*Message{
			{Role: "user", Content: "The following English text is a comment of Go language. Please translate it into Japanese:\n\n" + input},
		},
	})
	if err != nil {
		return "", err
	}

	// HTTP POSTリクエストを作成します
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	// ヘッダーを追加します
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// HTTPリクエストを送信します
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込みます
	dec := json.NewDecoder(resp.Body)
	var response Response
	if err := dec.Decode(&response); err != nil {
		return "", err
	}
	if len(response.Choices) == 0 {
		return "", nil
	}
	return response.Choices[0].Message.Content, nil
}
