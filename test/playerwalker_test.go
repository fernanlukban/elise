package test

import (
	"testing"

	"github.com/fernanlukban/elise"
)

// Stub test
func TestGetGameIDList(t *testing.T) {
	fakeClient := &TestClient{}
	playerWalker := elise.NewPlayerWalker(fakeClient)
	gameIDList := playerWalker.GetGameIDList(0)
	if gameIDList[0] != []int64{0}[0] {
		t.Errorf("Want []int64{0}, got %d", gameIDList)
	}
}
