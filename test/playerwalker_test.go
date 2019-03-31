package test

import (
	"testing"

	"github.com/fernanlukban/elise"
)

func TestGetGameIDList(t *testing.T) {
	fakeClient := &TestClient{}
	playerWalker := elise.NewPlayerWalker(fakeClient)
	resultGameIDList := []int64{0, 1, 2}

	// Passing in 0 since the fakeClient doesn't use accountID to
	// determine the gameIDList
	gameIDList, err := playerWalker.GetGameIDList("0")
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
