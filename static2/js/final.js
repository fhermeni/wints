function lang() {
	return {
		fr: [
			"Fiche d'évaluation finale",
			//
			"Maîtrise des domaines scientifiques et techniques",
			"Capacité d’analyse / compréhension des problèmes",
			"Mise en œuvre de ses connaissance",
			"Aptitude à acquérir  de nouvelles connaissances",
			//
			"Maîtrise des méthodes et des outils de l'ingénieur",
			"Méthodologie / organisation du travail, gestion de projet",
			"Synthèse et communication des résultats, maîtrise des outils de communication",
			//
			"Conduite de l’action et prise de décision",
			"Réalisation des objectifs - Qualité du travail réalisé",
			"Autonomie - initiative /créativité / ouverture d’esprit",
			//
			"Intégration dans une organisation et capacité d’animation",
			"Capacité à s’intégrer dans une équipe",
			"Communication sur ses activités et capacité à rendre compte",
			"Prise en compte des enjeux métiers et économiques - Respect des procédures (qualité, sécurité, santé, sécurité, ...)",
			//
			"Respect des valeurs sociétales, sociales et environnementales",
			"Appropriation  des valeurs, codes, et de la culture de l’équipe et de l’organisation",
			"Attitude / assiduité / ponctualité",
			//
			"Parmi les stagiaires du même niveau d’études que vous avez eu l’occasion d’accueillir, comment classez-vous cet élève ?",
			"Quels conseils donneriez-vous à ce futur ingénieur ?",
			"Observations supplémentaires éventuelles :",
			"Si vous aviez un emploi d’ingénieur à pourvoir, l’engageriez vous ?",
			"Note /20: ",
			"Guide pour la notation: <ul> <li> &lt;10: le travail réalisé ne correspond pas à ce qui est attendu de la part d'un jeune master / ingénieur.</li> <li> 10 à 12: le travail réalisé correspond au strict minimum </li> <li> 13 à 14: le stagiaire a réalisé un travail correct pour un jeune master / ingénieur </li> <li> 15 à 16: très bon stagiaire. Son travail a beaucoup compté pour l'entreprise </li> <li> 17 à 18: Il va être difficile de trouver un aussi bon remplaçant </li> <li> &gt; 18: Le stagiaire était un membre primordial de l 'équipe. Il ne sera pas possible de le remplacer sans une perte claire de connaissances / capacités</li> </ul>"
		],
		en: [
			"Final work placement",
			//
			"Scientific and technical knowledge",
			"Analytical skills and understanding problems ",
			"Application of knowledge ",
			"Willingness and capacity to learn",
			//
			"Mastery of methods and tools for  the engineer",
			"Organizational skills and project management",
			"Ability to summarize  and present results  clearly using communication tools",
			//
			"Task management and decision making",
			"Targets reached. Quality of work ",
			"Autonomy, initiative/creativity/open-mindedness",
			//
			"Aptitude to integrate into an organisation and manage a team ",
			"Aptitude to integrate into a team",
			"Communication skills : ability to report on work",
			"Understanding of economic and business challenges of the company. Following internal procedures (quality, health & safety)",
			//
			"Respect of corporate culture and environmental values",
			"Assimilation of values, codes and company culture",
			"Attitude, attendance  and punctuality",
			//
			"Among students of the same level of studies that you have previously supervised, how would you rate this student ?",
			"What advice if any would you give him/her ?",
			"Additional observations:",
			"If you had an engineer position to fill, would you employ him/her ?",
			"Grade /20: ",
			"Notation guide: <ul> <li> &lt; 10: the work done is below the expectations for a junior master degree / engineer. </li> <li> 10 to 12: the trainee did the bare minimum </li> <li> 13 to 14: the trainee did a correct job for a young master / engineer </li> <li> 15 to 16: Real good trainee.His work mattered significantly for the hosting company </li> <li> 17 to 18: It will be hard to find a valuable substitute for the trainee </li> <li> &gt; 18: The trainee succeeded as being a primordial team member. It will not be possible to replace him without a clear loss in terms of knowledge / skills. </li> </ul>"
		]
	};
}

$(document).ready(function() {
	expand();
	loadSurvey(function() {
		fullfill()
		ui();
	})
});