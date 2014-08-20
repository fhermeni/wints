/**
 * Created by fhermeni on 10/07/2014.
 */

var users;

var currentSession = 0;
var allChecked = false;

function showDefenses(t) {
    syncGetUsers(function (us) {
        users = us;
    });
    getEmbeddedDefenses(function(data) {
        defenses = JSON.parse(data);
        if (typeof(defenses.pool[0])=="string") {
            console.log("Conversion");
            defenses = embed2(defenses);
        } else {
            console.log("no need to convert");
        }
        console.log(defenses);
        //debugger;
        //console.log(embed2(defenses));
        if (t == "schedule") {
            showCoarseDefenseForm();
        } else {
            showDefensePlanningForm();
        }
    });
}

function nextSession(d) {
    var n;
    if (d.getHours() == 9) {
        n = new Date(d);
        n.setHours(14);
    }
    if (d.getHours() == 14) {
        n = new Date(d.getTime() + 86400000);
        n.setHours(9);
    }
    return n;
}

function storeDefenses() {
    console.log("saving the defenses");
    saveDefenses(defenses, publicDefenses(rmEmptyJuries(defenses)));
}

function saveSessionDate(i) {
    var id = $(i).attr("data-session");
    defenses.sessions[id].date = $(i).val();
    storeDefenses();
}

function saveRoom(i) {
    var idx = $(i).attr("data-jury");
    var r = $(i).val();
    defenses.sessions.forEach(function (s) {
        s.jury.forEach(function (x) {
            if (x.id == idx) {
                x.room = r;
                return false;
            }
        })
    });
    storeDefenses();
}

function saveTime(i) {
    var idx = $(i).attr("data-jury");
    var r = $(i).val();
    defenses.sessions.forEach(function (s) {
        s.jury.forEach(function (x) {
            if (x.id == idx) {
                x.date = r;
                return false;
            }
        })
    });
    storeDefenses();
}

function saveJury(i) {
    var idx = $(i).attr("data-jury");
    var jidx = $(i).attr("data-jury-idx");
    var m = $(i).val();
    defenses.sessions.forEach(function (s) {
        s.jury.forEach(function (x) {
            if (x.id == idx) {
                x.commission[jidx] = m;
                console.log(x);
                return true;
            }
        })
    });
    //Set the checkbox value accordingly
    storeDefenses();
}


function saveLists() {
    var raw = [];
    var pool = [];
    $(".students").each(function (idx, l) {
        var id = $(l).attr("data-jury");
        $(l).find("li").each(function (i, s) {
            var e = $(s).attr("data-email");
            if (id == -1) {
              //  console.log(e);
                pool.push(e);
            } else {
                if (!raw[id]) {
                    raw[id] = [];
                }
                raw[id].push(e);
            }
        });
    });
    defenses.sessions.forEach(function (s){
        s.jury.forEach(function (j) {
            j.students=raw[j.id];
        });
    });
    defenses.pool = pool;
    storeDefenses();
}

function prepareSchedule() {
    debugger;
    var nbSlots = $("#lbl-nbSlots").val();
    if (nbSlots == 0) {
        return;
    }
    defenses = {
        filter : $("#lbl-filter").val(),
        pool: [],
        tutors : extractTutors(conventions)
    };

    var byMajor = groupByMajor(filterConventions());
    var groups = [];
    var group = [];
    groups.push(group);
    Object.keys(byMajor).forEach(function (m) {
        var cc = byMajor[m];
        cc.forEach(function (c) {
            group.push(c.Stu.P.Email);
            if (group.length == 5) {
                group = [];
                groups.push(group);
            }
        });
    });

    var nbInParallel = Math.ceil(groups.length / nbSlots);

    var schedule = [];
    for (var t = 0; t < nbSlots; t++) {
        if (!schedule[t]) {
            schedule[t] = [];
        }
        for (var x = 0; x < nbInParallel; x++) {
            schedule[t][i] = [];
        }
    }
    for (var i = 0; i < groups.length; ) {
        for (var t = 0; t < nbSlots; t++) {
            //for each slot;
            if (schedule[t].length < nbInParallel) {
                schedule[t].push(groups[i++]);
                if (i == groups.length) {
                    break;
                }

            }
        }
    }


    defenses["sessions"] = [];
    var d = new Date();
    d = new Date(d.getTime() + 86400000);
    d.setMinutes(0);
    d.setSeconds(0);
    d.setHours(9);                                         //tomorrow 9am
    var idx = 0;
    for (var i = 0; i < nbSlots; i++) {
        var session = {
            date : moment(d).format("DD/MM/YYYY HH:mm"),
            jury : []
        };
        schedule[i].forEach(function (s) {
            //Middle of each students, we add a break
            s.splice((s.length + 1)/2, 0, undefined);
            var j = {
                room: "?",
                commission: ["","",""],
                students : s,
                id: idx
            };
            session.jury.push(j);
            idx++;
        });
        while (session.jury.length < nbInParallel) {
            session.jury.push({
                room : "?",
                commission: ["","",""],
                students: [undefined],
                id: idx++
            })
        }
        d = nextSession(d);
        defenses["sessions"].push(session);
    }
    //Padding of the jurys
    showCoarseDefenseForm();
}

