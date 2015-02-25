function lang() {
	return {
	fr : [
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
	"Si vous aviez un emploi d’ingénieur à pourvoir, l’engageriez vous ?"
	],
	en : [
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
	"If you had an engineer position to fill, would you employ him/her ?"		
	]
};
}

function choice() {
	$(".choice").each(function (idx, e) {		
		var buf = "<select name='" + e.id + "' class='col-md-3'>";		
		buf += "<option value=''> - </option>";		
		buf += "<option value='4' class='ex'></option>";
		buf += "<option value='3' class='go'></option>";
		buf += "<option value='2' class='sat'></option>";
		buf += "<option value='1' class='unsat'></option>";		
		buf += "<option value='0' class='na'></option>";
		buf += "</select>";
		$(e).after(buf);
	});
}

$( document ).ready(function () {	
	$(".alert").hide();
	yesno();
	textarea();
	choice();
	fill();
	setLang('fr');		
	$(":radio").iCheck();
	$("select").selecter();		
	fill();
});