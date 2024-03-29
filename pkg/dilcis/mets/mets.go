package mets

import (
	"encoding/xml"
)

// Mets ...
type Mets struct {
	XMLName           xml.Name `xml:"mets"`
	XMLNS             string   `xml:"xmlns,attr"`
	XMLNSXSI          string   `xml:"xmlns:xsi,attr"`
	XSISchemaLocation string   `xml:"xsi:schemaLocation,attr"`
	XMLXLinkNS        string   `xml:"xmlns:xlink,attr"`
	*MetsType
}

// Note ...
type Note struct {
	XMLName xml.Name `xml:"note"`
	Value   string   `xml:",chardata"`
}

// Agent is The element <name> can be used to record the full name of the document agent.
type Agent struct {
	XMLName       xml.Name `xml:"agent"`
	IDAttr        string   `xml:"ID,attr,omitempty"`
	ROLEAttr      string   `xml:"ROLE,attr"`
	OTHERROLEAttr string   `xml:"OTHERROLE,attr,omitempty"`
	TYPEAttr      string   `xml:"TYPE,attr,omitempty"`
	OTHERTYPEAttr string   `xml:"OTHERTYPE,attr,omitempty"`
	Name          string   `xml:"name"`
	Note          []*Note  `xml:"note"`
}

// AltRecordID ...
type AltRecordID struct {
	XMLName  xml.Name `xml:"altRecordID"`
	IDAttr   string   `xml:"ID,attr,omitempty"`
	TYPEAttr string   `xml:"TYPE,attr,omitempty"`
	Value    string   `xml:",chardata"`
}

// MetsDocumentID ...
type MetsDocumentID struct {
	XMLName  xml.Name `xml:"metsDocumentID"`
	IDAttr   string   `xml:"ID,attr,omitempty"`
	TYPEAttr string   `xml:"TYPE,attr,omitempty"`
	Value    string   `xml:",chardata"`
}

// MetsHdr ...
type MetsHdr struct {
	XMLName          xml.Name        `xml:"metsHdr"`
	IDAttr           string          `xml:"ID,attr,omitempty"`
	ADMIDAttr        []string        `xml:"ADMID,attr,omitempty"`
	CREATEDATEAttr   string          `xml:"CREATEDATE,attr,omitempty"`
	LASTMODDATEAttr  string          `xml:"LASTMODDATE,attr,omitempty"`
	RECORDSTATUSAttr string          `xml:"RECORDSTATUS,attr,omitempty"`
	Agent            []*Agent        `xml:"agent"`
	AltRecordID      []*AltRecordID  `xml:"altRecordID"`
	MetsDocumentID   *MetsDocumentID `xml:"metsDocumentID"`
}

// FileGrp ...
type FileGrp struct {
	XMLName xml.Name `xml:"fileGrp"`
	*FileGrpType
}

// FileSec ...
type FileSec struct {
	XMLName xml.Name   `xml:"fileSec"`
	IDAttr  string     `xml:"ID,attr,omitempty"`
	FileGrp []*FileGrp `xml:"fileGrp"`
}

// StructLink ...
type StructLink struct {
	XMLName xml.Name `xml:"structLink"`
	*StructLinkType
}

// MetsType is metsType: Complex Type for METS Sections
// A METS document consists of seven possible subsidiary sections: metsHdr (METS document header), dmdSec (descriptive metadata section), amdSec (administrative metadata section), fileGrp (file inventory group), structLink (structural map linking), structMap (structural map) and behaviorSec (behaviors section).
type MetsType struct {
	XMLName     xml.Name           `xml:"metsType"`
	IDAttr      string             `xml:"ID,attr,omitempty"`
	OBJIDAttr   string             `xml:"OBJID,attr,omitempty"`
	LABELAttr   string             `xml:"LABEL,attr,omitempty"`
	TYPEAttr    string             `xml:"TYPE,attr,omitempty"`
	PROFILEAttr string             `xml:"PROFILE,attr,omitempty"`
	MetsHdr     *MetsHdr           `xml:"metsHdr"`
	DmdSec      []*MdSecType       `xml:"dmdSec"`
	AmdSec      []*AmdSecType      `xml:"amdSec"`
	FileSec     *FileSec           `xml:"fileSec"`
	StructMap   []*StructMapType   `xml:"structMap"`
	StructLink  *StructLink        `xml:"structLink"`
	BehaviorSec []*BehaviorSecType `xml:"behaviorSec"`
}

