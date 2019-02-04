package typeplace

import vocab "github.com/go-fed/activity/streams/vocab"

// Represents a logical or physical location. See 5.3 Representing Places for
// additional information.
//
// Example 56 (https://www.w3.org/TR/activitystreams-vocabulary/#ex57-jsonld):
//   {
//     "name": "Work",
//     "type": "Place"
//   }
//
// Example 57 (https://www.w3.org/TR/activitystreams-vocabulary/#ex58-jsonld):
//   {
//     "latitude": 36.75,
//     "longitude": 119.7667,
//     "name": "Fresno Area",
//     "radius": 15,
//     "type": "Place",
//     "units": "miles"
//   }
type Place struct {
	Accuracy     vocab.AccuracyPropertyInterface
	Altitude     vocab.AltitudePropertyInterface
	Attachment   vocab.AttachmentPropertyInterface
	AttributedTo vocab.AttributedToPropertyInterface
	Audience     vocab.AudiencePropertyInterface
	Bcc          vocab.BccPropertyInterface
	Bto          vocab.BtoPropertyInterface
	Cc           vocab.CcPropertyInterface
	Content      vocab.ContentPropertyInterface
	Context      vocab.ContextPropertyInterface
	Duration     vocab.DurationPropertyInterface
	EndTime      vocab.EndTimePropertyInterface
	Generator    vocab.GeneratorPropertyInterface
	Icon         vocab.IconPropertyInterface
	Id           vocab.IdPropertyInterface
	Image        vocab.ImagePropertyInterface
	InReplyTo    vocab.InReplyToPropertyInterface
	Latitude     vocab.LatitudePropertyInterface
	Location     vocab.LocationPropertyInterface
	Longitude    vocab.LongitudePropertyInterface
	MediaType    vocab.MediaTypePropertyInterface
	Name         vocab.NamePropertyInterface
	Object       vocab.ObjectPropertyInterface
	Preview      vocab.PreviewPropertyInterface
	Published    vocab.PublishedPropertyInterface
	Radius       vocab.RadiusPropertyInterface
	Replies      vocab.RepliesPropertyInterface
	StartTime    vocab.StartTimePropertyInterface
	Summary      vocab.SummaryPropertyInterface
	Tag          vocab.TagPropertyInterface
	To           vocab.ToPropertyInterface
	Type         vocab.TypePropertyInterface
	Units        vocab.UnitsPropertyInterface
	Updated      vocab.UpdatedPropertyInterface
	Url          vocab.UrlPropertyInterface
	alias        string
	unknown      map[string]interface{}
}

