package entity

import "fmt"

type ViaCEPAddress struct {
    CEP          string `json:"cep"`
    Street       string `json:"logradouro"`
    Complement   string `json:"complemento"`
    Unit         string `json:"unidade"`
    Neighborhood string `json:"bairro"`
    City         string `json:"localidade"`
    State        string `json:"uf"`
    Estado       string `json:"estado"`
    Region       string `json:"regiao"`
    IBGE         string `json:"ibge"`
    GIA          string `json:"gia"`
    DDD          string `json:"ddd"`
    SIAFI        string `json:"siafi"`
}

func (a ViaCEPAddress) Print() {
    fmt.Printf("CEP: %s\n", a.CEP)
    fmt.Printf("Street: %s\n", a.Street)
    fmt.Printf("Complement: %s\n", a.Complement)
    fmt.Printf("Unit: %s\n", a.Unit)
    fmt.Printf("Neighborhood: %s\n", a.Neighborhood)
    fmt.Printf("City: %s\n", a.City)
    fmt.Printf("State: %s\n", a.State)
    fmt.Printf("Estado: %s\n", a.Estado)
    fmt.Printf("Region: %s\n", a.Region)
    fmt.Printf("IBGE: %s\n", a.IBGE)
    fmt.Printf("GIA: %s\n", a.GIA)
    fmt.Printf("DDD: %s\n", a.DDD)
    fmt.Printf("SIAFI: %s\n", a.SIAFI)
}
