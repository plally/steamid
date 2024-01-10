# steamid
Convert steamIDs between their types  
Supports
- steamID32
- SteamID3
- SteamID64

### Example
```go
steamID32, err := SteamID32("STEAM_0:0:450041545")
steamID64 = steamID32.SteamID64String()
```
