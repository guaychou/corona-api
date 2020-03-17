package corona_golang_api

import (
	"encoding/json"
	"io/ioutil"
	 log "github.com/sirupsen/logrus"
	"net/http"
)

type Confirmed struct {
	Value int `json:"Value"`
}

type Recovered struct {
	Value int `json:"Value"`
}

type Deaths struct {
	Value int `json:"Value"`
}

type CurrentCoronaStatus struct {
	Confirmed `json:"confirmed"`
	Recovered `json:"recovered"`
    Deaths `json:"deaths"`
}

func GetCorona(country string) CurrentCoronaStatus{
	var ccs CurrentCoronaStatus
	url := "https://covid19.mathdro.id/api/countries/"+country
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
			log.Fatal(err)
		}
		jsonErr := json.Unmarshal(body, &ccs)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
			if jsonErr != nil {
				log.Fatal(jsonErr)
			}
			return ccs
	} else {
		ccs.Confirmed.Value=-1
		ccs.Deaths.Value=-1
		ccs.Recovered.Value=-1
		return ccs
	}
}