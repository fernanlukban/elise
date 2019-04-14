package elise

import (
	"context"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/region"
)

// GameWalkerI is an interface to use for types that can
// walk through games
type GameWalkerI interface {
	GetPlayerIDList(ctx context.Context, region region.Region, matchID int64) ([]string, error)
}

// GameWalker walks over games and gets players
type GameWalker struct {
	client apiclient.Client
}

// NewGameWalker returns a new game walker
func NewGameWalker(client apiclient.Client) *GameWalker {
	return &GameWalker{client: client}
}

// GetAccountIDList gets the list of accountID's given a gameID
func (gw *GameWalker) GetAccountIDList(ctx context.Context, region region.Region, matchID int64) ([]string, error) {
	match, err := gw.client.GetMatch(ctx, region, matchID)
	if err != nil {
		return []string{""}, err
	}
	accountIDList := make([]string, 10)
	for i, participantIdentity := range match.ParticipantIdentities {
		accountID := participantIdentity.Player.AccountID
		accountIDList[i] = accountID
	}
	return accountIDList, nil
}
