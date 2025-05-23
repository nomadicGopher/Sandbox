package main

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

type SE struct {
	NumberOfIncludedSegments    int    `json:"number_of_included_segments"`    // SE01
	TransactionSetControlNumber string `json:"transaction_set_control_number"` // SE02
}

// HL Loop (can repeat for submitter, receiver, subscriber, dependent)
type HL struct {
	HierarchicalIDNumber  string `json:"hierarchical_id_number"`           // HL01
	HierarchicalParentID  string `json:"hierarchical_parent_id,omitempty"` // HL02 (optional, not present for first/parent HL)
	HierarchicalLevelCode string `json:"hierarchical_level_code"`          // HL03
	HierarchicalChildCode string `json:"hierarchical_child_code"`          // HL04
	NM1                   NM1    `json:"nm1"`                              // NM1 loop
	DMG                   *DMG   `json:"dmg,omitempty"`                    // DMG loop (optional)
	DependentLoops        []HL   `json:"dependent_loops,omitempty"`        // Nested HLs for dependents (optional)
}

// NM1: Name Information
type NM1 struct {
	EntityIdentifierCode        string `json:"entity_identifier_code"`                  // NM101
	EntityTypeQualifier         string `json:"entity_type_qualifier"`                   // NM102
	LastNameOrOrgName           string `json:"last_name_or_org_name"`                   // NM103
	FirstName                   string `json:"first_name,omitempty"`                    // NM104 (optional)
	MiddleName                  string `json:"middle_name,omitempty"`                   // NM105 (optional)
	NamePrefix                  string `json:"name_prefix,omitempty"`                   // NM106 (optional)
	NameSuffix                  string `json:"name_suffix,omitempty"`                   // NM107 (optional)
	IdentificationCodeQualifier string `json:"identification_code_qualifier,omitempty"` // NM108 (optional)
	IdentificationCode          string `json:"identification_code,omitempty"`           // NM109 (optional)
	EntityRelationshipCode      string `json:"entity_relationship_code,omitempty"`      // NM110 (optional)
	EntityIdentifierCode2       string `json:"entity_identifier_code_2,omitempty"`      // NM111 (optional)
}

// DMG: Demographic Information
type DMG struct {
	DateTimePeriodFormatQualifier string `json:"date_time_period_format_qualifier,omitempty"` // DMG01 (optional)
	DateOfBirth                   string `json:"date_of_birth,omitempty"`                     // DMG02 (optional)
	GenderCode                    string `json:"gender_code,omitempty"`                       // DMG03 (optional)
	MaritalStatusCode             string `json:"marital_status_code,omitempty"`               // DMG04 (optional)
	CitizenshipCode               string `json:"citizenship_code,omitempty"`                  // DMG05 (optional)
	CountryCode                   string `json:"country_code,omitempty"`                      // DMG06 (optional)
	BaselineResidenceCode         string `json:"baseline_residence_code,omitempty"`           // DMG07 (optional)
	Age                           string `json:"age,omitempty"`                               // DMG08 (optional)
	AgeQualifier                  string `json:"age_qualifier,omitempty"`                     // DMG09 (optional)
}

// HI: Health Care Information Codes (can be multiple per request)
type HI struct {
	HealthCareCodeInfo []HealthCareCodeInfo `json:"health_care_code_info"` // HI01-HI12 (each element = HIxx)
}

type HealthCareCodeInfo struct {
	CodeListQualifierCode string `json:"code_list_qualifier_code"` // HIxx-1
	IndustryCode          string `json:"industry_code"`            // HIxx-2
}

// UM: Utilization Management Information (can be multiple per request)
type UM struct {
	RequestCategoryCode        string `json:"request_category_code"`           // UM01
	CertificationTypeCode      string `json:"certification_type_code"`         // UM02
	ServicesReviewDecisionCode string `json:"services_review_decision_code"`   // UM03
	MonetaryAmount             string `json:"monetary_amount,omitempty"`       // UM04 (optional)
	Quantity                   string `json:"quantity,omitempty"`              // UM05 (optional)
	RejectReasonCode           string `json:"reject_reason_code,omitempty"`    // UM06 (optional)
	ActionCode                 string `json:"action_code,omitempty"`           // UM07 (optional)
	LevelOfServiceCode         string `json:"level_of_service_code,omitempty"` // UM08 (optional)
}

