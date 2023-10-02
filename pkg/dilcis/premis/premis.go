package premis

import (
	"encoding/xml"
)

// Premis ...
type Premis *PremisComplexType

// Object ...
type Object *ObjectComplexType

// Event ...
type Event *EventComplexType

// Agent ...
type Agent *AgentComplexType

// Rights ...
type Rights *RightsComplexType

// PremisComplexType ...
type PremisComplexType struct {
	XMLName     xml.Name             `xml:"premis"`
	VersionAttr string               `xml:"version,attr"`
	Object      []*File              `xml:"object"`
	Event       []*EventComplexType  `xml:"event"`
	Agent       []*AgentComplexType  `xml:"agent"`
	Rights      []*RightsComplexType `xml:"rights"`
}

// ObjectComplexType ...
type ObjectComplexType struct {
	XMLName xml.Name `xml:"object"`
}

// File ...
type File struct {
	XMLName                          xml.Name                                       `xml:"object"`
	XSIType                          string                                         `xml:"xsi:type,attr,omitempty"`
	XmlIDAttr                        string                                         `xml:"xmlID,attr,omitempty"`
	VersionAttr                      string                                         `xml:"version,attr,omitempty"`
	ObjectIdentifier                 []*ObjectIdentifierComplexType                 `xml:"objectIdentifier"`
	PreservationLevel                []*PreservationLevelComplexType                `xml:"preservationLevel"`
	SignificantProperties            []*SignificantPropertiesComplexType            `xml:"significantProperties"`
	ObjectCharacteristics            []*ObjectCharacteristicsComplexType            `xml:"objectCharacteristics"`
	OriginalName                     *OriginalNameComplexType                       `xml:"originalName"`
	Storage                          []*StorageComplexType                          `xml:"storage"`
	SignatureInformation             []*SignatureInformationComplexType             `xml:"signatureInformation"`
	Relationship                     []*RelationshipComplexType                     `xml:"relationship"`
	LinkingEventIdentifier           []*LinkingEventIdentifierComplexType           `xml:"linkingEventIdentifier"`
	LinkingRightsStatementIdentifier []*LinkingRightsStatementIdentifierComplexType `xml:"linkingRightsStatementIdentifier"`
	*ObjectComplexType
}

// Representation ...
type Representation struct {
	XMLName                          xml.Name                                       `xml:"representation"`
	XmlIDAttr                        string                                         `xml:"xmlID,attr,omitempty"`
	VersionAttr                      string                                         `xml:"version,attr,omitempty"`
	ObjectIdentifier                 []*ObjectIdentifierComplexType                 `xml:"objectIdentifier"`
	PreservationLevel                []*PreservationLevelComplexType                `xml:"preservationLevel"`
	SignificantProperties            []*SignificantPropertiesComplexType            `xml:"significantProperties"`
	OriginalName                     *OriginalNameComplexType                       `xml:"originalName"`
	Storage                          []*StorageComplexType                          `xml:"storage"`
	Relationship                     []*RelationshipComplexType                     `xml:"relationship"`
	LinkingEventIdentifier           []*LinkingEventIdentifierComplexType           `xml:"linkingEventIdentifier"`
	LinkingRightsStatementIdentifier []*LinkingRightsStatementIdentifierComplexType `xml:"linkingRightsStatementIdentifier"`
	*ObjectComplexType
}

// Bitstream ...
type Bitstream struct {
	XMLName                          xml.Name                                       `xml:"bitstream"`
	XmlIDAttr                        string                                         `xml:"xmlID,attr,omitempty"`
	VersionAttr                      string                                         `xml:"version,attr,omitempty"`
	ObjectIdentifier                 []*ObjectIdentifierComplexType                 `xml:"objectIdentifier"`
	SignificantProperties            []*SignificantPropertiesComplexType            `xml:"significantProperties"`
	ObjectCharacteristics            []*ObjectCharacteristicsComplexType            `xml:"objectCharacteristics"`
	Storage                          []*StorageComplexType                          `xml:"storage"`
	SignatureInformation             []*SignatureInformationComplexType             `xml:"signatureInformation"`
	Relationship                     []*RelationshipComplexType                     `xml:"relationship"`
	LinkingEventIdentifier           []*LinkingEventIdentifierComplexType           `xml:"linkingEventIdentifier"`
	LinkingRightsStatementIdentifier []*LinkingRightsStatementIdentifierComplexType `xml:"linkingRightsStatementIdentifier"`
	*ObjectComplexType
}

// IntellectualEntity ...
type IntellectualEntity struct {
	XMLName                          xml.Name                                       `xml:"intellectualEntity"`
	XmlIDAttr                        string                                         `xml:"xmlID,attr,omitempty"`
	VersionAttr                      string                                         `xml:"version,attr,omitempty"`
	ObjectIdentifier                 []*ObjectIdentifierComplexType                 `xml:"objectIdentifier"`
	PreservationLevel                []*PreservationLevelComplexType                `xml:"preservationLevel"`
	SignificantProperties            []*SignificantPropertiesComplexType            `xml:"significantProperties"`
	OriginalName                     *OriginalNameComplexType                       `xml:"originalName"`
	EnvironmentFunction              []*EnvironmentFunctionComplexType              `xml:"environmentFunction"`
	EnvironmentDesignation           []*EnvironmentDesignationComplexType           `xml:"environmentDesignation"`
	EnvironmentRegistry              []*EnvironmentRegistryComplexType              `xml:"environmentRegistry"`
	EnvironmentExtension             []*ExtensionComplexType                        `xml:"environmentExtension"`
	Relationship                     []*RelationshipComplexType                     `xml:"relationship"`
	LinkingEventIdentifier           []*LinkingEventIdentifierComplexType           `xml:"linkingEventIdentifier"`
	LinkingRightsStatementIdentifier []*LinkingRightsStatementIdentifierComplexType `xml:"linkingRightsStatementIdentifier"`
	*ObjectComplexType
}