// AmdSecType is A digital provenance metadata element <digiprovMD> can be used to record any preservation-related actions taken on the various files which comprise a digital object (e.g., those subsequent to the initial digitization of the files such as transformation or migrations) or, in the case of born digital materials, the files’ creation. In short, digital provenance should be used to record information that allows both archival/library staff and scholars to understand what modifications have been made to a digital object and/or its constituent parts during its life cycle. This information can then be used to judge how those processes might have altered or corrupted the object’s ability to accurately represent the original item. One might, for example, record master derivative relationships and the process by which those derivations have been created. Or the <digiprovMD> element could contain information regarding the migration/transformation of a file from its original digitization (e.g., OCR, TEI, etc.,)to its current incarnation as a digital object (e.g., JPEG2000). The <digiprovMD> element conforms to same generic datatype as the <dmdSec>,  <techMD>, <rightsMD>, and <sourceMD> elements, and supports the same sub-elements and attributes. A digital provenance metadata element can either wrap the metadata  (mdWrap) or reference it in an external location (mdRef) or both.  METS allows multiple <digiprovMD> elements; and digital provenance metadata can be associated with any METS element that supports an ADMID attribute. Digital provenance metadata can be expressed according to current digital provenance description standards (such as PREMIS) or a locally produced XML schema.
type AmdSecType struct {
	XMLName    xml.Name     `xml:"amdSec"`
	IDAttr     string       `xml:"ID,attr,omitempty"`
	TechMD     []*MdSecType `xml:"techMD"`
	RightsMD   []*MdSecType `xml:"rightsMD"`
	SourceMD   []*MdSecType `xml:"sourceMD"`
	DigiprovMD []*MdSecType `xml:"digiprovMD"`
}

// FileGrpType is The file element <file> provides access to the content files for the digital object being described by the METS document. A <file> element may contain one or more <FLocat> elements which provide pointers to a content file and/or a <FContent> element which wraps an encoded version of the file. Embedding files using <FContent> can be a valuable feature for exchanging digital objects between repositories or for archiving versions of digital objects for off-site storage. All <FLocat> and <FContent> elements should identify and/or contain identical copies of a single file. The <file> element is recursive, thus allowing sub-files or component files of a larger file to be listed in the inventory. Alternatively, by using the <stream> element, a smaller component of a file or of a related file can be placed within a <file> element. Finally, by using the <transformFile> element, it is possible to include within a <file> element a different version of a file that has undergone a transformation for some reason, such as format migration.
type FileGrpType struct {
	XMLName      xml.Name       `xml:"fileGrpType"`
	IDAttr       string         `xml:"ID,attr,omitempty"`
	VERSDATEAttr string         `xml:"VERSDATE,attr,omitempty"`
	ADMIDAttr    []string       `xml:"ADMID,attr,omitempty"`
	USEAttr      string         `xml:"USE,attr,omitempty"`
	FileGrp      []*FileGrpType `xml:"fileGrp"`
	File         []*FileType    `xml:"file"`
}

// StructMapType is The structural divisions of the hierarchical organization provided by a <structMap> are represented by division <div> elements, which can be nested to any depth. Each <div> element can represent either an intellectual (logical) division or a physical division. Every <div> node in the structural map hierarchy may be connected (via subsidiary <mptr> or <fptr> elements) to content files which represent that div's portion of the whole document.
type StructMapType struct {
	XMLName   xml.Name `xml:"structMap"`
	IDAttr    string   `xml:"ID,attr,omitempty"`
	TYPEAttr  string   `xml:"TYPE,attr,omitempty"`
	LABELAttr string   `xml:"LABEL,attr,omitempty"`
	Div       *DivType `xml:"div"`
}