// DeserializePlace creates a Place from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializePlace(m map[string]interface{}, aliasMap map[string]string) (*Place, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	this := &Place{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAccuracyPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Accuracy = p
	}
	if p, err := mgr.DeserializeAltitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Altitude = p
	}
	if p, err := mgr.DeserializeAttachmentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Attachment = p
	}
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.AttributedTo = p
	}
	if p, err := mgr.DeserializeAudiencePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Audience = p
	}
	if p, err := mgr.DeserializeBccPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Bcc = p
	}
	if p, err := mgr.DeserializeBtoPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Bto = p
	}
	if p, err := mgr.DeserializeCcPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Cc = p
	}
	if p, err := mgr.DeserializeContentPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Content = p
	}
	if p, err := mgr.DeserializeContextPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Context = p
	}
	if p, err := mgr.DeserializeDurationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Duration = p
	}
	if p, err := mgr.DeserializeEndTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.EndTime = p
	}
	if p, err := mgr.DeserializeGeneratorPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Generator = p
	}
	if p, err := mgr.DeserializeIconPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Icon = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Id = p
	}
	if p, err := mgr.DeserializeImagePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Image = p
	}
	if p, err := mgr.DeserializeInReplyToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.InReplyTo = p
	}
	if p, err := mgr.DeserializeLatitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Latitude = p
	}
	if p, err := mgr.DeserializeLocationPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Location = p
	}
	if p, err := mgr.DeserializeLongitudePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Longitude = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.MediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Name = p
	}
	if p, err := mgr.DeserializeObjectPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Object = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Preview = p
	}
	if p, err := mgr.DeserializePublishedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Published = p
	}
	if p, err := mgr.DeserializeRadiusPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Radius = p
	}
	if p, err := mgr.DeserializeRepliesPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Replies = p
	}
	if p, err := mgr.DeserializeStartTimePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.StartTime = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Summary = p
	}
	if p, err := mgr.DeserializeTagPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Tag = p
	}
	if p, err := mgr.DeserializeToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.To = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Type = p
	}
	if p, err := mgr.DeserializeUnitsPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Units = p
	}
	if p, err := mgr.DeserializeUpdatedPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Updated = p
	}
	if p, err := mgr.DeserializeUrlPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else {
		this.Url = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "accuracy" {
			continue
		} else if k == "altitude" {
			continue
		} else if k == "attachment" {
			continue
		} else if k == "attributedTo" {
			continue
		} else if k == "audience" {
			continue
		} else if k == "bcc" {
			continue
		} else if k == "bto" {
			continue
		} else if k == "cc" {
			continue
		} else if k == "content" {
			continue
		} else if k == "context" {
			continue
		} else if k == "duration" {
			continue
		} else if k == "endTime" {
			continue
		} else if k == "generator" {
			continue
		} else if k == "icon" {
			continue
		} else if k == "id" {
			continue
		} else if k == "image" {
			continue
		} else if k == "inReplyTo" {
			continue
		} else if k == "latitude" {
			continue
		} else if k == "location" {
			continue
		} else if k == "longitude" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "object" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "published" {
			continue
		} else if k == "radius" {
			continue
		} else if k == "replies" {
			continue
		} else if k == "startTime" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "tag" {
			continue
		} else if k == "to" {
			continue
		} else if k == "type" {
			continue
		} else if k == "units" {
			continue
		} else if k == "updated" {
			continue
		} else if k == "url" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// NewPlace creates a new Place type
func NewPlace() *Place {
	return &Place{
		alias:   "",
		unknown: make(map[string]interface{}, 0),
	}
}

// PlaceExtends returns true if the Place type extends from the other type.
func PlaceExtends(other vocab.Type) bool {
	extensions := []string{"Object"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// PlaceIsDisjointWith returns true if the other provided type is disjoint with
// the Place type.
func PlaceIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Link", "Mention"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// PlaceIsExtendedBy returns true if the other provided type extends from the
// Place type.
func PlaceIsExtendedBy(other vocab.Type) bool {
	// Shortcut implementation: is not extended by anything.
	return false
}

// GetAccuracy returns the "accuracy" property if it exists, and nil otherwise.
func (this Place) GetAccuracy() vocab.AccuracyPropertyInterface {
	return this.Accuracy
}

// GetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this Place) GetAltitude() vocab.AltitudePropertyInterface {
	return this.Altitude
}

// GetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this Place) GetAttachment() vocab.AttachmentPropertyInterface {
	return this.Attachment
}

// GetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Place) GetAttributedTo() vocab.AttributedToPropertyInterface {
	return this.AttributedTo
}

// GetAudience returns the "audience" property if it exists, and nil otherwise.
func (this Place) GetAudience() vocab.AudiencePropertyInterface {
	return this.Audience
}

// GetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this Place) GetBcc() vocab.BccPropertyInterface {
	return this.Bcc
}

// GetBto returns the "bto" property if it exists, and nil otherwise.
func (this Place) GetBto() vocab.BtoPropertyInterface {
	return this.Bto
}

// GetCc returns the "cc" property if it exists, and nil otherwise.
func (this Place) GetCc() vocab.CcPropertyInterface {
	return this.Cc
}

// GetContent returns the "content" property if it exists, and nil otherwise.
func (this Place) GetContent() vocab.ContentPropertyInterface {
	return this.Content
}

