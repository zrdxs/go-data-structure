package main

import (
	"fmt"
	"time"
)

// Calcular o valor mensal da assinatura de um servico para usuarios ativos
// mensal = 4 reais
// dia pode ser 30, 31 ou 28 (calcular o valor por dia)
// usuario pode iniciar a assinatura no começo como no meio do mes e pode cancelar no fim ou no meio do mês

// ------------- PONTOS DE ATENÇĀO -------------------
// tinha duas funçoes para pegar o inicio e o fim do mês de acordo com o timestamp
/*
   o timestamp era iniciado com o build completo
   era um array de usuario por um array de serviço
*/

type SubscribeType struct {
	Price int64
	Users []User
}

type User struct {
	SubscribeInit time.Time
	SubscribeEnd  time.Time
	Active        bool
}

func main() {

	toCalculate := SubscribeType{
		Price: 5,
		Users: []User{
			{
				SubscribeInit: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
				SubscribeEnd:  time.Date(2023, 7, 10, 0, 0, 0, 0, time.Local),
				Active:        true,
			},
			{
				SubscribeInit: time.Date(2023, 5, 1, 0, 0, 0, 0, time.Local),
				SubscribeEnd:  time.Date(2023, 6, 1, 0, 0, 0, 0, time.Local),
				Active:        true,
			},
		},
	}

	var totalToPay float64

	for _, user := range toCalculate.Users {
		pricePerDay := float64(toCalculate.Price) / AmountDaysOfMonth(user.SubscribeInit)
		daysOfSubscription := user.SubscribeEnd.Sub(user.SubscribeInit)
		totalToPay += daysOfSubscription.Hours() / 24 * pricePerDay
	}

	fmt.Println(totalToPay)
}

func FirstDayOfTheMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
}

func LastDayOfTheMonth(date time.Time) time.Time {
	year, month, _ := date.Date()
	lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, date.Location())
	return lastDay
}

func AmountDaysOfMonth(date time.Time) float64 {
	firstDay := FirstDayOfTheMonth(date)
	lastDay := LastDayOfTheMonth(date)
	return lastDay.Sub(firstDay).Hours()/24 + 1
}
