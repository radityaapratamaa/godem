package transaction

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	"bcg-test/domain/models/goods"
	"bcg-test/domain/models/transaction"
)

type purchase interface {
	CreateNew(ctx context.Context, requestData []*transaction.Purchase) (*transaction.PurchaseResponse, error)
}

func (uc *Usecase) CreateNew(ctx context.Context, requestData []*transaction.Purchase) (*transaction.PurchaseResponse, error) {
	// check request data is null or not
	if requestData == nil {
		errMsg := errors.New("the request data is null")
		return nil, errors.Wrap(errMsg, "usecase.transaction.purchase.CreateNew")
	}

	// check each of purchase data request
	// dummy goods_id for mackbook pro
	macbookProDummyID := 1
	googleHomeDummyID := 2
	alexaSpeakerDummyID := 3

	purchaseMsg := ""
	rowsAffected := 0
	for _, purchaseData := range requestData {
		detailData, err := uc.goodsDB.GetDetailByID(ctx, purchaseData.GoodsID)
		if err != nil {
			return nil, err
		}
		goodsDetailData := detailData.(*goods.Goods)
		if purchaseData.GoodsID == int64(macbookProDummyID) {
			macbookPurchaseData := &transaction.Purchase{
				GoodsID:      int64(macbookProDummyID),
				PurchaseDate: time.Now().String(),
				Qty:          purchaseData.Qty,
				TotalPrice:   goodsDetailData.Price * float64(purchaseData.Qty),
			}
			createMacbookTransaction, err := uc.db.CreateNew(ctx, macbookPurchaseData)
			if err != nil {
				return nil, err
			}
			rowsAffected += int(createMacbookTransaction.RowsAffected)
			raspberryDetailData, err := uc.goodsDB.GetDetailByName(ctx, "Raspberry Pi")
			usedRaspberryQty := 0
			if err != nil {
				log.Println("got an error when get raspberry qty, please check: ", err.Error())
			} else if purchaseData.Qty >= raspberryDetailData.Qty && raspberryDetailData.Qty > 0 {
				usedRaspberryQty = raspberryDetailData.Qty
			} else {
				usedRaspberryQty = purchaseData.Qty
			}
			raspberryPurchaseData := &transaction.Purchase{
				GoodsID:      raspberryDetailData.ID,
				Qty:          usedRaspberryQty,
				PurchaseDate: time.Now().String(),
				CreatedAt:    time.Now().String(),
				TotalPrice:   0,
			}
			createRaspberryTransaction, err := uc.db.CreateNew(ctx, raspberryPurchaseData)
			if err != nil {
				return nil, err
			}
			rowsAffected += int(createRaspberryTransaction.RowsAffected)
			purchaseMsg = fmt.Sprintf("you've got %d Raspberry Pi for free", usedRaspberryQty)
			continue
		} else if purchaseData.GoodsID == int64(googleHomeDummyID) {
			if purchaseData.Qty == 3 {
				purchaseData.Qty = 2
			}
			purchaseData.TotalPrice = float64(purchaseData.Qty) * goodsDetailData.Price
		} else if purchaseData.GoodsID == int64(alexaSpeakerDummyID) {
			if purchaseData.Qty >= 3 {
				purchaseData.TotalPrice = ((float64(purchaseData.Qty) * goodsDetailData.Price) * 0.1)
			}
		}

		storedData, err := uc.db.CreateNew(ctx, purchaseData)
		if err != nil {
			return nil, err
		}
		rowsAffected += int(storedData.RowsAffected)
	}

	result := new(transaction.PurchaseResponse)
	result.Message = purchaseMsg
	result.RowsAffected = int64(rowsAffected)
	return result, nil
}