// GetContext returns the "context" property if it exists, and nil otherwise.
func (this Place) GetContext() vocab.ContextPropertyInterface {
	return this.Context
}

// GetDuration returns the "duration" property if it exists, and nil otherwise.
func (this Place) GetDuration() vocab.DurationPropertyInterface {
	return this.Duration
}

// GetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this Place) GetEndTime() vocab.EndTimePropertyInterface {
	return this.EndTime
}

// GetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this Place) GetGenerator() vocab.GeneratorPropertyInterface {
	return this.Generator
}

// GetIcon returns the "icon" property if it exists, and nil otherwise.
func (this Place) GetIcon() vocab.IconPropertyInterface {
	return this.Icon
}

// GetId returns the "id" property if it exists, and nil otherwise.
func (this Place) GetId() vocab.IdPropertyInterface {
	return this.Id
}

// GetImage returns the "image" property if it exists, and nil otherwise.
func (this Place) GetImage() vocab.ImagePropertyInterface {
	return this.Image
}

// GetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this Place) GetInReplyTo() vocab.InReplyToPropertyInterface {
	return this.InReplyTo
}

// GetLatitude returns the "latitude" property if it exists, and nil otherwise.
func (this Place) GetLatitude() vocab.LatitudePropertyInterface {
	return this.Latitude
}

// GetLocation returns the "location" property if it exists, and nil otherwise.
func (this Place) GetLocation() vocab.LocationPropertyInterface {
	return this.Location
}

// GetLongitude returns the "longitude" property if it exists, and nil otherwise.
func (this Place) GetLongitude() vocab.LongitudePropertyInterface {
	return this.Longitude
}

// GetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Place) GetMediaType() vocab.MediaTypePropertyInterface {
	return this.MediaType
}

// GetName returns the "name" property if it exists, and nil otherwise.
func (this Place) GetName() vocab.NamePropertyInterface {
	return this.Name
}

// GetObject returns the "object" property if it exists, and nil otherwise.
func (this Place) GetObject() vocab.ObjectPropertyInterface {
	return this.Object
}

// GetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Place) GetPreview() vocab.PreviewPropertyInterface {
	return this.Preview
}

// GetPublished returns the "published" property if it exists, and nil otherwise.
func (this Place) GetPublished() vocab.PublishedPropertyInterface {
	return this.Published
}

// GetRadius returns the "radius" property if it exists, and nil otherwise.
func (this Place) GetRadius() vocab.RadiusPropertyInterface {
	return this.Radius
}

// GetReplies returns the "replies" property if it exists, and nil otherwise.
func (this Place) GetReplies() vocab.RepliesPropertyInterface {
	return this.Replies
}

// GetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this Place) GetStartTime() vocab.StartTimePropertyInterface {
	return this.StartTime
}

// GetSummary returns the "summary" property if it exists, and nil otherwise.
func (this Place) GetSummary() vocab.SummaryPropertyInterface {
	return this.Summary
}

// GetTag returns the "tag" property if it exists, and nil otherwise.
func (this Place) GetTag() vocab.TagPropertyInterface {
	return this.Tag
}

// GetTo returns the "to" property if it exists, and nil otherwise.
func (this Place) GetTo() vocab.ToPropertyInterface {
	return this.To
}

// GetType returns the "type" property if it exists, and nil otherwise.
func (this Place) GetType() vocab.TypePropertyInterface {
	return this.Type
}

// GetUnits returns the "units" property if it exists, and nil otherwise.
func (this Place) GetUnits() vocab.UnitsPropertyInterface {
	return this.Units
}

// GetUnknownProperties returns the unknown properties for the Place type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this Place) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// GetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this Place) GetUpdated() vocab.UpdatedPropertyInterface {
	return this.Updated
}

// GetUrl returns the "url" property if it exists, and nil otherwise.
func (this Place) GetUrl() vocab.UrlPropertyInterface {
	return this.Url
}

