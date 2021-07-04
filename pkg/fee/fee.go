package fee

var (
	feePairs = []feePair{
		{
			minVolume:     0.0,
			maxVolume:     500.00,
			feePercentage: 0.0085,
		},
		{
			minVolume:     500.00,
			maxVolume:     1000.00,
			feePercentage: 0.0083,
		},
		{
			minVolume:     1000.00,
			maxVolume:     3000.00,
			feePercentage: 0.0080,
		},
		{
			minVolume:     3000.00,
			maxVolume:     9000.00,
			feePercentage: 0.0075,
		},
		{
			minVolume:     9000.00,
			maxVolume:     18000.00,
			feePercentage: 0.0070,
		},
		{
			minVolume:     18000.00,
			maxVolume:     40000.00,
			feePercentage: 0.0065,
		},
		{
			minVolume:     40000.00,
			maxVolume:     60000.00,
			feePercentage: 0.0060,
		},
		{
			minVolume:     60000.00,
			maxVolume:     70000.00,
			feePercentage: 0.0055,
		},
		{
			minVolume:     70000.00,
			maxVolume:     80000.00,
			feePercentage: 0.0050,
		},
		{
			minVolume:     80000.00,
			maxVolume:     90000.00,
			feePercentage: 0.0045,
		},
		{
			minVolume:     90000.00,
			maxVolume:     115000.00,
			feePercentage: 0.0040,
		},
		{
			minVolume:     115000.00,
			maxVolume:     125000.00,
			feePercentage: 0.0035,
		},
		{
			minVolume:     125000.00,
			maxVolume:     200000.00,
			feePercentage: 0.0030,
		},
		{
			minVolume:     200000.00,
			maxVolume:     400000.00,
			feePercentage: 0.0025,
		},
		{
			minVolume:     400000.00,
			maxVolume:     650000.00,
			feePercentage: 0.0023,
		},
		{
			minVolume:     650000.00,
			maxVolume:     850000.00,
			feePercentage: 0.0020,
		},
		{
			minVolume:     850000.00,
			maxVolume:     1000000.00,
			feePercentage: 0.0018,
		},
		{
			minVolume:     1000000.00,
			maxVolume:     3000000.00,
			feePercentage: 0.0015,
		},
		{
			minVolume:     3000000.00,
			maxVolume:     5000000.00,
			feePercentage: 0.0013,
		},
		{
			minVolume:     5000000.00,
			maxVolume:     0.0,
			feePercentage: 0.0035,
		},
	}
)

// fees are based on 30 day volume. https://www.btcmarkets.net/fees
type feePair struct{
	minVolume     float64
	maxVolume     float64
	feePercentage float64
}

func getTradeFeePercentage(volume float64) float64 {
	for _, f := range feePairs {
		if volume > f.minVolume && f.maxVolume != 0.0 && volume <= f.maxVolume {
			return f.feePercentage
		}
	}

	// default to highest fee
	return 0.0085
}

func CalculateTradeFee(cost float64, volume float64) float64 {
	feePercentage := getTradeFeePercentage(volume)
	return cost * feePercentage
}
