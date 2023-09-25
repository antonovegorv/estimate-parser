package parser

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
)

const (
	_defaultDetailedEstimateSheetName   = "ПОДРОБНАЯ СМЕТА"
	_defaultSimplifiedEstimateSheetName = "ПРОСТАЯ СМЕТА"
)

type blockInfo struct {
	start, end int
}

// EstimateParser ...
type EstimateParser struct {
	file *excelize.File
	rows [][]string

	detailedEstimateSheetName   string
	simplifiedEstimateSheetName string

	estimateConfig EstimateConfig
}

// NewEstimateParser ...
func NewEstimateParser(options ...Option) *EstimateParser {
	e := &EstimateParser{
		detailedEstimateSheetName:   _defaultDetailedEstimateSheetName,
		simplifiedEstimateSheetName: _defaultSimplifiedEstimateSheetName,
		estimateConfig:              DefaultEstimateConfig,
	}

	for _, opt := range options {
		opt(e)
	}

	return e
}

// ParseFromReader ...
func (ep *EstimateParser) ParseFromReader(r io.Reader) (*Estimate, error) {
	var err error

	ep.file, err = excelize.OpenReader(r)
	if err != nil {
		return nil, fmt.Errorf("failed to open reader: %w;", err)
	}

	ep.rows, err = ep.file.GetRows(ep.detailedEstimateSheetName)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows: %w;", err)
	}

	e, err := ep.getEstimate()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (ep *EstimateParser) getEstimate() (*Estimate, error) {
	var (
		err error
		e   = &Estimate{}
	)

	e.Project, err = ep.getEstimateProject()
	if err != nil {
		return nil, err
	}

	e.Author, err = ep.getEstimateAuthor()
	if err != nil {
		return nil, err
	}

	e.CreateDt, err = ep.getEstimateCreateDt()
	if err != nil {
		return nil, err
	}

	e.Client, err = ep.getEstimateClient()
	if err != nil {
		return nil, err
	}

	e.Blocks, err = ep.getEstimateBlocks()
	if err != nil {
		return nil, err
	}

	return e, nil
}