function sortableOptions() {
    return {
        group:'students',
        pullPlaceholder: false,
        onDrop: function ($item, container, _super, event) {
            $item.removeClass("dragged").removeAttr("style");
            $("body").removeClass("dragging");
            saveLists();
        },
        // set item relative to cursor position
        onDragStart: function ($item, container, _super) {
            var offset = $item.offset(),
                pointer = container.rootGroup.pointer;

            adjustment = {
                left: pointer.left - offset.left,
                top: pointer.top - offset.top
            };

            _super($item, container)
        },
        onDrag: function ($item, position) {
            $item.css({
                left: position.left - adjustment.left,
                top: position.top - adjustment.top
            })
        }
    }
}

function showCoarseDefenseForm() {
    var html = Handlebars.getTemplate("defenses-coarse")(defenses);
    $("#cnt").html(html);
    $(".students").sortable(sortableOptions());
}

function generateJuryMailto(root) {
    var emails = [];
    var checked = root.find(".mail-checkbox-students.checked");
    checked.each(function (i, e) {
            emails.push($(e).attr("data-email"));
    });
    var checked = root.find("select").each(function (i, s) {
        var x =$(s);
        //Associated checkbox checked ?
        var id = $(s).attr("data-for");
        if ($("#"+ id).hasClass("checked")) {
            var v  = x.val();
            emails.push(v);
        }
    });
    if (emails.length > 0) {
        root.find(".mail-selection").attr("href","mailto:" + emails.join(","));
    } else {
        root.find(".mail-selection").attr("href","#");
    }
}

function showDefensePlanningForm() {
    getUsers(function (users) {
        defenses.tutors = [];
        users.forEach(function (u) {
            defenses.tutors.push(u);
        });
        var html = Handlebars.getTemplate("defenses-planning2")(rmEmptyJuries(defenses));
        var root = $("#cnt");
        root.html(html);
        root.find(':checkbox').checkbox().on('toggle', function (e) {
            generateJuryMailto(root);
        });
        show_session(0);
    });
}

function displayDefenses(i) {
    var html = Handlebars.getTemplate("defenses-show-" + i)(rmEmptyJuries(defenses));
    window.open( "data:x-application/external;charset=utf-8," + escape(html));
}


function rmEmptyJuries(defenses) {
    var d = {
        sessions: [],
        private: defenses.private,
        visio : defenses.visio,
        tutors : defenses.tutors
    };
    var i = 0;
    defenses.sessions.forEach(function (s) {
        cs = {
            date : s.date,
            jury : []
        };
        d.sessions.push(cs);
        s.jury.forEach(function (j) {
            var empty = true;
            j.students.forEach(function (s) {
                if (s != undefined && s != null && s != "") {
                    empty = false;
                }
            });
            if (!empty) {
                cs.jury.push(j);
            }
        });
    });
    return d;
}

//defenses stuff
function show_session(nb) {

    $(".session-header").each(function (idx, d) {
        if (nb  == idx) {
            $(d).addClass("active");
        }
        else {
            $(d).removeClass("active");
        }
    });
    $(".session-content").each(function (idx, d) {
        if (nb  == idx) {
            $(d).show();
        }
        else {
            $(d).hide()
        }
    });
    if (nb < $(".session-header").length - 1) {
        $(".pagination .next a").attr("onclick", "show_session(" + (nb + 1) + ")");
    } else {
        $(".pagination .next a").removeAttr("onclick");
    }
    if (nb != 0) {
        $(".pagination .previous a").attr("onclick", "show_session(" + (nb - 1) + ")");
    } else {
        $(".pagination .previous a").removeAttr("onclick");
    }
    $(".mail-selection").attr("href","#");
    $(':checkbox').checkbox('uncheck');
}

function switchVisibility(i) {
    var j = $(i);
    var mail = j.parent().attr("data-email");
    if (j.hasClass("glyphicon-eye-open")) {
        j.removeClass("glyphicon-eye-open").addClass("glyphicon-eye-close late");
        defenses.private[mail] = true;
    } else {
        j.removeClass("glyphicon-eye-close").removeClass("late").addClass("glyphicon-eye-open");
        defenses.private[mail] = false;
    }
}

