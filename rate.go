package corona_golang_api

import (
	"errors"
)

func caseRate(c *CurrentCoronaStatus) error {
	c.CaseFatalityRate=fatalityRate(&c.Confirmed , &c.Deaths)
	c.CaseRecoveryRate=recoveryRate(&c.Confirmed , &c.Recovered)
	if c.CaseRecoveryRate<0 {
		return errors.New("Recovery rate is less than zero")
	}else if c.CaseFatalityRate<0{
		return errors.New("Fatalitiy rate less than zero")
	}
	return nil
}

func fatalityRate(confirmed *Confirmed, deaths *Deaths) float64 {
	rate:=float64(deaths.Value)/float64(confirmed.Value)*100
	return rate
}

func recoveryRate(confirmed *Confirmed, recovered *Recovered) float64 {
	rate:=float64(recovered.Value)/float64(confirmed.Value)*100
	return rate
}