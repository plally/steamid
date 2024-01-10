package steamid

import (
	"testing"
)

func TestSteamID32(t *testing.T) {
	data := []struct {
		steamid32 string
		steamid64 string
		steamid3  string
		expected  int64
	}{
		{"STEAM_0:0:450041545", "76561198860348818", "[U:1:900083090]", 76561198860348818},
		{"STEAM_0:0:11101", "76561197960287930", "[U:1:22202]", 76561197960287930},
	}

	for _, v := range data {
		t.Run("SteamID32", func(t *testing.T) {
			out, err := SteamID32(v.steamid64)
			if err != nil {
				t.Fatal(err)
			}
			if int64(out) != v.expected {
				t.Fatalf("Expected %d, got %d", v.expected, out)
			}
		})

		t.Run("SteamID64", func(t *testing.T) {
			out, err := SteamID64(v.steamid64)
			if err != nil {
				t.Fatal(err)
			}
			if int64(out) != v.expected {
				t.Fatalf("Expected %d, got %d", v.expected, out)
			}
		})

		t.Run("SteamID3", func(t *testing.T) {
			out, err := SteamID3(v.steamid3)
			if err != nil {
				t.Fatal(err)
			}
			if int64(out) != v.expected {
				t.Fatalf("Expected %d, got %d", v.expected, out)
			}
		})
	}
}
