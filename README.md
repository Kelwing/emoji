# emoji

A tiny auto-generated Go library for working with Emojis.  Provides the ability to format strings containing shortcodes, convert shortcodes to the corresponding emojis, and convert an emoji to a list of shortcodes.

Data extracted from [Emojibase](https://emojibase.dev/).

## Examples

```go
fmt.Println(emoji.Emoji("wave"))

Output: 👋
```

```go
fmt.Println(emoji.Shortcodes("👋"))

Output: ["wave", "waving_hand"]
```

```go
fmt.Println(emoji.Format("Hello! :wave:"))

Output: Hello! 👋
```