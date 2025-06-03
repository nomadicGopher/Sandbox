package main

type Transaction278 struct {
	ISA  ISA    `json:"isa"`
	GS   GS     `json:"gs"`
	ST   ST     `json:"st"`
	BHT  BHT    `json:"bht"`
	Body string `json:"body,omitempty"`
	SE   SE     `json:"se"`
	GE   GE     `json:"ge"`
	IEA  IEA    `json:"iea"`
}

// ISA/GS Header Segments
type ISA struct {
	AuthorizationInfoQualifier  string `json:"authorization_info_qualifier"`  // ISA01
	AuthorizationInfo           string `json:"authorization_info"`            // ISA02
	SecurityInfoQualifier       string `json:"security_info_qualifier"`       // ISA03
	SecurityInfo                string `json:"security_info"`                 // ISA04
	InterchangeIDQualifier1     string `json:"interchange_id_qualifier_1"`    // ISA05
	InterchangeSenderID         string `json:"interchange_sender_id"`         // ISA06
	InterchangeIDQualifier2     string `json:"interchange_id_qualifier_2"`    // ISA07
	InterchangeReceiverID       string `json:"interchange_receiver_id"`       // ISA08
	InterchangeDate             string `json:"interchange_date"`              // ISA09
	InterchangeTime             string `json:"interchange_time"`              // ISA10
	InterchangeControlStandards string `json:"interchange_control_standards"` // ISA11
	InterchangeControlVersion   string `json:"interchange_control_version"`   // ISA12
	InterchangeControlNumber    string `json:"interchange_control_number"`    // ISA13
	AcknowledgmentRequested     string `json:"acknowledgment_requested"`      // ISA14
	UsageIndicator              string `json:"usage_indicator"`               // ISA15
	ComponentElementSeparator   string `json:"component_element_separator"`   // ISA16
}

type GS struct {
	FunctionalIDCode           string `json:"functional_id_code"`            // GS01
	ApplicationSenderCode      string `json:"application_sender_code"`       // GS02
	ApplicationReceiverCode    string `json:"application_receiver_code"`     // GS03
	Date                       string `json:"date"`                          // GS04
	Time                       string `json:"time"`                          // GS05
	GroupControlNumber         string `json:"group_control_number"`          // GS06
	ResponsibleAgencyCode      string `json:"responsible_agency_code"`       // GS07
	VersionReleaseIndustryCode string `json:"version_release_industry_code"` // GS08
}

// ST/SE Transaction Set Header/Trailer
type ST struct {
	TransactionSetIDCode        string `json:"transaction_set_id_code"`                 // ST01
	TransactionSetControlNumber string `json:"transaction_set_control_number"`          // ST02
	ImplementationConventionRef string `json:"implementation_convention_ref,omitempty"` // ST03 (optional)
}

type BHT struct {
	HeirarchStructCode  string `json:"heirarch_struct_code"`  // BHT01
	TSPurposeCode       string `json:"ts_purpose_code"`       // BHT02
	ReferenceID         string `json:"reference_id"`          // BHT03
	Date                string `json:"date"`                  // BHT04
	Time                string `json:"time"`                  // BHT05
	TransactionTypeCode string `json:"transaction_type_code"` // BHT06
}

type SE struct {
	NumberOfIncludedSegments    int    `json:"number_of_included_segments"`    // SE01
	TransactionSetControlNumber string `json:"transaction_set_control_number"` // SE02
}

// GE: Functional Group Trailer
type GE struct {
	NumberOfTransactionSetsIncluded string `json:"number_of_transaction_sets_included"` // GE01
	GroupControlNumber              string `json:"group_control_number"`                // GE02
}

// IEA: Interchange Control Trailer
type IEA struct {
	NumberOfIncludedGroups   string `json:"number_of_included_groups"`  // IEA01
	InterchangeControlNumber string `json:"interchange_control_number"` // IEA02
}
