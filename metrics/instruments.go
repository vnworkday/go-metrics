package metrics

import (
	"github.com/pkg/errors"
	tags2 "github.com/vnworkday/go-metrics/tags"
	"github.com/vnworkday/go-metrics/units"
)

type InstrumentOptions struct {
	unit units.Unit
	tags []tags2.Tag
	desc string
}

func (io InstrumentOptions) Unit() units.Unit {
	return io.unit
}

func (io InstrumentOptions) Tags() []tags2.Tag {
	return io.tags
}

func (io InstrumentOptions) Desc() string {
	return io.desc
}

func NewInstrumentOptions() InstrumentOptions {
	return InstrumentOptions{}
}

func (io InstrumentOptions) clone() InstrumentOptions {
	return InstrumentOptions{
		unit: io.unit,
		tags: io.tags,
		desc: io.desc,
	}
}

func (io InstrumentOptions) WithUnit(unit units.Unit) InstrumentOptions {
	clone := io.clone()
	clone.unit = unit
	return clone
}

func (io InstrumentOptions) WithTags(tags ...tags2.Tag) InstrumentOptions {
	clone := io.clone()
	clone.tags = append(clone.tags, tags...)
	return clone
}

func (io InstrumentOptions) WithDesc(desc string) InstrumentOptions {
	clone := io.clone()
	clone.desc = desc
	return clone
}

// MergeInstrumentOptions merges multiple InstrumentOptions into a single InstrumentOptions.
// If more than one unit is specified, an error is returned.
// If more than one description is specified, the last one is used.
func MergeInstrumentOptions(options ...InstrumentOptions) (InstrumentOptions, error) {
	var tags []tags2.Tag
	var unit units.Unit
	var desc string

	for _, option := range options {
		tags = append(tags, option.Tags()...)
		desc = option.Desc()
		if option.Unit() != unit {
			if unit != "" {
				return InstrumentOptions{}, errors.New("more than one unit specified")
			} else {
				unit = option.Unit()
			}
		}
	}
	return InstrumentOptions{
		unit: unit,
		tags: tags,
		desc: desc,
	}, nil
}