// getEstimateProject ...
func (ep *EstimateParser) getEstimateProject() (string, error) {
	v, ok := ep.estimateConfig[ConfigEstimateProject]
	if !ok {
		return "", fmt.Errorf("no key in estimate config: %v", ConfigEstimateProject)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(cell), nil
}

// getEstimateAuthor ...
func (ep *EstimateParser) getEstimateAuthor() (string, error) {
	v, ok := ep.estimateConfig[ConfigEstimateAuthor]
	if !ok {
		return "", fmt.Errorf("no key in estimate config: %v", ConfigEstimateProject)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(strings.TrimPrefix(cell, "Смету составил:")), nil
}

// getEstimateCreateDt ...
func (ep *EstimateParser) getEstimateCreateDt() (time.Time, error) {
	v, ok := ep.estimateConfig[ConfigEstimateCreateDt]
	if !ok {
		return time.Time{}, fmt.Errorf("no key in estimate config: %v", ConfigEstimateCreateDt)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse("01-02-06", cell)
}

func (ep *EstimateParser) getEstimateClient() (*Client, error) {
	var (
		err error
		c   = &Client{}
	)

	c.FullName, err = ep.getEstimateClientFullName()
	if err != nil {
		return nil, err
	}

	c.Phone, err = ep.getEstimateClientPhone()
	if err != nil {
		return nil, err
	}

	c.Email, err = ep.getEstimateClientEmail()
	if err != nil {
		return nil, err
	}

	c.Address, err = ep.getEstimateClientAddress()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// getEstimateClientFullName ...
func (ep *EstimateParser) getEstimateClientFullName() (string, error) {
	v, ok := ep.estimateConfig[ConfigEstimateClientFullName]
	if !ok {
		return "", fmt.Errorf("no key in estimate config: %v", ConfigEstimateProject)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(cell), nil
}

// getEstimateClientPhone ...
func (ep *EstimateParser) getEstimateClientPhone() (string, error) {
	var phone string

	v, ok := ep.estimateConfig[ConfigEstimateClientPhone]
	if !ok {
		return "", fmt.Errorf("no key in estimate config: %v", ConfigEstimateProject)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return "", err
	}

	phoneAndEmail := strings.Split(cell, ",")
	if len(phoneAndEmail) > 0 {
		phone = strings.TrimSpace(phoneAndEmail[0])

		// TODO: Add validation for the phone check.
	}

	return phone, nil
}

// getEstimateClientEmail ...
func (ep *EstimateParser) getEstimateClientEmail() (string, error) {
	var email string

	v, ok := ep.estimateConfig[ConfigEstimateClientPhone]
	if !ok {
		return "", fmt.Errorf("no key in estimate config: %v", ConfigEstimateProject)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return "", err
	}

	phoneAndEmail := strings.Split(cell, ",")
	if len(phoneAndEmail) > 1 {
		email = strings.TrimSpace(phoneAndEmail[1])

		// TODO: Add validation for the phone check.
	}

	return email, nil
}

// getEstimateClientAddress ...
func (ep *EstimateParser) getEstimateClientAddress() (string, error) {
	v, ok := ep.estimateConfig[ConfigEstimateClientAddress]
	if !ok {
		return "", fmt.Errorf("no key in estimate config: %v", ConfigEstimateProject)
	}

	cell, err := ep.file.GetCellValue(v.Sheet(), v.Cell())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(cell), nil
}

// getEstimateBlocks ...
func (ep *EstimateParser) getEstimateBlocks() ([]*Block, error) {
	blocksInfo, err := ep.getEstimateBlocksInfo()
	if err != nil {
		return nil, err
	}

	blocks := make([]*Block, 0, len(blocksInfo))

	for _, bi := range blocksInfo {
		blocks = append(blocks, &Block{})

		for i := bi.start; i <= bi.end; i++ {
			rt, err := ep.getEstimateRowType(i)
			if err != nil || rt == undefined {
				if rt == undefined {
					return nil, fmt.Errorf("undefine row type at %d", i)
				}

				return nil, err
			}

			switch rt {
			case title:
				err = ep.enrichBlocksWithTitle(blocks, i)
				if err != nil {
					return nil, err
				}
			case process:
				err = ep.enrichBlocksWithProcess(blocks, i)
				if err != nil {
					return nil, err
				}
			case material:
				err = ep.enrichBlocksWithMaterial(blocks, i)
				if err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("unexpected row type at %d", i)
			}
		}
	}

	return blocks, nil
}

// getEstimateBlocksInfo ...
func (ep *EstimateParser) getEstimateBlocksInfo() ([]blockInfo, error) {
	hasBlockStarted := false

	blocksInfo := make([]blockInfo, 0)

	for i := 0; i < len(ep.rows); i++ {
		row := ep.rows[i]

		if len(row) > 0 {
			if row[0] == DefaultBlockStartCellValue {
				if hasBlockStarted {
					return nil, fmt.Errorf("block has already started")
				}

				hasBlockStarted = true

				blocksInfo = append(blocksInfo, blockInfo{
					start: i + 2,
				})
			}

			if row[0] == DefaultBlockEndCellValue {
				if !hasBlockStarted {
					return nil, fmt.Errorf("block was not started")
				}

				hasBlockStarted = false

				blocksInfo[len(blocksInfo)-1].end = i - 1
			}
		}
	}

	// TODO: Add final validation for each block:
	// 1. blockInfo != nil;
	// 2. end >= start.

	return blocksInfo, nil
}

// getEstimateRowType ...
func (ep *EstimateParser) getEstimateRowType(row int) (RowType, error) {
	cell, err := excelize.CoordinatesToCellName(1, row+1)
	if err != nil {
		return undefined, err
	}

	styleID, err := ep.file.GetCellStyle(ep.detailedEstimateSheetName, cell)
	if err != nil {
		return undefined, err
	}

	style, err := ep.file.GetStyle(styleID)
	if err != nil {
		return undefined, err
	}

	if style != nil && style.Font != nil {
		switch {
		case style.Font.Bold && style.Font.Size == 14:
			return title, nil
		case style.Font.Bold && style.Font.Size == 11:
			return process, nil
		case style.Font.Italic:
			return material, nil
		}
	}

	return undefined, nil
}

// enrichBlocksWithTitle ...
func (ep *EstimateParser) enrichBlocksWithTitle(blocks []*Block, row int) error {
	mostRecentBlock := blocks[len(blocks)-1]

	mostRecentBlock.Title = strings.TrimSpace(ep.rows[row][0])

	return nil
}

// enrichBlocksWithProcess ...
func (ep *EstimateParser) enrichBlocksWithProcess(blocks []*Block, row int) error {
	mostRecentBlock := blocks[len(blocks)-1]

	name := strings.TrimSpace(ep.rows[row][0])
	unit := strings.TrimSpace(ep.rows[row][1])

	number, err := decimal.NewFromString(strings.ReplaceAll(ep.rows[row][2], ",", ""))
	if err != nil {
		// return err
	}

	workPrice, err := decimal.NewFromString(strings.ReplaceAll(ep.rows[row][3], ",", ""))
	if err != nil {
		// return err
	}

	process := Process{
		Name:      name,
		Unit:      unit,
		Number:    number,
		WorkPrice: workPrice,
	}

	mostRecentBlock.Processes = append(mostRecentBlock.Processes, &process)

	return nil
}

// enrichBlocksWithMaterial ...
func (ep *EstimateParser) enrichBlocksWithMaterial(blocks []*Block, row int) error {
	mostRecentBlock := blocks[len(blocks)-1]

	mostRecentProcess := mostRecentBlock.Processes[len(mostRecentBlock.Processes)-1]

	name := strings.TrimSpace(ep.rows[row][0])
	unit := strings.TrimSpace(ep.rows[row][1])

	number, err := decimal.NewFromString(strings.ReplaceAll(ep.rows[row][2], ",", ""))
	if err != nil {
		// return err
	}

	price, err := decimal.NewFromString(strings.ReplaceAll(ep.rows[row][4], ",", ""))
	if err != nil {
		// return err
	}

	material := Material{
		Name:   name,
		Unit:   unit,
		Number: number,
		Price:  price,
	}

	mostRecentProcess.Materials = append(mostRecentProcess.Materials, &material)

	return nil
}
