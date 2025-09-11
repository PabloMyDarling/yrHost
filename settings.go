package main

/*
    yrHost - a multi-service home server
    Copyright (C) 2025  PabloMyLove

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserPreferences struct {
	Port      int      `json:"port"`
	Users     []User   `json:"users"`
	Blacklist []string `json:"ip-blacklist"`
	Whitelist []string `json:"ip-whitelist"`
	Services  []string `json:"services"`
	YrFiles   struct {
		SavePath string `json:"save-path"`
	} `json:"yrFiles"`
	YrSound struct {
		SavePath       string `json:"save-path"`
		ArtistPictures []struct {
			Name string `json:"name"`
			Path string `json:"path"`
		} `json:"artist-pictures"`
	} `json:"yrSound"`
}

func get_datentime() string {
	var Time time.Time = time.Now()
	return Time.Format("2006-01-02 15:04:05")
}
func get_settings() UserPreferences {
	var settings UserPreferences
	var config_json, _ = os.Open(filepath.Join(filePath, "config.json"))
	defer config_json.Close()

	json.NewDecoder(config_json).Decode(&settings)
	return settings
}