function switchVisio(i) {
    var j = $(i);
    var mail = j.parent().attr("data-email");
    if (j.hasClass("glyphicon-facetime-video")) {
        j.removeClass("glyphicon-facetime-video").addClass("glyphicon-user");
        defenses.visio[mail] = false;
    } else {
        j.removeClass("glyphicon-user").addClass("glyphicon-facetime-video");
        defenses.visio[mail] = true;
    }
}

function publicDefenses(d) {
    var long = {};
    long.sessions = [];
    d.sessions.forEach(function (session) {
        console.log(session);
        var newSession = {date: session.date, jury:[]};
        long.sessions.push(newSession);
        session.jury.forEach(function (j) {
            var myJury = {
                students : [],
                commission : j.commission,
                room : j.room,
                id : j.id,
                time : j.date
            };

            newSession.jury.push(myJury);
            j.students.forEach(function (stu) {
                var c = getConvention(stu);
                if (c) {
                    //get its fn,ln
                    var myStudent = {
                        P: {
                            Firstname: c.Stu.P.Firstname,
                            Lastname: c.Stu.P.Lastname
                        },
                        Major: c.Stu.Major,
                        Promotion: c.Stu.Promotion,
                        Subject: c.Title,
                        CompanyWWW: c.CompanyWWW,
                        Company: c.Company,
                        private : d.private[stu],
                        visio: d.visio[stu]

                    };
                    myJury.students.push(myStudent);
                } else {
                    myJury.students.push({});
                }
            });
        });
    });

    var byDay = [];
    for (i = 0; i < long.sessions.length; i+=2) {
        byDay.push({
            date: long.sessions[i].date,
            period : [long.sessions[i], long.sessions[i+1]]
        });
    }
    return byDay;
}

function showJuryService() {
    getEmbeddedDefenses(function (data) {
        syncGetUsers(function (us) {
            users = us;
        });
        var d = JSON.parse(data);
        var html = Handlebars.getTemplate("juries")(jury_service(rmEmptyJuries(d)));
        var root = $("#cnt");
        root.html(html);
        root.find(':checkbox').checkbox();
        root.find('tbody').find(':checkbox').checkbox().on('toggle', function (e) {
            generateMailto(root);
        });
        root.find('.mail-checkbox').click(function (e) {
            shiftSelect(e, this, root,'.mail-checkbox-j');
        });

        root.find(".mailto").on('toggle', function (e) {
            toggleMailCheckboxes(e.currentTarget, ".mail-checkbox-jury", root);
        });

    });
}

function rawJuries() {
    getEmbeddedDefenses(function (data) {
        syncGetUsers(function (us) {
            users = us;
        });
        var d = JSON.parse(data);
        var txt = Handlebars.getTemplate("rawJuries")(jury_service(rmEmptyJuries(d)));
        $("#modal").html(txt).modal('show');
    });
}


function jury_service(defenses) {
    var count = {};
    defenses.sessions.forEach(function (s) {
        s.jury.forEach(function (j) {
            j.commission.forEach(function (e) {
                if (e.length > 1) {
                    if (!count[e]) {
                        count[e] = 1;
                    } else {
                        count[e]++;
                    }
                }
            })
        });
    });
    arr = [];
    Object.keys(count).forEach(function (e) {
            users.forEach(function (u) {
                if (u.Email == e) {
                    arr.push({user: u, count: count[e]});
                    return true;
                }
            })
        }
    );
    return arr;
}

/**
 * Substitute jury fullname by their email
 * @param def
 */
function embed2(def) {
   var d = {
       filter : def.filter,
       pool : [],
       tutors : def.tutors,
       private : def.private,
       visio: def.visio,
       sessions : []
   };
   def.pool.forEach(function (e) {
       d.pool.push(getConvention(e).Stu);
   });
   def.sessions.forEach(function (s) {
     var ns = {
         date: s.date,
         jury: []
     };
      d.sessions.push(ns);
       s.jury.forEach(function(j) {
           var nj = {
               room : j.room,
               students : [],
               id : j.id,
               date : j.date,
               commission : []
           };
           j.students.forEach(function (u) {
               if (u && u.length > 1) {
                   var c = getConvention(u);
                   nj.students.push(c.Stu);
               }
           });
           j.commission.forEach(function (fn) {
               if (fn.length > 1) {
                   nj.commission.push(fullname2Email(fn));
               } else {
                   nj.commission.push("");
               }
           });
           ns.jury.push(nj);
       });
   });
   return d;
}

function fullname2Email(fn) {
    var got = undefined;
    users.forEach(function (u) {
            if (fn == u.Firstname + " " + u.Lastname) {
                got = u.Email;
                return true;
            }
        });
    return got;
}