// EventComplexType ...
type EventComplexType struct {
	XMLName                 xml.Name                              `xml:"event"`
	XmlIDAttr               string                                `xml:"xmlID,attr,omitempty"`
	VersionAttr             string                                `xml:"version,attr,omitempty"`
	EventIdentifier         *EventIdentifierComplexType           `xml:"eventIdentifier"`
	EventType               *StringPlusAuthority                  `xml:"eventType"`
	EventDateTime           string                                `xml:"eventDateTime"`
	EventDetailInformation  []*EventDetailInformationComplexType  `xml:"eventDetailInformation"`
	EventOutcomeInformation []*EventOutcomeInformationComplexType `xml:"eventOutcomeInformation"`
	LinkingAgentIdentifier  []*LinkingAgentIdentifierComplexType  `xml:"linkingAgentIdentifier"`
	LinkingObjectIdentifier []*LinkingObjectIdentifierComplexType `xml:"linkingObjectIdentifier"`
}

// AgentComplexType ...
type AgentComplexType struct {
	XMLName                          xml.Name                                       `xml:"agent"`
	XmlIDAttr                        string                                         `xml:"xmlID,attr,omitempty"`
	VersionAttr                      string                                         `xml:"version,attr,omitempty"`
	AgentIdentifier                  []*AgentIdentifierComplexType                  `xml:"agentIdentifier"`
	AgentName                        []*StringPlusAuthority                         `xml:"agentName"`
	AgentType                        *StringPlusAuthority                           `xml:"agentType"`
	AgentVersion                     string                                         `xml:"agentVersion"`
	AgentNote                        []string                                       `xml:"agentNote"`
	AgentExtension                   []*ExtensionComplexType                        `xml:"agentExtension"`
	LinkingEventIdentifier           []*LinkingEventIdentifierComplexType           `xml:"linkingEventIdentifier"`
	LinkingRightsStatementIdentifier []*LinkingRightsStatementIdentifierComplexType `xml:"linkingRightsStatementIdentifier"`
	LinkingEnvironmentIdentifier     []*LinkingEnvironmentIdentifierComplexType     `xml:"linkingEnvironmentIdentifier"`
}

// RightsComplexType ...
type RightsComplexType struct {
	XMLName         xml.Name                      `xml:"rights"`
	XmlIDAttr       string                        `xml:"xmlID,attr,omitempty"`
	VersionAttr     string                        `xml:"version,attr,omitempty"`
	RightsStatement []*RightsStatementComplexType `xml:"rightsStatement"`
	RightsExtension []*ExtensionComplexType       `xml:"rightsExtension"`
}

// AgentIdentifierComplexType ...
type AgentIdentifierComplexType struct {
	XMLName              xml.Name             `xml:"agentIdentifier"`
	SimpleLinkAttr       string               `xml:"simpleLink,attr,omitempty"`
	AgentIdentifierType  *StringPlusAuthority `xml:"agentIdentifierType"`
	AgentIdentifierValue string               `xml:"agentIdentifierValue"`
}

// CompositionLevelComplexType ...
type CompositionLevelComplexType struct {
	XMLName     xml.Name `xml:"compositionLevel"`
	UnknownAttr string   `xml:"unknown,attr,omitempty"`
	Value       int      `xml:",chardata"`
}

// ContentLocationComplexType ...
type ContentLocationComplexType struct {
	XMLName              xml.Name             `xml:"contentLocation"`
	SimpleLinkAttr       string               `xml:"simpleLink,attr,omitempty"`
	ContentLocationType  *StringPlusAuthority `xml:"contentLocationType"`
	ContentLocationValue string               `xml:"contentLocationValue"`
}

// CopyrightDocumentationIdentifierComplexType ...
type CopyrightDocumentationIdentifierComplexType struct {
	XMLName                               xml.Name             `xml:"copyrightDocumentationIdentifier"`
	CopyrightDocumentationIdentifierType  *StringPlusAuthority `xml:"copyrightDocumentationIdentifierType"`
	CopyrightDocumentationIdentifierValue string               `xml:"copyrightDocumentationIdentifierValue"`
	CopyrightDocumentationRole            *StringPlusAuthority `xml:"copyrightDocumentationRole"`
}

// CopyrightInformationComplexType ...
type CopyrightInformationComplexType struct {
	XMLName                          xml.Name                                       `xml:"copyrightInformation"`
	CopyrightStatus                  *StringPlusAuthority                           `xml:"copyrightStatus"`
	CopyrightJurisdiction            *CountryCode                                   `xml:"copyrightJurisdiction"`
	CopyrightStatusDeterminationDate *EdtfSimpleType                                `xml:"copyrightStatusDeterminationDate"`
	CopyrightNote                    []string                                       `xml:"copyrightNote"`
	CopyrightDocumentationIdentifier []*CopyrightDocumentationIdentifierComplexType `xml:"copyrightDocumentationIdentifier"`
	CopyrightApplicableDates         *StartAndEndDateComplexType                    `xml:"copyrightApplicableDates"`
}

// CreatingApplicationComplexType ...
type CreatingApplicationComplexType struct {
	XMLName                      xml.Name                `xml:"creatingApplication"`
	CreatingApplicationName      *StringPlusAuthority    `xml:"creatingApplicationName"`
	CreatingApplicationVersion   string                  `xml:"creatingApplicationVersion"`
	DateCreatedByApplication     *EdtfSimpleType         `xml:"dateCreatedByApplication"`
	CreatingApplicationExtension []*ExtensionComplexType `xml:"creatingApplicationExtension"`
}

// EnvironmentFunctionComplexType ...
type EnvironmentFunctionComplexType struct {
	XMLName                  xml.Name             `xml:"environmentFunction"`
	EnvironmentFunctionType  *StringPlusAuthority `xml:"environmentFunctionType"`
	EnvironmentFunctionLevel string               `xml:"environmentFunctionLevel"`
}

