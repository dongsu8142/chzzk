package chzzk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c ChzzkClient) Fetch(url string) ([]byte, error) {
	url = "https://api.chzzk.naver.com" + url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2049.0 Safari/537.36")

	authCookie := &http.Cookie{
		Name:  "NID_AUT",
		Value: c.NidAuth,
	}
	sessionCookie := &http.Cookie{
		Name:  "NID_SES",
		Value: c.NidSession,
	}
	req.AddCookie(authCookie)
	req.AddCookie(sessionCookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c ChzzkClient) Search(searchType string, keyword string) ([]byte, error) {
	params := url.Values{
		"keyword": {keyword},
	}
	return c.Fetch(fmt.Sprintf("/service/v1/search/%s?%s", searchType, params.Encode()))
}

func (c ChzzkClient) SearchChannels(keyword string) ([]ChannelData, error) {
	body, err := c.Search("channels", keyword)
	if err != nil {
		return nil, err
	}
	var result ChannelsResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Content.Data, nil
}

func (c ChzzkClient) SearchLives(keyword string) ([]LivesData, error) {
	body, err := c.Search("lives", keyword)
	if err != nil {
		return nil, err
	}
	var result LivesResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Content.Data, nil
}

func (c ChzzkClient) SearchVideos(keyword string) ([]VideosData, error) {
	body, err := c.Search("videos", keyword)
	if err != nil {
		return nil, err
	}
	var result VideosResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Content.Data, nil
}
