## Possible surveys

### Midterm evaluation
Survey
#id, file

StudentSurvey
#id, #token

Questions
#(survey,q), type

Answer
#(survey,q,student), value(json)


Authentication
admin token.

CloseSurvey()


A: id A, value

### Final evaluation

### After the internship

1 Question, multiple choice

```
1 - recherche d'emploi - looking for a job

2 - poursuite d'étude - pursuit of higher education

- working in a company
	- CDD - Fixed-term contract
		3 - embauché dans l'entreprise d'accueil - hired in the hosting company
		4 - création d'entreprise - entrepreneurship

	- CDI - permanent contract
		5 - embauché dans l'entreprise d'accueil - hired in the hosting company
		6 - création d'entreprise - entrepreneurship

7 - congés sabbatique - sabbatical leave
```

## Data management

To store:
	- token to write the survey
	- token to read the survey
	- survey id
	- student
	- answers

- supervisor can write midterm & final
- tutor can read midterm, final, after
- student can write other

surveys(student, type, read_token, write_token, answers)
survey_accesses(email, token)


### After
-> just store the answer in the internship table
