package typelink

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// A Link is an indirect, qualified reference to a resource identified by a URL.
// The fundamental model for links is established by [ RFC5988]. Many of the
// properties defined by the Activity Vocabulary allow values that are either
// instances of Object or Link. When a Link is used, it establishes a
// qualified relation connecting the subject (the containing object) to the
// resource identified by the href. Properties of the Link are properties of
// the reference as opposed to properties of the resource.
//
// Example 2 (https://www.w3.org/TR/activitystreams-vocabulary/#ex2-jsonld):
//   {
//     "hreflang": "en",
//     "mediaType": "text/html",
//     "name": "An example link",
//     "type": "owl:Class",
//     "url": "http://example.org/abc"
//   }
type ActivityStreamsLink struct {
	ActivityStreamsAttributedTo vocab.ActivityStreamsAttributedToProperty
	ActivityStreamsHeight       vocab.ActivityStreamsHeightProperty
	ActivityStreamsHref         vocab.ActivityStreamsHrefProperty
	ActivityStreamsHreflang     vocab.ActivityStreamsHreflangProperty
	ActivityStreamsId           vocab.ActivityStreamsIdProperty
	ActivityStreamsMediaType    vocab.ActivityStreamsMediaTypeProperty
	ActivityStreamsName         vocab.ActivityStreamsNameProperty
	ActivityStreamsPreview      vocab.ActivityStreamsPreviewProperty
	ActivityStreamsRel          vocab.ActivityStreamsRelProperty
	ActivityStreamsSummary      vocab.ActivityStreamsSummaryProperty
	ActivityStreamsType         vocab.ActivityStreamsTypeProperty
	ActivityStreamsWidth        vocab.ActivityStreamsWidthProperty
	alias                       string
	unknown                     map[string]interface{}
}

// ActivityStreamsLinkExtends returns true if the Link type extends from the other
// type.
func ActivityStreamsLinkExtends(other vocab.Type) bool {
	// Shortcut implementation: this does not extend anything.
	return false
}

// DeserializeLink creates a Link from a map representation that has been
// unmarshalled from a text or binary format.
func DeserializeLink(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsLink, error) {
	alias := ""
	aliasPrefix := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
		aliasPrefix = a + ":"
	}
	this := &ActivityStreamsLink{
		alias:   alias,
		unknown: make(map[string]interface{}),
	}
	if typeValue, ok := m["type"]; !ok {
		return nil, fmt.Errorf("no \"type\" property in map")
	} else if typeString, ok := typeValue.(string); ok {
		typeName := strings.TrimPrefix(typeString, aliasPrefix)
		if typeName != "Link" {
			return nil, fmt.Errorf("\"type\" property is not of %q type: %s", "Link", typeName)
		}
		// Fall through, success in finding a proper Type
	} else if arrType, ok := typeValue.([]interface{}); ok {
		found := false
		for _, elemVal := range arrType {
			if typeString, ok := elemVal.(string); ok && strings.TrimPrefix(typeString, aliasPrefix) == "Link" {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("could not find a \"type\" property of value %q", "Link")
		}
		// Fall through, success in finding a proper Type
	} else {
		return nil, fmt.Errorf("\"type\" property is unrecognized type: %T", typeValue)
	}
	// Begin: Known property deserialization
	if p, err := mgr.DeserializeAttributedToPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsAttributedTo = p
	}
	if p, err := mgr.DeserializeHeightPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsHeight = p
	}
	if p, err := mgr.DeserializeHrefPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsHref = p
	}
	if p, err := mgr.DeserializeHreflangPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsHreflang = p
	}
	if p, err := mgr.DeserializeIdPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsId = p
	}
	if p, err := mgr.DeserializeMediaTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsMediaType = p
	}
	if p, err := mgr.DeserializeNamePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsName = p
	}
	if p, err := mgr.DeserializePreviewPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsPreview = p
	}
	if p, err := mgr.DeserializeRelPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsRel = p
	}
	if p, err := mgr.DeserializeSummaryPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsSummary = p
	}
	if p, err := mgr.DeserializeTypePropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsType = p
	}
	if p, err := mgr.DeserializeWidthPropertyActivityStreams()(m, aliasMap); err != nil {
		return nil, err
	} else if p != nil {
		this.ActivityStreamsWidth = p
	}
	// End: Known property deserialization

	// Begin: Unknown deserialization
	for k, v := range m {
		// Begin: Code that ensures a property name is unknown
		if k == "attributedTo" {
			continue
		} else if k == "height" {
			continue
		} else if k == "href" {
			continue
		} else if k == "hreflang" {
			continue
		} else if k == "id" {
			continue
		} else if k == "mediaType" {
			continue
		} else if k == "name" {
			continue
		} else if k == "nameMap" {
			continue
		} else if k == "preview" {
			continue
		} else if k == "rel" {
			continue
		} else if k == "summary" {
			continue
		} else if k == "summaryMap" {
			continue
		} else if k == "type" {
			continue
		} else if k == "width" {
			continue
		} // End: Code that ensures a property name is unknown

		this.unknown[k] = v
	}
	// End: Unknown deserialization

	return this, nil
}

