package kata

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGame_GetShortestPath(t *testing.T) {
	f, err := os.Open("../data/words.txt")
	if err != nil {
		panic(err)
	}

	game, err := NewGame(f)

	if err != nil {
		panic(err)
	}

	sp := game.GetShortestPath("cat", "dog")
	assert.Equal(t, sp, []string{"cat", "cag", "cog", "dog"})

	sp = game.GetShortestPath("lead", "gold")
	assert.Equal(t, sp, []string{"lead", "load", "goad", "gold"})

	sp = game.GetShortestPath("ruby", "code")
	assert.Equal(t, sp, []string{"ruby", "roby", "robe", "cobe", "code"})
}
