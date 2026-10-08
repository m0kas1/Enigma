package simpleconnect

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectSQL()  {
	ctx := context.Background()
	con, err := pgx.Connect(ctx, "postgres://postgres:awdqse1234@localhost:5432/postgres")

	if err != nil {
		panic(err)
	}

	if err := con.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Успешное подключение")

}