// IsExtending returns true if the Place type extends from the other type.
func (this Place) IsExtending(other vocab.Type) bool {
	return PlaceExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this Place) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.Accuracy, m)
	m = this.helperJSONLDContext(this.Altitude, m)
	m = this.helperJSONLDContext(this.Attachment, m)
	m = this.helperJSONLDContext(this.AttributedTo, m)
	m = this.helperJSONLDContext(this.Audience, m)
	m = this.helperJSONLDContext(this.Bcc, m)
	m = this.helperJSONLDContext(this.Bto, m)
	m = this.helperJSONLDContext(this.Cc, m)
	m = this.helperJSONLDContext(this.Content, m)
	m = this.helperJSONLDContext(this.Context, m)
	m = this.helperJSONLDContext(this.Duration, m)
	m = this.helperJSONLDContext(this.EndTime, m)
	m = this.helperJSONLDContext(this.Generator, m)
	m = this.helperJSONLDContext(this.Icon, m)
	m = this.helperJSONLDContext(this.Id, m)
	m = this.helperJSONLDContext(this.Image, m)
	m = this.helperJSONLDContext(this.InReplyTo, m)
	m = this.helperJSONLDContext(this.Latitude, m)
	m = this.helperJSONLDContext(this.Location, m)
	m = this.helperJSONLDContext(this.Longitude, m)
	m = this.helperJSONLDContext(this.MediaType, m)
	m = this.helperJSONLDContext(this.Name, m)
	m = this.helperJSONLDContext(this.Object, m)
	m = this.helperJSONLDContext(this.Preview, m)
	m = this.helperJSONLDContext(this.Published, m)
	m = this.helperJSONLDContext(this.Radius, m)
	m = this.helperJSONLDContext(this.Replies, m)
	m = this.helperJSONLDContext(this.StartTime, m)
	m = this.helperJSONLDContext(this.Summary, m)
	m = this.helperJSONLDContext(this.Tag, m)
	m = this.helperJSONLDContext(this.To, m)
	m = this.helperJSONLDContext(this.Type, m)
	m = this.helperJSONLDContext(this.Units, m)
	m = this.helperJSONLDContext(this.Updated, m)
	m = this.helperJSONLDContext(this.Url, m)

	return m
}

