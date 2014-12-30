var interns;
var currentPage;
var waitingBlock;
var myself;

$( document ).ready(function () {
waitingBlock = $("#cnt").clone().html();

	user(function(u) {
        myself = u;
        $("#fullname").html(u.Firstname + " " + u.Lastname);
        showMyServices(u.Role);
	});	
	
       internships(function(data) {
       if (!interns) {            
        	interns = data;
            if (myself.Role >= 1) {
                showPage(undefined, "conventions");
            } else {
                showPage(undefined, "myStudents");
            }
        } else {
            interns = data;
        }
    }
    );
});

function showPage(li, id) {
    $("#cnt").html(waitingBlock);

    $("#menu-pages").find("li").removeClass("active");
    if (li) {
        $(li.parentNode).addClass("active");
    }
    currentPage = id;
    refresh();
}

function refresh() {
    //reset the div
    if (currentPage == "myStudents") {
        displayMyStudents();
    } else if (currentPage == "conventions") {
        displayMyConventions();
    } else if (currentPage == "assignments") {
        displayTutors();
    } else if (currentPage == "privileges") {
        showPrivileges();
    } else if (currentPage == "defenses-schedule") {
        showDefenses("schedule");
    } else if (currentPage == "defenses-juries") {
        showDefenses("juries");
    } else if (currentPage == "pending") {
        pickOne();
    } else if (currentPage == "juries") {
        showJuryService();
    } else if (currentPage == "service") {
        showService();
    }
    else {
        console.log("Unsupported operation on '" + currentPage + "'");
    }
}

function displayMyConventions() {
    var html = Handlebars.getTemplate("watchlist")(interns);
    var root = $("#cnt");
    root.html(html);
	if (interns.length == 0) {
    	return true
    }
    root.find(':checkbox').checkbox();
    root.find('tbody').find(':checkbox').checkbox().on('toggle', function (e) {
        generateMailto(root);
    });
    root.find('.mail-checkbox-stu').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-stu');
    });
    root.find('.mail-checkbox-s').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-s');
    });
    root.find('.mail-checkbox-t').click(function (e) {
        shiftSelect(e, this, root,'.mail-checkbox-t');
    });

    root.find(".mailto-students").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-students", root);
    });
    root.find(".mailto-tutors").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-tutors", root);
    });
    root.find(".mailto-sups").on('toggle', function (e) {
        toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-sup", root);
    });
    $("#table-conventions").tablesorter({
        headers: {
            0: {sorter: false},
            4: {sorter: false},
            5: {sorter: false},
            6: {sorter: "grades"},
            7: {sorter: "grades"},
            8: {sorter: "grades"}
        },
        theme: 'bootstrap',
        widgets : ["columns"],
        headerTemplate : '{content} {icon}'
    });
}

function showPrivileges() {
    users(function(data) {
        var others = data.filter(function (u) {
            return u.Email != myself.Email;
        });
        //The base
        var html = Handlebars.getTemplate("privileges")(others);
        $("#cnt").html(html);
        $("#cnt").find("select").selecter();
        $('[data-toggle="deluser-confirmation"]').each(function (i, a) {
            var j = $(a);
            j.confirmation({onConfirm: function() {removeUser(j.attr("data-user"), j.parent().parent().parent())}});
        });
    });
}

function showNewUser() {
    var buf = Handlebars.getTemplate("new-user")({});
    $("#modal").html(buf).modal('show');
    $("#modal").find("select").selecter();
}

function addUser() {
    if (missing("lbl-nu-fn") || missing("lbl-nu-ln") || missing("lbl-nu-email")) {
        return false;
    }
    newUser($("#lbl-nu-fn").val(), $("#lbl-nu-ln").val(),
                $("#lbl-nu-tel").val(), $("#lbl-nu-email").val(),
                parseInt($("#lbl-nu-priv").val()),
                function() {
                    $("#modal").modal('hide');
                    reportSuccess("Account created");
                    syncGetUsers(function (us) {
                        users = us;
                    });
                    showPrivileges();                    
                }, function(o) {
                    if (o.status == 409) {
                        $("#lbl-nu-email").notify("This email is already registered");
                    }
                });
}

function removeUser(email, div) {
    rmUser(email, function() {
        div.remove();
        reportSuccess("Account deleted")
    });
}
