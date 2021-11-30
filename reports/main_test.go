package reports

import (
	"testing"
)

func Test_player_setPlayerAge(t *testing.T) {
	type fields struct {
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "collect player age",
			fields: fields{
				Name: "Emile Smith-Rowe",
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player := &player{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			player.setPlayerAge()
			if player.Age != tt.want {
				t.Errorf("player.setPlayerAge() = %v, want %v", player.Age, tt.want)
			}
		})
	}
}