// LessThan computes if this Place is lesser, with an arbitrary but stable
// determination.
func (this Place) LessThan(o vocab.PlaceInterface) bool {
	// Begin: Compare known properties
	// Compare property "accuracy"
	if this.Accuracy.LessThan(o.GetAccuracy()) {
		return true
	} else if o.GetAccuracy().LessThan(this.Accuracy) {
		return false
	}
	// Compare property "altitude"
	if this.Altitude.LessThan(o.GetAltitude()) {
		return true
	} else if o.GetAltitude().LessThan(this.Altitude) {
		return false
	}
	// Compare property "attachment"
	if this.Attachment.LessThan(o.GetAttachment()) {
		return true
	} else if o.GetAttachment().LessThan(this.Attachment) {
		return false
	}
	// Compare property "attributedTo"
	if this.AttributedTo.LessThan(o.GetAttributedTo()) {
		return true
	} else if o.GetAttributedTo().LessThan(this.AttributedTo) {
		return false
	}
	// Compare property "audience"
	if this.Audience.LessThan(o.GetAudience()) {
		return true
	} else if o.GetAudience().LessThan(this.Audience) {
		return false
	}
	// Compare property "bcc"
	if this.Bcc.LessThan(o.GetBcc()) {
		return true
	} else if o.GetBcc().LessThan(this.Bcc) {
		return false
	}
	// Compare property "bto"
	if this.Bto.LessThan(o.GetBto()) {
		return true
	} else if o.GetBto().LessThan(this.Bto) {
		return false
	}
	// Compare property "cc"
	if this.Cc.LessThan(o.GetCc()) {
		return true
	} else if o.GetCc().LessThan(this.Cc) {
		return false
	}
	// Compare property "content"
	if this.Content.LessThan(o.GetContent()) {
		return true
	} else if o.GetContent().LessThan(this.Content) {
		return false
	}
	// Compare property "context"
	if this.Context.LessThan(o.GetContext()) {
		return true
	} else if o.GetContext().LessThan(this.Context) {
		return false
	}
	// Compare property "duration"
	if this.Duration.LessThan(o.GetDuration()) {
		return true
	} else if o.GetDuration().LessThan(this.Duration) {
		return false
	}
	// Compare property "endTime"
	if this.EndTime.LessThan(o.GetEndTime()) {
		return true
	} else if o.GetEndTime().LessThan(this.EndTime) {
		return false
	}
	// Compare property "generator"
	if this.Generator.LessThan(o.GetGenerator()) {
		return true
	} else if o.GetGenerator().LessThan(this.Generator) {
		return false
	}
	// Compare property "icon"
	if this.Icon.LessThan(o.GetIcon()) {
		return true
	} else if o.GetIcon().LessThan(this.Icon) {
		return false
	}
	// Compare property "id"
	if this.Id.LessThan(o.GetId()) {
		return true
	} else if o.GetId().LessThan(this.Id) {
		return false
	}
	// Compare property "image"
	if this.Image.LessThan(o.GetImage()) {
		return true
	} else if o.GetImage().LessThan(this.Image) {
		return false
	}
	// Compare property "inReplyTo"
	if this.InReplyTo.LessThan(o.GetInReplyTo()) {
		return true
	} else if o.GetInReplyTo().LessThan(this.InReplyTo) {
		return false
	}
	// Compare property "latitude"
	if this.Latitude.LessThan(o.GetLatitude()) {
		return true
	} else if o.GetLatitude().LessThan(this.Latitude) {
		return false
	}
	// Compare property "location"
	if this.Location.LessThan(o.GetLocation()) {
		return true
	} else if o.GetLocation().LessThan(this.Location) {
		return false
	}
	// Compare property "longitude"
	if this.Longitude.LessThan(o.GetLongitude()) {
		return true
	} else if o.GetLongitude().LessThan(this.Longitude) {
		return false
	}
	// Compare property "mediaType"
	if this.MediaType.LessThan(o.GetMediaType()) {
		return true
	} else if o.GetMediaType().LessThan(this.MediaType) {
		return false
	}
	// Compare property "name"
	if this.Name.LessThan(o.GetName()) {
		return true
	} else if o.GetName().LessThan(this.Name) {
		return false
	}
	// Compare property "object"
	if this.Object.LessThan(o.GetObject()) {
		return true
	} else if o.GetObject().LessThan(this.Object) {
		return false
	}
	// Compare property "preview"
	if this.Preview.LessThan(o.GetPreview()) {
		return true
	} else if o.GetPreview().LessThan(this.Preview) {
		return false
	}
	// Compare property "published"
	if this.Published.LessThan(o.GetPublished()) {
		return true
	} else if o.GetPublished().LessThan(this.Published) {
		return false
	}
	// Compare property "radius"
	if this.Radius.LessThan(o.GetRadius()) {
		return true
	} else if o.GetRadius().LessThan(this.Radius) {
		return false
	}
	// Compare property "replies"
	if this.Replies.LessThan(o.GetReplies()) {
		return true
	} else if o.GetReplies().LessThan(this.Replies) {
		return false
	}
	// Compare property "startTime"
	if this.StartTime.LessThan(o.GetStartTime()) {
		return true
	} else if o.GetStartTime().LessThan(this.StartTime) {
		return false
	}
	// Compare property "summary"
	if this.Summary.LessThan(o.GetSummary()) {
		return true
	} else if o.GetSummary().LessThan(this.Summary) {
		return false
	}
	// Compare property "tag"
	if this.Tag.LessThan(o.GetTag()) {
		return true
	} else if o.GetTag().LessThan(this.Tag) {
		return false
	}
	// Compare property "to"
	if this.To.LessThan(o.GetTo()) {
		return true
	} else if o.GetTo().LessThan(this.To) {
		return false
	}
	// Compare property "type"
	if this.Type.LessThan(o.GetType()) {
		return true
	} else if o.GetType().LessThan(this.Type) {
		return false
	}
	// Compare property "units"
	if this.Units.LessThan(o.GetUnits()) {
		return true
	} else if o.GetUnits().LessThan(this.Units) {
		return false
	}
	// Compare property "updated"
	if this.Updated.LessThan(o.GetUpdated()) {
		return true
	} else if o.GetUpdated().LessThan(this.Updated) {
		return false
	}
	// Compare property "url"
	if this.Url.LessThan(o.GetUrl()) {
		return true
	} else if o.GetUrl().LessThan(this.Url) {
		return false
	}
	// End: Compare known properties

	// Begin: Compare unknown properties (only by number of them)
	if len(this.unknown) < len(o.GetUnknownProperties()) {
		return true
	} else if len(o.GetUnknownProperties()) < len(this.unknown) {
		return false
	} // End: Compare unknown properties (only by number of them)

	// All properties are the same.
	return false
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format.
func (this Place) Serialize() (interface{}, error) {
	m := make(map[string]interface{})
	// Begin: Serialize known properties
	// Maybe serialize property "accuracy"
	if i, err := this.Accuracy.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Accuracy.Name()] = i
	}
	// Maybe serialize property "altitude"
	if i, err := this.Altitude.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Altitude.Name()] = i
	}
	// Maybe serialize property "attachment"
	if i, err := this.Attachment.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Attachment.Name()] = i
	}
	// Maybe serialize property "attributedTo"
	if i, err := this.AttributedTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.AttributedTo.Name()] = i
	}
	// Maybe serialize property "audience"
	if i, err := this.Audience.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Audience.Name()] = i
	}
	// Maybe serialize property "bcc"
	if i, err := this.Bcc.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Bcc.Name()] = i
	}
	// Maybe serialize property "bto"
	if i, err := this.Bto.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Bto.Name()] = i
	}
	// Maybe serialize property "cc"
	if i, err := this.Cc.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Cc.Name()] = i
	}
	// Maybe serialize property "content"
	if i, err := this.Content.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Content.Name()] = i
	}
	// Maybe serialize property "context"
	if i, err := this.Context.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Context.Name()] = i
	}
	// Maybe serialize property "duration"
	if i, err := this.Duration.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Duration.Name()] = i
	}
	// Maybe serialize property "endTime"
	if i, err := this.EndTime.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.EndTime.Name()] = i
	}
	// Maybe serialize property "generator"
	if i, err := this.Generator.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Generator.Name()] = i
	}
	// Maybe serialize property "icon"
	if i, err := this.Icon.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Icon.Name()] = i
	}
	// Maybe serialize property "id"
	if i, err := this.Id.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Id.Name()] = i
	}
	// Maybe serialize property "image"
	if i, err := this.Image.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Image.Name()] = i
	}
	// Maybe serialize property "inReplyTo"
	if i, err := this.InReplyTo.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.InReplyTo.Name()] = i
	}
	// Maybe serialize property "latitude"
	if i, err := this.Latitude.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Latitude.Name()] = i
	}
	// Maybe serialize property "location"
	if i, err := this.Location.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Location.Name()] = i
	}
	// Maybe serialize property "longitude"
	if i, err := this.Longitude.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Longitude.Name()] = i
	}
	// Maybe serialize property "mediaType"
	if i, err := this.MediaType.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.MediaType.Name()] = i
	}
	// Maybe serialize property "name"
	if i, err := this.Name.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Name.Name()] = i
	}
	// Maybe serialize property "object"
	if i, err := this.Object.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Object.Name()] = i
	}
	// Maybe serialize property "preview"
	if i, err := this.Preview.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Preview.Name()] = i
	}
	// Maybe serialize property "published"
	if i, err := this.Published.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Published.Name()] = i
	}
	// Maybe serialize property "radius"
	if i, err := this.Radius.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Radius.Name()] = i
	}
	// Maybe serialize property "replies"
	if i, err := this.Replies.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Replies.Name()] = i
	}
	// Maybe serialize property "startTime"
	if i, err := this.StartTime.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.StartTime.Name()] = i
	}
	// Maybe serialize property "summary"
	if i, err := this.Summary.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Summary.Name()] = i
	}
	// Maybe serialize property "tag"
	if i, err := this.Tag.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Tag.Name()] = i
	}
	// Maybe serialize property "to"
	if i, err := this.To.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.To.Name()] = i
	}
	// Maybe serialize property "type"
	if i, err := this.Type.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Type.Name()] = i
	}
	// Maybe serialize property "units"
	if i, err := this.Units.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Units.Name()] = i
	}
	// Maybe serialize property "updated"
	if i, err := this.Updated.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Updated.Name()] = i
	}
	// Maybe serialize property "url"
	if i, err := this.Url.Serialize(); err != nil {
		return nil, err
	} else if i != nil {
		m[this.Url.Name()] = i
	}
	// End: Serialize known properties

	// Begin: Serialize unknown properties
	for k, v := range this.unknown {
		// To be safe, ensure we aren't overwriting a known property
		if _, has := m[k]; !has {
			m[k] = v
		}
	}
	// End: Serialize unknown properties

	return m, nil
}

