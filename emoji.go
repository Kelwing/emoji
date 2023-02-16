package emoji

import (
	"regexp"
	"strings"
)

//go:generate emoji_codegen

var shortCodeRe = regexp.MustCompile(`:([a-z]+[a-z_]*):`)

func Emoji(shortCode string) string {
	cp, ok := shortCodeToEmoji[shortCode]

	if !ok {
		return ""
	}

	return cp
}

func ShortCodes(emoji string) []string {
	sc, ok := emojiToShortCodes[emoji]

	if !ok {
		return nil
	}

	return sc
}

func Format(in string) string {
	return shortCodeRe.ReplaceAllStringFunc(in, func(s string) string {
		return Emoji(strings.Trim(s, ":"))
	})
}
