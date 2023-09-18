package estimateparser

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	// Estimate ...
	Estimate struct {
		Project  string    `json:"project"`
		Author   string    `json:"author"`
		CreateDt time.Time `json:"create_dt"`
		Client   *Client   `json:"client"`
		Blocks   []*Block  `json:"blocks"`
	}

	// Client ...
	Client struct {
		FullName string `json:"full_name"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Address  string `json:"address"`
	}

	// Block ...
	Block struct {
		Title     string     `json:"title"`
		Processes []*Process `json:"processes"`
	}

	// Process ...
	Process struct {
		Name      string          `json:"name"`
		Unit      string          `json:"unit"`
		Number    decimal.Decimal `json:"number"`
		WorkPrice decimal.Decimal `json:"work_price"`
		Materials []*Material     `json:"materials"`
	}

	// Material ...
	Material struct {
		Name   string          `json:"name"`
		Unit   string          `json:"unit_unit"`
		Number decimal.Decimal `json:"number"`
		Price  decimal.Decimal `json:"price"`
	}
)