// Mptr ...
type Mptr struct {
	XMLName         xml.Name `xml:"mptr"`
	LOCATION        *LOCATION
	XlinkSimpleLink *SimpleLink
	IDAttr          string `xml:"ID,attr,omitempty"`
	CONTENTIDSAttr  *URIs  `xml:"CONTENTIDS,attr,omitempty"`
}

// Fptr ...
type Fptr struct {
	XMLName        xml.Name  `xml:"fptr"`
	IDAttr         string    `xml:"ID,attr,omitempty"`
	FILEIDAttr     string    `xml:"FILEID,attr,omitempty"`
	CONTENTIDSAttr *URIs     `xml:"CONTENTIDS,attr,omitempty"`
	Par            *ParType  `xml:"par"`
	Seq            *SeqType  `xml:"seq"`
	Area           *AreaType `xml:"area"`
}

// DivType is divType: Complex Type for Divisions
// The METS standard represents a document structurally as a series of nested div elements, that is, as a hierarchy (e.g., a book, which is composed of chapters, which are composed of subchapters, which are composed of text).  Every div node in the structural map hierarchy may be connected (via subsidiary mptr or fptr elements) to content files which represent that div's portion of the whole document.
//
// SPECIAL NOTE REGARDING DIV ATTRIBUTE VALUES:
// to clarify the differences between the ORDER, ORDERLABEL, and LABEL attributes for the <div> element, imagine a text with 10 roman numbered pages followed by 10 arabic numbered pages. Page iii would have an ORDER of "3", an ORDERLABEL of "iii" and a LABEL of "Page iii", while page 3 would have an ORDER of "13", an ORDERLABEL of "3" and a LABEL of "Page 3".
type DivType struct {
	XMLName xml.Name `xml:"div"`
	*ORDERLABELS
	IDAttr         string     `xml:"ID,attr,omitempty"`
	DMDIDAttr      []string   `xml:"DMDID,attr,omitempty"`
	ADMIDAttr      []string   `xml:"ADMID,attr,omitempty"`
	TYPEAttr       string     `xml:"TYPE,attr,omitempty"`
	CONTENTIDSAttr *URIs      `xml:"CONTENTIDS,attr,omitempty"`
	XlinkLabelAttr *Label     `xml:"xlink:label,attr,omitempty"`
	Mptr           []*Mptr    `xml:"mptr"`
	Fptr           []*Fptr    `xml:"fptr"`
	Div            []*DivType `xml:"div"`
}

// ParType is parType: Complex Type for Parallel Files
// The <par> or parallel files element aggregates pointers to files, parts of files, and/or sequences of files or parts of files that must be played or displayed simultaneously to manifest a block of digital content represented by an <fptr> element.
type ParType struct {
	XMLName     xml.Name `xml:"par"`
	ORDERLABELS *ORDERLABELS
	IDAttr      string      `xml:"ID,attr,omitempty"`
	Area        []*AreaType `xml:"area"`
	Seq         []*SeqType  `xml:"seq"`
}

// SeqType is seqType: Complex Type for Sequences of Files
// The seq element should be used to link a div to a set of content files when those files should be played/displayed sequentially to deliver content to a user.  Individual <area> subelements within the seq element provide the links to the files or portions thereof.
type SeqType struct {
	XMLName     xml.Name `xml:"seq"`
	ORDERLABELS *ORDERLABELS
	IDAttr      string      `xml:"ID,attr,omitempty"`
	Area        []*AreaType `xml:"area"`
	Par         []*ParType  `xml:"par"`
}

