package testing

import (
	"context"
	"fmt"
	"log"
	"testing"

	belajargolangdatabase "github.com/MargaTzi/belajar-golang-database"
	"github.com/MargaTzi/belajar-golang-database/constructor"
	"github.com/MargaTzi/belajar-golang-database/entity"
	_ "github.com/lib/pq"
)

func TestCreateData(t *testing.T) {
	OrderRepo := constructor.NewOrder(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	order := entity.Orders{
		Prd_id: 6,
		Qty: 30,
	}

	data, err := OrderRepo.CreateData(ctx, order)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}

func TestGetdata(t *testing.T) {
	getdata := constructor.NewOrder(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	data, err := getdata.FindAllOrderDetail(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, all := range data{
		fmt.Println(all)
	}
}