package backends

type Datastore interface {
	CreateLedger()
	ListLedgers()
	GetLedger()
	UpdateLedger()

	CreateAsset()
	ListAssets()
	GetAsset()
	UpdateAsset()

	CreateAccount()
	ListAccounts()
	GetAccount()
	UpdateAccount()

	CreateTransaction()
	ListTransactions()
	GetTransaction()
	UpdateTransaction()
}

type Backend struct{}

func (b *Backend) CreateLedger() {}
func (b *Backend) ListLedgers()  {}
func (b *Backend) GetLedger()    {}
func (b *Backend) UpdateLedger() {}

func (b *Backend) CreateAsset() {}
func (b *Backend) ListAssets()  {}
func (b *Backend) GetAsset()    {}
func (b *Backend) UpdateAsset() {}

func (b *Backend) CreateAccount() {}
func (b *Backend) ListAccounts()  {}
func (b *Backend) GetAccount()    {}
func (b *Backend) UpdateAccount() {}

func (b *Backend) CreateTransaction() {}
func (b *Backend) ListTransactions()  {}
func (b *Backend) GetTransaction()    {}
func (b *Backend) UpdateTransaction() {}
