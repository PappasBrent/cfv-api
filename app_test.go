package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"cfv-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestLoadDSN(t *testing.T) {

	if _, err := config.LoadDSN(); err != nil {
		t.Error(err)
	}

	envKeys := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_DB", "DB_PORT"}

	for _, key := range envKeys {
		if _, exists := os.LookupEnv(key); exists == false {
			t.Errorf("Error looking up environment variable %s", key)
		}
	}
}

func TestConnection(t *testing.T) {

	dsn, err := config.LoadDSN()

	if err != nil {
		t.Error(err)
	}

	dialector := postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})

	if _, err := gorm.Open(dialector, &gorm.Config{}); err != nil {
		t.Error(err)
	}
}

func TestApp(t *testing.T) {
	// These test cases are nested since they have layered dependencies
	app, err := config.SetupApp()

	if err != nil {
		t.Error(err)
	}

	t.Run("api", func(t *testing.T) {
		ts := httptest.NewServer(app)
		defer ts.Close()

		// Have to set the port of the URL manually
		// TODO: Use regex to read and replace the default port
		baseURL := strings.Replace(ts.URL, "27751", "8080", 1)

		var apiPath string

		t.Run("v1", func(t *testing.T) {
			apiPath = "/api/v1"

			var endpointPath string

			t.Run("cards", func(t *testing.T) {
				endpointPath = "/cards"
				reqURL := fmt.Sprintf("%s%s%s", baseURL, apiPath, endpointPath)
				cardName := "Abyss Healer"

				t.Run(fmt.Sprintf("Get cards named %q", cardName),
					func(t *testing.T) {
						expected, err := getExpectedFromFile("./test-jsons/cards-named-abyss-healer.json")
						if err != nil {
							t.Error()
						}
						res, err := getHttpResponse(fmt.Sprintf("%s?name=%s", reqURL, url.PathEscape(cardName)))
						if err != nil {
							t.Error(err)
						}
						if bytes.Compare(res, expected) != 0 {
							t.Errorf("URL:%q\nResponse: %q\nExpected:%q", reqURL, res, expected)
						}
					})

			})

			t.Run("sets", func(t *testing.T) {
				endpointPath = "/sets"
				reqURL := fmt.Sprintf("%s%s%s", baseURL, apiPath, endpointPath)

				t.Run("Get all sets",
					func(t *testing.T) {
						expected, err := getExpectedFromFile("./test-jsons/all-sets.json")
						if err != nil {
							t.Error()
						}
						res, err := getHttpResponse(reqURL)
						if err != nil {
							t.Error(err)
						}
						if bytes.Compare(res, expected) != 0 {
							t.Errorf("URL:%q\nResponse: %q\nExpected:%q", reqURL, res, expected)
						}
					})

			})

			t.Run("set", func(t *testing.T) {
				endpointPath = "/set"
				reqURL := fmt.Sprintf("%s%s%s", baseURL, apiPath, endpointPath)
				setName := "Burst Deck"

				t.Run("Get all cards in set",
					func(t *testing.T) {
						expected, err := getExpectedFromFile("./test-jsons/cards-in-burst-deck.json")
						if err != nil {
							t.Error()
						}
						res, err := getHttpResponse(fmt.Sprintf("%s?name=%s", reqURL, url.PathEscape(setName)))
						if err != nil {
							t.Error(err)
						}
						if bytes.Compare(res, expected) != 0 {
							t.Errorf("URL:%q\nResponse: %q\nExpected:%q", reqURL, res, expected)
						}
					})
			})
		})
	})
}

func getExpectedFromFile(filepath string) ([]byte, error) {
	expected, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open testing file %q", filepath))
	}
	return expected, nil
}

func getHttpResponse(reqURL string) ([]byte, error) {
	res, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
