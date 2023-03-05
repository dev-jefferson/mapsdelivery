package route

import "errors"

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("id da rota não informado!")
	}
}
