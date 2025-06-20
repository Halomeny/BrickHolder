package services

import (
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"set-service/models"
)

type RebrickableClient struct {
	BaseURL    string
	APIKey     string
	HttpClient *http.Client
}

func NewRebrickableClient(apiKey string) *RebrickableClient {
	return &RebrickableClient{
		APIKey: apiKey,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Получить список наборов Lego (пример с ограничением 10)
func (c *RebrickableClient) GetLegoSets(pageSize int) (*models.SetsResponse, error) {
	url := fmt.Sprintf("https://rebrickable.com/api/v3/lego/sets/?page_size=%d", pageSize)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "key "+c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Rebrickable API returned status %d", resp.StatusCode)
	}

	var setsResponse models.SetsResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&setsResponse); err != nil {
		return nil, err
	}

	return &setsResponse, nil
}

// Получения списка всех серий
func FetchAllSeries(apiKey string) ([]models.Series, error) {
	var allSeries []models.Series
	url := "https://rebrickable.com/api/v3/lego/themes/"

	client := &http.Client{}

	for url != "" {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "key "+apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch series: %s", resp.Status)
		}

		var sr models.SeriesResponse
		if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
			return nil, err
		}

		allSeries = append(allSeries, sr.Results...)
		url = sr.Next // ссылка на следующую страницу или пустая строка, если последняя
	}
	return allSeries, nil
}

// Получения списка всех наборов
func FetchAllSets(apiKey string) ([]models.Set, error) {
	var allSets []models.Set
	url := "https://rebrickable.com/api/v3/lego/sets/"

	client := &http.Client{}

	for url != "" {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "key "+apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch set: %s", resp.Status)
		}

		var sr models.SetsResponse
		if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
			return nil, err
		}

		allSets = append(allSets, sr.Results...)
		url = sr.Next // ссылка на следующую страницу или пустая строка, если последняя
	}
	return allSets, nil
}


// Получения списка всех минифигурок
func FetchAllMinifigs(apiKey string) ([]models.Set, error) {
	var allMinigigs []models.Set
	
	
	return allMinigigs, nil
}
