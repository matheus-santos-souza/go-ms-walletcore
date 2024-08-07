package database

import (
	"database/sql"
	"testing"

	"github.com/matheus-santos-souza/go-ms-walletcore/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE)")
<<<<<<< HEAD
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), name VARCHAR(255), balance INT, created_at DATE)")
=======
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance INT, created_at DATE)")
>>>>>>> master
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount INT, created_at DATE)")
	s.client, _ = entity.NewClient("Matheus", "m@m.com")
	s.client2, _ = entity.NewClient("Santos", "s@s.com")
	s.accountFrom = entity.NewAccount(s.client)
	s.accountFrom.Balance = 1000
	s.accountTo = entity.NewAccount(s.client2)
	s.accountTo.Balance = 1000
	s.transactionDB = NewTransactionDB(db)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
