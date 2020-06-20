package customer

import (
	"fmt"

	"github.com/thexgithub777/finalexam/database"
)

//CreateTable create table customers if not exists
func CreateTable() error {
	createTb := `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`

	_, err := database.Conn().Exec(createTb)
	if err != nil {
		return fmt.Errorf("can't create table customers: %w", err)
	}

	fmt.Println("create table success.")
	return nil
}

//CreateCustomer create customer
func CreateCustomer(cust *Customer) (int, error) {
	var err error

	row := database.Conn().QueryRow("INSERT INTO customers (name,email,status) values ($1,$2,$3) RETURNING id", cust.Name, cust.Email, cust.Status)

	var id int
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("can't create customers: %w", err)
	}
	return id, nil
}

//FindCustomerByID find by id
func FindCustomerByID(rqid int) (Customer, error) {
	var cust Customer
	stmt, err := database.Conn().Prepare("select * from customers where id=$1")

	if err != nil {
		return cust, fmt.Errorf("can't prepare statement: %w", err)
	}

	row := stmt.QueryRow(rqid)

	var id int
	var name, email, status string

	err = row.Scan(&id, &name, &email, &status)
	if err != nil {
		return cust, fmt.Errorf("can't find customers by id: %w", err)
	}

	cust = Customer{id, name, email, status}
	return cust, nil
}

//FindAllCustomers find all customer
func FindAllCustomers() ([]Customer, error) {
	custs := []Customer{}
	stmt, err := database.Conn().Prepare("select * from customers order by id ")
	if err != nil {
		return custs, fmt.Errorf("can't prepare statement: %w", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		return custs, fmt.Errorf("can't find customers: %w", err)
	}

	// items := []*Customer{}
	for rows.Next() {
		var id int
		var name, email, status string

		err := rows.Scan(&id, &name, &email, &status)
		if err != nil {
			return custs, fmt.Errorf("can't find customers: %w", err)
		}
		c := Customer{id, name, email, status}
		custs = append(custs, c)
	}
	return custs, nil
}

//UpdateCustomerByID update customer data by id
func UpdateCustomerByID(cust *Customer) error {
	stmt, err := database.Conn().Prepare("UPDATE customers set name=$1,email=$2,status=$3 where id=$4 ")

	if err != nil {
		return fmt.Errorf("can't prepare statement: %w", err)
	}

	_, err = stmt.Exec(cust.Name, cust.Email, cust.Status, cust.ID)
	if err != nil {
		return fmt.Errorf("can't update customers: %w", err)
	}

	return nil
}

//DeleteCustomerByID delete customer by id
func DeleteCustomerByID(id int) error {
	stmt, err := database.Conn().Prepare("DELETE from customers where id=$1 ")
	if err != nil {
		return fmt.Errorf("can't prepare delete statement: %w", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("can't execute delete: %w", err)
	}

	return nil
}
