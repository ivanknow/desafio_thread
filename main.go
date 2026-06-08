package main

import (
	"context"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "os"
    "strings"
    "time"
    "desafio_thread/entity"
)

type apiResult struct {
    API   string
    Addr  entity.AddressPrinter
    Error error
}

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Usage: go run main.go CEP_NUMBER ")
        os.Exit(1)
    }

    cep := validateCEP(os.Args[1])
    if cep == "" {
        fmt.Fprintln(os.Stderr, "Invalid CEP. Use 8 digits")
        os.Exit(1)
    }

   

    result := getAddrByCEP(cep)
    if result.Error != nil {
        fmt.Fprintln(os.Stderr, "error:", result.Error)
        os.Exit(1)
    }

    printResult(result)
}

func validateCEP(raw string) string {
    cep := strings.TrimSpace(raw)
    cep = strings.ReplaceAll(cep, "-", "")
    if len(cep) != 8 {
        return ""
    }
    for _, ch := range cep {
        if ch < '0' || ch > '9' {
            return ""
        }
    }
    return cep
}

func getAddrByCEP( cep string) apiResult {
    resultCh := make(chan apiResult, 2)

    go fetchBrasilAPI( cep, resultCh)
    go fetchViaCEP( cep, resultCh)

    select {
    case res := <-resultCh:
        return res
    case <- time.After(1 * time.Second):
        return apiResult{Error: errors.New("timeout: no API responded within 1 second")}
    }
}

func fetchBrasilAPI( cep string, out chan<- apiResult) {
    url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
    var addr entity.BrasilAPIAddress
    if err := doRequest(url, &addr); err != nil {
        sendResult(out, apiResult{API: "BrasilAPI", Error: fmt.Errorf("BrasilAPI: %w", err)})
        return
    }
    sendResult(out, apiResult{API: "BrasilAPI", Addr: addr})
}

func fetchViaCEP( cep string, out chan<- apiResult) {
    url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
    var addr entity.ViaCEPAddress
    if err := doRequest(url, &addr); err != nil {
        sendResult(out, apiResult{API: "ViaCEP", Error: fmt.Errorf("ViaCEP: %w", err)})
        return
    }
    if addr.CEP == "" {
        sendResult(out, apiResult{API: "ViaCEP", Error: errors.New("ViaCEP returned empty result")})
        return
    }
    sendResult(out, apiResult{API: "ViaCEP", Addr: addr})
}

func doRequest(url string, v any) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
    if err != nil {
        cancel()
        return err
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        cancel()
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("status %d", resp.StatusCode)
    }

    if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
        return err
    }
    return nil
}

func sendResult(out chan<- apiResult, res apiResult) {
    select {
    case out <- res:
    default:
    }
}

func printResult(res apiResult) {
    fmt.Printf("API: %s\n", res.API)
    if res.Addr != nil {
        res.Addr.Print()
    }
}
