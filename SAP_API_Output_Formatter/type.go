package sap_api_output_formatter

type ProductionOrder struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	Product         string `json:"Product"`
	APISchema       string `json:"api_schema"`
	ProductionOrder string `json:"production_order"`
	Deleted         string `json:"deleted"`
}

type General struct {
	ManufacturingOrder            string `json:"ManufacturingOrder"`
	ManufacturingOrderCategory    string `json:"ManufacturingOrderCategory"`
	ManufacturingOrderType        string `json:"ManufacturingOrderType"`
	OrderIsCreated                string `json:"OrderIsCreated"`
	OrderIsReleased               string `json:"OrderIsReleased"`
	OrderIsPrinted                string `json:"OrderIsPrinted"`
	OrderIsConfirmed              string `json:"OrderIsConfirmed"`
	OrderIsPartiallyConfirmed     string `json:"OrderIsPartiallyConfirmed"`
	OrderIsDelivered              string `json:"OrderIsDelivered"`
	OrderIsDeleted                string `json:"OrderIsDeleted"`
	OrderIsPreCosted              string `json:"OrderIsPreCosted"`
	SettlementRuleIsCreated       string `json:"SettlementRuleIsCreated"`
	OrderIsPartiallyReleased      string `json:"OrderIsPartiallyReleased"`
	OrderIsLocked                 string `json:"OrderIsLocked"`
	OrderIsTechnicallyCompleted   string `json:"OrderIsTechnicallyCompleted"`
	OrderIsClosed                 string `json:"OrderIsClosed"`
	OrderIsPartiallyDelivered     string `json:"OrderIsPartiallyDelivered"`
	OrderIsMarkedForDeletion      string `json:"OrderIsMarkedForDeletion"`
	OrderIsScheduled              string `json:"OrderIsScheduled"`
	OrderHasGeneratedOperations   string `json:"OrderHasGeneratedOperations"`
	MaterialAvailyIsNotChecked    string `json:"MaterialAvailyIsNotChecked"`
	MfgOrderCreationDate          string `json:"MfgOrderCreationDate"`
	MfgOrderCreationTime          string `json:"MfgOrderCreationTime"`
	LastChangeDateTime            string `json:"LastChangeDateTime"`
	Material                      string `json:"Material"`
	StorageLocation               string `json:"StorageLocation"`
	GoodsRecipientName            string `json:"GoodsRecipientName"`
	UnloadingPointName            string `json:"UnloadingPointName"`
	MaterialGoodsReceiptDuration  string `json:"MaterialGoodsReceiptDuration"`
	OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
	ProductionPlant               string `json:"ProductionPlant"`
	Plant                         string `json:"Plant"`
	MRPArea                       string `json:"MRPArea"`
	MRPController                 string `json:"MRPController"`
	ProductionSupervisor          string `json:"ProductionSupervisor"`
	ProductionVersion             string `json:"ProductionVersion"`
	PlannedOrder                  string `json:"PlannedOrder"`
	SalesOrder                    string `json:"SalesOrder"`
	SalesOrderItem                string `json:"SalesOrderItem"`
	BasicSchedulingType           string `json:"BasicSchedulingType"`
	BusinessArea                  string `json:"BusinessArea"`
	CompanyCode                   string `json:"CompanyCode"`
	ProfitCenter                  string `json:"ProfitCenter"`
	FunctionalArea                string `json:"FunctionalArea"`
	MfgOrderPlannedStartDate      string `json:"MfgOrderPlannedStartDate"`
	MfgOrderPlannedStartTime      string `json:"MfgOrderPlannedStartTime"`
	MfgOrderPlannedEndDate        string `json:"MfgOrderPlannedEndDate"`
	MfgOrderPlannedEndTime        string `json:"MfgOrderPlannedEndTime"`
	MfgOrderScheduledStartDate    string `json:"MfgOrderScheduledStartDate"`
	MfgOrderScheduledStartTime    string `json:"MfgOrderScheduledStartTime"`
	MfgOrderScheduledEndDate      string `json:"MfgOrderScheduledEndDate"`
	MfgOrderScheduledEndTime      string `json:"MfgOrderScheduledEndTime"`
	MfgOrderActualReleaseDate     string `json:"MfgOrderActualReleaseDate"`
	ProductionUnit                string `json:"ProductionUnit"`
	TotalQuantity                 string `json:"TotalQuantity"`
	MfgOrderPlannedScrapQty       string `json:"MfgOrderPlannedScrapQty"`
	MfgOrderConfirmedYieldQty     string `json:"MfgOrderConfirmedYieldQty"`
	WBSElementExternalID          string `json:"WBSElementExternalID"`
	OrderLongText                 string `json:"OrderLongText"`
}
