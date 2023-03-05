package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat  float64
	Long float64
}

type PartialRoutePosition struct {
	ID       string    `json: routeId`
	ClientID string    `json: clientId`
	Position []float64 `json: position`
	Finished bool      `json: finished`
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("id da rota não informado!")
	}

	//criando variavel onde captura o arquivo
	f, err := os.Open("destinations/" + r.ID + ".txt")

	if err != nil {
		return err
	}

	//aguarda tudo da função ser executado pra executar por ultimo o fechamento do arquivo
	defer f.Close()

	// variavel para ler o conteudo da variavel f
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		// Adiciona os valores de lat e long do arquivo para a lista de posições
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})

	}
	return nil
}