// LinkIsDisjointWith returns true if the other provided type is disjoint with the
// Link type.
func LinkIsDisjointWith(other vocab.Type) bool {
	disjointWith := []string{"Accept", "Activity", "Add", "Announce", "Application", "Arrive", "Article", "Audio", "Block", "Collection", "CollectionPage", "Create", "Delete", "Dislike", "Document", "Event", "Flag", "Follow", "Group", "Ignore", "Image", "IntransitiveActivity", "Invite", "Join", "Leave", "Like", "Listen", "Move", "Note", "Object", "Offer", "OrderedCollection", "OrderedCollectionPage", "OrderedCollectionPage", "Organization", "Page", "Person", "Place", "Profile", "Question", "Read", "Reject", "Relationship", "Remove", "Service", "TentativeAccept", "TentativeReject", "Tombstone", "Travel", "Undo", "Update", "Video", "View"}
	for _, disjoint := range disjointWith {
		if disjoint == other.GetName() {
			return true
		}
	}
	return false
}

// LinkIsExtendedBy returns true if the other provided type extends from the Link
// type.
func LinkIsExtendedBy(other vocab.Type) bool {
	extensions := []string{"Mention"}
	for _, ext := range extensions {
		if ext == other.GetName() {
			return true
		}
	}
	return false
}

// NewActivityStreamsLink creates a new Link type
func NewActivityStreamsLink() *ActivityStreamsLink {
	typeProp := typePropertyConstructor()
	typeProp.AppendXMLSchemaString("Link")
	return &ActivityStreamsLink{
		ActivityStreamsType: typeProp,
		alias:               "",
		unknown:             make(map[string]interface{}, 0),
	}
}

// GetActivityStreamsAttributedTo returns the "attributedTo" property if it
// exists, and nil otherwise.
func (this ActivityStreamsLink) GetActivityStreamsAttributedTo() vocab.ActivityStreamsAttributedToProperty {
	return this.ActivityStreamsAttributedTo
}

// GetActivityStreamsHeight returns the "height" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsHeight() vocab.ActivityStreamsHeightProperty {
	return this.ActivityStreamsHeight
}

// GetActivityStreamsHref returns the "href" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsHref() vocab.ActivityStreamsHrefProperty {
	return this.ActivityStreamsHref
}

// GetActivityStreamsHreflang returns the "hreflang" property if it exists, and
// nil otherwise.
func (this ActivityStreamsLink) GetActivityStreamsHreflang() vocab.ActivityStreamsHreflangProperty {
	return this.ActivityStreamsHreflang
}

// GetActivityStreamsId returns the "id" property if it exists, and nil otherwise.
func (this ActivityStreamsLink) GetActivityStreamsId() vocab.ActivityStreamsIdProperty {
	return this.ActivityStreamsId
}

// GetActivityStreamsMediaType returns the "mediaType" property if it exists, and
// nil otherwise.
func (this ActivityStreamsLink) GetActivityStreamsMediaType() vocab.ActivityStreamsMediaTypeProperty {
	return this.ActivityStreamsMediaType
}

// GetActivityStreamsName returns the "name" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsName() vocab.ActivityStreamsNameProperty {
	return this.ActivityStreamsName
}

// GetActivityStreamsPreview returns the "preview" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsPreview() vocab.ActivityStreamsPreviewProperty {
	return this.ActivityStreamsPreview
}

// GetActivityStreamsRel returns the "rel" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsRel() vocab.ActivityStreamsRelProperty {
	return this.ActivityStreamsRel
}

// GetActivityStreamsSummary returns the "summary" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsSummary() vocab.ActivityStreamsSummaryProperty {
	return this.ActivityStreamsSummary
}

// GetActivityStreamsType returns the "type" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsType() vocab.ActivityStreamsTypeProperty {
	return this.ActivityStreamsType
}

