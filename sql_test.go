package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	ctx := context.Background()

	command := "INSERT INTO customer(id, name) VALUES ('joko', 'JOKO')"
	_, err := db.ExecContext(ctx, command)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	command := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	command := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt *time.Time
		var birthDate sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=========================")
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		} else {
			fmt.Println("Email: ")
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("BirthDate:", birthDate.Time)
		} else {
			fmt.Println("BirthDate: ")
		}
		fmt.Println("Married:", married)
		fmt.Println("CreatedAt:", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "akdmksmdkasm"

	command := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username sql.NullString
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	command := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, command, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username sql.NullString
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	ctx := context.Background()

	id := "rf"
	name := "Rafeyfa"

	command := "INSERT INTO customer(id, name) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, command, id, name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "akmal@gmail.com"
	comment := "Test Komen"

	command := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, command, email, comment)
	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id ", lastInsertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	command := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, command)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "akmal" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	command := "INSERT INTO comments (email, comment) VALUES (?, ?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "akmal" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, command, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", lastInsertId)
	}

	err = tx.Rollback()
	if err != nil {
		return
	}
}