// SetAccuracy returns the "accuracy" property if it exists, and nil otherwise.
func (this Place) SetAccuracy(i vocab.AccuracyPropertyInterface) {
	this.Accuracy = i
}

// SetAltitude returns the "altitude" property if it exists, and nil otherwise.
func (this Place) SetAltitude(i vocab.AltitudePropertyInterface) {
	this.Altitude = i
}

// SetAttachment returns the "attachment" property if it exists, and nil otherwise.
func (this Place) SetAttachment(i vocab.AttachmentPropertyInterface) {
	this.Attachment = i
}

// SetAttributedTo returns the "attributedTo" property if it exists, and nil
// otherwise.
func (this Place) SetAttributedTo(i vocab.AttributedToPropertyInterface) {
	this.AttributedTo = i
}

// SetAudience returns the "audience" property if it exists, and nil otherwise.
func (this Place) SetAudience(i vocab.AudiencePropertyInterface) {
	this.Audience = i
}

// SetBcc returns the "bcc" property if it exists, and nil otherwise.
func (this Place) SetBcc(i vocab.BccPropertyInterface) {
	this.Bcc = i
}

// SetBto returns the "bto" property if it exists, and nil otherwise.
func (this Place) SetBto(i vocab.BtoPropertyInterface) {
	this.Bto = i
}

