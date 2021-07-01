package shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.cyclingnews.com/races/tour-de-france-2021/stage-4/results/"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://arstechnica.com/gadgets/2021/06/the-linux-foundation-is-working-to-improve-voice-recognition-ethics/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://www.phoronix.com/scan.php?page=news_item&px=Linux-5.13-Features"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "jTa4L57P")
	assert.Equal(t, shortLink_2, "d66yfx7N")
	assert.Equal(t, shortLink_3, "dhZTayYQ")
}
