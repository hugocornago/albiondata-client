package client

import (
	"fmt"

	"github.com/ao-data/albiondata-client/log"
)

// Tipo 5: 12 (Hoz), 6 (Piedra), 23 (Mineral), 0 (Madera)
// 6: Ni idea
// 7: Tier
// 8: Coordenadas
// 10: Cargas
// 11: Enchant
type eventNewHarvestableObject struct {
	Tipo          int   `mapstructure:"5"`
	Tier          int   `mapstructure:"7"`
	Coordenadas   []int `mapstructure:"8"`
	Cargas        int   `mapstructure:"10"`
	Encantamiento int   `mapstructure:"11"`
}

func parseTipo(tipo int) string {
	switch tipo {
	case 0:
		return "Madera"
	case 6:
		return "Piedra"
	case 12:
		return "Hoz"
	case 23:
		return "Mineral"
	default:
		return fmt.Sprintf("Desconocido (%d)", tipo)
	}
}

func (event eventNewHarvestableObject) Process(state *albionState) {
	log.Infof("[RECURSO] Tipo: %v, Encantamiento: %v, Tier: %v, Coordenadas: (%v, %v), Cargas :%v", parseTipo(event.Tipo), event.Encantamiento, event.Tier, event.Coordenadas[0], event.Coordenadas[1], event.Cargas)
}