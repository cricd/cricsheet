package cricsheet

import (
	"reflect"
	"testing"
)

func TestGame_Flatten(t *testing.T) {
	game := Game{}
	err := game.Read("test.yaml")
	if err != nil {
		t.Errorf("Failed to open testing file - test.yaml")
	}
	tests := []struct {
		name       string
		inGame     Game
		wantEvents []Event
		wantErr    bool
	}{{
		name:   "Happy Case",
		inGame: game,
		wantEvents: []Event{
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 1, Batsman: "AC Gilchrist", NonStriker: "MJ Clarke", Bowler: "DR Tuffey", Runs: Runs{Batsman: 0, Extras: 1, Total: 1}, Wicket: Wicket{}, Extras: map[string]int{"wides": 1}}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 2, Batsman: "AC Gilchrist", NonStriker: "MJ Clarke", Bowler: "DR Tuffey", Runs: Runs{Batsman: 0, Extras: 1, Total: 1}, Wicket: Wicket{}, Extras: map[string]int{"legbyes": 1}}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 3, Batsman: "MJ Clarke", NonStriker: "AC Gilchrist", Bowler: "DR Tuffey", Runs: Runs{Batsman: 0, Extras: 0, Total: 0}, Wicket: Wicket{}, Extras: nil}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 4, Batsman: "MJ Clarke", NonStriker: "AC Gilchrist", Bowler: "DR Tuffey", Runs: Runs{Batsman: 1, Extras: 0, Total: 1}, Wicket: Wicket{}, Extras: nil}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 5, Batsman: "AC Gilchrist", NonStriker: "MJ Clarke", Bowler: "DR Tuffey", Runs: Runs{Batsman: 1, Extras: 0, Total: 1}, Wicket: Wicket{}, Extras: nil}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 6, Batsman: "MJ Clarke", NonStriker: "AC Gilchrist", Bowler: "DR Tuffey", Runs: Runs{Batsman: 6, Extras: 0, Total: 6}, Wicket: Wicket{}, Extras: nil}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 0, Ball: 7, Batsman: "MJ Clarke", NonStriker: "AC Gilchrist", Bowler: "DR Tuffey", Runs: Runs{Batsman: 0, Extras: 0, Total: 0}, Wicket: Wicket{Kind: "caught", Fielders: []string{"CD McMillan"}, PlayerOut: "MJ Clarke"}, Extras: nil}},
			Event{Info: game.Info, InningsNumber: 1, BattingTeam: "Australia",
				Delivery: Delivery{Over: 1, Ball: 1, Batsman: "A Symonds", NonStriker: "AC Gilchrist", Bowler: "KD Mills", Runs: Runs{Batsman: 0, Extras: 0, Total: 0}, Wicket: Wicket{}, Extras: nil}},
		},
		wantErr: false,
		// TODO: Add test cases.
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvents, err := tt.inGame.Flatten()
			if (err != nil) != tt.wantErr {
				t.Errorf("Game.Flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvents, tt.wantEvents) {
				t.Errorf("Game.Flatten() = %v, want %v", gotEvents, tt.wantEvents)
			}
		})
	}
}
