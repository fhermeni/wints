var interns;
var currentPage;
var waitingBlock;
var myself;
var currentConvention;
var shiftPressed = false;
var allMajors;
$( document ).ready(function () {
waitingBlock = $("#cnt").clone().html();

    majors(function(m) {
        allMajors = m;
    })
	user(document.cookie.split("=")[1], function(u) {
        myself = u;
        $("#fullname").html(u.Firstname + " " + u.Lastname);
        showMyServices(u.Role);
    if (myself.Role >= 2) {
        showPage(undefined, "conventions");
    } else {
        showPage(undefined, "myStudents");
    }

	});	
	
    $(document).keydown(function (e) {
        if (e.keyCode == 16) {shiftPressed = true;}
    });
    $(document).keyup(function (e) {
        
        if (e.keyCode == 16) {shiftPressed = false;}
    });
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
        showPendingConventions();
    } else if (currentPage == "juries") {
        showJuryService();
    } else if (currentPage == "service") {
        showService();
    }
    else {
        console.log("Unsupported operation on '" + currentPage + "'");
    }
}

function displayMyStudents() {
internships(function(data) {  
    mine = data.filter(function (i) {
        return i.Tutor.Email == myself.Email
    });        
    var html = Handlebars.getTemplate("tutoring")(mine);
    var root = $("#cnt");
    root.html(html);
    if (mine.length == 0) {
        return true
    }       
    $("#table-conventions").tablesorter({            
        theme: 'bootstrap',
        widgets : ["uitheme"],
        headerTemplate : '{content} {icon}',
        headers : {
            0: {sorter:false}
        }
    });    
    $('#cnt').find(":checkbox").iCheck()
    $('#cnt').find(".check_all").on("ifChecked", function (e) {
        $("#cnt").find("td .icheckbox").iCheck("check")        
    }).on("ifUnchecked", function (e) {
        $("#cnt").find("td .icheckbox").iCheck("unCheck")        
    });
    $("#cnt").find("td .icheckbox").on("ifChecked", shiftSelect)    
});
}

function displayMyConventions() { 
internships(function(data) {   
    interns = data;    
    var html = Handlebars.getTemplate("watchlist")(interns);
    var root = $("#cnt");
    root.html(html);
	if (interns.length == 0) {
    	return true
    }       
    $("#table-conventions").tablesorter({            
        theme: 'bootstrap',
        widgets : ["uitheme"],
        headerTemplate : '{content} {icon}',
        headers : {
            0: {sorter:false}
        }
    });
    //$("#cnt").find(":checkbox").iCheck("uncheck");    
    $('#cnt').find(":checkbox").iCheck()
    $('#cnt').find(".check_all").on("ifChecked", function (e) {
        $("#cnt").find("td .icheckbox").iCheck("check")        
    }).on("ifUnchecked", function (e) {
        $("#cnt").find("td .icheckbox").iCheck("unCheck")        
    });
    $("#cnt").find("td .icheckbox").on("ifChecked", shiftSelect)    
});
}


function shiftSelect(e) {  
    if (shiftPressed) {        
        var myTd = $(e.currentTarget).closest("td")
        var tr = myTd.parent()
        var col = tr.children().index(myTd)        
        var p = tr.prev();
        while (p.length > 0) {
            var chk = $(p.children().get(col)).find(".icheckbox")            
            if (chk.hasClass("checked")) {
                break;
            } else {
                chk.iCheck("check")
            }
            p = p.prev();
        }
    }
}
    
