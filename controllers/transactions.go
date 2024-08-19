package controllers

import (
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type FormTransactions struct{
	EventId         	int `json:"eventId" form:"eventId" db:"event_id"`
	PaymentMethodId 	int `json:"paymentMethodId" form:"paymentMethodId" db:"payment_method_id"`
	SectionId       	[]int `json:"sectionId" form:"sectionId[]" db:"section_id"`
	TicketQty       	[]int `json:"ticketQty" form:"ticketQty[]" db:"ticket_qty"`
}

func CreateTransaction(ctx *gin.Context) {
	form := FormTransactions{}

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Response{
			Success: false,
			Message: "invalid input data",
		})
		return
	}
	// fmt.Println(form)

	trx := models.CreateTransaction(models.Transactions{
		UserId: ctx.GetInt("userId"),
		PaymentMethodId: form.PaymentMethodId,
		EventId: form.EventId,
	})

	for i := range form.SectionId{
		models.CreateTransactionDetail(models.TransactionDetail{
				SectionId: form.SectionId[i],
				TicketQty: form.TicketQty[i],
				TransactionId: trx.Id,
		})
	}

	data := models.DetailTransactions(trx.Id)
	// if data == (models.DetailTransaction{}) {
	// 	ctx.JSON(http.StatusBadRequest, lib.Response{
	// 		Success: false,
	// 		Message: "Failed to create transaction",
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Transaction created successfully",
		Results: data,
	})
}

func FindTransactionByUserId(ctx *gin.Context){
	UserId := ctx.GetInt("userId")

	detailTransactionbyId := models.FindTransactionByUserId(UserId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Transaction User Id",
		Results: detailTransactionbyId,
	})
}