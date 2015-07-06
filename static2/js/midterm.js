function lang() {
	return {
		fr: [
			"Fiche d'évaluation intermédiaire",
			//
			"Ponctualité",
			"Est-il ponctuel ?",
			//
			"Intégration dans l'entreprise",
			"Cherche-t-il à communiquer avec les autres ?",
			"D’après vous, est-il déjà bien intégré parmi les membres de votre service ?",
			//
			"Travail",
			"Est-il intéressé par son travail ?",
			"S’est-il préoccupé des méthodes de travail de l’entreprise ? ",
			"La quantité de travail fournie est elle satisfaisante ?",
			"Respecte-t-il les délais ?",
			//
			"Compétence techniques",
			"Possède-t-il les compétences techniques nécessaires pour son travail ?",
			"Si non, précisez ses lacunes :",
			"A-t-il eu besoin d’apprendre une nouvelle technique ou un nouveau logiciel ?",
			"Si oui, a-t-il montré sa capacité à apprendre ?",
			"Cherche-t-il à améliorer ses compétences dans certains domaines ?",
			"Est-il autonome ?",
			"Cherche-t-il à aider les autres ?",
			//
			"Évaluation globale",
			"Êtes-vous globalement satisfait du début de ce stage ?",
			"Quels sont d’après vous les points forts de ce stagiaire (un par ligne) :",
			"Quels sont d’après vous les points faibles de ce stagiaire (un par ligne) :"
		],
		en: [
			"Midterm work placement",
			//
			"Punctuality",
			"Is the trainee punctual at work ?",
			//
			"Integration inside the company ?",
			"Does the trainee seek to communicate with the others ?",
			"Is the trainee well integrated within your team ?",
			//
			"Work",
			"Is the trainee interested by his work ?",
			"Was he concerned with the work methods of the company ?",
			"Is the quantity of work provided satisfactory?",
			"Does the trainee respect the deadlines ?",
			//
			"Technical skills",
			"Does the trainee have the technical skills necessary for his work ?",
			"If not, specify his shortcomings :",
			"Did the trainee need to learn a novel method or a new software ?",
			"If so, does he demonstrate his ability to learn ?",
			"Does he seek to improve his competences in certain fields ?",
			"Is he autonomous ?",
			"Does he seek to help the others ?",
			//
			"Overall assessment",
			"Are you satisfied with the beginning of this training course ?",
			"Which are the strong points of this trainee according to you (one per line) :",
			"Which are the week points of this trainee according to you (one par line) :"
		]
	};
}

$(document).ready(function() {

	//IE8 stuff
	if (typeof Array.prototype.forEach != 'function') {
		Array.prototype.forEach = function(callback) {
			for (var i = 0; i < this.length; i++) {
				callback.apply(this, [this[i], i, this]);
			}
		};
	}

	$(".alert").hide();
	yesno();
	textarea();
	setLang('fr');
	fill();
	$(":radio").iCheck();
});