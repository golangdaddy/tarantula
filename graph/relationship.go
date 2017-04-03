package graph

type Relationships []*Relationship

type Relationship struct {
	In        *Class
	Out       *Class
	Predicate string
}

func (db *Database) AddRelationship(relationship *Relationship) {

	db.Classes.Lock()

	db.Classes.Relationships = append(db.Classes.Relationships, relationship)

	db.Classes.Unlock()
}