// SetCc returns the "cc" property if it exists, and nil otherwise.
func (this Place) SetCc(i vocab.CcPropertyInterface) {
	this.Cc = i
}

// SetContent returns the "content" property if it exists, and nil otherwise.
func (this Place) SetContent(i vocab.ContentPropertyInterface) {
	this.Content = i
}

// SetContext returns the "context" property if it exists, and nil otherwise.
func (this Place) SetContext(i vocab.ContextPropertyInterface) {
	this.Context = i
}

// SetDuration returns the "duration" property if it exists, and nil otherwise.
func (this Place) SetDuration(i vocab.DurationPropertyInterface) {
	this.Duration = i
}

// SetEndTime returns the "endTime" property if it exists, and nil otherwise.
func (this Place) SetEndTime(i vocab.EndTimePropertyInterface) {
	this.EndTime = i
}

// SetGenerator returns the "generator" property if it exists, and nil otherwise.
func (this Place) SetGenerator(i vocab.GeneratorPropertyInterface) {
	this.Generator = i
}

// SetIcon returns the "icon" property if it exists, and nil otherwise.
func (this Place) SetIcon(i vocab.IconPropertyInterface) {
	this.Icon = i
}

// SetId returns the "id" property if it exists, and nil otherwise.
func (this Place) SetId(i vocab.IdPropertyInterface) {
	this.Id = i
}

