var internship = {
	Begin: "2015-03-16T01:00:00+01:00",
	Cpy: {
		Name: "air france",
		WWW: "http://www.airfrance.fr/cgi-bin/af/fr/fr/common/home/vols/billet-avion.do?wt.srch=1&wt.mc_id=c_fr_sea_google_brandname_null_null_null&gclid=cnhfuopmkmmcfu6wtaodaxsada"
	},
	End: "2015-09-11T02:00:00+02:00",
	Major: "caspar",
	Promotion: "master imafa",
	Student: {
		Email: "mohamedalibenayed1@gmail.com",
		Firstname: "mohamed ali",
		Lastname: "ben ayed",
		Tel: "0779497651",
	},
	Sup: {
		Email: "naelmazouzi@airfrance.fr",
		Firstname: "nadia",
		Lastname: "el mazouzi",
		Tel: "04 97 28 38 82"
	},
	Title: "(ingénieur développement j2ee et analyse de données",
	Tutor: {
		Email: "cedric.boulbe@unice.fr",
		Firstname: "cédric",
		Lastname: "boulbe",
		Tel: "04 92 07 64 93"
	}
}
var hdrFr = [
	"Ponctualité",
	"Intégration dans l'entreprise",
	"Travail",
	"Compétence techniques",
	"Évaluation globale"
]

var hdrEn = [
	"Punctuality",
	"Integration inside the company",
	"Work",
	"Technical skills",
	"Overall assessment"
]

var questions = {
	fr : [
"Est-il ponctuel",
"Cherche-t-il à communiquer avec les autres",
"D’après vous, est-il déjà bien intégré parmi les membres de votre service",
"Est-il intéressé par son travail",
"S’est-il préoccupé des méthodes de travail de l’entreprise",
"La quantité de travail fournie est elle satisfaisante",
"Respecte-t-il les délais",
"Possède-t-il les compétences techniques nécessaires pour son travail",
"Si non, précisez ses lacunes",
"A-t-il eu besoin d’apprendre une nouvelle technique ou un nouveau logiciel",
"Si oui, a-t-il montré sa capacité à apprendre",
"Cherche-t-il à améliorer ses compétences dans certains domaines",
"Est-il autonome",
"Cherche-t-il à aider les autres",
"Êtes-vous globalement satisfait du début de ce stage",
"Quels sont d’après vous les points forts de ce stagiaire (un par ligne)",
"Quels sont d’après vous les points faibles de ce stagiaire (un par ligne)"
]
,en : [
"Does the trainee seek to communicate with the others",
"Is the trainee well integrated within your team",
"Is the trainee interested by his work",
"Was he concerned with the work methods of the company",
"Is the quantity of work provided satisfactory", 
"Does the trainee respect the deadlines",
"Does the trainee have the technical skills necessary for his work",
"If not, specify his shortcomings",
"Did the trainee need to learn a novel method or a new software",
"If so, does he demonstrate his ability to learn",
"Does he seek to improve his competences in certain fields",
"Is he autonomous",
"Does he seek to help the others",
"Are you satisfied with the beginning of this training course",
"Which are the strong points of this trainee according to you (one per line)",
"Which are the week points of this trainee according to you (one par ligne)"
]
};

function tr(l) {	
	$("#title").html(l == "fr" ? "Rapport intermédiaire" : "Midterm report");
	$("#stu").html(l == "fr" ? "Stagiaire" : "Trainee");
	$("#tut").html(l == "fr" ? "Tuteur" : "Tutor");
	$("#submit").html(l == "fr" ? "Soumettre" : "Submit");

	$("h3").each(function (i, e) {		
		$(e).html(l == "fr" ? hdrFr[i] : hdrEn[i]);		
	});

	$(".q").each(function (i, e) {		
		$(e).html(questions[l][i] + " ?");		
	});

	$(".yes").html(l == "fr" ? "Oui" : "Yes")
	$(".no").html(l == "fr" ? "Non" : "No")
	$(".c").html(l == "fr" ? "Commentaires" : "Comments");
	$(":radio").iCheck({autoHide: false});
}

function checked(v) {
 var b = $("input[name=" + v + "]:checked").val(); 
 if (b == undefined) {
 	$("input[name=" + v + "]").closest("label").css.color = "red";
 	$("#" + v).notify("Required");
 }
 return b; 
}
function submit() {
	//required fields
	var fs = ["q1", "q3","q4","q6","q7","q8","q9","q11","q13","q15","q16","q17","q19"]
	var ok = true;
	fs.forEach(function (e) {
		if (!checked(e)) {
			ok = false;
		}
	});
	return ok;
}

function fn(p) {
	return "<a href='mailto:" + p.Email + "'>" + p.Firstname + " " + p.Lastname + "</a>"
} 
$( document ).ready(function () {
	tr("fr");

	var token = $.urlParam("token")
	console.log(token)
	$("#student").html(fn(internship.Student))
	$("#tutor").html(fn(internship.Tutor))
});