package ingest

type Match struct {
	Status int `json:"status"`
	Data   []struct {
		Metadata struct {
			Map              string `json:"map"`
			GameVersion      string `json:"game_version"`
			GameLength       int    `json:"game_length"`
			GameStart        int    `json:"game_start"`
			GameStartPatched string `json:"game_start_patched"`
			RoundsPlayed     int    `json:"rounds_played"`
			Mode             string `json:"mode"`
			SeasonID         string `json:"season_id"`
			Platform         string `json:"platform"`
			Matchid          string `json:"matchid"`
			Region           string `json:"region"`
			Cluster          string `json:"cluster"`
		} `json:"metadata"`
		Players struct {
			AllPlayers []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"all_players"`
			Red []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"red"`
			Blue []struct {
				Puuid              string `json:"puuid"`
				Name               string `json:"name"`
				Tag                string `json:"tag"`
				Team               string `json:"team"`
				Level              int    `json:"level"`
				Character          string `json:"character"`
				Currenttier        int    `json:"currenttier"`
				CurrenttierPatched string `json:"currenttier_patched"`
				PlayerCard         string `json:"player_card"`
				PlayerTitle        string `json:"player_title"`
				PartyID            string `json:"party_id"`
				SessionPlaytime    struct {
					Minutes      int `json:"minutes"`
					Seconds      int `json:"seconds"`
					Milliseconds int `json:"milliseconds"`
				} `json:"session_playtime"`
				Assets struct {
					Card struct {
						Small string `json:"small"`
						Large string `json:"large"`
						Wide  string `json:"wide"`
					} `json:"card"`
					Agent struct {
						Small    string `json:"small"`
						Full     string `json:"full"`
						Bust     string `json:"bust"`
						Killfeed string `json:"killfeed"`
					} `json:"agent"`
				} `json:"assets"`
				Behaviour struct {
					AfkRounds    int `json:"afk_rounds"`
					FriendlyFire struct {
						Incoming int `json:"incoming"`
						Outgoing int `json:"outgoing"`
					} `json:"friendly_fire"`
					RoundsInSpawn int `json:"rounds_in_spawn"`
				} `json:"behaviour"`
				Platform struct {
					Type string `json:"type"`
					Os   struct {
						Name    string `json:"name"`
						Version string `json:"version"`
					} `json:"os"`
				} `json:"platform"`
				AbilityCasts struct {
					CCast int `json:"c_cast"`
					QCast int `json:"q_cast"`
					ECast int `json:"e_cast"`
					XCast int `json:"x_cast"`
				} `json:"ability_casts"`
				Stats struct {
					Score     int `json:"score"`
					Kills     int `json:"kills"`
					Deaths    int `json:"deaths"`
					Assists   int `json:"assists"`
					Bodyshots int `json:"bodyshots"`
					Headshots int `json:"headshots"`
					Legshots  int `json:"legshots"`
				} `json:"stats"`
				Economy struct {
					Spent struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"spent"`
					LoadoutValue struct {
						Overall int `json:"overall"`
						Average int `json:"average"`
					} `json:"loadout_value"`
				} `json:"economy"`
				DamageMade     int `json:"damage_made"`
				DamageReceived int `json:"damage_received"`
			} `json:"blue"`
		} `json:"players"`
		Teams struct {
			Red struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
			} `json:"red"`
			Blue struct {
				HasWon     bool `json:"has_won"`
				RoundsWon  int  `json:"rounds_won"`
				RoundsLost int  `json:"rounds_lost"`
			} `json:"blue"`
		} `json:"teams"`
		Rounds []struct {
			WinningTeam string `json:"winning_team"`
			EndType     string `json:"end_type"`
			BombPlanted bool   `json:"bomb_planted"`
			BombDefused bool   `json:"bomb_defused"`
			PlantEvents struct {
				PlantLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"plant_location"`
				PlantedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"planted_by"`
				PlantSite              string `json:"plant_site"`
				PlantTimeInRound       int    `json:"plant_time_in_round"`
				PlayerLocationsOnPlant []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_plant"`
			} `json:"plant_events"`
			DefuseEvents struct {
				DefuseLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"defuse_location"`
				DefusedBy struct {
					Puuid       string `json:"puuid"`
					DisplayName string `json:"display_name"`
					Team        string `json:"team"`
				} `json:"defused_by"`
				DefuseTimeInRound       int `json:"defuse_time_in_round"`
				PlayerLocationsOnDefuse []struct {
					PlayerPuuid       string `json:"player_puuid"`
					PlayerDisplayName string `json:"player_display_name"`
					PlayerTeam        string `json:"player_team"`
					Location          struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
					ViewRadians float64 `json:"view_radians"`
				} `json:"player_locations_on_defuse"`
			} `json:"defuse_events"`
			PlayerStats []struct {
				AbilityCasts struct {
					CCasts int `json:"c_casts"`
					QCasts int `json:"q_casts"`
					ECasts int `json:"e_casts"`
					XCasts int `json:"x_casts"`
				} `json:"ability_casts"`
				PlayerPuuid       string        `json:"player_puuid"`
				PlayerDisplayName string        `json:"player_display_name"`
				PlayerTeam        string        `json:"player_team"`
				DamageEvents      []interface{} `json:"damage_events"`
				Damage            int           `json:"damage"`
				Bodyshots         int           `json:"bodyshots"`
				Headshots         int           `json:"headshots"`
				Legshots          int           `json:"legshots"`
				KillEvents        []interface{} `json:"kill_events"`
				Kills             int           `json:"kills"`
				Score             int           `json:"score"`
				Economy           struct {
					LoadoutValue int `json:"loadout_value"`
					Weapon       struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon  string `json:"display_icon"`
							KillfeedIcon string `json:"killfeed_icon"`
						} `json:"assets"`
					} `json:"weapon"`
					Armor struct {
						ID     string `json:"id"`
						Name   string `json:"name"`
						Assets struct {
							DisplayIcon string `json:"display_icon"`
						} `json:"assets"`
					} `json:"armor"`
					Remaining int `json:"remaining"`
					Spent     int `json:"spent"`
				} `json:"economy"`
				WasAfk        bool `json:"was_afk"`
				WasPenalized  bool `json:"was_penalized"`
				StayedInSpawn bool `json:"stayed_in_spawn"`
			} `json:"player_stats"`
		} `json:"rounds"`
	} `json:"data"`
}
