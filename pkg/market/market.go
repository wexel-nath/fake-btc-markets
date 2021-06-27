package market

func NewMarket(
	baseAsset string,
	quoteAsset string,
	minOrderAmount string,
	maxOrderAmount string,
	amountDecimals string,
	priceDecimals string,
) (Market, error) {
	id := baseAsset + "-" + quoteAsset

	row, err := insertMarket(
		id,
		baseAsset,
		quoteAsset,
		minOrderAmount,
		maxOrderAmount,
		amountDecimals,
		priceDecimals,
	)
	if err != nil {
		return Market{}, err
	}

	return newMarketFromRow(row)
}

func GetMarketByID(id string) (Market, error) {
	row, err := selectMarket(id)
	if err != nil {
		return Market{}, err
	}

	return newMarketFromRow(row)
}

func GetMarkets() ([]Market, error) {
	rows, err := selectMarkets()
	if err != nil {
		return nil, err
	}

	return newMarketsFromRows(rows), nil
}
