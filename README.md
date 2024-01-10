[![Go Reference](https://pkg.go.dev/badge/github.com/plally/steamid.svg)](https://pkg.go.dev/github.com/plally/steamid)
[![Go Report Card](https://goreportcard.com/badge/github.com/plally/steamid)](https://goreportcard.com/report/github.com/plally/steamid)
# steamid
Convert steamIDs between their types  
Supports
- steamID32
- SteamID3
- SteamID64

### Example
```go
steamID32, err := SteamID32("STEAM_0:0:12345")
steamID64 = steamID32.SteamID64String()
```