// GetActivityStreamsWidth returns the "width" property if it exists, and nil
// otherwise.
func (this ActivityStreamsLink) GetActivityStreamsWidth() vocab.ActivityStreamsWidthProperty {
	return this.ActivityStreamsWidth
}

// GetName returns the name of this type.
func (this ActivityStreamsLink) GetName() string {
	return "Link"
}

// GetUnknownProperties returns the unknown properties for the Link type. Note
// that this should not be used by app developers. It is only used to help
// determine which implementation is LessThan the other. Developers who are
// creating a different implementation of this type's interface can use this
// method in their LessThan implementation, but routine ActivityPub
// applications should not use this to bypass the code generation tool.
func (this ActivityStreamsLink) GetUnknownProperties() map[string]interface{} {
	return this.unknown
}

// IsExtending returns true if the Link type extends from the other type.
func (this ActivityStreamsLink) IsExtending(other vocab.Type) bool {
	return ActivityStreamsLinkExtends(other)
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// type and the specific properties that are set. The value in the map is the
// alias used to import the type and its properties.
func (this ActivityStreamsLink) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	m = this.helperJSONLDContext(this.ActivityStreamsAttributedTo, m)
	m = this.helperJSONLDContext(this.ActivityStreamsHeight, m)
	m = this.helperJSONLDContext(this.ActivityStreamsHref, m)
	m = this.helperJSONLDContext(this.ActivityStreamsHreflang, m)
	m = this.helperJSONLDContext(this.ActivityStreamsId, m)
	m = this.helperJSONLDContext(this.ActivityStreamsMediaType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsName, m)
	m = this.helperJSONLDContext(this.ActivityStreamsPreview, m)
	m = this.helperJSONLDContext(this.ActivityStreamsRel, m)
	m = this.helperJSONLDContext(this.ActivityStreamsSummary, m)
	m = this.helperJSONLDContext(this.ActivityStreamsType, m)
	m = this.helperJSONLDContext(this.ActivityStreamsWidth, m)

	return m
}

// LessThan computes if this Link is lesser, with an arbitrary but stable
// determination.
func (this ActivityStreamsLink) LessThan(o vocab.ActivityStreamsLink) bool {
	// Begin: Compare known properties
	// Compare property "attributedTo"
	if lhs, rhs := this.ActivityStreamsAttributedTo, o.GetActivityStreamsAttributedTo(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "height"
	if lhs, rhs := this.ActivityStreamsHeight, o.GetActivityStreamsHeight(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "href"
	if lhs, rhs := this.ActivityStreamsHref, o.GetActivityStreamsHref(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "hreflang"
	if lhs, rhs := this.ActivityStreamsHreflang, o.GetActivityStreamsHreflang(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "id"
	if lhs, rhs := this.ActivityStreamsId, o.GetActivityStreamsId(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "mediaType"
	if lhs, rhs := this.ActivityStreamsMediaType, o.GetActivityStreamsMediaType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "name"
	if lhs, rhs := this.ActivityStreamsName, o.GetActivityStreamsName(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "preview"
	if lhs, rhs := this.ActivityStreamsPreview, o.GetActivityStreamsPreview(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "rel"
	if lhs, rhs := this.ActivityStreamsRel, o.GetActivityStreamsRel(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "summary"
	if lhs, rhs := this.ActivityStreamsSummary, o.GetActivityStreamsSummary(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "type"
	if lhs, rhs := this.ActivityStreamsType, o.GetActivityStreamsType(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
	// Compare property "width"
	if lhs, rhs := this.ActivityStreamsWidth, o.GetActivityStreamsWidth(); lhs != nil && rhs != nil {
		if lhs.LessThan(rhs) {
			return true
		} else if rhs.LessThan(lhs) {
			return false
		}
	} else if lhs == nil && rhs != nil {
		// Nil is less than anything else
		return true
	} else if rhs != nil && rhs == nil {
		// Anything else is greater than nil
		return false
	} // Else: Both are nil
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
func (this ActivityStreamsLink) Serialize() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	typeName := "Link"
	if len(this.alias) > 0 {
		typeName = this.alias + ":" + "Link"
	}
	m["type"] = typeName
	// Begin: Serialize known properties
	// Maybe serialize property "attributedTo"
	if this.ActivityStreamsAttributedTo != nil {
		if i, err := this.ActivityStreamsAttributedTo.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsAttributedTo.Name()] = i
		}
	}
	// Maybe serialize property "height"
	if this.ActivityStreamsHeight != nil {
		if i, err := this.ActivityStreamsHeight.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsHeight.Name()] = i
		}
	}
	// Maybe serialize property "href"
	if this.ActivityStreamsHref != nil {
		if i, err := this.ActivityStreamsHref.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsHref.Name()] = i
		}
	}
	// Maybe serialize property "hreflang"
	if this.ActivityStreamsHreflang != nil {
		if i, err := this.ActivityStreamsHreflang.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsHreflang.Name()] = i
		}
	}
	// Maybe serialize property "id"
	if this.ActivityStreamsId != nil {
		if i, err := this.ActivityStreamsId.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsId.Name()] = i
		}
	}
	// Maybe serialize property "mediaType"
	if this.ActivityStreamsMediaType != nil {
		if i, err := this.ActivityStreamsMediaType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsMediaType.Name()] = i
		}
	}
	// Maybe serialize property "name"
	if this.ActivityStreamsName != nil {
		if i, err := this.ActivityStreamsName.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsName.Name()] = i
		}
	}
	// Maybe serialize property "preview"
	if this.ActivityStreamsPreview != nil {
		if i, err := this.ActivityStreamsPreview.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsPreview.Name()] = i
		}
	}
	// Maybe serialize property "rel"
	if this.ActivityStreamsRel != nil {
		if i, err := this.ActivityStreamsRel.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsRel.Name()] = i
		}
	}
	// Maybe serialize property "summary"
	if this.ActivityStreamsSummary != nil {
		if i, err := this.ActivityStreamsSummary.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsSummary.Name()] = i
		}
	}
	// Maybe serialize property "type"
	if this.ActivityStreamsType != nil {
		if i, err := this.ActivityStreamsType.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsType.Name()] = i
		}
	}
	// Maybe serialize property "width"
	if this.ActivityStreamsWidth != nil {
		if i, err := this.ActivityStreamsWidth.Serialize(); err != nil {
			return nil, err
		} else if i != nil {
			m[this.ActivityStreamsWidth.Name()] = i
		}
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

// SetActivityStreamsAttributedTo returns the "attributedTo" property if it
// exists, and nil otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsAttributedTo(i vocab.ActivityStreamsAttributedToProperty) {
	this.ActivityStreamsAttributedTo = i
}

