package models

import (
	"time"

	"shanker.com/expense-tracker/db"
)

type Expense struct {
	ExpenseID   int64      `json:"expense_id" db:"expense_id"` 
	Category    string    `json:"category" db:"category"`         
	Amount      float64   `json:"amount" db:"amount"`             
	Date        time.Time `json:"date" db:"date"`                 
	Description string    `json:"description,omitempty" db:"description"` 
}


func GetAllExpenses() ([]Expense,error){

	query := "SELECT * FROM expenses"

	rows,err := db.DB.Query(query)

	if err != nil{
		return nil,err
	}

	defer rows.Close()

	var expenses []Expense

	for rows.Next(){
		var expense Expense
		err := rows.Scan(&expense.ExpenseID,&expense.Category,&expense.Amount,&expense.Date,&expense.Description);
		if err!= nil{
			return nil,err
		}
		
		expenses = append(expenses, expense)
	}

	return expenses,nil;

}


func (expense *Expense) Save() error {
	query := `
	INSERT INTO expenses(category, amount, date, description)
	VALUES ($1, $2, $3, $4) RETURNING expense_id`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Use QueryRow to return the inserted expense_id
	err = stmt.QueryRow(expense.Category, expense.Amount, expense.Date, expense.Description).Scan(&expense.ExpenseID)
	if err != nil {
		return err
	}

	return nil
}


func (expense *Expense) Modify(eid int64) error {
	query := `
	UPDATE expenses
	SET category = $1, amount = $2, date = $3, description = $4
	WHERE expense_id = $5`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(expense.Category, expense.Amount, expense.Date, expense.Description, eid)
	if err != nil {
		return err
	}

	return nil
}


func Delete(eid int64) error {
	query := `
	DELETE FROM expenses
	WHERE expense_id = $1`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(eid)
	if err != nil {
		return err
	}

	return nil
}

