package spy

import (
	"log"
	"time"

	"github.com/fhermeni/wints/schema"
)

/*
* Precondition:
* - student is late
* - review deadline passed
 */

type Tutored struct {
	reviews map[string]time.Duration
	st      schema.Internshipser
}

func NewTutor(s schema.Internshipser, r map[string]time.Duration) *Tutored {
	return &Tutored{st: s, reviews: r}
}

func (tut *Tutored) Run() {
	log.Println("Running Tutored spy")
	assigns := make(map[string][]schema.Internship)

	ints, err := tut.st.Internships()
	if err != nil {
		return
	}
	for _, i := range ints {
		em := i.Convention.Tutor.Person.Email
		l, ok := assigns[em]
		if !ok {
			l = make([]schema.Internship, 0, 0)
		}
		l = append(l, i)
		assigns[em] = l
	}

	for _, l := range assigns {
		tut.spy(l)
	}
}

func (tut *Tutored) spy(ints []schema.Internship) {
	for _, i := range ints {
		if tut.late(i) {
			tut.report(ints)
		}
	}
}

func (tut *Tutored) report(ints []schema.Internship) {
	log.Println("Report for " + ints[0].Convention.Tutor.String())
	/**
	  Bonjour

	  Toto:
	  	Rapport 'midterm' (deadline ...) : - | en retard de Y jours
	  	Rapport 'foo' (deadline pour la review ...) : note en retard de Z jours

	*/
}

func (tut *Tutored) late(i schema.Internship) bool {
	now := time.Now()
	for _, r := range i.Reports {
		if r.Delivery == nil && r.Deadline.Before(now) {
			//late delivery
			return true
		}
		ps, _ := tut.reviews[r.Kind]
		if r.Delivery != nil && r.Delivery.Add(ps).Before(now) {
			//late review
			return true
		}
	}
	return false
}
