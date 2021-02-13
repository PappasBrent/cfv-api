package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
			cardName := "Abyss+Healer"

			var endpointPath string

			t.Run("cards", func(t *testing.T) {
				endpointPath = "/cards"
				url := fmt.Sprintf("%s%s%s", baseURL, apiPath, endpointPath)

				var expected string

				t.Run(fmt.Sprintf("Get cards named %q", cardName),
					func(t *testing.T) {
						expected = `[{"CardType":"Trigger Unit","Clan":"Shadow Paladin","Critical":1,"DesignIllus":"","Effect":"(You may only have up to four cards with \"HEAL\" in a deck.)","Flavor":"Not yet! I will not let you fall before we demolish the enemy!","Format":"Premium Standard","Grade":"0","Illust":"Azusa","IllustColor":"","Illust2":"","Illust3":"","Illust4":"","Illust5":"","ImageUrlEn":"https://static.wikia.nocookie.net/cardfight/images/7/79/BT04-053EN-C.jpg/revision/latest/scale-to-width-down/274?cb=20121214165408","ImageUrlJp":"https://static.wikia.nocookie.net/cardfight/images/b/b6/Abyss_Healer.png/revision/latest/scale-to-width-down/274?cb=20111029172311","ImaginaryGift":"","Italian":"Guaritrice dell'Abisso","Kana":"アビス・ヒーラー","Kanji":"","Korean":"","LimitationText":"","MangaIllust":"","Name":"Abyss Healer","Nation":"United Sanctuary","Note":"","OtherNames":"","Phonetic":"Abisu Hīrā","Power":5000,"Race":"Angel","RideSkill":"","Sets":["Booster Set 4: Eclipse of Illusionary Shadows"],"TournamentStatuses":{"En":"Unrestricted","Jp":"Unrestricted","Kr":"Unrestricted","Th":"Unrestricted","It":"Unrestricted"},"Shield":10000,"Skill":"Boost","Thai":"","Translation":"","TriggerEffect":"Heal / +5000"},{"CardType":"Trigger Unit","Clan":"Shadow Paladin","Critical":1,"DesignIllus":"Azusa / 天城望","Effect":"(You may only have up to four cards with \"HEAL\" in a deck.)","Flavor":"(V-TD04): Those with the will to fight will never give up!(V-BT04): I don't believe you wish to die like this!(V-BT06): Change the pain of your wounds into anger. And stand up once more!","Format":"Standard / Premium Standard","Grade":"0","Illust":"","IllustColor":"","Illust2":"","Illust3":"","Illust4":"","Illust5":"","ImageUrlEn":"https://static.wikia.nocookie.net/cardfight/images/9/90/V-TD04-015EN.png/revision/latest/scale-to-width-down/274?cb=20180915162720","ImageUrlJp":"https://static.wikia.nocookie.net/cardfight/images/a/a9/V-TD04-015.png/revision/latest/scale-to-width-down/274?cb=20180809140951","ImaginaryGift":"","Italian":"Guaritrice dell'Abisso","Kana":"アビス・ヒーラー","Kanji":"","Korean":"","LimitationText":"","MangaIllust":"","Name":"Abyss Healer","Nation":"United Sanctuary","Note":"","OtherNames":"","Phonetic":"Abisu Hīrā","Power":5000,"Race":"Angel","RideSkill":"","Sets":["V Booster Set 04: Vilest! Deletor","V Booster Set 06: Phantasmal Steed Restoration","V Extra Booster 12: Team Dragon's Vanity!","V Special Series 03: Start Deck Blaster Dark","V Trial Deck 04: Ren Suzugamori"],"TournamentStatuses":{"En":"Unrestricted","Jp":"Unrestricted","Kr":"Unrestricted","Th":"Unrestricted","It":"Unrestricted"},"Shield":20000,"Skill":"Boost","Thai":"","Translation":"","TriggerEffect":"Heal / +10000"}]`
						res, err := getHttpResponseString(fmt.Sprintf("%s?name=%s", url, cardName))
						if err != nil {
							t.Error(err)
						}
						if res != expected {
							t.Errorf("URL:%q\nResponse: %q\nExpected:%q", url, res, expected)
						}
					})
			})
		})
	})
}

func getHttpResponseString(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
