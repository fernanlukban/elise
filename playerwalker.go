package elise

import (
	"context"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
)

// Walker is an interface to use for types that can
// walk through players
type Walker interface {
	GetGameIDList(accountID string) []int64
}

// PlayerWalker walks for players
type PlayerWalker struct {
	client apiclient.Client
}

// NewPlayerWalker returns a new playerwalker
func NewPlayerWalker(client apiclient.Client) *PlayerWalker {
	return &PlayerWalker{client: client}
}

// GetGameIDList walks to get list of games
func (pw *PlayerWalker) GetGameIDList(accountID string) ([]int64, error) {
	var gameIDList []int64
	matchlist, err := pw.client.GetMatchlist(context.Background(), region.NA1, accountID, nil)
	if err != nil {
		return nil, nil
	}
	for _, match := range matchlist.Matches {
		gameIDList = append(gameIDList, match.GameID)
	}
	return gameIDList, nil
}
