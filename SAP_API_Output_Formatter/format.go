package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-order-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) (*General, error) {
	pm := &responses.General{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	general := &General{
		ManufacturingOrder:            data.ManufacturingOrder,
		ManufacturingOrderCategory:    data.ManufacturingOrderCategory,
		ManufacturingOrderType:        data.ManufacturingOrderType,
		OrderIsCreated:                data.OrderIsCreated,
		OrderIsReleased:               data.OrderIsReleased,
		OrderIsPrinted:                data.OrderIsPrinted,
		OrderIsConfirmed:              data.OrderIsConfirmed,
		OrderIsPartiallyConfirmed:     data.OrderIsPartiallyConfirmed,
		OrderIsDelivered:              data.OrderIsDelivered,
		OrderIsDeleted:                data.OrderIsDeleted,
		OrderIsPreCosted:              data.OrderIsPreCosted,
		SettlementRuleIsCreated:       data.SettlementRuleIsCreated,
		OrderIsPartiallyReleased:      data.OrderIsPartiallyReleased,
		OrderIsLocked:                 data.OrderIsLocked,
		OrderIsTechnicallyCompleted:   data.OrderIsTechnicallyCompleted,
		OrderIsClosed:                 data.OrderIsClosed,
		OrderIsPartiallyDelivered:     data.OrderIsPartiallyDelivered,
		OrderIsMarkedForDeletion:      data.OrderIsMarkedForDeletion,
		OrderIsScheduled:              data.OrderIsScheduled,
		OrderHasGeneratedOperations:   data.OrderHasGeneratedOperations,
		MaterialAvailyIsNotChecked:    data.MaterialAvailyIsNotChecked,
		MfgOrderCreationDate:          data.MfgOrderCreationDate,
		MfgOrderCreationTime:          data.MfgOrderCreationTime,
		LastChangeDateTime:            data.LastChangeDateTime,
		Material:                      data.Material,
		StorageLocation:               data.StorageLocation,
		GoodsRecipientName:            data.GoodsRecipientName,
		UnloadingPointName:            data.UnloadingPointName,
		MaterialGoodsReceiptDuration:  data.MaterialGoodsReceiptDuration,
		OrderInternalBillOfOperations: data.OrderInternalBillOfOperations,
		ProductionPlant:               data.ProductionPlant,
		Plant:                         data.Plant,
		MRPArea:                       data.MRPArea,
		MRPController:                 data.MRPController,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductionVersion:             data.ProductionVersion,
		PlannedOrder:                  data.PlannedOrder,
		SalesOrder:                    data.SalesOrder,
		SalesOrderItem:                data.SalesOrderItem,
		BasicSchedulingType:           data.BasicSchedulingType,
		BusinessArea:                  data.BusinessArea,
		CompanyCode:                   data.CompanyCode,
		ProfitCenter:                  data.ProfitCenter,
		FunctionalArea:                data.FunctionalArea,
		MfgOrderPlannedStartDate:      data.MfgOrderPlannedStartDate,
		MfgOrderPlannedStartTime:      data.MfgOrderPlannedStartTime,
		MfgOrderPlannedEndDate:        data.MfgOrderPlannedEndDate,
		MfgOrderPlannedEndTime:        data.MfgOrderPlannedEndTime,
		MfgOrderScheduledStartDate:    data.MfgOrderScheduledStartDate,
		MfgOrderScheduledStartTime:    data.MfgOrderScheduledStartTime,
		MfgOrderScheduledEndDate:      data.MfgOrderScheduledEndDate,
		MfgOrderScheduledEndTime:      data.MfgOrderScheduledEndTime,
		MfgOrderActualReleaseDate:     data.MfgOrderActualReleaseDate,
		ProductionUnit:                data.ProductionUnit,
		TotalQuantity:                 data.TotalQuantity,
		MfgOrderPlannedScrapQty:       data.MfgOrderPlannedScrapQty,
		MfgOrderConfirmedYieldQty:     data.MfgOrderConfirmedYieldQty,
		WBSElementExternalID:          data.WBSElementExternalID,
		OrderLongText:                 data.OrderLongText,
	}

	return general, nil
}
