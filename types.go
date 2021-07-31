package main

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	SourceNodeID   uuid.UUID `json:"SourceNodeId"`
	SourceNodePort string    `json:"SourceNodePort"`

	TargetNodeID   uuid.UUID `json:"TargetNodeId"`
	TargetNodePort string    `json:"TargetNodePort"`

	// not clear what the structure of these are, as I can't get GoNorth to
	// export them.
	Vertices interface{} `json:"Vertices"`
	Label    interface{} `json:"Label"`
}

type TextNode struct {
	ID uuid.UUID `json:"Id"`

	Text string `json:"Text"`

	// position on editor graph
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

type ChoiceNode struct {
	ID uuid.UUID `json:"Id"`

	Choices []ChoiceNodeChoice `json:"Choices"`

	// "default" choice id? or is this last-selected? who knows. couldn't get
	// it to change.
	//
	// there's two entries in the ones i've generated, but Id starts at 0,
	// so why is it 2 in those?
	CurrentChoiceID int `json:"CurrentChoiceId"`

	// position on editor graph
	X float64 `json:"X"`
	Y float64 `json:"Y"`
}

type ChoiceNodeChoice struct {
	ID int `json:"Id"`

	IsRepeatable bool           `json:"IsRepeatable"`
	Condition    *ConditionNode `json:"Condition"`

	Text string `json:"Text"`
}

type ConditionNode struct {
	ID int `json:"Id"`

	// can't figure out what this is, couldn't get GN to export it
	DependsOnObjects []interface{} `json:"DependsOnObjects"`

	// json-encoded string.
	// not going to figure out how to support this _for now_.
	// it's unlikely that _i_ will be using the cond editor in GN.
	//
	// that being said, it _would_ be incredibly convenient to map
	// conds in the editor.
	//
	// as an example, a rand value string (jsonified):
	// "[{\"conditionType\":19,\"conditionData\":{\"operator\":\"=\",\"minValue\":1,\"maxValue\":3,\"compareValue\":2}}]"
	//
	// the only problem i see is that i'd have to go all-in on GN to make
	// effective use of this - items, skills, etc would all need to be managed
	// in there.
	//
	// at least for now i only want to use it for dialogue trees, perhaps in
	// the future i'll expand that (+ this code)
	ConditionElements string `json:"ConditionElements"`
}

type Dialogue struct {
	ID                uuid.UUID `json:"Id"`
	RelatedObjectedID uuid.UUID `json:"RelatedObjectId"`

	Condition []ConditionNode `json:"Condition"`

	Link []Link `json:"Link"`

	PlayerText []TextNode
	NPCText    []TextNode
	Choice     []ChoiceNode

	Action []interface{} `json:"Action"`

	Reference []interface{} `json:"Reference"`

	IsImplemented bool `json:"IsImplemented"`

	ModifiedOn time.Time `json:"ModifiedOn"`
	ModifiedBy uuid.UUID `json:"ModifiedBy"`
}

type NPC struct {
	ID uuid.UUID `json:"Id"`

	ParentFolderID uuid.UUID `json:"ParentFolderId"`
	ProjectID      uuid.UUID `json:"ProjectId"`
	TemplateID     uuid.UUID `json:"TemplateId"`

	Name               string `json:"Name"`
	ImageFile          string `json:"ImageFile"`
	ThumbnailImageFile string `json:"ThumbnailImageFile"`

	Dialogue Dialogue `json:"Dialog"`

	// not working with this stuff currently, npc behaviors specified in the ai subsys
	StateMachine    interface{}   `json:"StateMachine"`
	ExportSnippets  []interface{} `json:"ExportSnippets"`
	IsPlayerNPC     bool          `json:"IsPlayerNpc"`
	NameGenTemplate string        `json:"NameGenTemplate"`
	Inventory       []interface{} `json:"Inventory"`
	Skills          []interface{} `json:"Skills"`
	DailyRoutine    []interface{} `json:"DailyRoutine"`
	Fields          []interface{} `json:"Fields"`
	Tags            []interface{} `json:"Tags"`

	IsImplemented bool `json:"IsImplemented"`

	ModifiedOn time.Time `json:"ModifiedOn"`
	ModifiedBy uuid.UUID `json:"ModifiedBy"`
}