// EnvironmentDesignationComplexType ...
type EnvironmentDesignationComplexType struct {
	XMLName                         xml.Name             `xml:"environmentDesignation"`
	EnvironmentName                 *StringPlusAuthority `xml:"environmentName"`
	EnvironmentVersion              string               `xml:"environmentVersion"`
	EnvironmentOrigin               string               `xml:"environmentOrigin"`
	EnvironmentDesignationNote      []string             `xml:"environmentDesignationNote"`
	EnvironmentDesignationExtension []string             `xml:"environmentDesignationExtension"`
}

// EnvironmentRegistryComplexType ...
type EnvironmentRegistryComplexType struct {
	XMLName                 xml.Name             `xml:"environmentRegistry"`
	EnvironmentRegistryName string               `xml:"environmentRegistryName"`
	EnvironmentRegistryKey  string               `xml:"environmentRegistryKey"`
	EnvironmentRegistryRole *StringPlusAuthority `xml:"environmentRegistryRole"`
}

// EventDetailInformationComplexType ...
type EventDetailInformationComplexType struct {
	XMLName              xml.Name                `xml:"eventDetailInformation"`
	EventDetail          string                  `xml:"eventDetail"`
	EventDetailExtension []*ExtensionComplexType `xml:"eventDetailExtension"`
}

// EventIdentifierComplexType ...
type EventIdentifierComplexType struct {
	XMLName              xml.Name             `xml:"eventIdentifier"`
	SimpleLinkAttr       string               `xml:"simpleLink,attr,omitempty"`
	EventIdentifierType  *StringPlusAuthority `xml:"eventIdentifierType"`
	EventIdentifierValue string               `xml:"eventIdentifierValue"`
}

// EventOutcomeDetailComplexType ...
type EventOutcomeDetailComplexType struct {
	XMLName                     xml.Name                `xml:"eventOutcomeDetail"`
	EventOutcomeDetailNote      string                  `xml:"eventOutcomeDetailNote"`
	EventOutcomeDetailExtension []*ExtensionComplexType `xml:"eventOutcomeDetailExtension"`
}

// EventOutcomeInformationComplexType ...
type EventOutcomeInformationComplexType struct {
	XMLName            xml.Name                         `xml:"eventOutcomeInformation"`
	EventOutcome       *StringPlusAuthority             `xml:"eventOutcome"`
	EventOutcomeDetail []*EventOutcomeDetailComplexType `xml:"eventOutcomeDetail"`
}

// FixityComplexType ...
type FixityComplexType struct {
	XMLName                 xml.Name             `xml:"fixity"`
	MessageDigestAlgorithm  *StringPlusAuthority `xml:"messageDigestAlgorithm"`
	MessageDigest           string               `xml:"messageDigest"`
	MessageDigestOriginator *StringPlusAuthority `xml:"messageDigestOriginator"`
}

// FormatComplexType ...
type FormatComplexType struct {
	XMLName           xml.Name                      `xml:"format"`
	FormatDesignation *FormatDesignationComplexType `xml:"formatDesignation"`
	FormatRegistry    *FormatRegistryComplexType    `xml:"formatRegistry"`
	FormatNote        []string                      `xml:"formatNote"`
}

// FormatDesignationComplexType ...
type FormatDesignationComplexType struct {
	XMLName       xml.Name             `xml:"formatDesignation"`
	FormatName    *StringPlusAuthority `xml:"formatName"`
	FormatVersion string               `xml:"formatVersion"`
}

// FormatRegistryComplexType ...
type FormatRegistryComplexType struct {
	XMLName            xml.Name             `xml:"formatRegistry"`
	SimpleLinkAttr     string               `xml:"simpleLink,attr,omitempty"`
	FormatRegistryName *StringPlusAuthority `xml:"formatRegistryName"`
	FormatRegistryKey  *StringPlusAuthority `xml:"formatRegistryKey"`
	FormatRegistryRole *StringPlusAuthority `xml:"formatRegistryRole"`
}

// InhibitorsComplexType ...
type InhibitorsComplexType struct {
	XMLName         xml.Name               `xml:"inhibitors"`
	InhibitorType   *StringPlusAuthority   `xml:"inhibitorType"`
	InhibitorTarget []*StringPlusAuthority `xml:"inhibitorTarget"`
	InhibitorKey    string                 `xml:"inhibitorKey"`
}

// LicenseDocumentationIdentifierComplexType ...
type LicenseDocumentationIdentifierComplexType struct {
	XMLName                             xml.Name             `xml:"licenseDocumentationIdentifier"`
	LicenseDocumentationIdentifierType  *StringPlusAuthority `xml:"licenseDocumentationIdentifierType"`
	LicenseDocumentationIdentifierValue string               `xml:"licenseDocumentationIdentifierValue"`
	LicenseDocumentationRole            *StringPlusAuthority `xml:"licenseDocumentationRole"`
}

// LicenseInformationComplexType ...
type LicenseInformationComplexType struct {
	XMLName                        xml.Name                                     `xml:"licenseInformation"`
	LicenseDocumentationIdentifier []*LicenseDocumentationIdentifierComplexType `xml:"licenseDocumentationIdentifier"`
	LicenseTerms                   string                                       `xml:"licenseTerms"`
	LicenseNote                    []string                                     `xml:"licenseNote"`
	LicenseApplicableDates         *StartAndEndDateComplexType                  `xml:"licenseApplicableDates"`
}

// LinkingAgentIdentifierComplexType ...
type LinkingAgentIdentifierComplexType struct {
	XMLName                     xml.Name               `xml:"linkingAgentIdentifier"`
	LinkAgentXmlIDAttr          string                 `xml:"LinkAgentXmlID,attr,omitempty"`
	SimpleLinkAttr              string                 `xml:"simpleLink,attr,omitempty"`
	LinkingAgentIdentifierType  *StringPlusAuthority   `xml:"linkingAgentIdentifierType"`
	LinkingAgentIdentifierValue string                 `xml:"linkingAgentIdentifierValue"`
	LinkingAgentRole            []*StringPlusAuthority `xml:"linkingAgentRole"`
}