// AreaType is areaType: Complex Type for Area Linking
// The area element provides for more sophisticated linking between a div element and content files representing that div, be they text, image, audio, or video files.  An area element can link a div to a point within a file, to a one-dimension segment of a file (e.g., text segment, image line, audio/video clip), or a two-dimensional section of a file (e.g, subsection of an image, or a subsection of the  video display of a video file.  The area element has no content; all information is recorded within its various attributes.
type AreaType struct {
	XMLName        xml.Name `xml:"area"`
	ORDERLABELS    *ORDERLABELS
	IDAttr         string   `xml:"ID,attr,omitempty"`
	FILEIDAttr     string   `xml:"FILEID,attr"`
	SHAPEAttr      string   `xml:"SHAPE,attr,omitempty"`
	COORDSAttr     string   `xml:"COORDS,attr,omitempty"`
	BEGINAttr      string   `xml:"BEGIN,attr,omitempty"`
	ENDAttr        string   `xml:"END,attr,omitempty"`
	BETYPEAttr     string   `xml:"BETYPE,attr,omitempty"`
	EXTENTAttr     string   `xml:"EXTENT,attr,omitempty"`
	EXTTYPEAttr    string   `xml:"EXTTYPE,attr,omitempty"`
	ADMIDAttr      []string `xml:"ADMID,attr,omitempty"`
	CONTENTIDSAttr *URIs    `xml:"CONTENTIDS,attr,omitempty"`
}

// SmLink ...
type SmLink struct {
	XMLName          xml.Name `xml:"smLink"`
	IDAttr           string   `xml:"ID,attr,omitempty"`
	XlinkArcroleAttr *Arcrole `xml:"xlink:arcrole,attr,omitempty"`
	XlinkTitleAttr   *Title   `xml:"xlink:title,attr,omitempty"`
	XlinkShowAttr    *Show    `xml:"xlink:show,attr,omitempty"`
	XlinkActuateAttr *Actuate `xml:"xlink:actuate,attr,omitempty"`
	XlinkToAttr      *To      `xml:"xlink:to,attr"`
	XlinkFromAttr    *From    `xml:"xlink:from,attr"`
}

// SmLocatorLink ...
type SmLocatorLink struct {
	XMLName          xml.Name `xml:"smLocatorLink"`
	XlinkLocatorLink *LocatorLink
	IDAttr           string `xml:"ID,attr,omitempty"`
}

// SmArcLink ...
type SmArcLink struct {
	XMLName      xml.Name `xml:"smArcLink"`
	XlinkArcLink *ArcLink
	IDAttr       string   `xml:"ID,attr,omitempty"`
	ARCTYPEAttr  string   `xml:"ARCTYPE,attr,omitempty"`
	ADMIDAttr    []string `xml:"ADMID,attr,omitempty"`
}

// SmLinkGrp ...
type SmLinkGrp struct {
	XMLName           xml.Name `xml:"smLinkGrp"`
	XlinkExtendedLink *ExtendedLink
	IDAttr            string           `xml:"ID,attr,omitempty"`
	ARCLINKORDERAttr  string           `xml:"ARCLINKORDER,attr,omitempty"`
	SmLocatorLink     []*SmLocatorLink `xml:"smLocatorLink"`
	SmArcLink         []*SmArcLink     `xml:"smArcLink"`
}

// StructLinkType is structLinkType: Complex Type for Structural Map Linking
// The Structural Map Linking section allows for the specification of hyperlinks between different components of a METS structure delineated in a structural map.  structLink contains a single, repeatable element, smLink.  Each smLink element indicates a hyperlink between two nodes in the structMap.  The structMap nodes recorded in smLink are identified using their XML ID attributevalues.
type StructLinkType struct {
	XMLName   xml.Name     `xml:"structLinkType"`
	IDAttr    string       `xml:"ID,attr,omitempty"`
	SmLink    []*SmLink    `xml:"smLink"`
	SmLinkGrp []*SmLinkGrp `xml:"smLinkGrp"`
}