function showPrivileges() {
    users(function(data) {
        var others = data.filter(function (u) {
            return u.Email != myself.Email && u.Role != 0;
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
                    refresh();                    
                }, function(o) {                	
                    if (o.status == 409) {
                        $("#lbl-nu-email").notify(o.responseText);
                    }
                });
}

function removeUser(email, div) {
    rmUser(email, function() {
        div.remove();
        reportSuccess("Account deleted")
    });
}


function pickBestMatching(tutor, users) {    
    var res = undefined;
    users.forEach(function (t) {
        var known_ln  = t.Lastname;
        if (tutor.Lastname.indexOf(known_ln) > -1 || known_ln.indexOf(tutor.Lastname) > -1) {
            res = t;
            return false;
        }
    });
    if (res == undefined) {
        var firstLetter = tutor.Lastname[0];
        users.forEach(function (t) {
            var knownFirstLetter = t.Lastname[0];
            if (firstLetter >= knownFirstLetter) {
                res =  t;
            }
        });
    }
    return res;
}

function showPendingConventions() {
    conventions(function(cc) {
        users(function (uss) {   
            uss = uss.filter(function (u) {
                return u.Role != 0;
            }); 
            var ignored = []
            cc = cc.filter(function (c) {
                if (c.Skip) {
                    ignored.push(c)
                }
                return !c.Skip
            }); 
            $("#pending-counter").html(" <span class='badge'>" + cc.length + (ignored.length > 0 ? "/" + ignored.length : "") + "</span>");
            if (cc.length == 0 && ignored.length == 0) {
                $("#pending-counter").html("");
                $("#cnt").html("<h5 class='text-center'>Nothing to import</h5>")
                return
            }               
            currentConvention = cc[0]     
            var html = Handlebars.getTemplate("pending")({
                c: currentConvention, users : uss, "Ignored": ignored
            });                        
            $("#cnt").html(html);    
            if (cc.length > 0) {        
            //Find the most appropriate predefined tutor            
            var best = pickBestMatching(currentConvention.Tutor, uss)                        
            $("#known-tutor-selector").val(best.Email)            
            $("#cnt").find("select").selecter();

            if (best.Lastname == currentConvention.Tutor.Lastname) {                
                d=$('#pick-theory');
                d.confirmation({onConfirm: pickTheory, placement: "right", title: "Sure ? It's a perfect match !", btnOkLabel: '<i class="icon-ok-sign icon-white"></i> Confirm'});
                d.removeAttr("onclick");
                k=$("#btn-choose-known");
                k.attr("onclick", "pickKnown()");
                k.confirmation('destroy');
            } else {
                d=$('#pick-theory');
                d.attr("onclick","pickTheory()");
                d.confirmation('destroy');

                k=$("#btn-choose-known");
                k.confirmation({onConfirm: pickKnown, placement: "left", title: "Sure ? They differ !", btnOkLabel: '<i class="icon-ok-sign icon-white"></i> Confirm'});
                k.removeAttr("onclick");
            }
        }
        });
    });
}

function sendSkipConvention(email, s) {
    skipConvention(email, s, function() {
        if (s) {
            reportSuccess("Convention ignored")    
        } else {
            reportSuccess("The convention is no longer ignored")
        }
        refresh()
    })
}
function pickTheory() {
     if (missing("th-tutor-fn") || missing("th-tutor-ln") || missing("th-tutor-email")  || missing("th-tutor-tel")) {
        return false;
    }
    var fn = $("#th-tutor-fn").val();
    var ln = $("#th-tutor-ln").val();
    var email = $("#th-tutor-email").val();
    var tel = $("#th-tutor-tel").val();
    newUser(fn, ln, tel, email, -1, function() {        
        currentConvention.Tutor = {Firstname: fn, Lastname: ln, Tel: tel, Email: email, Role: 1} //tutor     
        reportSuccess("Tutor account created")
        newInternship(currentConvention, function() {
            reportSuccess("internship added");       
            refresh()            
            })
    }, function(o) {        
        if (o.status == 409) {
            reportSuccess("Tutor account already exists.")
            newInternship(currentConvention, function() {
                reportSuccess("internship added");       
                refresh()            
            })
        } else {            
            reportError(jqr.responseText);
        }
    })    
}

function pickKnown() {         
    user($("#known-tutor-selector").val(), function(us) {        
        currentConvention.Tutor = us
        newInternship(currentConvention, function() {
            reportSuccess("internship added");       
            refresh()            
        })
    });    
}

function mailing(inCC) {
    var to = [];
    var cc = [];
    $(".icheckbox.checked").find(":checkbox").each(function(i, c) {
        var em = $(c).attr("data-email")                      
        to.push(em);        
        var i = getInternship(em)
        if (i && inCC == "tutor") {
            cc.push(i.Tutor.Email);               
        }                
    });
    if (to.length > 0) {
        window.location.href = "mailto:" + to.join(",") + (cc.length > 0 ? "?cc=" + cc.join(",") : "");
    }    
}

function getInternship(email) {
    var i = undefined;    
    interns.forEach(function (e) {        
        if (e.Student.Email == email) {
            i = e;
            return false;
        }
    })
    return i;
}

function showService() {
internships(function(interns) {
    var service = {};   
    interns.forEach(function (i) {
        if (!service[i.Tutor.Email]) {
            service[i.Tutor.Email] = {P : i.Tutor, Internships : [i]};
        } else {
            service[i.Tutor.Email].Internships.push(i)
        }        
    });
    var html = Handlebars.getTemplate("service")(service);
    $("#cnt").html(html);
    $('#cnt').find(":checkbox").iCheck()
    $('#cnt').find(".check_all").on("ifChecked", function (e) {
        $("#cnt").find("td .icheckbox").iCheck("check")        
    }).on("ifUnchecked", function (e) {
        $("#cnt").find("td .icheckbox").iCheck("unCheck")        
    });
    $("#cnt").find("td .icheckbox").on("ifChecked", shiftSelect)    

});    
}

function showRawService() {
internships(function(interns) {
    var service = {};   
    interns.forEach(function (i) {
        if (!service[i.Tutor.Email]) {
            service[i.Tutor.Email] = {P : i.Tutor, Internships : [i], Juries: {}};
        } else {
            service[i.Tutor.Email].Internships.push(i)
        }        
    });
    console.log(service)
    var html = Handlebars.getTemplate("raw-service")(service);
    $("#modal").html(html).modal('show');
});
}

function showRawFullname(root) {
    var checked = $(".icheckbox.checked").find(":checkbox")
        var fns = [];
        checked.each(function (i, e) {
            var em = $(e).attr("data-email");
            var i = getInternship(em);
            fns.push(i.Student.Firstname + " " + i.Student.Lastname);
        });
    if (fns.length > 0) {
        var html = Handlebars.getTemplate("raw");
        $("#modal").html(html);
        $("#rawContent").html(fns.join("\n"));
        $("#modal").modal('show');
    }
}

function selectText(elm) {
  // for Internet Explorer
  if(document.body.createTextRange) {
    var range = document.body.createTextRange();
    range.moveToElementText(elm);
    range.select();
  }
  else if(window.getSelection) {
    // other browsers
    var selection = window.getSelection();
    var range = document.createRange();
    range.selectNodeContents(elm);
    selection.removeAllRanges();
    selection.addRange(range);
  }
}

function showInternship(s) {
    internship(s, function (i) {
        users(function (uss) {
            uss = uss.filter(function (u) {
                return u.Role != 0; //get rid of students
            });             
            buf = Handlebars.getTemplate("student")({I: i, Admin: myself.Role >= 3, Tutors: uss})
            $("#modal").html(buf).modal('show');            
            var c = $("#modal").find("select.select-tutor");            
            c.val(i.Tutor.Email)
            $("#modal").find("select.select-major").selecter({callback: function(v) {sendMajor(i.Student.Email, v)}});
            $("#modal").find("select.select-tutor").selecter({callback: function(v) {sendTutor(i.Student.Email, v)}});
        });        
    });
}

function showReport(email, kind) {    
    reportHeader(email, kind, function(r) {            
            r.Reviewable = r.Grade != -2
            r.Email = email
            buf = Handlebars.getTemplate("reportEditor")(r)
            $("#modal").html(buf).modal('show');      
            $(':checkbox').iCheck()
                .on('ifChecked', function(){setReportPrivate(email, kind, true)})
                .on('ifUnchecked', function(){setReportPrivate(email, kind, false)})
            $(".date").datepicker({format:'d M yyyy', autoclose: true, minViewMode: 0, weekStart: 1}).on("changeDate", function (e) { setReportDeadline(email, kind, e.date)})             
    });
}

function sendMajor(e, m) {
    setMajor(e, m, function() {
        reportSuccess("Major updated")
        refresh()
    })
}

function sendTutor(e, s) {
    setTutor(e, s, function() {
        reportSuccess("Tutor updated")
        refresh();
    })
}