type Transaction278 struct {
	ISA          ISA    `json:"isa"`
	GS           GS     `json:"gs"`
	ST           ST     `json:"st"`
	BHT          BHT    `json:"bht"`
	SubmitterHL  HL     `json:"submitter_hl"`
	ReceiverHL   HL     `json:"receiver_hl"`
	SubscriberHL HL     `json:"subscriber_hl"`
	DependentHLs []HL   `json:"dependent_hls,omitempty"`
	HI           []HI   `json:"hi,omitempty"`
	UM           []UM   `json:"um,omitempty"`
	TRN          []TRN  `json:"trn,omitempty"`
	N3           []N3   `json:"n3,omitempty"`
	N4           []N4   `json:"n4,omitempty"`
	PER          []PER  `json:"per,omitempty"`
	DTP          []DTP  `json:"dtp,omitempty"`
	SV2          []SV2  `json:"sv2,omitempty"`
	HSD          []HSD  `json:"hsd,omitempty"`
	PWK          []PWK  `json:"pwk,omitempty"`
	SE           SE     `json:"se"`
	GE           GE     `json:"ge"`
	IEA          IEA    `json:"iea"`
	Comment      string `json:"comment,omitempty"`
}

// In structs.go, add these struct definitions:

// BHT: Beginning of Hierarchical Transaction
type BHT struct {
	HierarchicalStructureCode string `json:"hierarchical_structure_code"`     // BHT01
	TransactionSetPurposeCode string `json:"transaction_set_purpose_code"`    // BHT02
	ReferenceIdentification   string `json:"reference_identification"`        // BHT03
	Date                      string `json:"date"`                            // BHT04
	Time                      string `json:"time"`                            // BHT05
	TransactionTypeCode       string `json:"transaction_type_code,omitempty"` // BHT06 (optional)
}

// TRN: Trace Number
type TRN struct {
	TraceTypeCode            string `json:"trace_type_code"`                      // TRN01
	ReferenceIdentification  string `json:"reference_identification"`             // TRN02
	OriginatingCompanyID     string `json:"originating_company_id,omitempty"`     // TRN03 (optional)
	ReferenceIdentification2 string `json:"reference_identification_2,omitempty"` // TRN04 (optional)
}

// N3: Address Information
type N3 struct {
	AddressInformation1 string `json:"address_information_1"`           // N301
	AddressInformation2 string `json:"address_information_2,omitempty"` // N302 (optional)
}

// N4: Geographic Location
type N4 struct {
	CityName        string `json:"city_name"`              // N401
	StateOrProvince string `json:"state_or_province_code"` // N402
	PostalCode      string `json:"postal_code"`            // N403
}

// PER: Administrative Communications Contact
type PER struct {
	ContactFunctionCode  string   `json:"contact_function_code"`           // PER01
	Name                 string   `json:"name,omitempty"`                  // PER02
	CommunicationNumbers []string `json:"communication_numbers,omitempty"` // PER03-PER10, collect as array
}

// DTP: Date or Time Period
type DTP struct {
	DateTimeQualifier string `json:"date_time_qualifier"` // DTP01
	FormatQualifier   string `json:"format_qualifier"`    // DTP02
	DateTimePeriod    string `json:"date_time_period"`    // DTP03
}

// SV2: Service Line Information
type SV2 struct {
	ProductOrServiceID            string `json:"product_or_service_id"`              // SV201
	ProcedureModifier             string `json:"procedure_modifier,omitempty"`       // SV202 (optional)
	LineItemChargeAmount          string `json:"line_item_charge_amount"`            // SV203
	UnitOrBasisForMeasurementCode string `json:"unit_or_basis_for_measurement_code"` // SV204
	ServiceUnitCount              string `json:"service_unit_count"`                 // SV205
}

// HSD: Health Care Services Delivery
type HSD struct {
	QuantityQualifier string `json:"quantity_qualifier"` // HSD01
	Quantity          string `json:"quantity"`           // HSD02
}

// PWK: Paperwork
type PWK struct {
	ReportTypeCode         string `json:"report_type_code"`                // PWK01
	ReportTransmissionCode string `json:"report_transmission_code"`        // PWK02
	ReportCopies           string `json:"report_copies,omitempty"`         // PWK03 (optional)
	ReportControlNumber    string `json:"report_control_number,omitempty"` // PWK07 (optional)
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
