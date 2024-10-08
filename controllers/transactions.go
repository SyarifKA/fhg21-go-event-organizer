package controllers

import (
	"fmt"
	"net/http"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(ctx *gin.Context) {
	form := dtos.FormTransactions{}

	if err := ctx.ShouldBind(&form); err != nil {
		lib.HandlerBadReq(ctx, "Invalid input data")
		return
	}
	// fmt.Println(form)

	// ctx.Bind(&form)
	// user := ctx.GetInt("userId")

	trx := repository.CreateTransaction(models.Transactions{
		UserId:          ctx.GetInt("userId"),
		PaymentMethodId: form.PaymentMethodId,
		EventId:         form.EventId,
	})

	fmt.Println(form.SectionId)

	for i := range form.SectionId {
		repository.CreateTransactionDetail(models.TransactionDetail{
			TransactionId: trx.Id,
			SectionId:     form.SectionId[i],
			TicketQty:     form.TicketQty[i],
		})
	}

	data := repository.DetailTransactions(trx.Id)

	// if data == (models.DetailTransaction{}) {
	// 	ctx.JSON(http.StatusBadRequest, lib.Response{
	// 		Success: false,
	// 		Message: "Failed to create transaction",
	// 	})
	// 	return
	// }
	lib.HandlerOK(ctx, "Transaction created successfully", data, nil)
}

func FindTransactionByUserId(ctx *gin.Context) {
	UserId := ctx.GetInt("userId")

	detailTransactionbyId := repository.FindTransactionByUserId(UserId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Transaction User Id",
		Results: detailTransactionbyId,
	})
}
