package main

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

//go:embed emoji.gogo
var templates embed.FS

const emojiDataURL = "https://cdn.jsdelivr.net/npm/emojibase-data/en/shortcodes/emojibase.json"

func download(url string) (*EmojiData, error) {
	client := http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	emojiData, err := UnmarshalEmojiData(data)

	return &emojiData, err
}

func parseShortCode(datum *EmojiDatum) []string {
	if datum.String != nil {
		return []string{*datum.String}
	} else {
		return datum.StringArray
	}
}

func parseCodePoint(codePoint string) ([]rune, error) {
	parts := strings.Split(codePoint, "-")
	runes := make([]rune, 0)
	for _, p := range parts {
		intCode, err := strconv.ParseInt(p, 16, 64)
		if err != nil {
			return nil, err
		}
		runes = append(runes, rune(intCode))
	}

	return runes, nil
}

type Emoji struct {
	CodePoints []rune
	ShortCodes []string
}

func (e *Emoji) String() string {
	return string(e.CodePoints)
}

func (e *Emoji) Format() string {
	sb := &strings.Builder{}
	for _, cp := range e.CodePoints {
		sb.WriteString(fmt.Sprintf("\\U%08X", cp))
	}

	return sb.String()
}

type TemplateData struct {
	Emojis  []Emoji
	Command string
}

func main() {
	data, err := download(emojiDataURL)
	if err != nil {
		panic(err)
	}

	emojis := make([]Emoji, 0)

	for k, v := range *data {
		codePoints, err := parseCodePoint(k)
		if err != nil {
			fmt.Printf("invalid code point: %s\n", err.Error())
			continue
		}

		shortCodes := parseShortCode(v)

		emojis = append(emojis, Emoji{
			CodePoints: codePoints,
			ShortCodes: shortCodes,
		})
	}

	td := TemplateData{
		Emojis:  emojis,
		Command: strings.Join(os.Args, " "),
	}

	t, err := template.ParseFS(templates, "*")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("emoji_gen.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := t.Execute(f, td); err != nil {
		panic(err)
	}
}
