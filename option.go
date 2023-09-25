package parser

type Option func(*EstimateParser)

// DetailedEstimateSheetName ...
func DetailedEstimateSheetName(name string) Option {
	return func(ep *EstimateParser) {
		ep.detailedEstimateSheetName = name
	}
}

// SimplifiedEstimateSheetName ...
func SimplifiedEstimateSheetName(name string) Option {
	return func(ep *EstimateParser) {
		ep.simplifiedEstimateSheetName = name
	}
}
