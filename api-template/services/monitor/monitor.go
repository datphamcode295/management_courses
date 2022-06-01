package monitor

import (
	"context"
	"time"

	"github.com/MStation-io/api/config"
	"github.com/MStation-io/api/consts"
	"github.com/MStation-io/api/repo"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

type tokenPrice struct {
	USD float64 `json:"usd"`
}
type pancakeInfoTokenData struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	PriceBNB string `json:"price_BNB"`
}
type pancakeInfoToken struct {
	UpdatedAt int64                `json:"updated_at"`
	Data      pancakeInfoTokenData `json:"data"`
}

type Monitor struct {
	store       repo.DBRepo
	repo        *repo.Repo
	cfg         config.Config
	client      *ethclient.Client
	tokenPrices map[string]tokenPrice
}

func NewMonitorService(cfg config.Config, store repo.DBRepo, repo *repo.Repo) MonitorService {
	ctx, cancel := context.WithTimeout(context.Background(), consts.DefaultContextTimeout*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, cfg.BscRPCEndPoint)
	if err != nil {
		zap.L().Fatal("cannot init client", zap.Error(err))
	}

	var monitor = Monitor{
		store:       store,
		repo:        repo,
		cfg:         cfg,
		client:      client,
		tokenPrices: map[string]tokenPrice{},
	}
	// go monitor.UpdatePricesInRoutine(consts.DefaultGetPriceInterval)
	// go monitor.UpdateCharactersInRoutine(consts.DefaultVentureTimeout)
	return &monitor
}

type MonitorService interface {
	// UnstuckCharacterVenture()
	// UpdateCharactersInRoutine(seconds int)
	// GetPriceOverUSD(tokenSymbol string) float64
	// UpdatePricesInRoutine(seconds int)
	// UpdateCurrencyTranslationPriceUSD()
	// UpdateTokenPriceMap()
}

// func (m *Monitor) UnstuckCharacterVenture() {
// 	characters, err := m.repo.Character.GetCharactersByStatusTime(m.store, consts.CharacterStatusVenture, consts.DefaultVentureTimeout)
// 	if err != nil {
// 		zap.L().Sugar().Errorf("m.repo.Character.GetCharactersByStatusTime(): %v", err)
// 		return
// 	}
// 	characterIds := []string{}
// 	for _, character := range characters {
// 		characterIds = append(characterIds, character.Id)
// 	}
// 	updateCharacterModel := model.Character{
// 		Status: consts.CharacterStatusIdle,
// 	}

// 	_, err = m.repo.Character.UpdateByIds(m.store, characterIds, updateCharacterModel)
// 	if err != nil {
// 		zap.L().Sugar().Errorf("m.repo.Character.GetCharactersByStatusTime(): %v", err)
// 		return
// 	}
// }

// func (m *Monitor) UpdateCharactersInRoutine(seconds int) {
// 	for {
// 		m.UnstuckCharacterVenture()
// 		time.Sleep(time.Second * time.Duration(seconds))
// 	}
// }

// func (m *Monitor) GetPriceOverUSD(address string) float64 {
// 	value, valueExists := m.tokenPrices[address]
// 	if valueExists {
// 		return value.USD
// 	}
// 	return 0
// }

// func (m *Monitor) UpdatePricesInRoutine(seconds int) {
// 	for {
// 		m.UpdateTokenPriceMap()
// 		m.UpdateCurrencyTranslationPriceUSD()
// 		time.Sleep(time.Second * time.Duration(seconds))
// 	}
// }

// func (m *Monitor) UpdateCurrencyTranslationPriceUSD() {
// 	currencies := []string{}
// 	prices := []float64{}

// 	currencyMap, err := m.repo.CurrencyTranslation.GetAllMap(m.store)
// 	if err != nil {
// 		zap.L().Sugar().Errorf("m.repo.CurrencyTranslation.GetAll(): %v", err)
// 		return
// 	}

// 	for _, currency := range currencyMap {
// 		priceUSD := m.GetPriceOverUSD(currency.CurrencyAddress)
// 		if currency.PriceUsd != priceUSD {
// 			currencies = append(currencies, currency.Currency)
// 			prices = append(prices, priceUSD)
// 		}
// 	}

// 	if len(currencies) > 0 {
// 		m.repo.CurrencyTranslation.BatchUpdatePriceUSD(m.store, currencyMap, currencies, prices)
// 	}
// }

// func (m *Monitor) UpdateTokenPriceMap() {
// 	var tokens []string
// 	tokenPriceMap := map[string]string{}

// 	currencies, err := m.repo.CurrencyTranslation.GetAll(m.store)
// 	if err != nil {
// 		zap.L().Sugar().Errorf("m.repo.CurrencyTranslation.GetAll(): %v", err)
// 		return
// 	}

// 	for _, currency := range currencies {
// 		tokens = append(tokens, currency.CurrencyAddress)
// 		client := &http.Client{}
// 		req, err := http.NewRequest("GET", "https://api.pancakeswap.info/api/v2/tokens/"+currency.CurrencyAddress, nil)
// 		if err != nil {
// 			zap.L().Sugar().Error("http.NewRequest()", zap.Error(err))
// 			return
// 		}

// 		resInfo, err := client.Do(req)
// 		if err != nil {
// 			zap.L().Sugar().Errorf("client.Do(): %v", err)
// 			return
// 		}
// 		defer resInfo.Body.Close()

// 		bodyBytesInfo, _ := ioutil.ReadAll(resInfo.Body)
// 		var resp pancakeInfoToken
// 		err = json.Unmarshal(bodyBytesInfo, &resp)
// 		if err != nil {
// 			zap.L().Sugar().Errorf("json.Unmarshal(): %v", err)
// 			return
// 		}

// 		tokenPriceMap[currency.CurrencyAddress] = resp.Data.Price
// 	}

// 	for _, token := range tokens {
// 		value, valueExists := tokenPriceMap[token]
// 		if valueExists {
// 			price, err := strconv.ParseFloat(value, 64)
// 			if err != nil {
// 				zap.L().Sugar().Errorf("strconv.ParseFloat(): %v", err)
// 				return
// 			}
// 			m.tokenPrices[token] = tokenPrice{USD: price}
// 		}
// 	}
// }
