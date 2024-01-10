package steamid

import (
	"testing"
)

func TestSteamID64(t *testing.T) {
	data := []struct {
		steamid  string
		expected int64
	}{
		{"STEAM_0:0:450041545", 76561198860348818},
	}

	for _, v := range data {
		out, err := SteamID32(v.steamid)
		if err != nil {
			t.Fatal(err)
		}
		if int64(out) != v.expected {
			t.Fatalf("Expected %d, got %d", v.expected, out)
		}
	}
}
