package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Universe represents a set of configuration, often refered as data or database.
type Universe struct {
	Backgrounds     map[string][]Background `json:"backgrounds"`
	Aptitudes       []Aptitude              `json:"aptitudes"`
	Characteristics []Characteristic        `json:"characteristics"`
	Gauges          []Gauge                 `json:"gauges"`
	Skills          []Skill                 `json:"skills"`
	Talents         []Talent                `json:"talents"`
	Costs           CostMatrix              `json:"costs"`
}

// ParseUniverse load an from a plain JSON file.
// It returns a well-formed universe that describe all the components of a game setting.
func ParseUniverse(file io.Reader) (Universe, error) {

	// Open and parse universe.
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return Universe{}, err
	}
	universe := Universe{}
	err = json.Unmarshal(raw, &universe)
	if err != nil {
		return Universe{}, err
	}

	// Add the type value to each history defined.
	for typ, backgrounds := range universe.Backgrounds {
		for i := range backgrounds {
			universe.Backgrounds[typ][i].Type = typ
		}
	}

	return universe, nil
}

func (u Universe) FindCoster(label string) (Coster, bool) {
	characteristic, found := u.FindCharacteristic(label)
	if found {
		return characteristic, true
	}

	skill, found := u.FindSkill(label)
	if found {
		return skill, true
	}

	talent, found := u.FindTalent(label)
	if found {
		return talent, true
	}

	aptitude, found := u.FindAptitude(label)
	if found {
		return aptitude, true
	}

	gauge, found := u.FindGauge(label)
	if found {
		return gauge, true
	}

	return nil, false
}

// FindCharacteristic returns the characteristic correponding to the given label or a zero-value, and a boolean indicating if it was found.
func (u Universe) FindCharacteristic(label string) (Characteristic, bool) {

	// Characteristics upgrades are defined by a name and a value, separated by a space, so we need to look for the first
	// part of the label.
	// Examples: STR +5, FEL -1, TOU 40.
	fields := split(label, ' ')
	name := fields[0]

	for _, characteristic := range u.Characteristics {
		if characteristic.Name == name {
			return characteristic, true
		}
	}

	return Characteristic{}, false
}

// FindSkill returns the skill corresponding to the given label or a zero-value, and a boolean indicating if it was found.
func (u Universe) FindSkill(label string) (Skill, bool) {

	// Skills upgrades are defined by a name and eventually a speciality, separated by a colon.
	// Examples: Common Lore: Dark Gods
	fields := split(label, ':')
	name := fields[0]

	for _, skill := range u.Skills {
		if skill.Name == name {

			if len(fields) == 2 {
				skill.Speciality = fields[1]
			}

			return skill, true
		}
	}

	return Skill{}, false
}

// FindTalent returns the talent corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindTalent(label string) (Talent, bool) {

	// Talents upgrades are defined by a name and eventually a speciality, separated by a colon.
	// Examples: Psychic Resistance: Fear
	fields := split(label, ':')
	name := fields[0]

	for _, talent := range u.Talents {
		if talent.Name == name {

			if len(fields) == 2 {
				talent.Speciality = fields[1]
			}

			return talent, true
		}
	}

	return Talent{}, false
}

// FindAptitude returns the aptitude corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindAptitude(label string) (Aptitude, bool) {

	for _, aptitude := range u.Aptitudes {
		if string(aptitude) == label {
			return aptitude, true
		}
	}

	return Aptitude(""), false
}

// FindGauge returns the gauge corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindGauge(label string) (Gauge, bool) {

	// Gauges upgrades are defined by a name and a value, separated by a space.
	name := split(label, ' ')[0]

	for _, gauge := range u.Gauges {
		if gauge.Name == name {
			return gauge, true
		}
	}

	return Gauge{}, false
}

// FindBackground returns the background corresponding to the given label, and a boolean indicating if it was found.
func (u Universe) FindBackground(typ string, label string) (Background, bool) {

	backgrounds, found := u.Backgrounds[typ]
	if !found {
		return Background{}, false
	}

	for _, background := range backgrounds {
		if background.Name == label {
			return background, true
		}
	}

	return Background{}, false
}
