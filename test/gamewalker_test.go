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
	expectedAccountIDList := []string{""}

	if err != nil {
		t.Errorf("Error is nil")
	}
	if len(accountIDList) != len(expectedAccountIDList) {
		t.Errorf("Expected 1 item, got %d", len(accountIDList))
	}
	if accountIDList[0] != "" {
		t.Errorf("Got wrong value for accountIDList, wanted \"\", got %s", accountIDList[0])
	}
}