// BehaviorSecType is A behavior element <behavior> can be used to associate executable behaviors with content in the METS document. This element has an interface definition <interfaceDef> element that represents an abstract definition of a set of behaviors represented by a particular behavior. A <behavior> element also has a behavior mechanism <mechanism> element, a module of executable code that implements and runs the behavior defined abstractly by the interface definition.
type BehaviorSecType struct {
	XMLName     xml.Name           `xml:"behaviorSecType"`
	IDAttr      string             `xml:"ID,attr,omitempty"`
	CREATEDAttr string             `xml:"CREATED,attr,omitempty"`
	LABELAttr   string             `xml:"LABEL,attr,omitempty"`
	BehaviorSec []*BehaviorSecType `xml:"behaviorSec"`
	Behavior    []*BehaviorType    `xml:"behavior"`
}

// BehaviorType is A mechanism element <mechanism> contains a pointer to an executable code module that implements a set of behaviors defined by an interface definition. The <mechanism> element will be a pointer to another object (a mechanism object). A mechanism object could be another METS object, or some other entity (e.g., a WSDL file). A mechanism object should contain executable code, pointers to executable code, or specifications for binding to network services (e.g., web services).
type BehaviorType struct {
	XMLName      xml.Name    `xml:"behaviorType"`
	IDAttr       string      `xml:"ID,attr,omitempty"`
	STRUCTIDAttr []string    `xml:"STRUCTID,attr,omitempty"`
	BTYPEAttr    string      `xml:"BTYPE,attr,omitempty"`
	CREATEDAttr  string      `xml:"CREATED,attr,omitempty"`
	LABELAttr    string      `xml:"LABEL,attr,omitempty"`
	GROUPIDAttr  string      `xml:"GROUPID,attr,omitempty"`
	ADMIDAttr    []string    `xml:"ADMID,attr,omitempty"`
	InterfaceDef *ObjectType `xml:"interfaceDef"`
	Mechanism    *ObjectType `xml:"mechanism"`
}

// ObjectType is objectType: complexType for interfaceDef and mechanism elements
// The mechanism and behavior elements point to external objects--an interface definition object or an executable code object respectively--which together constitute a behavior that can be applied to one or more <div> elements in a <structMap>.
type ObjectType struct {
	XMLName         xml.Name `xml:"objectType"`
	LOCATION        *LOCATION
	XlinkSimpleLink *SimpleLink
	IDAttr          string `xml:"ID,attr,omitempty"`
	LABELAttr       string `xml:"LABEL,attr,omitempty"`
}

// MdRef ...
type MdRef struct {
	XMLName   xml.Name `xml:"mdRef"`
	IDAttr    string   `xml:"ID,attr,omitempty"`
	LABELAttr string   `xml:"LABEL,attr,omitempty"`
	XPTRAttr  string   `xml:"XPTR,attr,omitempty"`
	// SimpleLink
	TypeAttr         string `xml:"xlink:type,attr,omitempty"`
	XlinkHrefAttr    string `xml:"xlink:href,attr,omitempty"`
	XlinkRoleAttr    string `xml:"xlink:role,attr,omitempty"`
	XlinkArcroleAttr string `xml:"xlink:arcrole,attr,omitempty"`
	XlinkTitleAttr   string `xml:"xlink:title,attr,omitempty"`
	XlinkShowAttr    string `xml:"xlink:show,attr,omitempty"`
	XlinkActuateAttr string `xml:"xlink:actuate,attr,omitempty"`
	// LOCTYPE
	LOCTYPEAttr      string `xml:"LOCTYPE,attr"`
	OTHERLOCTYPEAttr string `xml:"OTHERLOCTYPE,attr,omitempty"`
	// METADATA
	MDTYPEAttr        string `xml:"MDTYPE,attr"`
	OTHERMDTYPEAttr   string `xml:"OTHERMDTYPE,attr,omitempty"`
	MDTYPEVERSIONAttr string `xml:"MDTYPEVERSION,attr,omitempty"`
	// FILECORE
	MIMETYPEAttr     string `xml:"MIMETYPE,attr,omitempty"`
	SIZEAttr         int64  `xml:"SIZE,attr,omitempty"`
	CREATEDAttr      string `xml:"CREATED,attr,omitempty"`
	CHECKSUMAttr     string `xml:"CHECKSUM,attr,omitempty"`
	CHECKSUMTYPEAttr string `xml:"CHECKSUMTYPE,attr,omitempty"`
}

