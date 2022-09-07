package sap_api_input_reader

type EC_MC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	ProuctionOrder struct {
		ManufacturingOrder             string `json:"document_no"`
		DeliverTo                      string `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string `json:"quantity"`
		MfgOrderItemActualDeviationQty string `json:"picked_quantity"`
		Price                          string `json:"price"`
		Batch                          string `json:"batch"`
	} `json:"document"`
	ManufacturingOrder struct {
		ManufacturingOrder             string `json:"document_no"`
		StatusCode                     string `json:"status"`
		DeliverTo                      string `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string `json:"quantity"`
		MfgOrderItemActualDeviationQty string `json:"completed_quantity"`
		PlannedStartDate               string `json:"planned_start_date"`
		MfgOrderItemPlndDeliveryDate   string `json:"planned_validated_date"`
		ActualStartDate                string `json:"actual_start_date"`
		MfgOrderItemActualDeliveryDate string `json:"actual_validated_date"`
		Batch                          string `json:"batch"`
		BillOfOperations               struct {
			OrderIntBillOfOperationsItem string `json:"work_no"`
			OpPlannedTotalQuantity       string `json:"quantity"`
			OpTotalConfirmedYieldQty     string `json:"completed_quantity"`
			ErroredQuantity              string `json:"errored_quantity"`
			Component                    string `json:"component"`
			MaterialCompOriginalQuantity string `json:"planned_component_quantity"`
			OpErlstSchedldExecStrtDte    string `json:"planned_start_date"`
			OpErlstSchedldExecStrtTme    string `json:"planned_start_time"`
			OpErlstSchedldExecEndDte     string `json:"planned_validated_date"`
			OpErlstSchedldExecEndTme     string `json:"planned_validated_time"`
			OpActualExecutionStartDate   string `json:"actual_start_date"`
			OpActualExecutionStartTime   string `json:"actual_start_time"`
			OpActualExecutionEndDate     string `json:"actual_validated_date"`
			OpActualExecutionEndTime     string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                      string `json:"api_schema"`
	MaterialCode                   string `json:"material_code"`
	ProductionPlant                string `json:"plant/supplier"`
	Stock                          string `json:"stock"`
	ManufacturingOrderType         string `json:"document_type"`
	ManufacturingOrderNo           string `json:"document_no"`
	MfgOrderItemPlndDeliveryDate   string `json:"planned_date"`
	MfgOrderItemActualDeliveryDate string `json:"validated_date"`
	Deleted                        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	ProductionOrder struct {
		ManufacturingOrder            string  `json:"ManufacturingOrder"`
		ManufacturingOrderCategory    string `json:"ManufacturingOrderCategory"`
		ManufacturingOrderType        string `json:"ManufacturingOrderType"`
		OrderIsCreated                *string `json:"OrderIsCreated"`
		OrderIsReleased               *string `json:"OrderIsReleased"`
		OrderIsPrinted                *string `json:"OrderIsPrinted"`
		OrderIsConfirmed              *string `json:"OrderIsConfirmed"`
		OrderIsPartiallyConfirmed     *string `json:"OrderIsPartiallyConfirmed"`
		OrderIsDelivered              *string `json:"OrderIsDelivered"`
		OrderIsDeleted                *string `json:"OrderIsDeleted"`
		OrderIsPreCosted              *string `json:"OrderIsPreCosted"`
		SettlementRuleIsCreated       *string `json:"SettlementRuleIsCreated"`
		OrderIsPartiallyReleased      *string `json:"OrderIsPartiallyReleased"`
		OrderIsLocked                 *string `json:"OrderIsLocked"`
		OrderIsTechnicallyCompleted   *string `json:"OrderIsTechnicallyCompleted"`
		OrderIsClosed                 *string `json:"OrderIsClosed"`
		OrderIsPartiallyDelivered     *string `json:"OrderIsPartiallyDelivered"`
		OrderIsMarkedForDeletion      *string `json:"OrderIsMarkedForDeletion"`
		OrderIsScheduled              *string `json:"OrderIsScheduled"`
		OrderHasGeneratedOperations   *string `json:"OrderHasGeneratedOperations"`
		MaterialAvailyIsNotChecked    *string `json:"MaterialAvailyIsNotChecked"`
		MfgOrderCreationDate          *string `json:"MfgOrderCreationDate"`
		MfgOrderCreationTime          *string `json:"MfgOrderCreationTime"`
		LastChangeDateTime            *string `json:"LastChangeDateTime"`
		Material                      *string `json:"Material"`
		StorageLocation               *string `json:"StorageLocation"`
		GoodsRecipientName            *string `json:"GoodsRecipientName"`
		UnloadingPointName            *string `json:"UnloadingPointName"`
		MaterialGoodsReceiptDuration  *string `json:"MaterialGoodsReceiptDuration"`
		OrderInternalBillOfOperations *string `json:"OrderInternalBillOfOperations"`
		ProductionPlant               *string `json:"ProductionPlant"`
		Plant                         *string `json:"Plant"`
		MRPArea                       *string `json:"MRPArea"`
		MRPController                 *string `json:"MRPController"`
		ProductionSupervisor          *string `json:"ProductionSupervisor"`
		ProductionVersion             *string `json:"ProductionVersion"`
		PlannedOrder                  *string `json:"PlannedOrder"`
		SalesOrder                    *string `json:"SalesOrder"`
		SalesOrderItem                *string `json:"SalesOrderItem"`
		BasicSchedulingType           *string `json:"BasicSchedulingType"`
		BusinessArea                  *string `json:"BusinessArea"`
		CompanyCode                   *string `json:"CompanyCode"`
		ProfitCenter                  *string `json:"ProfitCenter"`
		FunctionalArea                *string `json:"FunctionalArea"`
		MfgOrderPlannedStartDate      *string `json:"MfgOrderPlannedStartDate"`
		MfgOrderPlannedStartTime      *string `json:"MfgOrderPlannedStartTime"`
		MfgOrderPlannedEndDate        *string `json:"MfgOrderPlannedEndDate"`
		MfgOrderPlannedEndTime        *string `json:"MfgOrderPlannedEndTime"`
		MfgOrderScheduledStartDate    *string `json:"MfgOrderScheduledStartDate"`
		MfgOrderScheduledStartTime    *string `json:"MfgOrderScheduledStartTime"`
		MfgOrderScheduledEndDate      *string `json:"MfgOrderScheduledEndDate"`
		MfgOrderScheduledEndTime      *string `json:"MfgOrderScheduledEndTime"`
		MfgOrderActualReleaseDate     *string `json:"MfgOrderActualReleaseDate"`
		ProductionUnit                *string `json:"ProductionUnit"`
		TotalQuantity                 *string `json:"TotalQuantity"`
		MfgOrderPlannedScrapQty       *string `json:"MfgOrderPlannedScrapQty"`
		MfgOrderConfirmedYieldQty     *string `json:"MfgOrderConfirmedYieldQty"`
		WBSElementExternalID          *string `json:"WBSElementExternalID"`
		OrderLongText                 *string `json:"OrderLongText"`
		Status                        struct {
			StatusCode      string  `json:"StatusCode"`
			IsUserStatus    bool    `json:"IsUserStatus"`
			StatusShortName *string `json:"StatusShortName"`
			StatusName      *string `json:"StatusName"`
		} `json:"Status"`
		Component struct {
			ManufacturingOrderSequence     string `json:"ManufacturingOrderSequence"`
			ManufacturingOrderOperation    string `json:"ManufacturingOrderOperation"`
			OrderInternalBillOfOperations  string `json:"OrderInternalBillOfOperations"`
			Reservation                    *string `json:"Reservation"`
			ReservationItem                *string `json:"ReservationItem"`
			MaterialGroup                  *string `json:"MaterialGroup"`
			Material                       *string `json:"Material"`
			Plant                          *string `json:"Plant"`
			ProductionPlant                *string `json:"ProductionPlant"`
			MatlCompRequirementDate        *string `json:"MatlCompRequirementDate"`
			MatlCompRequirementTime        *string `json:"MatlCompRequirementTime"`
			ReservationIsFinallyIssued     *bool   `json:"ReservationIsFinallyIssued"`
			IsBulkMaterialComponent        *bool   `json:"IsBulkMaterialComponent"`
			MatlCompIsMarkedForBackflush   *bool   `json:"MatlCompIsMarkedForBackflush"`
			MaterialCompIsCostRelevant     *string `json:"MaterialCompIsCostRelevant"`
			SalesOrder                     *string `json:"SalesOrder"`
			SalesOrderItem                 *string `json:"SalesOrderItem"`
			SortField                      *string `json:"SortField"`
			BillOfMaterialCategory         *string `json:"BillOfMaterialCategory"`
			BOMItem                        string `json:"BOMItem"`
			BOMItemCategory                *string `json:"BOMItemCategory"`
			BillOfMaterialItemNumber       string `json:"BillOfMaterialItemNumber"`
			BOMItemDescription             *string `json:"BOMItemDescription"`
			StorageLocation                *string `json:"StorageLocation"`
			Batch                          *string `json:"Batch"`
			BatchSplitType                 *string `json:"BatchSplitType"`
			GoodsMovementType              *string `json:"GoodsMovementType"`
			SupplyArea                     *string `json:"SupplyArea"`
			GoodsRecipientName             *string `json:"GoodsRecipientName"`
			UnloadingPointName             *string `json:"UnloadingPointName"`
			MaterialCompIsAlternativeItem  *bool   `json:"MaterialCompIsAlternativeItem"`
			AlternativeItemGroup           *string `json:"AlternativeItemGroup"`
			AlternativeItemStrategy        *string `json:"AlternativeItemStrategy"`
			AlternativeItemPriority        *string `json:"AlternativeItemPriority"`
			UsageProbabilityPercent        *string `json:"UsageProbabilityPercent"`
			MaterialComponentIsPhantomItem *bool   `json:"MaterialComponentIsPhantomItem"`
			LeadTimeOffset                 *string `json:"LeadTimeOffset"`
			QuantityIsFixed                *bool   `json:"QuantityIsFixed"`
			IsNetScrap                     *bool   `json:"IsNetScrap"`
			ComponentScrapInPercent        *string `json:"ComponentScrapInPercent"`
			OperationScrapInPercent        *string `json:"OperationScrapInPercent"`
			BaseUnit                       *string `json:"BaseUnit"`
			BaseUnitISOCode                *string `json:"BaseUnitISOCode"`
			BaseUnitSAPCode                *string `json:"BaseUnitSAPCode"`
			RequiredQuantity               *string `json:"RequiredQuantity"`
			WithdrawnQuantity              *string `json:"WithdrawnQuantity"`
			ConfirmedAvailableQuantity     *string `json:"ConfirmedAvailableQuantity"`
			MaterialCompOriginalQuantity   *string `json:"MaterialCompOriginalQuantity"`
			Currency                       *string `json:"Currency"`
			WithdrawnQuantityAmount        *string `json:"WithdrawnQuantityAmount"`
		} `json:"Component"`
		Operation struct {
			ManufacturingOrderSequence     string `json:"ManufacturingOrderSequence"`
			ManufacturingOrderOperation    string `json:"ManufacturingOrderOperation"`
			ManufacturingOrderSubOperation *string `json:"ManufacturingOrderSubOperation"`
			OrderInternalBillOfOperations  string  `json:"OrderInternalBillOfOperations"`
			OrderIntBillOfOperationsItem   string  `json:"OrderIntBillOfOperationsItem"`
			MfgOrderSequenceText           *string `json:"MfgOrderSequenceText"`
			MfgOrderOperationText          *string `json:"MfgOrderOperationText"`
			OperationIsCreated             *string `json:"OperationIsCreated"`
			OperationIsReleased            *string `json:"OperationIsReleased"`
			OperationIsPrinted             *string `json:"OperationIsPrinted"`
			OperationIsConfirmed           *string `json:"OperationIsConfirmed"`
			OperationIsPartiallyConfirmed  *string `json:"OperationIsPartiallyConfirmed"`
			OperationIsDeleted             *string `json:"OperationIsDeleted"`
			OperationIsTechlyCompleted     *string `json:"OperationIsTechlyCompleted"`
			OperationIsClosed              *string `json:"OperationIsClosed"`
			OperationIsScheduled           *string `json:"OperationIsScheduled"`
			OperationIsPartiallyDelivered  *string `json:"OperationIsPartiallyDelivered"`
			OperationIsDelivered           *string `json:"OperationIsDelivered"`
			ProductionPlant                *string `json:"ProductionPlant"`
			WorkCenterInternalID           *string `json:"WorkCenterInternalID"`
			WorkCenterTypeCode             *string `json:"WorkCenterTypeCode"`
			WorkCenter                     *string `json:"WorkCenter"`
			OperationControlProfile        *string `json:"OperationControlProfile"`
			OpErlstSchedldExecStrtDte      *string `json:"OpErlstSchedldExecStrtDte"`
			OpErlstSchedldExecStrtTme      *string `json:"OpErlstSchedldExecStrtTme"`
			OpErlstSchedldExecEndDte       *string `json:"OpErlstSchedldExecEndDte"`
			OpErlstSchedldExecEndTme       *string `json:"OpErlstSchedldExecEndTme"`
			OpActualExecutionStartDate     *string `json:"OpActualExecutionStartDate"`
			OpActualExecutionStartTime     *string `json:"OpActualExecutionStartTime"`
			OpActualExecutionEndDate       *string `json:"OpActualExecutionEndDate"`
			OpActualExecutionEndTime       *string `json:"OpActualExecutionEndTime"`
			ErlstSchedldExecDurnInWorkdays *int    `json:"ErlstSchedldExecDurnInWorkdays"`
			OpActualExecutionDays          *int    `json:"OpActualExecutionDays"`
			OperationUnit                  *string `json:"OperationUnit"`
			OpPlannedTotalQuantity         *string `json:"OpPlannedTotalQuantity"`
			OpTotalConfirmedYieldQty       *string `json:"OpTotalConfirmedYieldQty"`
			LastChangeDateTime             *string `json:"LastChangeDateTime"`
		} `json:"Operation"`
		ProductionOrderItem struct {
			ManufacturingOrderItem         string  `json:"ManufacturingOrderItem"`
			IsCompletelyDelivered          *bool   `json:"IsCompletelyDelivered"`
			Material                       *string `json:"Material"`
			ProductionPlant                *string `json:"ProductionPlant"`
			Plant                          *string `json:"Plant"`
			MRPArea                        *string `json:"MRPArea"`
			QuantityDistributionKey        *string `json:"QuantityDistributionKey"`
			MaterialGoodsReceiptDuration   *string `json:"MaterialGoodsReceiptDuration"`
			StorageLocation                *string `json:"StorageLocation"`
			Batch                          *string `json:"Batch"`
			InventoryUsabilityCode         *string `json:"InventoryUsabilityCode"`
			GoodsRecipientName             *string `json:"GoodsRecipientName"`
			UnloadingPointName             *string `json:"UnloadingPointName"`
			MfgOrderItemPlndDeliveryDate   *string `json:"MfgOrderItemPlndDeliveryDate"`
			MfgOrderItemActualDeliveryDate *string `json:"MfgOrderItemActualDeliveryDate"`
			ProductionUnit                 *string `json:"ProductionUnit"`
			MfgOrderItemPlannedTotalQty    *string `json:"MfgOrderItemPlannedTotalQty"`
			MfgOrderItemPlannedScrapQty    *string `json:"MfgOrderItemPlannedScrapQty"`
			MfgOrderItemGoodsReceiptQty    *string `json:"MfgOrderItemGoodsReceiptQty"`
			MfgOrderItemActualDeviationQty *string `json:"MfgOrderItemActualDeviationQty"`
		} `json:"ProductionOrderItem"`
	} `json:"ProductionOrder"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	ProductionOrderNo string   `json:"production_order"`
	Deleted           bool     `json:"deleted"`
}
