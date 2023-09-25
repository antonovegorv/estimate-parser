package parser

const (
	// ConfigEstimateProject ...
	ConfigEstimateProject = "ESTIMATE.PROJECT"

	// ConfigEstimateAuthor ...
	ConfigEstimateAuthor = "ESTIMATE.AUTHOR"

	// ConfigEstimateCreateDt ...
	ConfigEstimateCreateDt = "ESTIMATE.CREATE_DT"

	// ConfigEstimateClientFullName ...
	ConfigEstimateClientFullName = "ESTIMATE.CLIENT.FULL_NAME"

	// ConfigEstimateClientPhone ...
	ConfigEstimateClientPhone = "ESTIMATE.CLIENT.PHONE"

	// ConfigEstimateClientEmail ...
	ConfigEstimateClientEmail = "ESTIMATE.CLIENT.EMAIL"

	// ConfigEstimateClientAddress ...
	ConfigEstimateClientAddress = "ESTIMATE.CLIENT.ADDRESS"

	// DefaultDetailedEstimateSheetName ...
	DefaultDetailedEstimateSheetName = "ПОДРОБНАЯ СМЕТА"

	// DefaultSimplifiedEstimateSheetName ...
	DefaultSimplifiedEstimateSheetName = "ПРОСТАЯ СМЕТА"

	// DefaultBlockStartCellValue ...
	DefaultBlockStartCellValue = "НАИМЕНОВАНИЕ"

	// DefaultBlockEndCellValue ...
	DefaultBlockEndCellValue = "Строительная компания «Приват-Строй»,   +7(495)7210366,    www.private-stroy.ru"
)

// DefaultEstimateConfig ...
var DefaultEstimateConfig = EstimateConfig{
	ConfigEstimateProject: EstimateConfigItem{
		sheet: DefaultSimplifiedEstimateSheetName,
		cell:  "C4",
	},
	ConfigEstimateAuthor: EstimateConfigItem{
		sheet: DefaultSimplifiedEstimateSheetName,
		cell:  "A7",
	},
	ConfigEstimateCreateDt: EstimateConfigItem{
		sheet: DefaultDetailedEstimateSheetName,
		cell:  "I9",
	},
	ConfigEstimateClientFullName: EstimateConfigItem{
		sheet: DefaultSimplifiedEstimateSheetName,
		cell:  "C1",
	},
	ConfigEstimateClientPhone: EstimateConfigItem{
		sheet: DefaultSimplifiedEstimateSheetName,
		cell:  "C3",
	},
	ConfigEstimateClientEmail: EstimateConfigItem{
		sheet: DefaultSimplifiedEstimateSheetName,
		cell:  "C3",
	},
	ConfigEstimateClientAddress: EstimateConfigItem{
		sheet: DefaultSimplifiedEstimateSheetName,
		cell:  "C2",
	},
}

// EstimateConfig ...
type EstimateConfig map[string]EstimateConfigItem

// EstimateConfigItem ...
type EstimateConfigItem struct {
	sheet string
	cell  string
}

// Sheet ...
func (ci *EstimateConfigItem) Sheet() string {
	return ci.sheet
}

// Cell ...
func (ci *EstimateConfigItem) Cell() string {
	return ci.cell
}
