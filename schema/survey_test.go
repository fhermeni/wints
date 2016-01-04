package schema

import (
	"log"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

var buffer = `
en = "Midterm survey"
fr = "Évaluation intermédiaire"

[[Categories]]
en = "Ponctuality"
fr = "Ponctualité"

[[Categories.Q]]
en = "Is the trainee punctual at work ?"
fr = "Est-il ponctuel ?"
type = "yesno"

[[Categories.Q]]
en = "Comments"
fr = "Commentaires"
type = "comment"

[[Categories]]
en = "Integration inside the company"
fr = "Intégration dans l'entreprise"

[[Categories.Q]]
en = "Does the trainee seek to communicate with the others ?"
fr = "Cherche-t-il à communiquer avec les autres ?"
type = "textarea"
`

func TestParsing(t *testing.T) {

	var s Survey
	if _, err := toml.Decode(buffer, &s); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 2, len(s.Categories))
	assert.Equal(t, 2, len(s.Categories[0].Q))
	assert.Equal(t, 1, len(s.Categories[1].Q))
	log.Println(s)

}
