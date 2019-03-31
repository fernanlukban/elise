package test

import (
	"context"
	"testing"

	"github.com/yuhanfang/riot/constants/region"

	"github.com/fernanlukban/elise"
)

func TestGetGameIDList(t *testing.T) {
	fakeClient := &TestClient{}
	playerWalker := elise.NewPlayerWalker(fakeClient)
	resultGameIDList := []int64{0, 1, 2}
	ctx := context.Background()
	const region = region.NA1

	// Passing in 0 since the fakeClient doesn't use accountID to
	// determine the gameIDList
	gameIDList, err := playerWalker.GetGameIDList(ctx, region, "0", nil)
	if err != nil {
		t.Errorf("Got error: %s after calling playerWalker.GetGameIDList(\"0\")", err)
	}
	if len(resultGameIDList) != len(gameIDList) {
		t.Errorf("Size of expected resultGameIDList and actual gameIDList not the same %d != %d", len(resultGameIDList), len(gameIDList))
	}
	for i := 0; i < len(resultGameIDList); i++ {
		if resultGameIDList[i] != gameIDList[i] {
			t.Errorf("Element at index: %d is not the same between expected resultGameIDList and gameIDList %d != %d", i, resultGameIDList[i], gameIDList[i])
		}
	}
}

func TestGetGameIDListAsync(t *testing.T) {
	fakeClient := &TestClient{}
	playerWalker := elise.NewPlayerWalker(fakeClient)
	resultGameIDList := []int64{0, 1, 2}
	ctx := context.Background()
	const region = region.NA1

	gameIDList := make(chan int64, 4)
	gameIDListErr := make(chan error, 4)

	// Passing in 0 since the fakeClient doesn't use accountID to
	// determine the gameIDList
	go playerWalker.GetGameIDListAsync(ctx, region, "0", nil, gameIDList, gameIDListErr)
	gameIDFound := make(map[int64]bool)

outerLoop:
	for {
		select {
		case resultGameID, ok := <-gameIDList:
			if ok {
				gameIDFound[resultGameID] = true
			} else {
				break outerLoop
			}
		case err := <-gameIDListErr:
			t.Errorf("Got error: %s after calling playerWalker.GetGameIDList(\"0\")", err)
		}
	}
	// TODO: Maybe show the list and map?
	if len(resultGameIDList) != len(gameIDFound) {
		t.Errorf("Size of expected gameIDFound map and actual gameIDList not the same %d != %d", len(resultGameIDList), len(gameIDFound))
	}
	for _, gameID := range resultGameIDList {
		if !gameIDFound[gameID] {
			t.Errorf("Element: %d was not found", gameID)
		}
	}
}
