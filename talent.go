package main

import "fmt"

// Talent is a character's trait.
type Talent struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Aptitudes    []Aptitude    `json:"aptitudes"`
	Tier         int           `json:"tier"`
	Requirements []Requirement `json:"requirements"`
	Speciality   string        `json:"-"`
	Value        int           `json:"-"`
	Stackable    bool          `json:"stackable"`
}

// Cost returns the cost of the talent given the character's aptitudes and the current tier.
func (t Talent) Cost(universe Universe, character Character) (int, error) {

	// Return the price as determined by the cost matrix.
	return universe.Costs.Price("talent", character.CountMatchingAptitudes(t.Aptitudes), t.Tier)
}

// FullName return the name of the talent and it's speciality if defined.
func (t Talent) FullName() string {
	if len(t.Speciality) == 0 {
		return t.Name
	}
	return fmt.Sprintf("%s: %s", t.Name, t.Speciality)
}
