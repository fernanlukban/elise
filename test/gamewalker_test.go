package test

import (
	"context"
	"testing"

	"github.com/yuhanfang/riot/constants/region"

	"github.com/fernanlukban/elise"
)

func TestGetAccountIDList(t *testing.T) {
	fakeClient := &TestClient{}
	playerWalker := elise.NewGameWalker(fakeClient)
	ctx := context.Background()
	const reg = region.NA1

	// Passing in 0 since we need placeholder value
	accountIDList, err := playerWalker.GetAccountIDList(ctx, reg, 0)
	expectedAccountIDList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	if err != nil {
		t.Errorf("Error is not nil")
	}
	if len(accountIDList) != len(expectedAccountIDList) {
		t.Errorf("Expected %d items, got %d", len(expectedAccountIDList), len(accountIDList))
	}
	for i, accountID := range accountIDList {
		if accountID != expectedAccountIDList[i] {
			t.Errorf("Got wrong value for accountID, wanted %s, got %s", expectedAccountIDList[i], accountID)
		}
	}
}
