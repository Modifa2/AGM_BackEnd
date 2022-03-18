package services

import (
	"context"
	"errors"
	"fmt"
	"gopls-workspace/models"
	"os"
	"reflect"
	"strconv"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

//
/*DB ... */
type DB struct {
}

//
func (db *DB) SaveOnDB(functionnamewithschema string, m interface{}) ([]models.UserRegisted, error) {
	userReg := []models.UserRegisted{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}

func (db *DB) AssignUserToQRCode(functionnamewithschema string, m interface{}) ([]models.AssignUserToQECode, error) {
	userReg := []models.AssignUserToQECode{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &userReg, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return userReg, nil
}


func (db *DB) GetUserByID(functionnamewithschema string, m interface{}) ([]models.GetUserByDetails, error) {
	user := []models.GetUserByDetails{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &user, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return user, nil
}

//
func (db *DB) CheckInUser(functionnamewithschema string, m interface{}) ([]models.CheckInUserReturn, error) {
	checkIn := []models.CheckInUserReturn{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &checkIn, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return checkIn, nil
}

//agm.display_user
func (db *DB) MarkAsDisplayed(functionnamewithschema string, m interface{}) ([]models.UserDisplayed, error) {
	checkIn := []models.UserDisplayed{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &checkIn, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return checkIn, nil
}
//AllUser
func (db *DB) GetAllUsers(functionnamewithschema string, m interface{}) ([]models.AllUser, error) {
	users := []models.AllUser{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &users, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return users, nil
}
//
func (db *DB) ReturnUser(functionnamewithschema string, m interface{}) ([]models.GetUser, error) {
	User := []models.GetUser{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601	
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return User, nil
}
//AllUser
func (db *DB) ReturnAssocuations(functionnamewithschema string, m interface{}) ([]models.Associations, error) {
	User := []models.Associations{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601	
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return User, nil
}

func (db *DB) ReturnAssociations(functionnamewithschema string, m interface{}) ([]models.GetAssociations, error) {
	Associations := []models.GetAssociations{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("DBTshwane"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &Associations, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601	
		}
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}

	return Associations, nil
}

//Convert Interface and return Query string
func ConVertInterface(funcstr string, m interface{}) string {
	q := "select * from " + funcstr + "("

	if m != nil {

		v := reflect.ValueOf(m)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			switch typeOfS.Field(i).Type.Name() {
			case "int", "int16", "int32", "int64", "int8":
				str := v.Field(i).Interface().(int64)
				strInt64 := strconv.FormatInt(str, 10)
				q += strInt64 + ","
			case "float64":
				str := v.Field(i).Interface().(float64)
				s := fmt.Sprintf("%f", str)
				q += s + ","
			case "bool":
				q += "'" + strconv.FormatBool(v.Field(i).Interface().(bool)) + "',"
			default:
				if v.Field(i).Interface().(string) == "" {
					q += "null,"
				} else {
					q += "'" + v.Field(i).Interface().(string) + "',"
				}
			}
		}

		q = q[0 : len(q)-len(",")]
	}

	q += ")"

	return q
}