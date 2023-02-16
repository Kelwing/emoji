package emoji

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmoji(t *testing.T) {
	e := Emoji("santa")
	require.Equal(t, "🎅", e)
}

func TestShortCodes(t *testing.T) {
	scs := ShortCodes("🪩")
	require.Contains(t, scs, "disco")
	require.Contains(t, scs, "disco_ball")
	require.Contains(t, scs, "mirror_ball")
}

func TestFormat(t *testing.T) {
	s := Format("Hello :wave:")
	require.Equal(t, "Hello \U0001F44B", s)
}
