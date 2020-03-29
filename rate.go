package corona_golang_api

func caseRate(confirmed *Confirmed, recovered *Recovered, deaths *Deaths) (float64,float64) {
	return fatalityRate(confirmed,deaths),recoveryRate(confirmed,recovered)
}

func fatalityRate(confirmed *Confirmed, deaths *Deaths) float64 {
	rate:=float64(deaths.Value)/float64(confirmed.Value)*100
	return rate
}

func recoveryRate(confirmed *Confirmed, recovered *Recovered) float64 {
	rate:=float64(recovered.Value)/float64(confirmed.Value)*100
	return rate
}