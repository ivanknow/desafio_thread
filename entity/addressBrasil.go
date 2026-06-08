package entity

import "fmt"

type BrasilAPIAddress struct {
    CEP          string `json:"cep"`
    State        string `json:"state"`
    City         string `json:"city"`
    Neighborhood string `json:"neighborhood"`
    Street       string `json:"street"`
    Service      string `json:"service"`
}

func (a BrasilAPIAddress) Print() {
    fmt.Printf("CEP: %s\n", a.CEP)
    fmt.Printf("Street: %s\n", a.Street)
    fmt.Printf("Neighborhood: %s\n", a.Neighborhood)
    fmt.Printf("City: %s\n", a.City)
    fmt.Printf("State: %s\n", a.State)
    fmt.Printf("Service: %s\n", a.Service)
}
