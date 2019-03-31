package test

import (
	"context"

	"github.com/yuhanfang/riot/apiclient"
	"github.com/yuhanfang/riot/constants/champion"
	"github.com/yuhanfang/riot/constants/queue"
	"github.com/yuhanfang/riot/constants/region"
)

type TestClient struct {
}

func (tc *TestClient) GetAllChampionMasteries(ctx context.Context, r region.Region, summonerID string) ([]apiclient.ChampionMastery, error) {
	return []apiclient.ChampionMastery{}, nil
}

func (tc *TestClient) GetChampionMastery(ctx context.Context, r region.Region, summonerID string, champ champion.Champion) (*apiclient.ChampionMastery, error) {
	return &apiclient.ChampionMastery{}, nil
}

func (tc *TestClient) GetChampionMasteryScore(ctx context.Context, r region.Region, summonerID string) (int, error) {
	return 0, nil
}

func (tc *TestClient) GetChampions(ctx context.Context, r region.Region) (*apiclient.ChampionList, error) {
	return &apiclient.ChampionList{}, nil
}

func (tc *TestClient) GetChampionByID(ctx context.Context, r region.Region, champ champion.Champion) (*apiclient.Champion, error) {
	return &apiclient.Champion{}, nil
}

func (tc *TestClient) GetChallengerLeague(context.Context, region.Region, queue.Queue) (*apiclient.LeagueList, error) {
	return &apiclient.LeagueList{}, nil
}

func (tc *TestClient) GetGrandmasterLeague(ctx context.Context, r region.Region, q queue.Queue) (*apiclient.LeagueList, error) {
	return &apiclient.LeagueList{}, nil
}

func (tc *TestClient) GetMasterLeague(context.Context, region.Region, queue.Queue) (*apiclient.LeagueList, error) {
	return &apiclient.LeagueList{}, nil
}

func (tc *TestClient) GetAllLeaguePositionsForSummoner(ctx context.Context, r region.Region, summonerID string) ([]apiclient.LeaguePosition, error) {
	return []apiclient.LeaguePosition{}, nil
}

func (tc *TestClient) GetLeagueByID(ctx context.Context, r region.Region, leagueID string) (*apiclient.LeagueList, error) {
	return &apiclient.LeagueList{}, nil
}

func (tc *TestClient) GetMatch(ctx context.Context, r region.Region, matchID int64) (*apiclient.Match, error) {
	return &apiclient.Match{}, nil
}

func (tc *TestClient) GetMatchTimeline(ctx context.Context, r region.Region, matchID int64) (*apiclient.MatchTimeline, error) {
	return &apiclient.MatchTimeline{}, nil
}

func (tc *TestClient) GetMatchlist(ctx context.Context, r region.Region, accountID string, opts *apiclient.GetMatchlistOptions) (*apiclient.Matchlist, error) {
	return &apiclient.Matchlist{}, nil
}

func (tc *TestClient) GetRecentMatchlist(ctx context.Context, r region.Region, accountID string) (*apiclient.Matchlist, error) {
	return &apiclient.Matchlist{}, nil
}

func (tc *TestClient) GetFeaturedGames(ctx context.Context, r region.Region) (*apiclient.FeaturedGames, error) {
	return &apiclient.FeaturedGames{}, nil
}

func (tc *TestClient) GetCurrentGameInfoBySummoner(ctx context.Context, r region.Region, summonerID string) (*apiclient.CurrentGameInfo, error) {
	return &apiclient.CurrentGameInfo{}, nil
}

func (tc *TestClient) GetByAccountID(ctx context.Context, r region.Region, accountID string) (*apiclient.Summoner, error) {
	return &apiclient.Summoner{}, nil
}

func (tc *TestClient) GetBySummonerName(ctx context.Context, r region.Region, name string) (*apiclient.Summoner, error) {
	return &apiclient.Summoner{}, nil
}

func (tc *TestClient) GetBySummonerPUUID(ctx context.Context, r region.Region, puuid string) (*apiclient.Summoner, error) {
	return &apiclient.Summoner{}, nil
}

func (tc *TestClient) GetBySummonerID(ctx context.Context, r region.Region, summonerID string) (*apiclient.Summoner, error) {
	return &apiclient.Summoner{}, nil
}

func (tc *TestClient) GetThirdPartyCodeByID(ctx context.Context, r region.Region, summonerID string) (string, error) {
	return "", nil
}