// XmlData ...
type XmlData struct {
	XMLName xml.Name `xml:"xmlData"`
}

// MdWrap ...
type MdWrap struct {
	XMLName   xml.Name `xml:"mdWrap"`
	METADATA  *METADATA
	FILECORE  *FILECORE
	IDAttr    string   `xml:"ID,attr,omitempty"`
	LABELAttr string   `xml:"LABEL,attr,omitempty"`
	BinData   string   `xml:"binData"`
	XmlData   *XmlData `xml:"xmlData"`
}

// MdSecType is mdSecType: Complex Type for Metadata Sections
// A generic framework for pointing to/including metadata within a METS document, a la Warwick Framework.
type MdSecType struct {
	//	XMLName     xml.Name `xml:"mdSecType"`
	IDAttr      string   `xml:"ID,attr"`
	GROUPIDAttr string   `xml:"GROUPID,attr,omitempty"`
	ADMIDAttr   []string `xml:"ADMID,attr,omitempty"`
	CREATEDAttr string   `xml:"CREATED,attr,omitempty"`
	STATUSAttr  string   `xml:"STATUS,attr,omitempty"`
	MdRef       *MdRef   `xml:"mdRef"`
	MdWrap      *MdWrap  `xml:"mdWrap"`
}

// FLocat ...
type FLocat struct {
	*LOCATION
	*SimpleLink
	IDAttr  string `xml:"ID,attr,omitempty"`
	USEAttr string `xml:"USE,attr,omitempty"`
}

// FContent ...
type FContent struct {
	IDAttr  string   `xml:"ID,attr,omitempty"`
	USEAttr string   `xml:"USE,attr,omitempty"`
	BinData string   `xml:"binData"`
	XmlData *XmlData `xml:"xmlData"`
}

// Stream ...
type Stream struct {
	XMLName        xml.Name `xml:"stream"`
	IDAttr         string   `xml:"ID,attr,omitempty"`
	StreamTypeAttr string   `xml:"streamType,attr,omitempty"`
	OWNERIDAttr    string   `xml:"OWNERID,attr,omitempty"`
	ADMIDAttr      []string `xml:"ADMID,attr,omitempty"`
	DMDIDAttr      []string `xml:"DMDID,attr,omitempty"`
	BEGINAttr      string   `xml:"BEGIN,attr,omitempty"`
	ENDAttr        string   `xml:"END,attr,omitempty"`
	BETYPEAttr     string   `xml:"BETYPE,attr,omitempty"`
}

// TransformFile ...
type TransformFile struct {
	XMLName                xml.Name `xml:"transformFile"`
	IDAttr                 string   `xml:"ID,attr,omitempty"`
	TRANSFORMTYPEAttr      string   `xml:"TRANSFORMTYPE,attr"`
	TRANSFORMALGORITHMAttr string   `xml:"TRANSFORMALGORITHM,attr"`
	TRANSFORMKEYAttr       string   `xml:"TRANSFORMKEY,attr,omitempty"`
	TRANSFORMBEHAVIORAttr  string   `xml:"TRANSFORMBEHAVIOR,attr,omitempty"`
	TRANSFORMORDERAttr     int      `xml:"TRANSFORMORDER,attr"`
}