// SetImage returns the "image" property if it exists, and nil otherwise.
func (this Place) SetImage(i vocab.ImagePropertyInterface) {
	this.Image = i
}

// SetInReplyTo returns the "inReplyTo" property if it exists, and nil otherwise.
func (this Place) SetInReplyTo(i vocab.InReplyToPropertyInterface) {
	this.InReplyTo = i
}

// SetLatitude returns the "latitude" property if it exists, and nil otherwise.
func (this Place) SetLatitude(i vocab.LatitudePropertyInterface) {
	this.Latitude = i
}

// SetLocation returns the "location" property if it exists, and nil otherwise.
func (this Place) SetLocation(i vocab.LocationPropertyInterface) {
	this.Location = i
}

// SetLongitude returns the "longitude" property if it exists, and nil otherwise.
func (this Place) SetLongitude(i vocab.LongitudePropertyInterface) {
	this.Longitude = i
}

// SetMediaType returns the "mediaType" property if it exists, and nil otherwise.
func (this Place) SetMediaType(i vocab.MediaTypePropertyInterface) {
	this.MediaType = i
}

// SetName returns the "name" property if it exists, and nil otherwise.
func (this Place) SetName(i vocab.NamePropertyInterface) {
	this.Name = i
}

// SetObject returns the "object" property if it exists, and nil otherwise.
func (this Place) SetObject(i vocab.ObjectPropertyInterface) {
	this.Object = i
}

// SetPreview returns the "preview" property if it exists, and nil otherwise.
func (this Place) SetPreview(i vocab.PreviewPropertyInterface) {
	this.Preview = i
}

// SetPublished returns the "published" property if it exists, and nil otherwise.
func (this Place) SetPublished(i vocab.PublishedPropertyInterface) {
	this.Published = i
}

// SetRadius returns the "radius" property if it exists, and nil otherwise.
func (this Place) SetRadius(i vocab.RadiusPropertyInterface) {
	this.Radius = i
}

// SetReplies returns the "replies" property if it exists, and nil otherwise.
func (this Place) SetReplies(i vocab.RepliesPropertyInterface) {
	this.Replies = i
}

// SetStartTime returns the "startTime" property if it exists, and nil otherwise.
func (this Place) SetStartTime(i vocab.StartTimePropertyInterface) {
	this.StartTime = i
}

// SetSummary returns the "summary" property if it exists, and nil otherwise.
func (this Place) SetSummary(i vocab.SummaryPropertyInterface) {
	this.Summary = i
}

// SetTag returns the "tag" property if it exists, and nil otherwise.
func (this Place) SetTag(i vocab.TagPropertyInterface) {
	this.Tag = i
}

// SetTo returns the "to" property if it exists, and nil otherwise.
func (this Place) SetTo(i vocab.ToPropertyInterface) {
	this.To = i
}

// SetType returns the "type" property if it exists, and nil otherwise.
func (this Place) SetType(i vocab.TypePropertyInterface) {
	this.Type = i
}

// SetUnits returns the "units" property if it exists, and nil otherwise.
func (this Place) SetUnits(i vocab.UnitsPropertyInterface) {
	this.Units = i
}

// SetUpdated returns the "updated" property if it exists, and nil otherwise.
func (this Place) SetUpdated(i vocab.UpdatedPropertyInterface) {
	this.Updated = i
}

// SetUrl returns the "url" property if it exists, and nil otherwise.
func (this Place) SetUrl(i vocab.UrlPropertyInterface) {
	this.Url = i
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this Place) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
	if i == nil {
		return toMerge
	}
	for k, v := range i.JSONLDContext() {
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		toMerge[k] = v
	}
	return toMerge
}