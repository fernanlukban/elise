package elise

import (
	"context"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
)

// PlayerWalkerI is an interface to use for types that can
// walk through players
type PlayerWalkerI interface {
	GetGameIDList(ctx context.Context, region region.Region, accountID string, opts *apiclient.GetMatchlistOptions) ([]int64, error)
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
func (pw *PlayerWalker) GetGameIDList(ctx context.Context, region region.Region, accountID string, opts *apiclient.GetMatchlistOptions) ([]int64, error) {
	var gameIDList []int64
	matchlist, err := pw.client.GetMatchlist(ctx, region, accountID, nil)
	if err != nil {
		return nil, err
	}
	for _, match := range matchlist.Matches {
		gameIDList = append(gameIDList, match.GameID)
	}
	return gameIDList, nil
}

// GetGameIDListAsync is the async version of GetGameIDList
func (pw *PlayerWalker) GetGameIDListAsync(ctx context.Context, region region.Region, accountID string, opts *apiclient.GetMatchlistOptions, gameIDList chan int64, errChannel chan error) {
	matchlist, err := pw.client.GetMatchlist(ctx, region, accountID, nil)
	if err != nil {
		errChannel <- err
		return
	}
	for _, match := range matchlist.Matches {
		gameIDList <- match.GameID
	}
	close(gameIDList)
}