// LinkingEnvironmentIdentifierComplexType ...
type LinkingEnvironmentIdentifierComplexType struct {
	XMLName                           xml.Name               `xml:"linkingEnvironmentIdentifier"`
	LinkEventXmlIDAttr                string                 `xml:"LinkEventXmlID,attr,omitempty"`
	SimpleLinkAttr                    string                 `xml:"simpleLink,attr,omitempty"`
	LinkingEnvironmentIdentifierType  string                 `xml:"linkingEnvironmentIdentifierType"`
	LinkingEnvironmentIdentifierValue string                 `xml:"linkingEnvironmentIdentifierValue"`
	LinkingEnvironmentRole            []*StringPlusAuthority `xml:"linkingEnvironmentRole"`
}

// LinkingEventIdentifierComplexType ...
type LinkingEventIdentifierComplexType struct {
	XMLName                     xml.Name             `xml:"linkingEventIdentifier"`
	LinkEventXmlIDAttr          string               `xml:"LinkEventXmlID,attr,omitempty"`
	SimpleLinkAttr              string               `xml:"simpleLink,attr,omitempty"`
	LinkingEventIdentifierType  *StringPlusAuthority `xml:"linkingEventIdentifierType"`
	LinkingEventIdentifierValue string               `xml:"linkingEventIdentifierValue"`
}

// LinkingObjectIdentifierComplexType ...
type LinkingObjectIdentifierComplexType struct {
	XMLName                      xml.Name               `xml:"linkingObjectIdentifier"`
	LinkObjectXmlIDAttr          string                 `xml:"LinkObjectXmlID,attr,omitempty"`
	SimpleLinkAttr               string                 `xml:"simpleLink,attr,omitempty"`
	LinkingObjectIdentifierType  *StringPlusAuthority   `xml:"linkingObjectIdentifierType"`
	LinkingObjectIdentifierValue string                 `xml:"linkingObjectIdentifierValue"`
	LinkingObjectRole            []*StringPlusAuthority `xml:"linkingObjectRole"`
}

// LinkingRightsStatementIdentifierComplexType ...
type LinkingRightsStatementIdentifierComplexType struct {
	XMLName                               xml.Name             `xml:"linkingRightsStatementIdentifier"`
	LinkPermissionStatementXmlIDAttr      string               `xml:"LinkPermissionStatementXmlID,attr,omitempty"`
	SimpleLinkAttr                        string               `xml:"simpleLink,attr,omitempty"`
	LinkingRightsStatementIdentifierType  *StringPlusAuthority `xml:"linkingRightsStatementIdentifierType"`
	LinkingRightsStatementIdentifierValue string               `xml:"linkingRightsStatementIdentifierValue"`
}

// ObjectCharacteristicsComplexType ...
type ObjectCharacteristicsComplexType struct {
	XMLName                        xml.Name                          `xml:"objectCharacteristics"`
	CompositionLevel               *CompositionLevelComplexType      `xml:"compositionLevel"`
	Fixity                         []*FixityComplexType              `xml:"fixity"`
	Size                           int64                             `xml:"size"`
	Format                         []*FormatComplexType              `xml:"format"`
	CreatingApplication            []*CreatingApplicationComplexType `xml:"creatingApplication"`
	Inhibitors                     []*InhibitorsComplexType          `xml:"inhibitors"`
	ObjectCharacteristicsExtension []*ExtensionComplexType           `xml:"objectCharacteristicsExtension"`
}

// ObjectIdentifierComplexType ...
type ObjectIdentifierComplexType struct {
	XMLName               xml.Name             `xml:"objectIdentifier"`
	SimpleLinkAttr        string               `xml:"simpleLink,attr,omitempty"`
	ObjectIdentifierType  *StringPlusAuthority `xml:"objectIdentifierType"`
	ObjectIdentifierValue string               `xml:"objectIdentifierValue"`
}

// OriginalNameComplexType ...
type OriginalNameComplexType struct {
	XMLName        xml.Name `xml:"originalName"`
	SimpleLinkAttr string   `xml:"simpleLink,attr,omitempty"`
	Value          string   `xml:",chardata"`
}

// OtherRightsDocumentationIdentifierComplexType ...
type OtherRightsDocumentationIdentifierComplexType struct {
	XMLName                                 xml.Name             `xml:"otherRightsDocumentationIdentifier"`
	OtherRightsDocumentationIdentifierType  *StringPlusAuthority `xml:"otherRightsDocumentationIdentifierType"`
	OtherRightsDocumentationIdentifierValue string               `xml:"otherRightsDocumentationIdentifierValue"`
	OtherRightsDocumentationRole            *StringPlusAuthority `xml:"otherRightsDocumentationRole"`
}

// OtherRightsInformationComplexType ...
type OtherRightsInformationComplexType struct {
	XMLName                            xml.Name                                         `xml:"otherRightsInformation"`
	OtherRightsDocumentationIdentifier []*OtherRightsDocumentationIdentifierComplexType `xml:"otherRightsDocumentationIdentifier"`
	OtherRightsBasis                   *StringPlusAuthority                             `xml:"otherRightsBasis"`
	OtherRightsApplicableDates         *StartAndEndDateComplexType                      `xml:"otherRightsApplicableDates"`
	OtherRightsNote                    []string                                         `xml:"otherRightsNote"`
}

// PreservationLevelComplexType ...
type PreservationLevelComplexType struct {
	XMLName                       xml.Name             `xml:"preservationLevel"`
	PreservationLevelType         *StringPlusAuthority `xml:"preservationLevelType"`
	PreservationLevelValue        *StringPlusAuthority `xml:"preservationLevelValue"`
	PreservationLevelRole         *StringPlusAuthority `xml:"preservationLevelRole"`
	PreservationLevelRationale    []string             `xml:"preservationLevelRationale"`
	PreservationLevelDateAssigned *EdtfSimpleType      `xml:"preservationLevelDateAssigned"`
}

