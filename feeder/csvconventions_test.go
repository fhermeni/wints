package feeder

import (
	"io"
	"strings"
	"testing"

	"github.com/fhermeni/wints/schema"
	"github.com/stretchr/testify/assert"
)

var buf = `
"timestamp";"Filiere";"EstEtranger";"EstLabo";"GenreEtu";"PrenomEtu";"NomEtu";"NumSS";"NumEtu";"Adresse1Etu";"Adresse2Etu";"Adresse3Etu";"Adresse4Etu";"CPEtu";"VilleEtu";"PaysEtu";"DomEtu";"EmailEtu";"TelEtu";"DateNaissEtu";"Assurance";"AssuranceNo";"NomEntreprise";"SiteWebEnt";"Adresse1SiegeSocial";"Adresse2SiegeSocial";"Adresse3SiegeSocial";"Adresse4SiegeSocial";"CPSiegeSocial";"VilleSiegeSocial";"PaysSiegeSocial";"AdrSiegeSocial";"GenreDirEnt";"PrenomDirEnt";"NomDirEnt";"EmailDirEnt";"TelDirEnt";"QualiteDirEnt";"DateDebut";"DateFin";"DureeHebdo";"DureeStageJour";"DureeStageSemaine";"HoraireSpecif";"FerieEventuel";"MontantGrat";"Devise";"MontantGratEuros";"ModaliteGrat";"AvantagesGrat";"TitreStage";"DescrStage";"Adresse1Stage";"Adresse2Stage";"Adresse3Stage";"Adresse4Stage";"CPStage";"VilleStage";"PaysStage";"AdrStage";"GenreEncadreur";"PrenomEncadreur";"NomEncadreur";"EmailEncadreur";"TelEncadreur";"FctEncadreur";"GenreEnsResp";"PrenomEnsResp";"NomEnsResp";"EmailEnsResp";"TelEnsResp";"InfoComp";"convsignee";"Parcours";
"2015-03-12 10:45";"SI 5";"oui";"oui";"M.";"Romain";"ALEXANDRE";"192040602706694";"21001838";"235 boulevard Andre Breton";"les jonquilles batiment 9";"";"";"06600";"ANTIBES (France)";"FRANCE";"235 boulevard Andre Breton--les jonquilles batiment 9--06600 ANTIBES (France)";"alexandre.romain06@gmail.com";"+33668138450";"27/04/1992";"Matmut";"980000269722Y";"McGill University";"";"845 Sherbrooke Street West";"";"";"";"H3A 0G4";"MONTREAL";"CANADA";"845 Sherbrooke Street West--H3A 0G4 MONTREAL--CANADA";"M.";"Gregory";"DUDEK";"dudek@cs.mcgill.ca";"+1 (514) 398-7071";"Head of Department";"23/03/2015";"21/09/2015";"35";"127";"26";"9:00 - 17:00 each day, with 1 hour lunch break";"";"1500";"CAD";"1115,5";"Direct deposit";"";"Concern-Driven Software Development with TouchCORE";"TouchCORE is a multi-touch enabled software design modelling tool aimed at developing scalable and reusable software design models following the concern-driven software development paradigm. To prepare TouchCORE for the tool demonstration session of the 18th International Conference on Model-Driven Engineering Languages and Systems (MODELS 2015) in October 2015, several areas need to be improved, including model verification and model checking, model tracing, support for instantiation cardinalities and support for state diagram modelling. Upon arrival, the student will be assigned to one of the areas depending on her/his expertise and preference. In all cases, the student will work on changing the representation of models within the tool (metamodelling), adapting the manipulation of models by the tool (model transformations and weaving), as well as extending the graphical user interface (OpenGL, multi-touch gestures).";"School of Computer Science, 3480 University";"x";"x";"x";"H3A 0G4";"MONTREAL";"Canada";"School of Computer Science, 3480 University--H3A 0G4 MONTREAL--Canada";"M.";"Jörg";"KIENZLE";"Joerg.Kienzle@mcgill.ca";"+1 (514) 398-2049";"Associate Professor, Head of the Software Engineering Laboratory";"Mme";"Mireille";"BLAY-FORNARINO";"blay@unice.fr";"(33)4 92 96 51 61 /(33) 4 97 25 82 15";"";"";"";
"2015-03-18 20:21";"SI 5";"non";"non";"M.";"Martin";"ALFONSI";"1 90 04 67 482 350 34";"20808334";"Résidence THESA";"210 Avenue ROUMANILLE";"";"";"06410";"BIOT";"FRANCE";"Résidence THESA--210 Avenue ROUMANILLE--06410 BIOT";"alfonsi@polytech.unice.fr";"06 33 76 37 29";"16/04/1990";"GMF";"88.705343.65E";"INNO TSD";"http://www.inno-group.com/";"Place Joseph Bermond";"Ophira 1 – BP63";"";"";"06902";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Place Joseph Bermond--Ophira 1 – BP63--06902 SOPHIA ANTIPOLIS CEDEX";"M.";"Marc ";"PATTINSON";"m.pattinson@inno-group.com";"04 92 38 84 10";"Gérant ";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"600";"Euros";"600";"";"";"Développement d’applications web";"Inno labs est la cellule TIC de Inno Group. Dans le cadre de ses activités techniques, un stagiaire ingénieur en développement logiciel est recherché. Il aura pour principales activités :
    • D’assister le chef de projet dans la gestion quotidienne des projets ;
    • Participer à toutes les phases des projets (conception, réalisation, exploitation, support) ;
    • Participer aux études de conseil (analyse de besoin client, spécifications, benchmarks technologiques) ;
    • Participer à la rédaction des propositions (réponses à appel d’offres / appels à projets).
";"Place Joseph Bermond";"Ophira 1 – BP63";"";"";"06902";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Place Joseph Bermond--Ophira 1 – BP63--06902 SOPHIA ANTIPOLIS CEDEX";"M.";"Fabrice ";"CLARI";"f.clari@inno-group.com";"+33 4 92 38 84 19";"Senior ICT consultant";"M.";"Christian";"BREL";"brel@polytech.unice.fr";"06 52 48 53 67";"";"";"";
"2015-02-19 17:41";"SI 5";"non";"non";"M.";"Quentin";"BITSCHENÉ";"191058305010171";"20901433";"391 chemin des bruyères";"";"";"";"83550";"Vidauban";"FRANCE";"391 chemin des bruyères--83550 Vidauban";"bitschene.quentin@yahoo.fr";"+33672131657";"31/05/1991";"MAIF";"1300020R";"Atos";"atos.net";"80 quai Voltaire";"";"";"";"95877";"Bezons";"FRANCE";"80 quai Voltaire--95877 Bezons";"M.";"Cédric";"Couget";"cedric.couget@atos.net";"+33(0)786284830";"Responsable d’agence";"30/03/2015";"30/09/2015";"35";"129";"26";"";"";"1000";"Euros";"1000";"";"";"Immersion opérationnelle dans un projet ATOS pour un client d’envergure internationale. ";"Immersion opérationnelle dans un projet ATOS pour un client d’envergure internationale. Vous intégrez notre division Technology Services d‘Atos Intégration, Région Midi-Pyrénées, site de Toulouse St Martin, sous la responsabilité de Monsieur Cédric Couget, Responsable d’agence.
 
Objectifs pédagogiques :
 
Sous le tutorat de Monsieur Jean-Bernard Pugnet, Directeur de projet, vous intégrez l’équipe de projet PARAMETRAGE ORANGE dans les locaux Atos à Toulouse (6, impasse Alice Guy) sur des activités de conception et développement d’un module (ORPEO) d’une application WEB.
 
Vous ferez l‘apprentissage des méthodes de travail opérationnelles dans un environnement structuré. En fin de stage, vous aurez développé les compétences techniques requises dans le cadre d‘un projet de conception/développement en informatique de gestion.
 
Objectifs du stage :
 
-       Découverte des projets, du client, du contexte, du sujet, de l’organisation et des process.
-       Prise de connaissance des projets PARAMETRAGE, COCPIT et du module ORPEO.
-       Montée en compétence sur les technologies de conception et développement utilisées :
o    Développement : PHP5, HTML, CSS, JAVASCRIPT
o    Base de données : MYSQL, PostgreSQL
o    Méthodes : UML, AGILE (SCRUM)
-       L’objectif du stage est la refonte du module ORPEO :
o    conception d’un nouveau modèle de base de données relationnelle,
o    migration des données,
o    adaptation de l‘application à ce nouveau référentiel
o    développement de nouvelles fonctionnalités (dans la mesure du possible).
-       Vous serez amenée à réaliser à l’ensemble des activités : étude, conception, maquettage, développement, tests, intégration en fonction du niveau de votre niveau de montée en compétence.
 
Objectifs pédagogiques :
 
-       Intégration dans une équipe de projet,
-       Mise en pratique vos compétences techniques,
-       Assimilation de l‘organisation, le fonctionnement et le travail sur un projet en ESN,
-       Travail en équipe et esprit de service.";"6, impasse Alice Guy";"";"";"";"31300";"Toulouse";"FRANCE";"6, impasse Alice Guy--31300 Toulouse";"M.";"Jean-Bernard";"Pugnet";"jean-bernard.pugnet@atos.net";"+33 6 84 68 73 88";"Directeur de projet";"M.";"Philippe";"Renevier";"Philippe.RENEVIER@unice.fr";"+33 4 9296 5167";"Vous bénéficierez du RIE sur place pour prendre vos repas et d’un remboursement des transports en commun à hauteur de 50% sur justificatif.";"";"";
"2015-03-06 10:25";"SI 5";"non";"non";"M.";"Pierre";"BOUILLET";"191067425622144";"21210034";"48 chemin du Battieu";"";"";"";"74190";"PASSY";"FRANCE";"48 chemin du Battieu--74190 PASSY";"pierre_bouillet@hotmail.fr";"0659008593";"06/06/1991";"mae";"C003345329";"SOPRA BANKING";"";"3 RUE DU PRE FAUCON";"PAE DES GLAISINS";"";"";"74940";"ANNECY LE VIEUX";"FRANCE";"3 RUE DU PRE FAUCON--PAE DES GLAISINS--74940 ANNECY LE VIEUX";"Mme";"Hélène";"RIPPERT";"helene.rippert@soprabanking.com";"04.50.33.31.49";"Directeur d’agence";"30/03/2015";"30/09/2015";"35";"129";"26";"";"";"1097,74";"Euros";"1097,74";"";"Tickets restaurant, et éventuellement, indemnisation 50% abonnement transport en commun sur justificatif en note de frais.";"Développement d’un progiciel de Reporting Bancaire";"Intégré à une équipe de développement, dans le cadre du projet de refonte de l’outil de développement de modèles de reporting, au sein de l’entité Sopra Banking Compliance R&D.

L’outil est architecturé en couches :

-          Une IHM low cost base sur Ms Excel, remplaçable à terme par une future IHM
-          Une couche Métier, service & batch en java JEE
-          Une couche data en XML/SGBD

Il s’agit donc de développement dans un environnement java, JEE, XML, Ms Excel, Eclipse, maven, JUnit, CI (Jenkins/Sonar).

Vous aurez donc d’abord à appréhender l’architecture technique et logique du produit, et contribuer à sa mise en œuvre/amélioration. ";"3 RUE DU PRE FAUCON";"PAE DES GLAISINS";"";"";"74940";"ANNECY LE VIEUX";"FRANCE";"3 RUE DU PRE FAUCON--PAE DES GLAISINS--74940 ANNECY LE VIEUX";"Mme";"Stéphane";"POLICET";" stephane.policet@soprabanking.com";"04.50.33.32.61 ";"Ingénieur d’études";"Mme";"Audrey ";"OCCELLO";"occello@polytech.unice.fr";"04 92 96 51 02 ";"";"";"";
"2015-03-10 11:00";"SI 5";"oui";"oui";"Mlle";"Cécile";"CAMILLIERI";"292086822461427";"21004323";"4 rue de Roppe";"";"";"";"68200";"MULHOUSE";"FRANCE";"4 rue de Roppe--68200 MULHOUSE";"cecile.camillieri@gmail.com";"+33680900327";"13/08/1992";"GMF";"88.341406.65K";"McGill University";"";"845 Sherbrooke Street West";"";"";"";"H3A 0G4";"MONTREAL";"CANADA";"845 Sherbrooke Street West--H3A 0G4 MONTREAL--CANADA";"M.";"Gregory";"DUDEK";"dudek@cs.mcgill.ca";"+1 (514) 398-7071";"Head of Department";"23/03/2015";"21/09/2015";"35";"127";"26";"9:00 - 17:00 each day, with 1 hour lunch break";"";"1500";"CAD";"1099,197";"Direct deposit";"";"Concern-Driven Software Development with TouchCORE";"TouchCORE is a multi-touch enabled software design modelling tool aimed at developing scalable and reusable software design models following the concern-driven software development paradigm. To prepare TouchCORE for the tool demonstration session of the 18th International Conference on Model-Driven Engineering Languages and Systems (MODELS 2015) in October 2015, several areas need to be improved, including model verification and model checking, model tracing, support for instantiation cardinalities and support for state diagram modelling. Upon arrival, the student will be assigned to one of the areas depending on her/his expertise and preference. 
In all cases, the student will work on changing the representation of models within the tool (metamodelling), adapting the manipulation of models by the tool (model transformations and weaving), as well as extending the graphical user interface (OpenGL, multi-touch gestures).";"School of Computer Science";"3480 University";"";"";" H3A 0E9";"MONTREAL";"CANADA";"School of Computer Science--3480 University--H3A 0E9 MONTREAL--CANADA";"M.";"Jörg";"Kienzle  KIENZLE";"Joerg.Kienzle@mcgill.ca";"+1 (514) 398-2049";"Associate Professor, Head of the Software Engineering Laboratory";"Mme";"Mireille";"BLAY-FORNARINO";"blay@unice.fr";"(+33)4 92 96 51 61";"";"";"";
"2014-12-16 22:00";"SI 5";"oui";"non";"M.";"Adrien";"CASANOVA";"192050608870764";"21000600";"4 rue barla";"";"";"";"06300";"NICE";"FRANCE";"4 rue barla--06300 NICE";"acasanov@polytech.unice.fr";"0661755892";"27/05/1992";"MAPA";"2225574B";"Garagesocial, Inc.";"https://www.garagesocial.com/#!home";"20 Park Plaza 4th Floor";"";"";"";"02116";"BOSTON, MA";"ETATS-UNIS";"20 Park Plaza 4th Floor--02116 BOSTON, MA--ETATS-UNIS";"M.";"Maxime";"RASSI";"maxime@garagesocial.com";"617-948-2530";"President";"23/03/2015";"18/09/2015";"40";"124";"26";"";"";"2700";"Dollars";"2 370";"";"";"Front End Development Internship";"Development of user interfaces working with Javascript MVCs frameworks and mockup HTML templates.

- Company: Garagesocial, Inc. is a new online community and marketplace for the automotive industry. Garagesocial, Inc. develops tools that allow users and companies to showcase vehicles, parts and services online and allow them to engage in networking and commercial activities.

- The technologies may or will include:
 Languages: HTML, CSS, Javascript, PHP, Python, Ruby, Objective C, SQL
 Processors: Compass, Coffeescript, Less
 Frameworks: Laravel, ROR, EmberJS, BackboneJS, MarionetteJS
 Infrastructure: Amazon Web Services (S3, EC2, Route53, RDS, ElasticCache, ElasticBeanstalk), ElasticSearch, Hadoop, Memcache, Solr

- The type of work involved may or will include:
 Prototyping of new features including UX Mockups & Graphic Creation
 Developing, Testing and Deploying New Features
 Cross browser testing
 Performance Monitoring and Optimization
 Mobile Development
 Data Store Schema design and optimization
 API Development";"20 Park Plaza 4th Floor";"";"";"";"02116";"BOSTON, MA";"Etats-Unis";"20 Park Plaza 4th Floor--02116 BOSTON, MA--Etats-Unis";"M.";"Maxime";"RASSI";"maxime@garagesocial.com";"617-948-2530";"President";"M.";"Anne-Marie";"PINNA-DERY";"pinna@polytech.unice.fr";"(+33) 4 92 96 51 62";"";"";"";
"2014-12-08 10:24";"SI 5";"non";"non";"M.";"Guy";"CHAMPOLLION";"190047511317802";"20808172";"8 rue Saint-Antoine";"";"";"";"06600";"ANTIBES";"FRANCE";"8 rue Saint-Antoine--06600 ANTIBES";"champoll@polytech.unice.fr";"0625899748";"15/04/1990";"Macif";"13057068";"Reador.NET";"reador.net";"Business Pôle, Bat.B - Entrée A - 2ème étage";"1047 Route des Dolines";"";"";"06560";"VALBONNE";"FRANCE";"Business Pôle, Bat.B - Entrée A - 2ème étage--1047 Route des Dolines--06560 VALBONNE";"M.";"Christophe";"DESCLAUX";"christophe@reador.net";"0762596417";"Fondateur de la startup Reador.NET";"01/12/2014";"16/06/2015";"35";"138";"29";"";"";"436,05";"Euros";"436,05";"Chèque ou virement bancaire";"";"Stage développement client web";"L’objectif de ce stage est de proposer une optimisation de l’IHM existante en l’adaptant aux lecteurs et/ou rédacteurs d’informations. L’IHM refondue sera utilisée comme client officiel de l‘application web Reador. C’est une étape critique du projet car elle permettra de fournir une interface graphique en adéquation avec les attentes des utilisateurs (lecteurs, rédacteurs d’informations…) du service. Elle permettra de valoriser le travail d’annotation des news effectué en amont. Vous devrez donc proposer et concevoir au cours de votre stage des améliorations de l’IHM actuellement mise en place à l’aide du framework RubyOnRails et de JQuery. Vous aurez une grande liberté dans le choix des fonctionnalités à implémenter et ferez preuve d’initiative.";"Business Pôle, Bat.B - Entrée A - 2ème étage";"1047 Route des Dolines";"";"";"06560";"VALBONNE";"FRANCE";"Business Pôle, Bat.B - Entrée A - 2ème étage--1047 Route des Dolines--06560 VALBONNE";"M.";"Christophe";"Desclaux";"christophe@reador.net";"0762596417";"Fondateur de la startup Reador.NET";"Mme";"Anne-Marie";"Dery-Pinna";"pinna@polytech.unice.fr";"0661029387";"";"";"";
"2015-01-14 13:40";"SI 5";"non";"non";"M.";"Zhang";"CHEN";"190109921600030";"21210445";"18 avenue docteur fabre";"";"";"";"06160";"JUAN LES PINS";"FRANCE";"18 avenue docteur fabre--06160 JUAN LES PINS";"zchen@polytech.unice.fr";"0750358934";"01/10/1990";"Société Courtage d‘Assurance";"07001878";"ATOS";"";"80 quai Voltaire";"";"";"";"95877";"BEZONS CEDEX";"FRANCE";"80 quai Voltaire--95877 BEZONS CEDEX";"Mme";"Guiselene";"CIEUTAT";"guiselene.cieutat@atos.net";"04 97 15 79 11";"Rersponsable des Ressources Humaines";"09/03/2015";"09/09/2015";"35";"129";"26";"";"";"1000";"Euros";"1000";"";"";"Stage d’analyse d’impact à partir de bases de données NoSQL -orientées graphes";"Pour des besoins d’intégration dans un système de releasing et de tests, nous vous confierons l’étude, la conception et le développement d‘un outil d‘Analyse d‘impact et de Root Cause.
Cet outil sera utilisé à des fins de gestion, d’aide à la décision, de traçabilité et de reporting,
 Pour mener ce projet une étude en deux phases est nécessaire :
?Etablir l’état de l’art dans les domaines d’Analyse d’impact et de recherche de Root Cause.
?Concevoir et développer un outil d‘Analyse d‘impact et de Root Cause à base de Base de Données Orientée Graphes.";"Le Millenium 150 allée Pierre Ziller";"";"";"";"06560";"SOPHIA ANTIPOLIS";"FRANCE";"Le Millenium 150 allée Pierre Ziller--06560 SOPHIA ANTIPOLIS";"M.";"Salim";"AINOUCHE";"salim.ainouche@atos.net";"06 20 43 55 86";"Senior Software Designer";"M.";"Lionel";"FILLATRE";"lionel.fillatre@i3s.unice.fr";"0492942785";"";"";"";
"2015-04-30 12:32";"SI 5";"non";"non";"M.";"Julien";"CHIARAMELLO";"1 94 02 06 088 271 63";"21209527";"4 rue de Dijon";"";"";"";"06000";"NICE";"FRANCE";"4 rue de Dijon--06000 NICE";"chiarame@polytech.unice.fr";"0659455961";"04/02/1994";"ACM-IARD SA";"4228532";"QuantifiCare SA";"www.quantificare.com";"1180 route des Dolines";"Bâtiment Athena B";"BP 40051";"";"06901";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"1180 route des Dolines--Bâtiment Athena B--BP 40051--06901 SOPHIA ANTIPOLIS CEDEX";"Mme";"Pascale";"BUISSON";"pbuisson@quantificare.com";"04 92 91 54 40";"Directrice";"11/05/2015";"30/09/2015";"35";"98";"21";"";"";"700";"Euros";"700";"Virement";"Accès tickets restau de 8€, prise en charge société de 4.60€";"Administrateur Systèmes et Réseaux";"QuantifiCare est une entreprise dynamique en croissance régulière qui est parmi les
leader mondiaux dans le développement d‘applications et de services visant le marché
de la recherche médicale. 

Nous travaillons en relation avec des clients dans le monde entier et nous possédons
des bureaux à Sophia Antipolis et à San Mateo (US). 

Une partie de nos activités nous oblige à mettre en place toutes les solutions de
sécurité à notre disposition pour récupérer des données sensibles envoyées par nos
clients.

Nous recherchons actuellement un stagiaire qui, sous la responsabilité du
responsable informatique et réseaux, participera à : 

- L‘installation, l‘administration et la maintenance des machines clientes de
l‘entreprise (majorité de client Windows 7/8/8.1, quelques Linux)

- La gestion des utilisateurs, en particulier la création de leurs comptes, de leurs
droits d’accès sécurisés ainsi que de leurs certificats électroniques personnels
(SSL) et/ou clefs SSH

La maintenance et l‘administration des services aux utilisateurs avec une attention
- particulière portée sur les outils permettant leurs accès aux ressources de
l’entreprise de manière sécurisée (OpenVPN, SSL/TLS, SSH)

- La surveillance des sondes de sécurité des serveurs et l’analyse des éventuelles
alertes

- La maintenance des serveurs pour assurer leur intégrité opérationnelle (tests
opérationnels, analyse et application des mises à jour de sécurité)

- La maintenance et la mise à jour des actifs réseaux, en particulier les solutions de routage et de sécurité installées entre les différents réseaux

- La gestion des achats du matériel informatique 

- Éventuellement des interventions de maintenance à distance sur des machines chez nos
clients

En dehors de ces tâches de fond, le stagiaire participera également, en relation
avec divers départements de l‘entreprise, à l‘étude, l‘amélioration et l‘évolution
de l’infrastructure « serveur » de l‘entreprise fournissant les services sécurisés
destinés à nos employés et à nos clients (serveurs exclusivement sous systèmes
OpenSource Linux/FreeBSD). 

Force de proposition, il participera de manière active à l‘évolution de nos
différentes plate-formes, en particulier en apportant son expertise en matière de
sécurisation.

Patience, écoute, ouverture d‘esprit, réactivité « réfléchie » sont des qualités
indéniables pour cette mission.

De bonnes notions de « hardware » seraient un plus.

QuantifiCare étant une société tournée vers l‘international, nous recherchons une
personne possédant un bon niveau d‘anglais capable de lire / rédiger des procédures
techniques dans un anglais correct mais aussi capable d‘éventuellement communiquer
oralement avec nos employés américains ainsi qu‘avec notre clientèle.";"1180 route des Dolines";"Bâtiment Athena B";"BP 40051";"";"06901";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"1180 route des Dolines--Bâtiment Athena B--BP 40051--06901 SOPHIA ANTIPOLIS CEDEX";"M.";"Matthieu";"MEURILLON";"mmeurillon@quantificare.com";"04 92 91 54 41";"Responsable informatique & réseaux";"Mme";"Tamara";"REZK";"tamara.rezk@inria.fr";"04 97 15 53 37";"";"";"";
"2015-02-06 13:58";"SI 5";"oui";"oui";"Mlle";"Geneviève";"CIRERA";"292023155564686";"21206834";"1 rue la fontaine";"";"";"";"31450";"Donneville";"FRANCE";"1 rue la fontaine--31450 Donneville";"genevieve.cirera@gmail.com";"0647968604";"19/02/1992";"AXA";"3792084404";"NII";"";"2-1-2 Hitotsubashi Chiyoda-ku";"";"";"";"101-8430";"Tokyo";"JAPAN";"2-1-2 Hitotsubashi Chiyoda-ku--101-8430 Tokyo--JAPAN";"M.";"Masaru";"KITSUREGAWA";"nii-internship@nii.ac.jp";"03-4212-2000";"Director General";"15/03/2015";"30/08/2015";"35";"118";"24";"";"";"171000";"Yen";"1273,28";"";"";"AI system that solve physics problems of entrance exam for university";"General intelligent robots in daily life environment must observe and model external world, understand users’ instructions and intentions. In such situations, robots always face to ambiguity of information. Conventional approaches to solve the ambiguity were using common sense database/ontology, or asking users; however developers’/users’ cost is too huge.
You will engage on the software development for connection between 
natural language processing and physics simulator.
The languages used are perl, Modelica, C++.";"2-1-2 Hitotsubashi Chiyoda-ku";"";"";"";"101-8430";"Tokyo";"JAPAN";"2-1-2 Hitotsubashi Chiyoda-ku--101-8430 Tokyo--JAPAN";"M.";"Tetsunari";"Inamura";"inamura@nii.ac.jp";"03-4212-1605";"Associate Prof";"M.";"Frédéric";"Precioso";"frederic.precioso@polytech.unice.fr";"+33 (0)4 92 96 51 43";"";"";"";
"2015-03-24 10:07";"SI 5";"oui";"oui";"M.";"Hélène";"COLLAVIZZA";"111112";"11112";"tttt";"";"";"";"06130";"mouans";"FRANCE";"tttt--06130 mouans";"helene.collavizza@gmail.com";"0606060606";"03/03/2015";"MAIF";"1234";"Essai Hélène";"";"fdgdfg";"";"";"";"06370";"Mouans";"FRANCE";"fdgdfg--06370 Mouans";"M.";"ffff";"ffff";"tt@tt.fr";"060606";"sfsfsf";"03/06/2015";"30/07/2015";"20";"40";"9";"";"";"100";"Euros";"100";"";"";"le beau stage";"dsdgsdgdgsdgsdgg";"sdsdgd";"";"";"";"02136";"kolp";"FRANCE";"sdsdgd--02136 kolp";"M.";"ddfdd";"dfdfdf";"ddd@tt.fr";"03030";"patron";"M.";"tttt";"pppp";"jj@ff.fr";"02 03 03 03";"";"";"";
"2015-02-26 11:32";"SI 5";"non";"non";"M.";"Clément";"CRISTIN";"192054209506311";"21210627";"140 rue Albert Einstein appt C05";"";"";"";"06560";"Valbonne";"FRANCE";"140 rue Albert Einstein appt C05--06560 Valbonne";"cristin.clement@gmail.com";"+336 72 34 33 18";"05/05/1992";"MAIF";"1367936P";"Hewlett-Packard";"";"Z.I. de Courtaboeuf 1, avenue du Canada";"";"";"";"91947";"Les Ulis";"FRANCE";"Z.I. de Courtaboeuf 1, avenue du Canada--91947 Les Ulis";"M.";"Virginie";"MARESCHAL";"Virginie.mareschal@hp.com";"+334 80 32 14 87";"Chargée d’administration RH";"16/03/2015";"15/09/2015";"35";"128";"26";"";"";"1220";"Euros";"1220";"";"";"HP R&D – Participer au développement de notre GUI Framework Telecom HP CMS ‘Unified OSS Console’";"Vous rejoignez une équipe R&D dynamique, compétente et expérimentée dans le domaine de la
Télécommunication et du développement logiciel. 
Cette équipe basée à Sophia Antipolis est en charge d’un Web Framework UI nouvelle génération destiné à intégrer les différentes applications du Portfolio HP (Fault Management, Service Quality Management, Service Level Agreement Management, Events Correlation, Customer Experience, OSS Analytics...) et potentiellement les applications de nos clients dans un environnement graphique unifié et sécurisé.

Ce Framework UI est totalement ouvert à travers un SDK et est implémenté sur des technologies web modernes (JavaScript MVVM à base de NodeJS / AngularJS) qui permettent un haut niveau d’intégration, de customisations et de sécurité pour les différentes applications télécoms. 
Il propose une expérience utilisateur riche et efficace sur différents hardware (desktop, tablette...) en intégrant les concepts de web design « responsive ». 
Un outil ‘designer’ WYSIWYG permettra aux opérateurs de construire graphiquement des tableaux de bord et des écrans opérationnels à partir des briques prédéfinies disponibles.

Le stagiaire sera immergé dans une équipe R&D et participera au développement du Framework
commun et à l’enrichissement de ses différents modules métiers en rajoutant de nouvelles
fonctionnalités en relation avec les responsables produits, les architectes et les équipes de R&D.

Il devra notamment:

- Enrichir le GUI Framework existant avec de nouvelles fonctionnalités et des services communs.
- Spécifier, désigner et implémenter des modules fonctionnels en relation avec des besoins clients Telecom.
- Fournir aux intégrateurs et aux clients d’HP une infrastructure de développement et de
déploiement de leurs propres modules, mise en plages et leurs widgets graphiques. (SDK, API, RestAPI, générateurs Yeoman...).
- Participer au développement du designer graphique WYSIWYG et à son intégration dans le
produit.
- Enrichir notre bibliothèque de widgets graphiques. Certains widgets sont basées sur Highcharts et HighMaps (www.highcharts.com ).
- Concevoir et proposer des écrans HTML responsives pour une expérience utilisateur efficace.
- Participer aux activités liées à la mise en production en respectant les règles de qualité de l’équipe R&D.
- Participer à la mise en place de tests unitaires et fonctionnels
- Participer à la documentation des produits et des différentes APIs (REST, JavaScript...)
- Assister à la gestion de projet Agile basé sur SCRUM
- Fixer les problèmes remontés par nos utilisateurs.
- Participer aux présentations techniques et aux démonstrations clientes";"Marco Polo – Bâtiment B – Entrée B1 ZAC du Font de l’Orme 1 BP1220 790 Avenue du Docteur Donat";"";"";"";"06254";"Mougins";"FRANCE";"Marco Polo – Bâtiment B – Entrée B1 ZAC du Font de l’Orme 1 BP1220 790 Avenue du Docteur Donat--06254 Mougins";"M.";"Jean-Charles";"Picard";"jean-charles.picard@hp.com";"+33 4 22390127";"Responsable Developpement Logiciel";"M.";"Sébastien ";"Mosser";"mosser@i3s.unice.fr";"+334 92 96 50 58";"";"";"";
"2014-12-08 10:38";"SI 5";"non";"non";"M.";"Anthony";"DA MOTA";"191078313740995";"20901891";"30 Lotissement le Bois de la Combe";"";"";"";"83720";"Trans en Provence";"FRANCE";"30 Lotissement le Bois de la Combe--83720 Trans en Provence";"anthony.damota06@gmail.com";"0686072630";"24/10/1991";"LMDE";"3283564704";"Atos Toulouse";"";"6 Impasse Alice Guy";"";"";"";"31024  ";"TOULOUSE";"FRANCE";"6 Impasse Alice Guy--31024 TOULOUSE";"M.";"Julien";"BIREBENT";"julien.birebent@atos.net";"+33 5 34 36 32 23";"Delivery Manager – BPS / Telco IS";"30/03/2015";"30/09/2015";"35";"129";"26";"";"";"1062,20";"Euros";"1062,20";"";"";"Stage de développement Firefox OS";"Mission :
Il s’agira de réaliser un prototype d’application mobile pour FirefoxOS servant dans un second temps de support d’étude aux fonctionnalités de ce type d’application. Le stagiaire mettra notamment en relief ses résultats avec ce qu’il est possible de faire sur le système d’exploitation Android.
Cette étude pourra porter sur les fonctionnalités disponibles mais également les caractéristiques telles que la sécurité, la performance, l’outillage disponible dans cet environnement ou encore la fiabilité des applications développées sur FirefoxOS.
 
La réalisation du prototype ainsi que l’étude comparative sera réalisée en utilisant la méthodologie Scrum. Le périmètre du prototype et de l’étude sera donc affiné tout au long du stage.
 
Le stagiaire sera intégré au sein d’une équipe projet de l’agence toulousaine participant au développement d’applications similaires en utilisant la méthode Scrum également. ";"6 Impasse Alice Guy";"";"";"";"31024 ";"TOULOUSE";"FRANCE";"6 Impasse Alice Guy--31024 TOULOUSE";"M.";"Julien";"BIREBENT";"julien.birebent@atos.net";"+33 5 34 36 32 23";"Delivery Manager – BPS / Telco IS - Toulouse";"M.";"Michel";"BUFFA";"Michel.Buffa@unice.fr";"(33)-92-07-66-60";"";"";"";
"2015-03-26 07:44";"SI 5";"oui";"non";"Mlle";"Huinan";"DONG";"289099921600013";"dh210451";"Cite U jean medecin 25rue robert latouche";"";"";"";"06200";"Nice";"FRANCE";"Cite U jean medecin 25rue robert latouche--06200 Nice";"dlutdong@gmail.com";"+33605573389";"28/09/1989";"BNP";"177900968";"Price waterhouseCoopers";"http://www.pwccn.com/home/eng/index.html";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District";"";"";"";"100020";"Beijing ";"CHINE";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District--100020 Beijing--CHINE";"M.";"Depei";"SHEN";"offeree.support@cn.pwc.com";" +86106533 8888";"Directeur ";"23/03/2015";"23/09/2015";"40";"129";"26";"";"";"3000";"Yuan";"440";"";"";"Asset Liability Management";"1. ALM(Asset Liability Management) advisory.In this project PwC do the advisory work for the interest rate risk of banking book and liquidity risk management for the client, including gap analysis, optimizing the management structure, risk measurement for interest rate risk of banking book and liquidity risk management, and the related reporting roles. 

2. ALM system validation.PwC help the client propose system implementation roadmap, design ALM database and ALM models, set up ALM data standards and conduct the ALM system validation.

The role of the intern:
1.Using SQL to prepare the data for system validation;
2.Design the system validation program;
3.Verify report developed by sub-contractor;
4.Data requirement analysis, data gap analysis;
?";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District";"";"";"";"100020";"Beijing";"CHINE";"26/F., Office Tower A, Beijing Fortune Plaza,  7 Dongsanhuan Zhong Road, Chaoyang District--100020 Beijing--CHINE";"M.";"Yong";"LU";"pwc_recruit@foxmail.com";"+8613811630156";"directeur du projet";"M.";"Ioan";"BOND";"bond@polytech.unice.fr ";"+33678110112";"";"";"";
"2014-12-09 15:03";"SI 5";"non";"non";"M.";"Clément";"DUFFAU";"192102403710940";"21002723";"13 rue des Petits Ponts";"";"";"";"06250";"MOUGINS LE HAUT";"FRANCE";"13 rue des Petits Ponts--06250 MOUGINS LE HAUT";"duffau@polytech.unice.fr";"0633796793";"02/10/1992";"PACIFICA";"2649317906";"Axonic";"http://www.axonic.fr/";"2720 Chemin Saint-Bernard ";"";"";"";"06224";"VALLAURIS";"FRANCE";"2720 Chemin Saint-Bernard--06224 VALLAURIS";"M.";"Marc";"ROUGE";"mrouge@axonic.fr";"0497213040";"PDG";"16/03/2015";"13/09/2015";"35";"126";"26";"";"";"1200";"Euros";"1200";"";"";"Stage R&D en développement logiciel";"Axonic, filiale du groupe MXM, est une startup basée à Sophia-Antipolis, 
spécialisée dans le développement d‘appareils médicaux actifs implantables 
dédiés à la neuro-stimulation ayant pour objectif d‘améliorer les 
conditions de vie de patients atteints de maladies chroniques sévères ou 
dégénératives. 

Axonic recherche un stagiaire (Ingénieur / Master 2) en développement logiciel pour intégrer 
son équipe et travailler sur la conception et le développement du 
""Framework Logiciel Axonic"". 

Ce Framework permet de gérer, de piloter, et de contrôler les produits 
Axonic, et notamment les paramètres de la stimulation. Il s‘adresse à des 
utilisateurs qui ont pour métier la recherche clinique, la médecine, la 
chirurgie mais aussi à terme directement aux patients. 

L‘équipe de développement logiciel suit une méthode de développement Agile 
basée sur Scrum et l‘intégration continue. 

Dans un contexte de recherche et développement et d‘innovation, le stagiaire 
sera en interface permanente avec tous les métiers de notre entreprise 
(micro-électronique, électronique, mécanique, logiciel, clinique) et il 
découvrira les spécificités et les exigences liées au domaine médical 
(respect de normes, sécurité, couverture de tests, risques patients, 
éthique). 

Dans le cadre de ses attributions le stagiaire prendra en charge le 
développement complet des ""User Stories"" dont les priorités sont fixées par 
nos ""Product Owners"" (équipe de recherche clinique). 

Il pourra intervenir sur l‘ensemble de l‘application, les couches basses 
(ex: drivers permettant la communication avec les stimulateurs), le code 
métier (support et gestion de la stimulation), et/ou sur le code de 
l‘interface homme-machine (IHM) en fonction des priorités et de ses 
attentes. 

Sur chacun des domaines, le stagiaire participera à la définition de 
l’architecture, à la conception, à l’implémentation, à ses tests, et à la 
rédaction de la documentation. 

Il pourra également être amené à étudier, proposer et mettre en place des 
solutions d‘outillage permettant d‘améliorer nos pratiques de 
développement. ";"2720 Chemin Saint-Bernard ";"";"";"";"06224";"VALLAURIS";"FRANCE";"2720 Chemin Saint-Bernard--06224 VALLAURIS";"M.";"Pierrick";"Perret";"pperret@axonic.fr";"0497213040";"Head of Project Management Office";"M.";"Sebastien";"Mosser";"mosser@i3s.unice.fr";"0492965058";"J‘ai un doute sur le tuteur enseignant entre le responsable spécialité AL et le responsable stage SI5";"";"";
"2015-01-08 18:05";"SI 5";"non";"non";"M.";"Thibaut";"DUFOUR";"193047815823714";"21002820";"37 Avenue Balzac";"";"";"";"92410";"VILLE d‘AVRAY";"FRANCE";"37 Avenue Balzac--92410 VILLE d‘AVRAY";"dufour@polytech.unice.fr";"0677910408";"09/04/1993";"Macif";"5176896";"Rivage Investment SAS";"www.rivageinvestment.com";"26 Rue du Quatre Septembre";"";"";"";"75002";"PARIS";"FRANCE";"26 Rue du Quatre Septembre--75002 PARIS";"M.";"Thierry";"SENG";"thierry.seng@rivageinvestment.com";"01 70 91 25 94";"Directeur Systèmes Informatiques";"16/03/2015";"16/09/2015";"35";"129";"26";"";"";"1200";"Euros";"1200";"Virement bancaire";"Remboursement de frais professionnels sur présentation des factures acquittées";"Ingénieur Développeur";"Le collaborateur fera partie intégrante de l’équipe de développement de la plateforme propriétaire de gestion
et à ce titre aura les responsabilités suivantes :
- Responsabilités principales :
o développer et maintenir la plateforme de gestion et ses outils (C++)
o développer et maintenir la base de données et ses outils (SQL Server)
o garantir la qualité, la robustesse et l’adaptabilité de la plateforme de gestion
o être force de proposition dans l’amélioration des briques logicielles existantes
- Autres missions :
o soutenir la gestion courante des fonds (assistant de gestion front office : Excel)
o maintenir le site web de la société (mises à jour, corrections, améliorations)
o assurer une veille technologique sur les architectures concurrentes 
";"26 Rue du Quatre Septembre";"";"";"";"75002";"PARIS";"FRANCE";"26 Rue du Quatre Septembre--75002 PARIS";"M.";"Thierry";"SENG";"thierry.seng@rivageinvestment.com";"01 70 91 25 94";"Directeur Systèmes Informatiques";"Mme";"Anne-Marie";"HUGUES";"hugues@unice.fr";"06 84 04 59 30";"";"";"";
"2015-03-24 21:47";"SI 5";"non";"non";"M.";"Fabien";"FOERSTER";"192068311802118";"21003654";"22 place des arcades";"";"";"";"06250";"MOUGINS";"FRANCE";"22 place des arcades--06250 MOUGINS";"fabienfoerster@gmail.com";"+33664513101";"05/06/1992";"Crédit Agricole";"931932908";"ATOS";"http://atos.net/";"River Ouest";"80 quai Voltaire";"";"";"95877";"BEZONS CEDEX";"FRANCE";"River Ouest--80 quai Voltaire--95877 BEZONS CEDEX";"Mme";"Guiselenne";"CIEUTAT";"anne.aime@atos.net";"04 97 15 79 11";" Responsable des Ressources Humaines";"07/04/2015";"20/09/2015";"35";"116";"24";"";"";"1000";"Euros";"1000";"virement bancaire";"";"Cartographie et traçabilité";"Le stage a pour objectif de définir une méthode de calcul d’impact dans une cartographie multi-vue. 
En détail l’étudiant doit :
étudier les offres d’outils de modélisation multi paradigmes (UML, BPMN, MCD, autres)
implémenter, dans un outil choisit suite à l’étude, une cartographie multi-vue
définir une méthode (voir une automatisation) du calcul d’impact dans un contexte de cartographie multi-vue
(si on a le temps) attacher à la méthode une pondération en points de fonction des périmètres d’impact détectés";"Le Millénium";"150, allée Pierre Ziller";"B.P. 279";"";"06905";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Le Millénium--150, allée Pierre Ziller--B.P. 279--06905 SOPHIA ANTIPOLIS CEDEX";"Mme";"Clémentine";"NEMO";"clementine.nemo@atos.net";"04 93 95 46 44";"Consultante Solution";"Mme";"Mireille";"BLAY-FORNARINO";"blay@i3s.unice.fr";"O4 92 96 51 61";"Alors premièrement la position de mon encadrant au sein de l‘entreprise peut etre sujette à modification suivant la réponse de mon encadrant.
Deuxièmement mon sujet de stage peut être un peu plus étoffé au besoin.";"";"";
"2015-02-25 22:08";"SI 5";"oui";"non";"Mlle";"Nancy";"FONG";"292069941713060";"21002581";"475 rue Evariste Galois, Résidence Les Calades, n°3206";"";"";"";"06410";"Biot";"FRANCE";"475 rue Evariste Galois, Résidence Les Calades, n°3206--06410 Biot";"nancy.fong06@gmail.com";"0610703448";"26/06/1992";"Matmut";"769 9040 72338 P 50";"mycs GmbH";"";"Friedrichstr. 123";"";"";"";"10117";"Berlin";"ALLEMAGNE";"Friedrichstr. 123--10117 Berlin--ALLEMAGNE";"M.";"Ka Chun";"To";"kachun@mycs.com";" +49 176 6118 1645";"Managing Director";"01/04/2015";"31/08/2015";"40";"107";"22";"";"";"1000";"Euros";"1000";"";"";"Web Development Internship";"- Learn how to code using the mycs technology stack
  - Javascript / CoffeeScript
  - AngularJS
  - React & Flux (still to be implemented)
  - node.js / hapijs
  - Postgres
  - Redis, Knex, Bootstrap, Grunt, Gulp, Bower, Jasmine, Karma

- Work using agile methodologies on a daily basis
  - Scrum
  - Continous Integration & Deployment
  - Self-organizing teams
  - Pair programming
  - MVP approach

- Get to know how a fully automated Microservice Architecture works on a production environment
  - RESTful
  - Amazon AWS
  - Docker
  - Ansible
  - Varnish

- Work on own Frontend projects

- Work on own Backend projects";"Friedrichstr. 123";"";"";"";"10117";"Berlin";"Allemagne";"Friedrichstr. 123--10117 Berlin--Allemagne";"M.";"Claudio";"Bredfeldt";"claudio@mycs.com";"+49 (0) 30/24333736";"Technology Director";"M.";"Michel";"Buffa";"buffa@unice.fr";"0662659345";"";"";"";
"2015-04-13 14:54";"SI 5";"non";"oui";"M.";"Luis";"GIOANNI";"191097511877859";"20905615";"45 rue Barberis";"";"";"";"06300";"Nice";"FRANCE";"45 rue Barberis--06300 Nice";"lgioanni@gmail.com";"0679306991";"26/09/1991";"12345";"12345";"Laboratoire I3S";"http://www.i3s.unice.fr";"2000 Route des Lucioles";"Les Algorithmes";"Bâtiment Euclide B";"";"06900";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"2000 Route des Lucioles--Les Algorithmes--Bâtiment Euclide B--06900 SOPHIA ANTIPOLIS CEDEX";"M.";"Michel";"RIVEILL";"sabine.barrere@i3s.unice.fr";"0492942705";"Directeur du Laboratoire I3S";"01/05/2015";"30/09/2015";"35";"106";"22";"";"";"500,51";"Euros";"500,51";"";"";" Sélection et composition opportuniste de capteurs pour la reconnaissance d’activités";"Contexte du stage :

Christel Dartigues Pallez et Frédéric Precioso : Equipe Mind, Laboratoire I3S
Stéphane Lavirotte et Jean-Yves Tigli : Equipe Rainbow, Laboratoire I3S

Nous disposons d’un ensemble de données issues de différents capteurs portés (accéléromètres, magnétomètres, gyroscopes, centrales inertielles, …) et équipant un environnement physique (détecteur de présence, contacteurs, …). Notre objectif est de reconnaître les activités d’un utilisateur à partir de ces données.
La problématique de l’étude proposée est que nous ne pouvons pas faire de présuppositions sur la disponibilité des capteurs ; on peut même avoir de nouveaux capteurs non connus a priori qui vont concourir à la reconnaissance de l’activité. L’approche souhaitée est donc de réaliser cette reconnaissance de manière totalement opportuniste, en fonction des capteurs disponibles à un instant donné.
Nous avons déjà mis en place différents algorithmes d’apprentissage permettant de reconnaître l’activité d’un utilisateur. L’apprentissage de l’activité a été réalisé avec l’ensemble des données de tous capteurs, puis pour chaque capteur indépendamment puis avec des configurations de capteurs en nombre variables. Les résultats obtenus sont cohérents avec les autres travaux du domaine et il a été possible d’améliorer ces résultats grâce à l’utilisation d’algorithmes d’apprentissage de type Random Forest.
Objectifs du stage
Le premier objectif de ce stage est de repartir des résultats obtenus précédemment et de confirmer et justifier les améliorations apportées par l’apprentissage à base de Random Forest. Le but de cette première partie de stage est de finaliser l’ensemble des informations nécessaires pour publier les résultats obtenus dans une conférence scientifique.
Dans un second temps, il faudra étudier les approches possibles pour le mécanisme de sélection des capteurs disponibles pour la reconnaissance d’une activité. Ce deuxième objectif est donc d’améliorer les résultats de reconnaissance d’activité tout en garantissant une approche opportuniste (à savoir une sélection parmi les capteurs disponibles à un instant donné). Le mécanisme de sélection pourra être influencé par différents critères de recherche (consommation énergétique minimisée, optimisation de la reconnaissance, …).
Le troisième objectif de ce stage sera d’étudier la composition opportuniste des capteurs sélectionnés pour réaliser une application auto-adaptative de reconnaissance d’un ensemble d’activités. Cette troisième partie devra débuter par une étude bibliographique et proposer une approche pour la composition dynamique des données des capteurs ou des résultats des apprentissages dans le but de reconnaître un ensemble d’activités.
Compétences requises
Les compétences attendues pour traiter ce sujet sont :
-	Des connaissances sur les techniques d’apprentissage et la maîtrise d’un outil comme Matlab
-	Des connaissances sur le domaine de l’Intelligent Ambiante
-	Des compétences personnelles comme : initiative et force de proposition, autonomie, …

Références
D. Roggen, A. Calatroni, K. Fröster, G. Tröster, P. Lukowicz, D. Bannach, A. Ferscha, M. Kurz, G. Hölzl, H. Sagha, H. Bayati, J. Millán and R. Chavarriaga. « Activity recognition in opportunistic sensor environments », Procedia Computer Science, vol 7, pp 173-174, 2011
G. Hölzl, M. Kurz and A. Ferscha. « Goal oriented opportunistic recognition of high-level composed activities using dynamically configured hidden markov models ». In The 3rd International Conference on Ambient Systems, Networks and Technologies (ANT2012), 2012
M. Kurz, G. Hölzl, and Alois Ferscha. « Dynamic adaptation of opportunistic sensor configurations for continuous and accurate activity recognition ». In Fourth International Conference on Adaptive and Self-Adaptive Systems and Applications (ADAPTIVE2012), July 22-27, Nice, France, July 2012
D. Roggen, A. Calatroni, K. Förster, G. Tröster, P. Lukowicz et al. « Activity Recognition in Opportunistic Sensor Environments ». The European Future Technologies Conference and Exhibition 2011, 2011.
";"Equipe Sparks";"";"";"";"06900";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"Equipe Sparks--06900 SOPHIA ANTIPOLIS CEDEX";"M.";"Stéphane";"LAVIROTTE";"stephane.lavirotte@unice.fr";"0679672827";"Enseignant chercheur";"M.";"Jean-Yves";"TIGLI";"tigli@unice.fr";"0684245567";"";"";"";
"2015-05-15 10:13";"SI 5";"non";"non";"M.";"David";"GUERRERO";"191046444522936";"20906154";"210 Avenue Roumanille";"Résidence Thésa";"Appartement 222 B1";"";"06410";"BIOT";"FRANCE";"210 Avenue Roumanille--Résidence Thésa--Appartement 222 B1--06410 BIOT";"heldroe@gmail.com";"0629089011";"22/04/1991";"GAN";"A01385/091552079";"Ignilife France SAS";"www.ignilife.com";"27 rue Professeur Delvalle";"";"";"";"06300";"NICE";"FRANCE";"27 rue Professeur Delvalle--06300 NICE";"M.";"Fabrice";"Pakin";"fabrice@ignilife.com";"0673464911";"CEO";"11/05/2015";"28/08/2015";"35";"76";"16";"";"";"1000";"Euros";"1000";"";"20€ par mois de frais de transport";"Plate-forme objets connectés";"La plate-forme des objets connectés est une application à part entière indépendante de l‘application Ignilife. L‘idée est d‘externaliser les connecteurs actuels vers les API de Fitbit, Runkeeper et Withings (et d‘autres à venir) vers cette application tout en repensant le modèle d‘intégration. 
L‘objectif est d‘optimiser le temps et la facilité d‘intégration de nouveaux objets connectés en assouplissant et en rendant générique le modèle actuel. De plus, l‘application doit permettre de délivrer une donnée formalisée et ce quelque soit le provider source.  

Aspects techniques : l‘application doit proposer un modèle générique de provider et doit permettre de plugger de la manière la plus simple possible tous nouveaux providers en limitant au maximum le code spécifique. Les données utilisateurs sont stockées sur la plateforme selon un modèle unique et doivent être accessibles via une API sécurisée. ";"27 rue Professeur Delvalle";"";"";"";"06300";"NICE";"FRANCE";"27 rue Professeur Delvalle--06300 NICE";"M.";"David";"BESSOUDO";"david@ignilife.com";"0689248157";"CTO";"M.";"Frédéric";"PRECIOSO";"frederic.precioso@polytech.unice.fr";"0492965143";"";"";"";
"2015-03-18 15:27";"SI 5";"non";"non";"M.";"Jamal";"HENNANI";"192013417256369";"21209240";"2255 route des Dolines, Residence les Dolines APT 74";"";"";"";"06560";"VALBONNE";"FRANCE";"2255 route des Dolines, Residence les Dolines APT 74--06560 VALBONNE";"jamal.hennani@gmail.com";"0663795666";"27/01/1992";"ADH - SEGIA";"AC482864";"AMADEUS SAS";"";"485 Route du Pin Montard- Les Bouillides BP 69-";"";"";"";" 06902";"SOPHIA ANTIPOLIS CEDEX";"FRANCE";"485 Route du Pin Montard- Les Bouillides BP 69---06902 SOPHIA ANTIPOLIS CEDEX";"M.";"Pierre ";"PUIG";"stephanie.gasperini@amadeus.com";"04 97 15 45 81";"Talent Acquisition Manager";"01/04/2015";"30/09/2015";"37";"128";"26";"";"";"1100";"Euros";"1100";"Virement bancaire";"Le stagiaire bénéficiera de l’accès au restaurant d‘entreprise et d‘une prime de transport prenant en compte le trajet quotidien entre le domicile du stagiaire pendant l’exécution de son stage et l‘entreprise (versement de cette prime selon les procédures définies par l’entreprise).";"Monitor and better control unexpected overbookings";"Our functional areas include the management of the bookings in airline inventory system (counters, bookings data, etc), the waitlist clearance process and the data reconciliation between reservation and inventory.
Main Responsibilities:
Our teams are responsible for sustaining critical traffic, reaching more than 2500 transactions per seconds at peak time.
This traffic is made of various transactions. The diversity of the flow, the throughput and the continuous evolution of our applications can be a source of problems impacting the accuracy of our booking counters.
In this complex environment, we implemented a system to automatically recover most of the problems. Indeed, inaccurate counters is a source of dissatisfaction for our partners. This can lead to overbookings, with potential impact on flight departure, or to empty seats which is a loss of revenue for the airlines.
We need to:
 -improve our visibility on the counter corrections done every day in our system. How many bookings were updated? Is there a serious impact for the flight? Is the flight close to departure? What are the tools to put in place to notify relevant people? What‘s the cause of the problem?
 -have an efficient way to schedule some quality checks of our counters
The internship will participate in the implementation of a new application providing these functionalities to the teams.
The trainee will participate in following activities:
 -Get knowledge about airline inventory through existing documentation and presentations done by the team
 -Participate in defining the requirements with development and product definition teams
 -Help the team to design and implement the tools
 -Participate in the presentation to the teams";"2 rue du Vallon, Amadeus";"";"";"";"06560";"VALBONNE";"FRANCE";"2 rue du Vallon, Amadeus--06560 VALBONNE";"M.";"Maxime ";"ARMAND";"marmand@amadeus.com";"+33 4 9715 4971";"Manageur";"M.";"Francoise";"BAUDE";"baude@unice.fr";"+33 4 92 38 76 71";"";"";"";
"2015-03-19 12:14";"SI 5";"non";"non";"M.";"Swan";"JUMELLE-DUPUY";"189110608807665";"20904568";"3 rue guiglia";"";"";"";"06000";"NICE";"FRANCE";"3 rue guiglia--06000 NICE";"swan.jumelle@gmail.com";"+33618811975";"04/11/1989";"FILIA-MAIF";"5075611 N";"Ojingo Labs";"www.ojingolabs.com";"2101 23rd St";"";"";"";"94107";"San Francisco CA";"ETATS-UNIS";"2101 23rd St--94107 San Francisco CA--ETATS-UNIS";"M.";"Thomas";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Directeur";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"3600";"Euros";"3600";"Virement";"";"Dévoloppement d‘interfaces admin";"Les développeurs serveur se trouvant en Russie, et vu la complexité des technologies mises en place, il est actuellement difficile de faire le lien entres les bugs client/serveur.
Le but du stage va consister à donner plus de visibilité aux développeurs Client, testeurs et administrateurs.
Il faudra mettre en place plusieurs interfaces web permettant:
- De connaitre l‘état de santé du cluster de serveur (Monitoring OS, DB, APP)
- De visualiser les logs de l‘application, éventuellement avoir des statistiques
- De visualiser la base de données, et éventuellement interagir avec celle-ci

D‘autres solutions pourront être exploré afin d‘aider aux mieux les non-développeurs serveur à comprendre ce qu‘il s‘y passe.

L‘implémentation de certaines fonctionnalités et résolution de bugs coté serveur feront aussi parti des tâches assignées pendant le stage.";"1 place Masséna";"";"";"";"06000";"NICE";"FRANCE";"1 place Masséna--06000 NICE";"M.";"Thomas";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Directeur";"M.";"Christian";"BREL";"brel@i3s.unice.fr";"+33 (0)4 92 96 50 74";"";"";"";
"2015-03-06 12:24";"SI 5";"non";"non";"M.";"Antoine";"LAVAIL";"191106938237790";"20900385";"247 impasse des Genévriers";"";"";"";"83100";"TOULON";"FRANCE";"247 impasse des Genévriers--83100 TOULON";"antoine.lavail@icloud.com";"0650558350";"23/10/1991";"MMA";"116363311";"Ojingo Labs";"http://www.ojingolabs.com";"2101 23rd St";"";"";"";"94107";"San Francisco, California";"USA";"2101 23rd St--94107 San Francisco, California--USA";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"4000";"Euros";"4000";"";"";"Développeur d‘applications iOS";"Vous intégrerez l‘équipe composée de quatre développeurs iOS, de deux développeurs serveurs Vert.X, de trois graphistes et de deux project manager. Durant votre stage, vous participerez à la phase de reflexion et de création des projets en apportant votre vision technique sur la réalisation d‘applications aussi bien en terme d‘ossature, d‘ergonomie et de design de ces dernières. Vous pourrez aussi conseiller le graphiste lors de la réalisation des wireframes, et/ou maquettes en matière de format, de taille des fichiers, de compression et de portabilité mais aussi des chefs de projet lors de la rédaction du cahier des charges.";"1, place Masséna";"";"";"";"06000";"NICE";"FRANCE";"1, place Masséna--06000 NICE";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"M.";"Anne Marie";"DERY-PINNA";"pinna@polytech.unice.fr";"06 61 02 93 87";"";"";"";
"2015-02-18 11:40";"SI 5";"non";"non";"M.";"Paul";"LAVOINE";"19126748240459";"lp206655";"2255 route des Dolines";"";"";"";"06560";"Valbonne";"FRANCE";"2255 route des Dolines--06560 Valbonne";"paul.lavoine@hotmail.fr";"+33650789978";"03/12/1991";"maaf";"167148625K003";"Big Boss Studio";"http://www.bigbossstudio.com/";"4 rue de la liberté";"";"";"";"06000";"Nice";"FRANCE";"4 rue de la liberté--06000 Nice";"M.";"Eric";"Di Filippo";"edf@bigbossstudio.com";"+33 480805570";"CEO, Partner";"16/03/2015";"16/09/2015";"35";"129";"26";"10h - 12h30
14h-18h30";"";"508,20";"Euros";"508,20";"";"Tickets restaurant (20 X 8,50)";"Ingénieur développeur d’applications mobiles";"• participer à l’étude et au développement de nouvelles applications mobiles en
prenant soin de respecter les étapes et les bonnes pratiques relatives au
développement logiciel, ce qui consistera en :

- participer aux spécifications en apportant l’expertise technique aux clients

- concevoir et modéliser l‘application

- développer l’application et mettre en place des tests automatisés

- participer à la recette de l‘application

- mettre en production sur les stores d‘applications

• participer à la maintenance et à l’amélioration des applications développées par
Big Boss Studio

• rechercher de nouveaux axes de développement et de nouvelles méthodes par la
veille technologique et la réalisation de prototypes ou par la participation aux
réponses aux appels d’offres

En parallèle et dans une démarche d’amélioration continue, participer au
développement des frameworks et outils internes permettant la mutualisation de code
et l’amélioration de la qualité des applications développées par Big Boss Studio.
";"4 rue de la liberté";"";"";"";"06000";"Nice";"FRANCE";"4 rue de la liberté--06000 Nice";"M.";"Cyril";"Chandelier";"cchandelier@big-boss-studio.com";"+33 683996977";"iOS Developer";"M.";"Philippe";"Renevier";"Philippe.RENEVIER@unice.fr";"+ 33 6 1844 9350";"";"";"";
"2015-03-10 10:19";"SI 5";"oui";"oui";"M.";"Hadrien";"LUTTIAU";"191047919128814";"21209982";"28 impasse des Grillons";"";"";"";"34540";"Balaruc les Bains";"FRANCE";"28 impasse des Grillons--34540 Balaruc les Bains";"hadrien.luttiau@gmail.com";"0782266556";"09/04/1991";"MAIF";"1112916 M";"National Institut of Informatics";"www.nii.ac.jp/en/";"2-1-2 Hitotsubashi, Chiyoda-ku";"";"";"";"101-8430";"TOKYO";"JAPAN";"2-1-2 Hitotsubashi, Chiyoda-ku--101-8430 TOKYO--JAPAN";"M.";"Henri";"ANGELINO";"nii-internship@nii.ac.jp";"+81-3-4212-2165";"Acting Director Global Liaison Office";"16/03/2015";"10/09/2015";"35";"123";"26";"The NII Internship program is not an employment, so it‘s not appropriate to write as ""working hours"".
Instead, you could put like ""on a full-time basis""";"";"171000";"Yen";"1304,15";"";"";"Smart Service Compositions/Mashups in the City and the Web";"Intern students are expected to learn further knowledge and skills through collaborative activities for promoting our research activities and/or exploring further research topic. Depending on the status of the intern students (e.g., already have active master/phd topics or not), actual work in the internship can be determined flexibly through discussion. The outputs are typically joint papers and/or implemented software, but depend on the topic and the duration of the internship.

Subject:
Smart Service Compositions/Mashups in the City and the Web
Adaptation and Evolution Techniques in Service-based Systems (with Paradigms of ""Models at runtime"" / ""Requirements at runtime"")
Extension of Service Composition Techniques for Internet-of-Things (IoT) Services";"2-1-2 Hitotsubashi, Chiyoda-ku";"";"";"";"101-8430";"TOKYO";"JAPAN";"2-1-2 Hitotsubashi, Chiyoda-ku--101-8430 TOKYO--JAPAN";"M.";"Fuyuki";"ISHIKAWA";"f-ishikawa@nii.ac.jp ";"+81-3-4212-2675";"Associate Professor";"M.";"Gaëtan ";"REY";"gaetan.rey@unice.fr";" +33 4 92 96 51 44";"";"";"";
"2015-02-12 20:56";"SI 5";"oui";"non";"M.";"Nicolas";"MARQUEZ";"191079913816274";"20901829";"36 A val des castagnins";"les jardins de Ste Agnès";"";"";"06500";"Menton";"FRANCE";"36 A val des castagnins--les jardins de Ste Agnès--06500 Menton";"nmarquez@polytech.unice.fr";"0666664070";"10/07/1991";"PACIFICA";"1544166904";"Monaco Asset Management";"";"27 Boulevard Princesse Charlotte";"";"";"";"98000";"Monaco";"MONACO";"27 Boulevard Princesse Charlotte--98000 Monaco--MONACO";"M.";"Anthony";"Torriani";"atorriani@monacoasset.com";"+377 97 97 64 00";"CEO";"16/03/2015";"16/09/2015";"39";"129";"26";"8h30-12h30
14h-17h48";"";"1000";"Euros";"1000";"";"";"Développement d‘un outil de gestion de portefeuille";"Notre société développe actuellement en interne un logiciel de gestion de portefeuille.
Le stagiaire sera en charge de consolider l’équipe de développement et sera missionné sur deux axes importants :

- Améliorer l’interface (côté client – Visual C++)

- Développer de nouvelles fonctionnalités pour l’outil (Dont développement de procédures 

stockées / fonctions sous PostgreSQL)";"27 Boulevard Princesse Charlotte";"";"";"";"98000";"Monaco";"Monaco";"27 Boulevard Princesse Charlotte--98000 Monaco--Monaco";"M.";"Vincent";"Froidefond";"VFroidefond@monacoasset.com";"(+377) 97 97 64 14";"IT Manager & Administration";"M.";"Anne-Marie";"HUGUES";"amhc@wanadoo.fr";"0684045930";"";"";"";
"2015-02-25 12:41";"SI 5";"non";"non";"M.";"Hugo";"MARTINEZ";"192045951258782";"20905088";"10 domaine de l‘oratoire";"";"";"";"83300";"Draguignan";"FRANCE";"10 domaine de l‘oratoire--83300 Draguignan";"hugo.83300@gmail.com";"0629698652";"08/04/1992";"LMDE";"5625166";"Worldline";"http://worldline.com";"80 Quai Voltaire ";"";"";"";"95870 ";"BEZONS";"FRANCE";"80 Quai Voltaire--95870 BEZONS";"M.";"Alice ";"WAUCQUIER";"alice.waucquier@atos.net";"03.20.60.79.42";"Recruteuse";"23/03/2015";"18/09/2015";"37";"124";"26";"";"";"1300";"Euros";"1300";"Virement banquaires";"Accès au restaurant d’entreprise

prise en charge des transports en commun à hauteur de 50%";"ADAPTIVE WEB DESIGN ET WEBSOCKET";"Le stage démarre par l‘étude des solutions d‘identification des 

informations de l‘environnement de l‘utilisateur et de son contexte 

d‘utilisation (mobile dernière génération, réseau edge/3g/4g/wifi, ...), 

ainsi que l’étude de la mise en place d‘une solution de communication 

serveur vers client via les technologies WebSocket, Server-Sen events.

Le stagiaire proposera ensuite le cadre d‘utilisation de ces données dans 

nos services existants.

Puis il procèdera à l’implémentation des propositions retenues en 

veillant toujours à optimiser les performances et à améliorer l’expérience 

utilisateur.

Les technologies abordées et produits utilisés seront :

• HTML5 : Websocket / eventSource

• jQuery+Backbone.js

• JSON

• Java / SQL

• Apache / Tomcat

• Eclipse / Maven / Mercurial / Jenkins";"ZIA Rue de la Pointe";"";"";"";"59113 ";"SECLIN";"FRANCE";"ZIA Rue de la Pointe--59113 SECLIN";"M.";"Sylvain ";"POLLET-VILLARD";"sylvain.polletvillard@worldline.com";"+33 3 20 60 68 84";"Ingénieur Etudes & Developpement";"M.";"Michel";"Buffa";"buffa@unice.fr";"(33)-92-07-66-60";"";"";"";
"2015-04-09 14:22";"SI 5";"non";"non";"M.";"Eric";"MASOERO";"191030608844255";"20905415";"171, chemin Barella";"";"";"";"06390";"Contes";"FRANCE";"171, chemin Barella--06390 Contes";"eric.masoero@gmail.com";"0621423027";"15/03/1991";"Aviva";"74486834";"Orange";"";"78 rue Olivier de Serres";"";"";"";"75505";"Paris Cedex";"FRANCE";"78 rue Olivier de Serres--75505 Paris Cedex";"Mme";"Fabienne";"Patinet";"fabienne.patinet@orange.com";"02 38 42 91 71";"DRH Site";"13/04/2015";"25/09/2015";"35";"115";"24";"";"";"1311,77";"Euros";"1311,77";"";"";"Nouveaux enablers pour le Mail et les Contacts Orange";"Mission :Orange propose à ses clients une suite d’outils de communication (Mail, SMS, MMS, Instant Messaging, Carnet
d’adresses, Calendrier) accessible depuis le Portail Orange mais également depuis différents supports : mobiles, tablettes, PC,
TV. Ces produits évoluent rapidement, en fonction des avancées technologiques et marketing (exemple : Cloud, réseau sociaux,
webIM...) mais également en fonction des supports et des OS.
Au sein d’une équipe pluridisciplinaire et d’un projet de grande envergure, l’étudiant devra contribuer à l’amélioration des
services Mail et Carnet d’adresses d’Orange, permettant à l’opérateur de rester leader ou à la pointe du marché sur ces
produits stratégiques. L’étudiant sera donc amené à étudier, intégrer, développer et produire de nouvelles fonctionnalités
pertinentes pour l’utilisateur final.
En particulier, il s’intéressera dans le cadre de ce stage aux Enablers, API et Backends Mail et Carnet d’adresses PIM. Il devra
étudier, améliorer ou développer, puis intégrer de nouvelles API (REST), Enablers ou Backends. Il participera également à la
mise en production de ces nouvelles API pour l’amélioration des services Mail et Contacts Orange.
";"DSI DIGITAL FACTORY";"790 AV DR MAURICE DONAT";"BATIMENT MARCO POLO C2";"";"06250";"Mougins";"FRANCE";"DSI DIGITAL FACTORY--790 AV DR MAURICE DONAT--BATIMENT MARCO POLO C2--06250 Mougins";"M.";"Benoit";"Mercier";"benoit1.mercier@orange.com";"04 97 46 28 61";"Chef de projet SI";"M.";"Fabien";"Hermenier";"fabien.hermenier@unice.fr";"04 92 38 76 36";"";"";"";
"2015-03-29 21:08";"SI 5";"non";"non";"M.";"Marouan";"MESDOURI";"1880999000000093";"21410536";"25 , rue Robert Latouche Résidence Jean Medecin";"";"25 , rue Robert Latouche Résidence Jean Medecin";"";"06200";"NICE";"FRANCE";"25 , rue Robert Latouche Résidence Jean Medecin--25 , rue Robert Latouche Résidence Jean Medecin--06200 NICE";"marouanmesdouri@gmail.com";"0667302769";"02/09/1988";"lcl";"2006767539200";"APS int";"www.onduleurs.fr";"15 bis corniche André de Joly";"";"";"";" 06300";"NICE";"FRANCE";"15 bis corniche André de Joly--06300 NICE";"M.";"Jean Claude ";" MATHIEU";"jcmathieu@onduleurs.fr";"06 11 35 53 73";" Gérant";"09/04/2015";"09/09/2015";"35";"107";"22";"";"";"508.20";"Euros";"508.20";"";"";"DEVELOPPEMENT SITES WEB MARCHANDS";"SITE ONDULEURS.FR
Développement et amélioration du site de vente en ligne onduleurs.fr
Développement d‘un système de veille concurrentielle sur les sites concurrents
Synchronisation des deux systèmes entre eux en fonction de la politique tarifaire de la société.
Système utilisé PHP MYSQL 
ITE EVENTBOUTIK.COM Conception d’un Back office / CRM pour gérer le site ";"15 bis corniche André de Joly";"";"";"";"06300";"NICE";"FRANCE";"15 bis corniche André de Joly--06300 NICE";"M.";" Philippe";" MATHIEU";"pmathieu@onduleurs.fr";"04 93 89 19 29";"Gestionnaire informatique";"M.";"Igor";"LITOVSKY";"lito@polytech.unice.fr";"04 92 96 51 24";"";"";"";
"2015-02-17 11:53";"SI 5";"oui";"non";"M.";"Thomas";"MONTANA";"1 93 10 75 112 632 12";"mt003150";"37c, Boulevard Gorbella";"Palais Johnny";"";"";"06100";"Nice";"FRANCE";"37c, Boulevard Gorbella--Palais Johnny--06100 Nice";"thmsmontana@gmail.com";"0674067141";"11/10/1993";"BPCE Assurances";"006108944";"GoEuro Corp.";"http://www.goeuro.com/";"Sonnenburger Str. 73";"";"";"";"10437";"Berlin";"GERMANY";"Sonnenburger Str. 73--10437 Berlin--GERMANY";"M.";"Naren";"Shaam";"naren.shaam@goeuro.com";"+4915234071141";"Chief Executive Officer";"01/04/2015";"31/08/2015";"40";"107";"22";"";"";"800";"Euros";"800";"";"—";"Web Frontend - Software Engineering Internship";"Due to the nature of our business and small length of release cycles, there will not be any long lasting projects nor projections on upcoming tasks. Oppositely, the internship will consist of a large number of small/medium length tasks. The whole picture of high velocity web and mobile web frontends will approached during this internship, as well as best practices of software development. 

We set up aggressive goals in terms of velocity, scalability, testability and maintainability. Thomas will be involved, from an abstract manner as of now, in all this.

From a technological standpoint, JavaScript will be the main language, in addition to HTML5, CSS3 and many others.";"Sonnenburger Str. 73";"";"";"";"10437";"Berlin";"Germany";"Sonnenburger Str. 73--10437 Berlin--Germany";"M.";"Louis";"Hache";"louis.have@goeuro.com";"+49 162 770 69 37";"Software Engineer / Web Frontend";"M.";"Michel";"Buffa";"buffa@polytech.unice.fr";"0662659345";"";"";"";
"2015-01-23 18:12";"SI 5";"oui";"oui";"M.";"Jérôme";"RANCATI";"192040602701944";"rj001117";"Traverse du vieux four cidex 434";"06330 Roquefort les Pins";"";"";"06330";"Roquefort les Pins";"FRANCE";"Traverse du vieux four cidex 434--06330 Roquefort les Pins--06330 Roquefort les Pins";"jerome.rancati@gmail.com";"+33628518057";"02/04/1992";"PROTEC BTP";"106170540 Y 003";"SnT Research Center";"http://wwwfr.uni.lu/snt";"4 Rue Alphonse Weicker, L-2721 Luxembourg";"";"";"";"L-2721";"LUXEMBOURG";"LUXEMBOURG";"4 Rue Alphonse Weicker, L-2721 Luxembourg--L-2721 LUXEMBOURG--LUXEMBOURG";"M.";"Björn";"OTTERSTEN";"bjorn.ottersten@uni.lu";"(+352) 46 66 44 5665";"Directeur du Centre Interdisciplinaire ""Security, Reliability and Trust""";"15/03/2015";"15/09/2015";"40";"129";"26";"";"";"850";"Euros";"850";"";"";"Internship at the SnT Research Center working on a SmartHomes project";"In SmartHomes and SmartCities in general, various sensors and actuators are installed in the environmentto sense users activites. This not only allows to collect information about activities, but also to remotelyactivate devices to automate some scenario.
The ""Smart"" aspect of such environment is, nowadays, limited to reactive actions often encoded as sets ofrules interpreted by a specific engine to perform the automation. But the complexity to develop andconfigure such environment, specialized for each users, is today a blocking aspect. We are studying theuse of ""live machine learning"" to create a novel SmartHome engine, which should learn from users habits toalign the home settings to users‘ everyday activity (i.e., by tuning the luminosity automatically based on the
regular manual setting of user).
This internship will offer two main aspects in collaboration with our industrial partner.
Firstly, the internship student will collaborate with the research team in order to learn and develop amachine learning algorithm to correlate sensors data in live. The outcome of this step will be integratedinto the Kevoree Modelling Framework.
Secondly, the student will have to develop an innovative web based dashboard (technology : GooglePolymer + Kevoree Modeling Framework), leveraging this machine learning algorithm. This dashboard should help users to understand and anticipate actions taken autonomously by the home with thissmart non-rules based engine (i.e., by displaying prediction and potential actions in future).
Because this project is included in a long run research activities, this subject has a great potential to becontinued as a PhD.
The requirements for this internship are:
High programming skills(Java/JS mainly).
Some knowledge about IoT/Sensors environments to understand how to select and process data.
Modeling knowledge to represent complex graph of dataThe machine learning aspect can be learnt during the internship, a first knowledge is a plus, but is not
mandatory.
Useful Links:
Google Polymer Framework: https://www.polymer-project.org/
Kevoree Modeling Framework: http://kevoree.org/kmf/";"6 rue Richard Coudenhove-Kalergi L-1359 Luxembourg-Kirchberg";"";"";"";"L-1359 ";"Luxembourg-Kirchberg";"LUXEMBOURG";"6 rue Richard Coudenhove-Kalergi L-1359 Luxembourg-Kirchberg--L-1359 Luxembourg-Kirchberg--LUXEMBOURG";"M.";"François";"FOUQUET";"francois.fouquet@uni.lu";"+352 466644 5387 ";"Research Associate";"M.";"Sébastien";"MOSSER";"mosser@i3s.unice.fr";"0493925058";"";"";"";
"2015-03-19 13:29";"SI 5";"non";"non";"Mlle";"Roberta";"ROBERT";"000000000000";"21406870";"2255 Route des Dolines";"Chambre 203";"";"";"06560";"VALBONNE";"FRANCE";"2255 Route des Dolines--Chambre 203--06560 VALBONNE";"fleur.robert@gmail.com";"0695226334";"22/01/1992";"LCL";"1514341904";"SAP";"www.sap.com";"Tour SAP";"35 rue d‘Alsace";"";"";"92309";"LEVALLOIS-PERRET";"FRANCE";"Tour SAP--35 rue d‘Alsace--92309 LEVALLOIS-PERRET";"Mme";"Agnès";"DESPLECHIN";"sylvine.eusebi@sap.com";"04 92 28 62 00";"Responsable du Service Recrutement";"01/04/2015";"31/07/2015";"35";"85";"18";"";"";"1 174";"Euros";"1 174";"";"";"Proxified Dynamic Security Testing";"This internship is based in the SAP Labs France Research Lab, in Sophia-Antipolis. The work will be
performed in the context of the Research Program “Security & Trust”, and deals with dynamic security testing
and secure coding.
SAP has released a powerful platform called SAP HANA, which comprises two main components: the HANA
DB, an extremely efficient in-memory database, and the HANA XS engine, a server-side javascript based
application server. HANA application development is being done through an Integrated Development Editor
called WebIDE.
The development of an SAP application requires going through an internal security development lifecycle,
aiming at reducing the likeliness of introducing vulnerability at each development stage. When it comes to
the implementation phase, detecting code mistakes as they are being typed is the best way to ensure secure
code. For HANA development, a security guide was developed with requirements and best practices for
ensuring good code quality. The next step is to automate the verification of what is recommended in the
guide by the WebIDE itself.
A first prototype was already developed to this end, with a focus on static code analysis and on some
dynamic analysis. The next step is to extend on the dynamic analysis aspect, by keeping in mind the strong
requirements which come with it on performance and security. Dynamic security is on a thin edge between
defensive and offensive security. Offensive tools are a great asset for verifying certain security properties of
the application but can have detrimental effect if misused.
The goal of this internship is thus to assess different dynamic security testing scenarios and to implement a
proof of concept for those which comply to the performance and security requirements, in an effort to
enhance the WebIDE with a full set of offensive tools which can be used in a controlled fashion. One such
scenario will be to assess the possibility to develop and run an extension of or a tool similar to sqlmap
against SAP HANA applications through the WebIDE.
The candidate should be an expert in SQL, JavaScript and Web technologies. He should have a strong
background in (defensive) IT security and in penetration testing. He should be familiar with penetration
testing tools such as sqlmap and metasploit.
We expect that 20% of time will be dedicated to research activities, and 80% to development.";"805 Avenue Maurice Donat  Le Font de l’Orme";"";"";"";"06259";"MOUGINS";"FRANCE";"805 Avenue Maurice Donat  Le Font de l’Orme--06259 MOUGINS";"M.";"Cédric";"HERBERT";"cedric.hebert@sap.com";"04 92 28 62 00";"Team Leader";"M.";"Karima";"BOUDAOUD";"karima.boudaoud@polytech.unice.fr";"04 92 96 51 72";"";"";"";
"2015-03-20 14:23";"SI 5";"non";"non";"M.";"Victor";"SALLÉ";"193010502304469";"21206825";"2400 route des Dolines";"Res. Newton, Apt. A219";"";"";"06560";"VALBONNE";"FRANCE";"2400 route des Dolines--Res. Newton, Apt. A219--06560 VALBONNE";"victorsalle@outlook.com";"07.89.60.41.85";"28/01/1993";"Maaf";"159196140J002";"AVISTO";"www.avisto.com";"Space Antipolis 1";"2323, Porte 15, Chemin Saint-Bernard";"";"";"06220";"VALLAURIS";"FRANCE";"Space Antipolis 1--2323, Porte 15, Chemin Saint-Bernard--06220 VALLAURIS";"M.";"Christophe";"PARMENTIER";"christophe.parmentier@avisto.com";"04.92.38.74.71";"Responsable Régional";"25/03/2015";"25/09/2015";"35";"129";"26";"9h-12h / 14h-18h";"";"915";"Euros";"915";"Mensuellement par virement avec fiche de paie";"- 1 ticket restaurant (8,60€) par jour
- 16,67€ par mois de frais de déplacement";"Développement Android - MIDI DockStation & Android App";"Dans le cadre de développement d’une station d’accueil (dock station) MIDI et audio pour smartphone Android, vous participerez au développement (conception), au prototypage et à une étude de coût préliminaire d’industrialisation de cette station d’accueil.

En supplément de la fonction de chargement d’un smartphone, cette station d’accueil, destinée avant tout aux audiophiles et compositeurs/musiciens, permet :
  - Le mixage et le traitement de sources sonores externes et du flux audio d’un smartphone Android,
  - La génération sonore grâce à une puce sonore propriétaire Elsys Design,
  - L’enregistrement du flux audio provenant de la station d’accueil sur carte SD du smartphone via l’application Android dédiée,
  - Le contrôle d’instruments MIDI via toute application Android dédiée,
  - La transmission d’information MIDI via WiFi (MIDI over WiFi)


Pour cela, vous devrez:

1. Valider les fonctionnalités MIDI et audio USB de la station d’accueil :
  - Mise en place et prise en main des outils de développement (SDK) pour Android,
  - Etude de la norme USB et des différentes classes associées,
  - Conception et développement d’une application Android permettant d’envoyer/recevoir un flux audio vers/de la station - accueil via le port USB,
  - Conception et codage d’une application Android permettant d’envoyer/recevoir un flux MIDI vers/de la station d’accueil via le port USB,
  - Test et validation des fonctionnalités audio et MIDI USB de la station d’accueil,
  - Documents de conception et de validation

2. Créer une application audio media player/recorder permettant la lecture et l’enregistrement d’un fichier audio .WAV de/vers une carte SD à partir de la station d’accueil :
  - Définir le cahier des charges de l’application,
  - Conception et codage de l’application en JAVA,
  - Test et validation de l’application avec la station d’accueil,
  - Documents de conception et de validation

3. Créer une application spécifique permettant le séquençage d’évènements MIDI, la lecture et l’enregistrement audio, le contrôle de la station d’accueil :
  - Définir le cahier des charges de l’application,
  - Définir des différents formats de messages exclusifs MIDI entre l’application et la station d’accueil,
  - Conception et développement de l’application en JAVA,
  - Test et validation de l’application avec la station d’accueil,
  - Documents de conception et de validation";"Space Antipolis 1";"2323, Porte 15, Chemin Saint-Bernard";"";"";"06220";"VALLAURIS";"FRANCE";"Space Antipolis 1--2323, Porte 15, Chemin Saint-Bernard--06220 VALLAURIS";"M.";"Pierre";"PACCHIONI";"pierre.pacchioni@avisto.com";"04.92.38.74.72";"Directeur Technique";"M.";"Jean-Yves";"TIGLI";"jean-yves.tigli@unice.fr";"04.92.96.51.81";"";"";"";
"2015-03-12 01:19";"SI 5";"non";"non";"M.";"Rodrigo Augusto";"SCHELLER BOOS";"01B0JP00M31";"21407590";"45 Boulevard Pape Jean XXIII";"logement 43";"";"";"06300";"NICE";"FRANCE";"45 Boulevard Pape Jean XXIII--logement 43--06300 NICE";"rsboos@hotmail.com";"0782744908";"02/03/1995";"LCL";"1234556789";"Cadence Design Systems";"http://cadence.com/";"2655 Seely Avenue";"";"";"";"95134";"SAN JOSE";"USA";"2655 Seely Avenue--95134 SAN JOSE--USA";"M.";"Carolle";"PLANAS";"opicot@cadence.com";"04.89.87.30.00";"Human Resources Manager";"20/03/2015";"21/08/2015";"35";"108";"22";"";"";"1200";"Euros";"1200";"";"Droit au conges payes
Droit au titres restaurant
Droit au remboursement transport";"Développement du Virtuoso Dashboard - environment graphique pour les outils Virtuoso";"Les outils Virtuoso sont des logiciels graphiques très complexes (centaines de commandes, menus, langages d’extensions) manipulant des quantités importantes de données. Cadence souhaite développer un environnement graphique pour ses clients (Virtuoso Dashboard) permettant aux responsables des équipes de conception de mesurer les performances du logiciel (temps d’affichage, temps d’accès à la base de données, etc.) et les taux d’utilisation des fonctionnalités les plus avancées. 
Ce logiciel accèdera entre autres aux fichiers logs de Virtuoso, les organisera pour un accès rapide (data mining) et affichera des indicateurs de performance et d’utilisation.";"1080 route des dolines";"";"";"";"06560";"VALBONNE";"FRANCE";"1080 route des dolines--06560 VALBONNE";"M.";"Olivier";"PICOT";"opicot@cadence.com";"04.89.87.30.00";"Project Manager";"M.";"Michel";"RUHER";"ruher@i3s.unice.fr";"06";"";"";"";
"2015-03-17 19:27";"SI 5";"oui";"non";"M.";"Qihao";"TANG";"193029921600046";"21210465";"14 Rue de le Republique";"";"";"";"06600";"Antibes";"FRANCE";"14 Rue de le Republique--06600 Antibes";"qihao.tang@aliyun.com";"+33760456025";"21/02/1993";"LMDE";"06291980";"Everbright Securities";"http://www.ebscn.com/";"25 Taipingqiao St. Xicheng";"";"";"";"100034";"Beijing";"CHINA";"25 Taipingqiao St. Xicheng--100034 Beijing--CHINA";"M.";"Xinyuan";"Ge";"gexy@ebscn.com";"+8613528809518";"directeur général";"30/03/2015";"30/09/2015";"40";"129";"26";"";"";"24000";"RMB";"3500";"";"";"Les produits dérivés financiers";"Le design et le pricing sur les produits dérivés financiers.
Amélioration sur l‘algorithme existant et dévéloppement avec les programme.";"580 Nanjing West Rd. Jing‘an";"";"";"";"200041";"Shanghai";"China";"580 Nanjing West Rd. Jing‘an--200041 Shanghai--China";"M.";"Xinyuan";"Ge";"gexy@ebscn.com";"+8613528809518";"president général";"M.";"Ioan";"Bond";"bond@polytech.unice.fr";"+33(0)492965141";"";"";"";
"2015-03-02 11:04";"SI 5";"non";"non";"M.";"William";"TASSOUX";"191103417258051";"21206992";"210 Avenue roumanille";"";"";"";"06410";"BIOT";"FRANCE";"210 Avenue roumanille--06410 BIOT";"william.tassoux@gmail.com";"0651894242";"30/10/1991";"Cabinet Gril Assurances";"025899952";"GFI Informatique";"http://www.gfi.fr/";"145 boulevard Victor Hugo";"";"";"";"93400";"SAINT-OUEN";"FRANCE";"145 boulevard Victor Hugo--93400 SAINT-OUEN";"M.";"Stéphane";"Jourjon";"stephane.jourjon@gfi.fr";"0497155523";"Directeur de division";"09/03/2015";"09/09/2015";"35";"129";"26";"";"";"1200";"Euros";"1200";"Virement bancaire";"Carte avantage tickets restaurants
Prise en charge à hauteur de 50% des transports en commun sur justificatifs";"Evaluation, design and deployment of a distributed access control system";"The final purpose of the stage is to deploy a SSO (single-sign on) mechanism in the server-side application of a Graphical User Interface, in order to distribute authentication credentials to a plenty of possible systems the GUI is supposed to interact with.

The candidate will develop his project interacting with a team of 5/10 people, in GFI buildings located in Les Emeralds in Sophia Antipolis. 

The stage covers the requirements of a portion of a bigger project, meant to build-up a GUI able to communicate with multiple applications, to give the final users the possibility to interact with them through a single entry point.";"Emerald Square, Bâtiment B, Avenue Evariste Galois";"BP 199";"";"";"06904";"Sophia Antipolis";"FRANCE";"Emerald Square, Bâtiment B, Avenue Evariste Galois--BP 199--06904 Sophia Antipolis";"M.";"Emmanuel";"Jovenin";"emmanuel.jovenin@gfi.fr";"0489611288 ";"Delivery manager";"M.";"Jean-Paul";"Rigault";"jpr@polytech.unice.fr";"0492965133";"";"";"";
"2015-01-08 22:52";"SI 5";"non";"non";"Mlle";"Marie-Catherine";"TURCHINI";"291012B03319789";"21105993";"31 Boulevard PAOLI";"";"";"";"20200";"BASTIA";"FRANCE";"31 Boulevard PAOLI--20200 BASTIA";"turchini@polytech.unice.fr";"0615255288";"23/01/1991";"MAIF";"1530301D";"GoodBarber";"http://fr.goodbarber.com/";"12 Rue Général Fiorella";"";"";"";"20000";"AJACCIO";"FRANCE";"12 Rue Général Fiorella--20000 AJACCIO";"M.";"Dominique";"SIACCI";"siacci@duoapps.com";"0972134065";"Directeur général";"30/03/2015";"27/09/2015";"35";"126";"26";"";"";"1500 NET";"Euros";"1500";"";"";"Développement iOS (Objective-C) - Développement d‘un ensemble de fonctionnalités à destination des commerces de proximité";"La plate-forme GoodBarber consiste en un ensemble de fonctionnalités permettant la création d‘applications mobiles. Elle est le fruit de plusieurs mois de R&D ayant conduit à la création des différents moteurs d‘applications mobiles (pour iOS, Android, HTML5), qui constituent aujourd‘hui une base technique très solide pour l‘ajout de nouvelles fonctionnalités.
Pour l‘année 2015, les développements au sein de GoodBarber vont s‘orienter vers les interactions avec les utilisateurs finaux, et l‘ajout de fonctionnalités du produit à destination des commerces de proximité. Les objets connectés et leur technologies associées (iBeacon, NFC, …) seront une composante majeur de ces interactions. 

Ces fonctionnalités seront axées autour de la fidélisation, et du concept drive-to-store. Dans ce cadre, nous imaginons une adaptation des moteurs techniques pour permettre le déclenchement événementiel d‘actions en fonction des profils utilisateurs, ou d’informations transmises par des objets connectés, pour permettre la mise en place de scenarii marketing notamment à travers le geofencing.

Cet ensemble de fonctionnalités doit être pensé pour s‘intégrer et s‘inter-connecter au moteur GoodBarber, tant au niveau font-end natif (objet du stage) qu‘au niveau back-end.";"12 Rue Général Fiorella";"";"";"";"20000";"AJACCIO";"FRANCE";"12 Rue Général Fiorella--20000 AJACCIO";"M.";"Dominique";"SIACCI";"siacci@duoapps.com";"0972134065";"Directeur général";"M.";"Gaëtan";"REY";"Gaetan.Rey@unice.fr";"0492965144";"";"";"";
"2015-02-19 21:37";"SI 5";"non";"oui";"M.";"Michel";"VEDRINE";"1 91 06 78 646 509 73";"21209258";"41 rue Clément Roassal";"";"";"";"06000";"Nice";"FRANCE";"41 rue Clément Roassal--06000 Nice";"mvedrine@gmail.com";"0683158650";"23/06/1991";"MEP";"850445733";"CNRS Délégation Côte d‘Azur";"";"Les Lucioles 1 – Campus Azur - 250 rue Albert Einstein CS 10 269";"";"";"";"06905";"Sophia Antipolis cedex";"FRANCE";"Les Lucioles 1 – Campus Azur - 250 rue Albert Einstein CS 10 269--06905 Sophia Antipolis cedex";"Mme";"Béatrice";"SAINT-CRICQ";"beatrice.saint-cricq@dr20.cnrs.fr";"0493954256";"déléguée Régionale du CNRS Côte d‘Azur";"16/03/2015";"16/09/2015";"35";"129";"26";"";"";"508,20";"Euros";"508,20";"Virements bancaires mensuels";"";"Sécurisation des données et de l’échange de données dans un environnement mobile dans le cadre du projet PadDoc";"L‘objectif de ce stage d’une durée de 6 mois sera de poursuivre un travail initié lors de 2 PFEs de dernière année d’ingénieur concernant la sécurité de l’échange de données en utilisant les protocoles de communication NFC et Bluetooth et la sécurité du stockage des données sur un smartphone Android.

L’étudiant devra donc :

- Finaliser l’implémentation des composants logiciels de sécurité nécessaires pour sécuriser l’échange de données en NFC et Bluetooth.
- Finaliser l’implémentation des composants logiciels nécessaires à la sécurité du stockage des données sur le mobile
- Faire une évaluation de performances.
";"Laboratoire I3S UMR 7271 - GLC - Equipe RAINBOW ";"Bâtiment les Templiers 1 - Campus Sophi@Tech - 930 Route des Colles - BP 145";"";"";"06903";"Sophia Antipolis cedex";"FRANCE";"Laboratoire I3S UMR 7271 - GLC - Equipe RAINBOW--Bâtiment les Templiers 1 - Campus Sophi@Tech - 930 Route des Colles - BP 145--06903 Sophia Antipolis cedex";"Mme";"Karima";"Boudaoud";"karima@polytech.unice.fr";"0492965172 ";"Maître de Conférences";"M.";"Jean-Yves";"Tigli";"tigli@polytech.unice.fr";"0684245567";"";"";"";
"2015-02-08 21:17";"SI 5";"non";"non";"M.";"Damien";"VIANO";"191090608832970";"21004752";"195";"chemin des Caillades";"";"";"06480";"La Colle sur Loup";"FRANCE";"195--chemin des Caillades--06480 La Colle sur Loup";"damien.viano06@gmail.com";"0667507301";"07/09/1991";"Maif";"2316612 M";"Sopra Steria";"http://www.soprasteria.com/";"3 RUE DU PRE FAUCON PAE DES GLAISINS 	 ";"";"";"";"74940";" ANNECY LE VIEUX";"FRANCE";"3 RUE DU PRE FAUCON PAE DES GLAISINS--74940 ANNECY LE VIEUX";"M.";"Frédéric ";"Letot";"frederic.letot@soprasteria.com";" 0483150141";"Directeur d’agence";"09/03/2015";"30/09/2015";"35";"144";"29";"";"";"1120";"Euros";"1120";"";"Un accès à tous les avantages CE et une participation de Sopra aux frais de déjeuner d’un montant de 5,28€ par jour travaillé.";"Mise en place d‘une solution de statistiques et de supervision réseaux";"Au sein de l’équipe de développement de solutions de supervisions télécom, vous concevez
et élaborez une solution de statistiques et de reporting réseaux basée sur des produits
open source.
Vous êtes en charge de :
-> L’étude des besoins et le cadrage du projet, pour les Network Operations Center
d’opérateurs télécom,
-> Les spécifications fonctionnelles,
-> L’étude de faisabilité et la proposition des solutions sur des briques Open Sources,
-> La rédaction du dossier d’architecture,
-> La mise en œuvre des briques du projet sur un démonstrateur,
-> L’intégration des différents composants sur un portail.";"Athéna C, 1180 route des Dolines";"";"";"";"06904 ";"Sophia Antipolis";"FRANCE";"Athéna C, 1180 route des Dolines--06904 Sophia Antipolis";"M.";"Bruno";"Cruanes";"bruno.cruanes@soprasteria.com";"04 83 15 00 00";"Directeur de projet";"M.";"Frédéric";"Précioso";"frederic.precioso@polytech.unice.fr";"0492965143";"";"";"";
"2015-03-17 16:57";"SI 5";"non";"non";"M.";"Robin";"VIVANT";"191113919828331";"21003782";"200 chemin du Beal";"";"";"";"06480";"La Colle sur Loup";"FRANCE";"200 chemin du Beal--06480 La Colle sur Loup";"robin.vivant@gmail.com";"0622699358";"23/11/1991";"MAIF";"1273527D";"OjingoLabs";"http://www.ojingolabs.com/";"2101 23rd Street ";"";"";"";"94107";"San Francisco";"USA - CALIFORNIE";"2101 23rd Street--94107 San Francisco--USA - CALIFORNIE";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"23/03/2015";"22/09/2015";"35";"128";"26";"";"";"3600";"Euros";"3600";"Virement banquaire";"";"Fullstack web engineer";"Vous intègrerez l‘équipe composée de quatre développeurs iOS, de deux développeurs serveurs Vert.X, de trois graphistes et de deux project manager. Durant votre stage, vous participerez à la phase de réflexion et de création des projets en apportant votre vision technique sur la réalisation d‘applications aussi bien en terme d‘ossature, d‘ergonomie et de design de ces dernières. Vous pourrez aussi conseiller le graphiste lors de la réalisation des wireframes, et/ou maquettes.
Il y a besoin d‘une solution monitoring pour les serveurs. Vous effectuerez une recherche sur ce qui se fait et participerez à la mise en place des clusters de monitoring et si besoin est, créerez des composants graphiques pour afficher des metrics personnalisées.
La réalisation de vues web sera également demandé pour certains besoins de l‘application mobile et l‘envoi de mails.";"1 place Massena";"";"";"";"06000";"NICE";"FRANCE";"1 place Massena--06000 NICE";"M.";"TJ";"MARBOIS";"tj@ojingolabs.com";"0667317118";"Fondateur";"M.";"Christian";"BREL";"brel@polytech.unice.fr";"04 92 96 51 62";"";"";"";`

type BufferReader struct {
	input string
}

type MockJournal struct{}

func (m *MockJournal) UserLog(u schema.User, msg string, err error)       {}
func (m *MockJournal) Log(em, msg string, err error)                      {}
func (m *MockJournal) Wipe()                                              {}
func (m *MockJournal) Access(method, url string, statusCode, latency int) {}
func (m *MockJournal) Logs() ([]string, error)                            { return []string{}, nil }
func (m *MockJournal) StreamLog(k string) (io.ReadCloser, error)          { return nil, nil }

func (b BufferReader) Reader(year int, promotion string) (io.Reader, error) {
	return strings.NewReader(b.input), nil
}

func TestCSVParsing(t *testing.T) {
	r := BufferReader{input: buf}
	x := NewCsvConventions(r, []string{"si5"}, &MockJournal{})
	conventions, errors := x.Import()
	assert.Nil(t, errors)
	assert.Equal(t, 40, len(conventions))
}
