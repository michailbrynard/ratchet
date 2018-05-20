package engine

type Ledger struct {
	Id    string `db:"id" json:"id"`
	Label string `db:"label" json:"label"`
	Tags  []*Tag `db:"tags" json:"tags"`
	// TODO : keys?
}

type Asset struct {
	Id            string `db:"id" json:"id"`
	Label         string `db:"label" json:"label"`
	Tags          []*Tag `db:"tags" json:"tags"`
	Symbol        string `db:"symbol" json:"symbol"`
	Unit          string `db:"unit" json:"unit"`
	DecimalPlaces int    `db:"decimal_places" json:"decimalPlaces"`
	// TODO : keys?
}

type Account struct {
	Id    string `db:"id" json:"id"`
	Label string `db:"label" json:"label"`
	Tags  []*Tag `db:"tags" json:"tags"`
	// TODO : keys?
}

type Transaction struct {
	Id      string    `db:"id" json:"id"`
	Label   string    `db:"label" json:"label"`
	Tags    []*Tag    `db:"tags" json:"tags"`
	Actions []*Action `db:"actions" json:"actions"`
}

type Actiion struct {
	Id                     string  `db:"id" json:"id"`
	Label                  string  `db:"label" json:"label"`
	Tags                   []*Tag  `db:"tags" json:"tags"`
	Asset                  Asset   `db:"asset" json:"asset"`
	SourceAccount          Account `db:"source_account" json:"sourceAccount"`
	DestinationAccount     Account `db:"destination_source" json:"destinationAccount"`
	SourceAccountTags      []*Tag  `db:"source_account_tags" json:"sourceAccountTags"`
	DestinationAccountTags []*Tag  `db:"destination_account_tags" json:"destinationAccountTags"`
	Amount                 int     `db:"amount" json:"amount"`
}

type Tag struct {
	string
}

type Engine struct {
}

func NewEngine() *Engine {
	e := Engine{}
	return &e
}

func (e *Engine) CreateLedger() {}
func (e *Engine) ListLedgers()  {}
func (e *Engine) GetLedger()    {}
func (e *Engine) UpdateLedger() {}

func (e *Engine) CreateAsset() {}
func (e *Engine) ListAssets()  {}
func (e *Engine) GetAsset()    {}
func (e *Engine) UpdateAsset() {}

func (e *Engine) CreateAccount() {}
func (e *Engine) ListAccounts()  {}
func (e *Engine) GetAccount()    {}
func (e *Engine) UpdateAccount() {}

func (e *Engine) CreateTransaction() {}
func (e *Engine) ListTransactions()  {}
func (e *Engine) GetTransaction()    {}
func (e *Engine) UpdateTransaction() {}