// RelatedEventIdentifierComplexType ...
type RelatedEventIdentifierComplexType struct {
	XMLName                     xml.Name             `xml:"relatedEventIdentifier"`
	RelEventXmlIDAttr           string               `xml:"RelEventXmlID,attr,omitempty"`
	SimpleLinkAttr              string               `xml:"simpleLink,attr,omitempty"`
	RelatedEventIdentifierType  *StringPlusAuthority `xml:"relatedEventIdentifierType"`
	RelatedEventIdentifierValue string               `xml:"relatedEventIdentifierValue"`
	RelatedEventSequence        int                  `xml:"relatedEventSequence"`
}

// RelatedObjectIdentifierComplexType ...
type RelatedObjectIdentifierComplexType struct {
	XMLName                      xml.Name             `xml:"relatedObjectIdentifier"`
	RelObjectXmlIDAttr           string               `xml:"RelObjectXmlID,attr,omitempty"`
	SimpleLinkAttr               string               `xml:"simpleLink,attr,omitempty"`
	RelatedObjectIdentifierType  *StringPlusAuthority `xml:"relatedObjectIdentifierType"`
	RelatedObjectIdentifierValue string               `xml:"relatedObjectIdentifierValue"`
	RelatedObjectSequence        int                  `xml:"relatedObjectSequence"`
}

// RelationshipComplexType ...
type RelationshipComplexType struct {
	XMLName                          xml.Name                              `xml:"relationship"`
	RelationshipType                 *StringPlusAuthority                  `xml:"relationshipType"`
	RelationshipSubType              *StringPlusAuthority                  `xml:"relationshipSubType"`
	RelatedObjectIdentifier          []*RelatedObjectIdentifierComplexType `xml:"relatedObjectIdentifier"`
	RelatedEventIdentifier           []*RelatedEventIdentifierComplexType  `xml:"relatedEventIdentifier"`
	RelatedEnvironmentPurpose        []*StringPlusAuthority                `xml:"relatedEnvironmentPurpose"`
	RelatedEnvironmentCharacteristic *StringPlusAuthority                  `xml:"relatedEnvironmentCharacteristic"`
}

// RightsGrantedComplexType ...
type RightsGrantedComplexType struct {
	XMLName           xml.Name                    `xml:"rightsGranted"`
	Act               *StringPlusAuthority        `xml:"act"`
	Restriction       []*StringPlusAuthority      `xml:"restriction"`
	TermOfGrant       *StartAndEndDateComplexType `xml:"termOfGrant"`
	TermOfRestriction *StartAndEndDateComplexType `xml:"termOfRestriction"`
	RightsGrantedNote []string                    `xml:"rightsGrantedNote"`
}

// RightsStatementComplexType ...
type RightsStatementComplexType struct {
	XMLName                   xml.Name                              `xml:"rightsStatement"`
	RightsStatementIdentifier *RightsStatementIdentifierComplexType `xml:"rightsStatementIdentifier"`
	RightsBasis               *StringPlusAuthority                  `xml:"rightsBasis"`
	CopyrightInformation      *CopyrightInformationComplexType      `xml:"copyrightInformation"`
	LicenseInformation        *LicenseInformationComplexType        `xml:"licenseInformation"`
	StatuteInformation        []*StatuteInformationComplexType      `xml:"statuteInformation"`
	OtherRightsInformation    *OtherRightsInformationComplexType    `xml:"otherRightsInformation"`
	RightsGranted             []*RightsGrantedComplexType           `xml:"rightsGranted"`
	LinkingObjectIdentifier   []*LinkingObjectIdentifierComplexType `xml:"linkingObjectIdentifier"`
	LinkingAgentIdentifier    []*LinkingAgentIdentifierComplexType  `xml:"linkingAgentIdentifier"`
}

// RightsStatementIdentifierComplexType ...
type RightsStatementIdentifierComplexType struct {
	XMLName                        xml.Name             `xml:"rightsStatementIdentifier"`
	SimpleLinkAttr                 string               `xml:"simpleLink,attr,omitempty"`
	RightsStatementIdentifierType  *StringPlusAuthority `xml:"rightsStatementIdentifierType"`
	RightsStatementIdentifierValue string               `xml:"rightsStatementIdentifierValue"`
}

// SignatureComplexType ...
type SignatureComplexType struct {
	XMLName                  xml.Name                `xml:"signature"`
	SignatureEncoding        *StringPlusAuthority    `xml:"signatureEncoding"`
	Signer                   *StringPlusAuthority    `xml:"signer"`
	SignatureMethod          *StringPlusAuthority    `xml:"signatureMethod"`
	SignatureValue           string                  `xml:"signatureValue"`
	SignatureValidationRules *StringPlusAuthority    `xml:"signatureValidationRules"`
	SignatureProperties      []string                `xml:"signatureProperties"`
	KeyInformation           []*ExtensionComplexType `xml:"keyInformation"`
}

// SignatureInformationComplexType ...
type SignatureInformationComplexType struct {
	XMLName                       xml.Name                `xml:"signatureInformation"`
	Signature                     *SignatureComplexType   `xml:"signature"`
	SignatureInformationExtension []*ExtensionComplexType `xml:"signatureInformationExtension"`
}

// SignificantPropertiesComplexType ...
type SignificantPropertiesComplexType struct {
	XMLName                        xml.Name                `xml:"significantProperties"`
	SignificantPropertiesType      *StringPlusAuthority    `xml:"significantPropertiesType"`
	SignificantPropertiesValue     string                  `xml:"significantPropertiesValue"`
	SignificantPropertiesExtension []*ExtensionComplexType `xml:"significantPropertiesExtension"`
}

// StartAndEndDateComplexType ...
type StartAndEndDateComplexType struct {
	XMLName   xml.Name        `xml:"startAndEndDate"`
	StartDate *EdtfSimpleType `xml:"startDate"`
	EndDate   *EdtfSimpleType `xml:"endDate"`
}

