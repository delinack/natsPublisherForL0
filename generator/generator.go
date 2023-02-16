package generator

import (
	"github.com/brianvoe/gofakeit/v6"
	"natsPublisher/model"
	"time"
)

func GenerateOrder() *model.Order {
	gofakeit.Seed(time.Now().UnixNano())

	goods := gofakeit.IntRange(1, 5)

	order := &model.Order{
		OrderUID:          gofakeit.HexUint128(),
		TrackNumber:       gofakeit.HexUint128(),
		Entry:             gofakeit.HexUint128(),
		Locale:            gofakeit.Language(),
		InternalSignature: gofakeit.HexUint128(),
		CustomerID:        gofakeit.UUID(),
		DeliveryService:   gofakeit.UUID(),
		Shardkey:          gofakeit.DigitN(2),
		SmID:              gofakeit.IntRange(1, 100),
		DateCreated:       gofakeit.DateRange(time.Now().AddDate(0, 0, -7), time.Now()),
		OofShard:          gofakeit.DigitN(2),

		Payment: &model.Payment{
			Transaction:  gofakeit.UUID(),
			RequestID:    gofakeit.UUID(),
			Currency:     gofakeit.Currency().Short,
			Provider:     gofakeit.CelebrityBusiness(),
			Amount:       gofakeit.IntRange(300, 5000),
			PaymentDt:    int(gofakeit.DateRange(time.Now().AddDate(0, 0, -7), time.Now()).Unix()),
			Bank:         gofakeit.CreditCardType(),
			DeliveryCost: gofakeit.RandomInt([]int{0, 100, 150}),
			GoodsTotal:   goods,
			CustomFee:    gofakeit.IntRange(0, 100),
		},

		Delivery: &model.Delivery{
			Name:    gofakeit.CelebrityBusiness(),
			Phone:   gofakeit.Phone(),
			Zip:     gofakeit.Zip(),
			City:    gofakeit.City(),
			Address: gofakeit.Street(),
			Region:  gofakeit.State(),
			Email:   gofakeit.Email(),
		},

		Items: func() []*model.Item {
			var items []*model.Item
			for i := 0; i < goods; i++ {
				items = append(items, &model.Item{
					ChrtID:      gofakeit.IntRange(1, 100),
					TrackNumber: gofakeit.UUID(),
					Price:       gofakeit.IntRange(300, 5000),
					Rid:         gofakeit.UUID(),
					Name:        gofakeit.Snack(),
					Sale:        gofakeit.IntRange(0, 100),
					Size:        gofakeit.DigitN(2),
					TotalPrice:  gofakeit.IntRange(300, 5000),
					NmID:        gofakeit.IntRange(1, 100),
					Brand:       gofakeit.CelebrityBusiness(),
					Status:      gofakeit.IntRange(0, 100),
				})
			}
			return items
		}(),
	}

	return order
}