// SetActivityStreamsHeight returns the "height" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsHeight(i vocab.ActivityStreamsHeightProperty) {
	this.ActivityStreamsHeight = i
}

// SetActivityStreamsHref returns the "href" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsHref(i vocab.ActivityStreamsHrefProperty) {
	this.ActivityStreamsHref = i
}

// SetActivityStreamsHreflang returns the "hreflang" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsHreflang(i vocab.ActivityStreamsHreflangProperty) {
	this.ActivityStreamsHreflang = i
}

// SetActivityStreamsId returns the "id" property if it exists, and nil otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsId(i vocab.ActivityStreamsIdProperty) {
	this.ActivityStreamsId = i
}

// SetActivityStreamsMediaType returns the "mediaType" property if it exists, and
// nil otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsMediaType(i vocab.ActivityStreamsMediaTypeProperty) {
	this.ActivityStreamsMediaType = i
}

// SetActivityStreamsName returns the "name" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsName(i vocab.ActivityStreamsNameProperty) {
	this.ActivityStreamsName = i
}

// SetActivityStreamsPreview returns the "preview" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsPreview(i vocab.ActivityStreamsPreviewProperty) {
	this.ActivityStreamsPreview = i
}

// SetActivityStreamsRel returns the "rel" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsRel(i vocab.ActivityStreamsRelProperty) {
	this.ActivityStreamsRel = i
}

// SetActivityStreamsSummary returns the "summary" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsSummary(i vocab.ActivityStreamsSummaryProperty) {
	this.ActivityStreamsSummary = i
}

// SetActivityStreamsType returns the "type" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsType(i vocab.ActivityStreamsTypeProperty) {
	this.ActivityStreamsType = i
}

// SetActivityStreamsWidth returns the "width" property if it exists, and nil
// otherwise.
func (this *ActivityStreamsLink) SetActivityStreamsWidth(i vocab.ActivityStreamsWidthProperty) {
	this.ActivityStreamsWidth = i
}

// VocabularyURI returns the vocabulary's URI as a string.
func (this ActivityStreamsLink) VocabularyURI() string {
	return "https://www.w3.org/TR/activitystreams-vocabulary"
}

// helperJSONLDContext obtains the context uris and their aliases from a property,
// if it is not nil.
func (this ActivityStreamsLink) helperJSONLDContext(i jsonldContexter, toMerge map[string]string) map[string]string {
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