// StatuteDocumentationIdentifierComplexType ...
type StatuteDocumentationIdentifierComplexType struct {
	XMLName                             xml.Name             `xml:"statuteDocumentationIdentifier"`
	StatuteDocumentationIdentifierType  *StringPlusAuthority `xml:"statuteDocumentationIdentifierType"`
	StatuteDocumentationIdentifierValue string               `xml:"statuteDocumentationIdentifierValue"`
	StatuteDocumentationRole            *StringPlusAuthority `xml:"statuteDocumentationRole"`
}

// StatuteInformationComplexType ...
type StatuteInformationComplexType struct {
	XMLName                             xml.Name                                     `xml:"statuteInformation"`
	StatuteJurisdiction                 *CountryCode                                 `xml:"statuteJurisdiction"`
	StatuteCitation                     *StringPlusAuthority                         `xml:"statuteCitation"`
	StatuteInformationDeterminationDate *EdtfSimpleType                              `xml:"statuteInformationDeterminationDate"`
	StatuteNote                         []string                                     `xml:"statuteNote"`
	StatuteDocumentationIdentifier      []*StatuteDocumentationIdentifierComplexType `xml:"statuteDocumentationIdentifier"`
	StatuteApplicableDates              *StartAndEndDateComplexType                  `xml:"statuteApplicableDates"`
}

// StorageComplexType ...
type StorageComplexType struct {
	XMLName         xml.Name                    `xml:"storage"`
	ContentLocation *ContentLocationComplexType `xml:"contentLocation"`
	StorageMedium   *StringPlusAuthority        `xml:"storageMedium"`
}

// AgentIdentifierValue ...
type AgentIdentifierValue string

// AgentNote ...
type AgentNote string

// AgentVersion ...
type AgentVersion string

// ContentLocationValue ...
type ContentLocationValue string

// CopyrightDocumentationIdentifierValue ...
type CopyrightDocumentationIdentifierValue string

// CopyrightNote ...
type CopyrightNote string

// CreatingApplicationVersion ...
type CreatingApplicationVersion string

// EnvironmentDesignationExtension ...
type EnvironmentDesignationExtension string

// EnvironmentDesignationNote ...
type EnvironmentDesignationNote string

// EnvironmentFunctionLevel ...
type EnvironmentFunctionLevel string

// EnvironmentNote ...
type EnvironmentNote string

// EnvironmentOrigin ...
type EnvironmentOrigin string

// EnvironmentRegistryKey ...
type EnvironmentRegistryKey string

// EnvironmentRegistryName ...
type EnvironmentRegistryName string

// EnvironmentVersion ...
type EnvironmentVersion string

// EventDetail ...
type EventDetail string

// EventIdentifierValue ...
type EventIdentifierValue string

// EventOutcomeDetailNote ...
type EventOutcomeDetailNote string

// FormatNote ...
type FormatNote string

// FormatVersion ...
type FormatVersion string

// HwOtherInformation ...
type HwOtherInformation string

// InhibitorKey ...
type InhibitorKey string

// LicenseDocumentationIdentifierValue ...
type LicenseDocumentationIdentifierValue string

// LicenseIdentifierValue ...
type LicenseIdentifierValue string

// LicenseNote ...
type LicenseNote string

// LicenseTerms ...
type LicenseTerms string

// LinkingAgentIdentifierValue ...
type LinkingAgentIdentifierValue string

// LinkingEnvironmentIdentifierType ...
type LinkingEnvironmentIdentifierType string

// LinkingEnvironmentIdentifierValue ...
type LinkingEnvironmentIdentifierValue string

// LinkingEventIdentifierValue ...
type LinkingEventIdentifierValue string

// LinkingObjectIdentifierValue ...
type LinkingObjectIdentifierValue string

// LinkingRightsStatementIdentifierValue ...
type LinkingRightsStatementIdentifierValue string

// MessageDigest ...
type MessageDigest string

// ObjectIdentifierValue ...
type ObjectIdentifierValue string

// OtherRightsDocumentationIdentifierValue ...
type OtherRightsDocumentationIdentifierValue string

// OtherRightsNote ...
type OtherRightsNote string

// PreservationLevelRationale ...
type PreservationLevelRationale string

// RelatedEventIdentifierValue ...
type RelatedEventIdentifierValue string

// RelatedObjectIdentifierValue ...
type RelatedObjectIdentifierValue string

// RightsGrantedNote ...
type RightsGrantedNote string

// RightsStatementIdentifierValue ...
type RightsStatementIdentifierValue string

// SignatureProperties ...
type SignatureProperties string

// SignatureValue ...
type SignatureValue string

// SignificantPropertiesValue ...
type SignificantPropertiesValue string

// StatuteDocumentationIdentifierValue ...
type StatuteDocumentationIdentifierValue string

// StatuteNote ...
type StatuteNote string

// SwVersion ...
type SwVersion string

// SwOtherInformation ...
type SwOtherInformation string

// Act ...
type Act *StringPlusAuthority

// AgentIdentifierType ...
type AgentIdentifierType *StringPlusAuthority

// AgentName ...
type AgentName *StringPlusAuthority

// AgentType ...
type AgentType *StringPlusAuthority

// ContentLocationType ...
type ContentLocationType *StringPlusAuthority

// CopyrightDocumentationIdentifierType ...
type CopyrightDocumentationIdentifierType *StringPlusAuthority

// CopyrightDocumentationRole ...
type CopyrightDocumentationRole *StringPlusAuthority

// CopyrightStatus ...
type CopyrightStatus *StringPlusAuthority

// CreatingApplicationName ...
type CreatingApplicationName *StringPlusAuthority

// EnvironmentCharacteristic ...
type EnvironmentCharacteristic *StringPlusAuthority

// EnvironmentFunctionType ...
type EnvironmentFunctionType *StringPlusAuthority

// EnvironmentName ...
type EnvironmentName *StringPlusAuthority

// EnvironmentRegistryRole ...
type EnvironmentRegistryRole *StringPlusAuthority

