package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/do-app/model"
)

type TransactionRepository interface {
	AddBalance(addBallance int, registeredCustomers *model.User) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func (tx *transactionRepository) AddBalance(addBallance int, registeredCustomers *model.User) error {
	balance := &model.Balance{
		UserId:  registeredCustomers.Id,
		Balance: addBallance,
	}
	_, err := tx.db.NamedExec("INSERT INTO balance (user_id, balance) VALUES (:user_id, :balance)", &balance)
	if err != nil {
		return err
	}
	return nil

}

func NewTransactionRepository(dbArg *sqlx.DB) TransactionRepository {
	return &transactionRepository{
		db: dbArg,
	}
}

/*

kondisi ketika user udah create acount
syarat penggunaan aplikasi

1. Buka dompet dulu buat ngisi balance
Jika user ingin membuaka dompet digital wajib top up
so, user add balance > 10.000 user punya dompet digital dan bisa menggunakan fitur



2. Transaksi hela supaya dapet riwayat karena transaksi mencakup semua query di Db
kalo kita transaksi otomatis kita harus buat

flownya :

minimal transaksi di atas 10.000 if true
insert di tabel transaction
if true transaksi (amount) di atas 10.000 true

Insert to receiver
if acount number dan bank_name dan name == 0 maka false if true

insert transaction_category
if category_name / nama transaksi (tidak menyebutkan transaksi buat apa) maka false if ngisi maka true

Update tabel Balances
sesuai user.id



Menampilkan Saldo User

sqlx

3.









// bikin query


*/