// FileType is fileType: Complex Type for Files
// The file element provides access to content files for a METS object.  A file element may contain one or more FLocat elements, which provide pointers to a content file, and/or an FContent element, which wraps an encoded version of the file. Note that ALL FLocat and FContent elements underneath a single file element should identify/contain identical copies of a single file.
type FileType struct {
	XMLName xml.Name `xml:"file"`
	*FILECORE
	IDAttr        string           `xml:"ID,attr"`
	SEQAttr       int              `xml:"SEQ,attr,omitempty"`
	OWNERIDAttr   string           `xml:"OWNERID,attr,omitempty"`
	ADMIDAttr     []string         `xml:"ADMID,attr,omitempty"`
	DMDIDAttr     []string         `xml:"DMDID,attr,omitempty"`
	GROUPIDAttr   string           `xml:"GROUPID,attr,omitempty"`
	USEAttr       string           `xml:"USE,attr,omitempty"`
	BEGINAttr     string           `xml:"BEGIN,attr,omitempty"`
	ENDAttr       string           `xml:"END,attr,omitempty"`
	BETYPEAttr    string           `xml:"BETYPE,attr,omitempty"`
	FLocat        []*FLocat        `xml:"FLocat"`
	FContent      *FContent        `xml:"FContent"`
	Stream        []*Stream        `xml:"stream"`
	TransformFile []*TransformFile `xml:"transformFile"`
	File          []*FileType      `xml:"file"`
}

// URIs ...
type URIs []string

// ORDERLABELS is LABEL (string/O): An attribute used, for example, to identify a <div> to an end user viewing the document. Thus a hierarchical arrangement of the <div> LABEL values could provide a table of contents to the digital content represented by a METS document and facilitate the users’ navigation of the digital object. Note that a <div> LABEL should be specific to its level in the structural map. In the case of a book with chapters, the book <div> LABEL should have the book title and the chapter <div>; LABELs should have the individual chapter titles, rather than having the chapter <div> LABELs combine both book title and chapter title . For further of the distinction between LABEL and ORDERLABEL see the description of the ORDERLABEL attribute.
type ORDERLABELS struct {
	ORDERAttr      int    `xml:"ORDER,attr,omitempty"`
	ORDERLABELAttr string `xml:"ORDERLABEL,attr,omitempty"`
	LABELAttr      string `xml:"LABEL,attr,omitempty"`
}

// METADATA is MDTYPEVERSION(string/O): Provides a means for recording the version of the type of metadata (as recorded in the MDTYPE or OTHERMDTYPE attribute) that is being used.  This may represent the version of the underlying data dictionary or metadata model rather than a schema version.
type METADATA struct {
	MDTYPEAttr        string `xml:"MDTYPE,attr"`
	OTHERMDTYPEAttr   string `xml:"OTHERMDTYPE,attr,omitempty"`
	MDTYPEVERSIONAttr string `xml:"MDTYPEVERSION,attr,omitempty"`
}

// LOCATION is OTHERLOCTYPE (string/O): Specifies the locator type when the value OTHER is used in the LOCTYPE attribute. Although optional, it is strongly recommended when OTHER is used.
type LOCATION struct {
	LOCTYPEAttr      string `xml:"LOCTYPE,attr"`
	OTHERLOCTYPEAttr string `xml:"OTHERLOCTYPE,attr,omitempty"`
}

// FILECORE is CHECKSUMTYPE (enumerated string/O): Specifies the checksum algorithm used to produce the value contained in the CHECKSUM attribute.  CHECKSUMTYPE must contain one of the following values:
// Adler-32
// CRC32
// HAVAL
// MD5
// MNP
// SHA-1
// SHA-256
// SHA-384
// SHA-512
// TIGER
// WHIRLPOOL
type FILECORE struct {
	MIMETYPEAttr     string `xml:"MIMETYPE,attr,omitempty"`
	SIZEAttr         int64  `xml:"SIZE,attr,omitempty"`
	CREATEDAttr      string `xml:"CREATED,attr,omitempty"`
	CHECKSUMAttr     string `xml:"CHECKSUM,attr,omitempty"`
	CHECKSUMTYPEAttr string `xml:"CHECKSUMTYPE,attr,omitempty"`
}