// EnvironmentPurpose ...
type EnvironmentPurpose *StringPlusAuthority

// EventIdentifierType ...
type EventIdentifierType *StringPlusAuthority

// EventOutcome ...
type EventOutcome *StringPlusAuthority

// EventType ...
type EventType *StringPlusAuthority

// FormatName ...
type FormatName *StringPlusAuthority

// FormatRegistryName ...
type FormatRegistryName *StringPlusAuthority

// FormatRegistryKey ...
type FormatRegistryKey *StringPlusAuthority

// FormatRegistryRole ...
type FormatRegistryRole *StringPlusAuthority

// HwName ...
type HwName *StringPlusAuthority

// HwType ...
type HwType *StringPlusAuthority

// InhibitorTarget ...
type InhibitorTarget *StringPlusAuthority

// InhibitorType ...
type InhibitorType *StringPlusAuthority

// LicenseDocumentationIdentifierType ...
type LicenseDocumentationIdentifierType *StringPlusAuthority

// LicenseDocumentationRole ...
type LicenseDocumentationRole *StringPlusAuthority

// LicenseIdentifierType ...
type LicenseIdentifierType *StringPlusAuthority

// LinkingAgentIdentifierType ...
type LinkingAgentIdentifierType *StringPlusAuthority

// LinkingAgentRole ...
type LinkingAgentRole *StringPlusAuthority

// LinkingEventIdentifierType ...
type LinkingEventIdentifierType *StringPlusAuthority

// LinkingEnvironmentRole ...
type LinkingEnvironmentRole *StringPlusAuthority

// LinkingObjectIdentifierType ...
type LinkingObjectIdentifierType *StringPlusAuthority

// LinkingObjectRole ...
type LinkingObjectRole *StringPlusAuthority

// LinkingRightsStatementIdentifierType ...
type LinkingRightsStatementIdentifierType *StringPlusAuthority

// MessageDigestAlgorithm ...
type MessageDigestAlgorithm *StringPlusAuthority

// MessageDigestOriginator ...
type MessageDigestOriginator *StringPlusAuthority

// ObjectIdentifierType ...
type ObjectIdentifierType *StringPlusAuthority

// OtherRightsBasis ...
type OtherRightsBasis *StringPlusAuthority

// OtherRightsDocumentationRole ...
type OtherRightsDocumentationRole *StringPlusAuthority

// OtherRightsDocumentationIdentifierType ...
type OtherRightsDocumentationIdentifierType *StringPlusAuthority

// PreservationLevelType ...
type PreservationLevelType *StringPlusAuthority

// PreservationLevelValue ...
type PreservationLevelValue *StringPlusAuthority

// PreservationLevelRole ...
type PreservationLevelRole *StringPlusAuthority

// RelatedEventIdentifierType ...
type RelatedEventIdentifierType *StringPlusAuthority

// RelatedEnvironmentPurpose ...
type RelatedEnvironmentPurpose *StringPlusAuthority

// RelatedEnvironmentCharacteristic ...
type RelatedEnvironmentCharacteristic *StringPlusAuthority

// RelatedObjectIdentifierType ...
type RelatedObjectIdentifierType *StringPlusAuthority

// RelationshipType ...
type RelationshipType *StringPlusAuthority

// RelationshipSubType ...
type RelationshipSubType *StringPlusAuthority

// Restriction ...
type Restriction *StringPlusAuthority

// RightsBasis ...
type RightsBasis *StringPlusAuthority

// RightsStatementIdentifierType ...
type RightsStatementIdentifierType *StringPlusAuthority

// SignatureEncoding ...
type SignatureEncoding *StringPlusAuthority

// SignatureMethod ...
type SignatureMethod *StringPlusAuthority

// SignatureValidationRules ...
type SignatureValidationRules *StringPlusAuthority

// Signer ...
type Signer *StringPlusAuthority

// SignificantPropertiesType ...
type SignificantPropertiesType *StringPlusAuthority

// StorageMedium ...
type StorageMedium *StringPlusAuthority

// StatuteCitation ...
type StatuteCitation *StringPlusAuthority

// StatuteDocumentationIdentifierType ...
type StatuteDocumentationIdentifierType *StringPlusAuthority

// StatuteDocumentationRole ...
type StatuteDocumentationRole *StringPlusAuthority

// SwName ...
type SwName *StringPlusAuthority

// SwType ...
type SwType *StringPlusAuthority

// SwDependency ...
type SwDependency *StringPlusAuthority

// CopyrightJurisdiction ...
type CopyrightJurisdiction *CountryCode

// StatuteJurisdiction ...
type StatuteJurisdiction *CountryCode

// AgentIdentifier ...
type AgentIdentifier *AgentIdentifierComplexType

// ContentLocation ...
type ContentLocation *ContentLocationComplexType

// CompositionLevel ...
type CompositionLevel *CompositionLevelComplexType

// CopyrightDocumentationIdentifier ...
type CopyrightDocumentationIdentifier *CopyrightDocumentationIdentifierComplexType

// CopyrightInformation ...
type CopyrightInformation *CopyrightInformationComplexType

// CreatingApplication ...
type CreatingApplication *CreatingApplicationComplexType

// EnvironmentFunction ...
type EnvironmentFunction *EnvironmentFunctionComplexType

// EnvironmentDesignation ...
type EnvironmentDesignation *EnvironmentDesignationComplexType

// EnvironmentRegistry ...
type EnvironmentRegistry *EnvironmentRegistryComplexType

// EventDetailInformation ...
type EventDetailInformation *EventDetailInformationComplexType

// EventIdentifier ...
type EventIdentifier *EventIdentifierComplexType

// EventOutcomeDetail ...
type EventOutcomeDetail *EventOutcomeDetailComplexType

// EventOutcomeInformation ...
type EventOutcomeInformation *EventOutcomeInformationComplexType

// Fixity ...
type Fixity *FixityComplexType

