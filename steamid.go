package steamid

import (
	"fmt"
	"strconv"
	"strings"
)

type SteamID uint64

const steamIDBase = 76561197960265728

// SteamID3 parses steam ids in the form '[U:1:22202]'
func SteamID3(steamString string) (SteamID, error) {
	split := strings.Split(steamString[:len(steamString)-1], ":")
	if len(split) != 3 {
		return 0, fmt.Errorf("invalid SteamID3")
	}

	w, err := strconv.Atoi(split[2])
	if err != nil {
		return 0, fmt.Errorf("invalid SteamID3: %w", err)
	}

	return SteamID(w + steamIDBase), nil
}

// SteamID32 parses steam ids in the form 'STEAM_0:0:11101'
func SteamID32(steamString string) (SteamID, error) {
	Y, err := strconv.Atoi(steamString[8:9])
	if err != nil {
		return 0, err
	}

	Z, err := strconv.Atoi(steamString[10:])
	if err != nil {
		return 0, err
	}
	i := int64((Z * 2) + steamIDBase + Y)

	return SteamID(i), nil
}

// SteamID64 parses steam ids in the form '76561197960287930'
func SteamID64(steamString string) (SteamID, error) {
	i, err := strconv.ParseInt(steamString, 10, 64)
	if err != nil {
		return 0, err
	}

	if i < steamIDBase {
		return 0, fmt.Errorf("SteamID64 is too small")
	}

	return SteamID(i), nil
}

// SteamID32String returns the steam id in the form 'STEAM_0:0:11101'
func (s SteamID) SteamID32String() string {
	s = s - steamIDBase
	remainder := s % 2
	s = s / 2
	return fmt.Sprintf("STEAM_0:%d:%d", remainder, s)
}

// SteamID64String returns the steam id in the form '76561197960287930'
func (s SteamID) SteamID64String() string {
	return strconv.FormatInt(int64(s), 10)
}

// SteamID3String returns the steam id in the form '[U:1:22202]'
func (s SteamID) SteamID3String() string {
	s = s - steamIDBase
	return fmt.Sprintf("[U:1:%d]", s)
}
