package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/snowie2000/livetv/global"
)

func TXTGenerate() (string, error) {
	baseUrl, err := global.GetConfig("base_url")
	if err != nil {
		log.Println(err)
		return "", err
	}
	channels, err := GetAllChannel()
	if err != nil {
		log.Println(err)
		return "", err
	}
	var txt strings.Builder
	groupedChannels := make(map[string]map[string][]string)

	for _, v := range channels {
		category := "LiveTV"
		if v.Category != "" {
			category = v.Category
		}
		url := fmt.Sprintf("%s/live.m3u8?token=%s&c=%d", baseUrl, v.Token, v.ID)
		if groupedChannels[category] == nil {
			groupedChannels[category] = make(map[string][]string)
		}
		groupedChannels[category][v.Name] = append(groupedChannels[category][v.Name], url)
	}

	for category, channels := range groupedChannels {
		txt.WriteString(fmt.Sprintf("%s,#genre#\n", category))
		for name, urls := range channels {
			txt.WriteString(fmt.Sprintf("%s,%s\n", name, strings.Join(urls, " #")))
		}
	}

	return txt.String(), nil
}