// Format ...
type Format *FormatComplexType

// FormatDesignation ...
type FormatDesignation *FormatDesignationComplexType

// FormatRegistry ...
type FormatRegistry *FormatRegistryComplexType

// Inhibitors ...
type Inhibitors *InhibitorsComplexType

// LicenseDocumentationIdentifier ...
type LicenseDocumentationIdentifier *LicenseDocumentationIdentifierComplexType

// LicenseInformation ...
type LicenseInformation *LicenseInformationComplexType

// LinkingAgentIdentifier ...
type LinkingAgentIdentifier *LinkingAgentIdentifierComplexType

// LinkingEnvironmentIdentifier ...
type LinkingEnvironmentIdentifier *LinkingEnvironmentIdentifierComplexType

// LinkingEventIdentifier ...
type LinkingEventIdentifier *LinkingEventIdentifierComplexType

// LinkingObjectIdentifier ...
type LinkingObjectIdentifier *LinkingObjectIdentifierComplexType

// LinkingRightsStatementIdentifier ...
type LinkingRightsStatementIdentifier *LinkingRightsStatementIdentifierComplexType

// ObjectCharacteristics ...
type ObjectCharacteristics *ObjectCharacteristicsComplexType

// ObjectIdentifier ...
type ObjectIdentifier *ObjectIdentifierComplexType

// OriginalName ...
type OriginalName *OriginalNameComplexType

// OtherRightsDocumentationIdentifier ...
type OtherRightsDocumentationIdentifier *OtherRightsDocumentationIdentifierComplexType

// OtherRightsInformation ...
type OtherRightsInformation *OtherRightsInformationComplexType

// PreservationLevel ...
type PreservationLevel *PreservationLevelComplexType

// RelatedEventIdentifier ...
type RelatedEventIdentifier *RelatedEventIdentifierComplexType

// RelatedObjectIdentifier ...
type RelatedObjectIdentifier *RelatedObjectIdentifierComplexType

// Relationship ...
type Relationship *RelationshipComplexType

// RightsGranted ...
type RightsGranted *RightsGrantedComplexType

// RightsStatement ...
type RightsStatement *RightsStatementComplexType

// RightsStatementIdentifier ...
type RightsStatementIdentifier *RightsStatementIdentifierComplexType

// Signature ...
type Signature *SignatureComplexType

// SignatureInformation ...
type SignatureInformation *SignatureInformationComplexType

// SignificantProperties ...
type SignificantProperties *SignificantPropertiesComplexType

// StatuteDocumentationIdentifier ...
type StatuteDocumentationIdentifier *StatuteDocumentationIdentifierComplexType

// StatuteInformation ...
type StatuteInformation *StatuteInformationComplexType

// Storage ...
type Storage *StorageComplexType

// RelatedEventSequence ...
type RelatedEventSequence int

// RelatedObjectSequence ...
type RelatedObjectSequence int

// Size ...
type Size int64

// DateCreatedByApplication ...
type DateCreatedByApplication string

// EndDate ...
type EndDate string

// CopyrightApplicableDates ...
type CopyrightApplicableDates *StartAndEndDateComplexType

// CopyrightStatusDeterminationDate ...
type CopyrightStatusDeterminationDate string

// EventDateTime ...
type EventDateTime string

// LicenseApplicableDates ...
type LicenseApplicableDates *StartAndEndDateComplexType

// PreservationLevelDateAssigned ...
type PreservationLevelDateAssigned string

// StartDate ...
type StartDate string

// OtherRightsApplicableDates ...
type OtherRightsApplicableDates *StartAndEndDateComplexType

// StatuteApplicableDates ...
type StatuteApplicableDates *StartAndEndDateComplexType

// StatuteInformationDeterminationDate ...
type StatuteInformationDeterminationDate string

// TermOfGrant ...
type TermOfGrant *StartAndEndDateComplexType

// TermOfRestriction ...
type TermOfRestriction *StartAndEndDateComplexType

// AgentExtension ...
type AgentExtension *ExtensionComplexType

// CreatingApplicationExtension ...
type CreatingApplicationExtension *ExtensionComplexType

// EnvironmentExtension ...
type EnvironmentExtension *ExtensionComplexType

// EventDetailExtension ...
type EventDetailExtension *ExtensionComplexType

// EventOutcomeDetailExtension ...
type EventOutcomeDetailExtension *ExtensionComplexType

// KeyInformation ...
type KeyInformation *ExtensionComplexType

// ObjectCharacteristicsExtension ...
type ObjectCharacteristicsExtension *ExtensionComplexType

// RightsExtension ...
type RightsExtension *ExtensionComplexType

// SignatureInformationExtension ...
type SignatureInformationExtension *ExtensionComplexType

// SignificantPropertiesExtension ...
type SignificantPropertiesExtension *ExtensionComplexType

// CountryCode ...
type CountryCode struct {
	XMLName xml.Name `xml:"countryCode"`
	*StringPlusAuthority
}

// Version3 ...
type Version3 string

// ExtensionComplexType ...
type ExtensionComplexType struct {
	XMLName xml.Name `xml:"extension"`
}

// EdtfSimpleType ...
type EdtfSimpleType string

// AuthorityAttributeGroup ...
type AuthorityAttributeGroup struct {
	XMLName          xml.Name `xml:"authorityAttributeGroup"`
	AuthorityAttr    string   `xml:"authority,attr,omitempty"`
	AuthorityURIAttr string   `xml:"authorityURI,attr,omitempty"`
	ValueURIAttr     string   `xml:"valueURI,attr,omitempty"`
}

// StringPlusAuthority ...
type StringPlusAuthority struct {
	//	XMLName                 xml.Name `xml:"stringPlusAuthority"`
	AuthorityAttr    string `xml:"authority,attr,omitempty"`
	AuthorityURIAttr string `xml:"authorityURI,attr,omitempty"`
	ValueURIAttr     string `xml:"valueURI,attr,omitempty"`
	Value            string `xml:",chardata"`
}
