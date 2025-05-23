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
	SubmitterHL  HL     `json:"submitter_hl"`
	ReceiverHL   HL     `json:"receiver_hl"`
	SubscriberHL HL     `json:"subscriber_hl"`
	DependentHLs []HL   `json:"dependent_hls,omitempty"`
	HI           []HI   `json:"hi,omitempty"`
	UM           []UM   `json:"um,omitempty"`
	SE           SE     `json:"se"`
	Comment      string `json:"comment,omitempty"`
}
