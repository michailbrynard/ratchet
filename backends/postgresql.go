package backends

type PostgresSqlBackend struct{}

func (b *PostgresSqlBackend) CreateLedger() {}
func (b *PostgresSqlBackend) ListLedgers()  {}
func (b *PostgresSqlBackend) GetLedger()    {}
func (b *PostgresSqlBackend) UpdateLedger() {}

func (b *PostgresSqlBackend) CreateAsset() {}
func (b *PostgresSqlBackend) ListAssets()  {}
func (b *PostgresSqlBackend) GetAsset()    {}
func (b *PostgresSqlBackend) UpdateAsset() {}

func (b *PostgresSqlBackend) CreateAccount() {}
func (b *PostgresSqlBackend) ListAccounts()  {}
func (b *PostgresSqlBackend) GetAccount()    {}
func (b *PostgresSqlBackend) UpdateAccount() {}

func (b *PostgresSqlBackend) CreateTransaction() {}
func (b *PostgresSqlBackend) ListTransactions()  {}
func (b *PostgresSqlBackend) GetTransaction()    {}
func (b *PostgresSqlBackend) UpdateTransaction() {}
