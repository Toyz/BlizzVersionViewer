package btapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func callBlizzTrack(call string, target interface{}) error {
	r, err := myClient.Get(fmt.Sprintf("https://blizztrack.com/api/%s", call))
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func AllGames() ([]Game, error) {
	var games []Game

	err := callBlizzTrack("all_games/json?mode=none", &games)

	if err != nil {
		return nil, err
	}

	return games, nil
}

func (channel Channel) Versions() ([]RegionInfo, error) {
	var info []RegionInfo

	err := callBlizzTrack(fmt.Sprintf("%s/versions/json", strings.ToLower(channel.Code)), &info)

	if err != nil {
		return nil, err
	}

	return info, nil
}

func (channel Channel) PatchNotes(page, size int) (PatchNotes, error) {
	if len(channel.NotesCode) <= 0 {
		return PatchNotes{}, errors.New(fmt.Sprintf("%s doesn't support patch notes", channel.Name))
	}

	var notes PatchNotes

	err := callBlizzTrack(fmt.Sprintf("notes/%s/%d/%d/json", strings.ToLower(channel.Code), page, size), &notes)

	if err != nil {
		return PatchNotes{}, err
	}

	return notes, nil
}
