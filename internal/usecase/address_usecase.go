package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/alonsofritz/cep-service-go/internal/entity"
	"github.com/alonsofritz/cep-service-go/internal/repository"
)

type AddressUseCase struct {
	repo *repository.AddressRepository
}

var providers = []string{
	"https://viacep.com.br/ws/%s/json/",
	"https://brasilapi.com.br/api/cep/v1/%s",
	"https://api.postmon.com.br/v1/cep/%s",
}

func NewAddressUseCase(repo *repository.AddressRepository) *AddressUseCase {
	return &AddressUseCase{repo: repo}
}

func (uc *AddressUseCase) GetAddress(cep string) (*entity.Address, error) {
	cep = strings.ReplaceAll(cep, "-", "")
	if addr, err := uc.repo.FetchFromDB(cep); err == nil {
		return addr, nil
	}
	for _, provider := range providers {
		if addr, err := uc.fetchFromProvider(cep, provider); err == nil {
			uc.repo.SaveToDB(*addr)
			return addr, nil
		}
	}
	return nil, fmt.Errorf("CEP not found")
}

func (uc *AddressUseCase) fetchFromProvider(cep, url string) (*entity.Address, error) {
	resp, err := http.Get(fmt.Sprintf(url, cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch from provider")
	}
	body, _ := io.ReadAll(resp.Body)
	var addr entity.Address
	if err = json.Unmarshal(body, &addr); err != nil {
		return nil, err
	}
	addr.CEP = cep
	addr.Provider = url
	addr.FetchedAt = time.Now()
	return &addr, nil
}
