var active = undefined;
var teachers = []
var students = []
Array.prototype.remove = function(val) {
    var i = this.indexOf(val);
         return i>-1 ? this.splice(i, 1) : [];
};

var allDefenses = {}
var defenseSessions = []


function hash(s) {
    return s.Date.format("DDhha")+s.Room.replace("+","plus").replace("-","minus")
}

function getSession(sid) {
    var s = undefined;
    defenseSessions.forEach(function (ss) {
        if (hash(ss) == sid) {
            s = ss;
            return false;
        }
    });
    return s;
}

function sortByMajor(a, b) {
    if (a.Major < b.Major) return -1;
    if (a.Major > b.Major) return 1;
    return 0;
}

function showDefenses() {
    users(function(data) {
        teachers = [];        
        data.filter(function (u) {            
            if (myself.Email != u.Email && u.Role != 0) {
                teachers.push(u);
            }
        });
        students = interns.slice()
        students.sort(sortByMajor);        
        defenses(function (defs) {    
            console.log(defs)        
            var x = {Defenses : defs, Teachers: teachers, Students: students};        
            var html = Handlebars.getTemplate("defense-editor")(x);    
            var root = $("#cnt");
            root.html(html);
            root.find(".sortable").sortable().bind('sortupdate', updateSessionOrder);             

            //prepare the modal
            var html = Handlebars.getTemplate("defenseSessionEditor")({Room: "TBA"});    
            $('.date').datetimepicker();
            root.find(".date").datepicker({format:'d M yyyy', autoclose: true, minViewMode: 0, weekStart: 1}).on("changeDate", function (e) { setDefenseDay(id, e.date)})       
            var root = $("#modal");         
            root.html(html)
            load(defs)
        })
    });
}

function updateSessionOrder(e, li) {            
    var sid = $(li.item).closest(".session").attr("id")        
    s = getSession(sid)   
    //debugger 
    s.Defenses = []
    $(li.item).closest("ul").find("li").each(function(i, input) {
        em = $(input).find("input").data("email")                
        if (!em) {
            s.Pause = i
        } else {
            s.Defenses.push(em);
        }       
    } )
    //Get the emails
}

function addDefenseSession() {
    var r = $("#room").val();
    var d = $("#date").data("DateTimePicker").date()
    var s = {Pause: -1, Room : r, Date : d, Defenses : [], Juries : []};    
    if ($("#" + hash(s)).length > 0) {
        $("#room").closest(".form-group").addClass("has-error");
    } else {
        $("#room").closest(".form-group").removeClass("has-error");        
        $("#modal").modal('hide');
        drawSession(s);            
        defenseSessions.push(s);  
        activeSession(hash(s))          
    }
}

function rmSession() {    
    s = getSession(active);
    activeSession(active)
    $("#"+hash(s)).remove()
    defenseSessions.remove(s)    
}


function drawSession(s) {    
    var html = "<div id='" +  hash(s) + "' class='col-md-3 session'>"+
    "<div class='defense-panel panel'>"+ 
    "<div class='panel-heading' onclick='activeSession(\""+ hash(s) + "\")'>" +
    "<div class='panel-title'><i class='glyphicon glyphicon-calendar'> </i> <span class='date'>" + s.Date.format("D MMM - HH:mm") + "</span> &nbsp;"+                                            
    "<span class='where'><i class='glyphicon glyphicon-map-marker'> </i> <span class='room'>" + s.Room + "</span></span></div>"+                                            
    "</div>" +
    "<div class='panel-body'>" +                                                       
        "<label>Agenda</label> "+
    "<ul class='fn students sortable list-unstyled'></ul>"+
    "<label>Jury</label> "+
    "<ul class='fn juries list-unstyled'></ul>"+    
    "</div>"+
    "</div>";  
    var rId = s.Date.month() + "" + s.Date.date()
    var period = s.Date.hour() < 12 ? "am" : "pm";

    var day =  $("#" + rId);        
    if (day.length == 0) {     
        console.log("new day: " + rId);
        day = $("#days").append("<div id='" + rId + "' class='day'><div class='row am'></div><div class='row pm'></div></div>").children().last();
        console.log(day)
    }
    var row = day.find("."+period) 
    console.log(row)
    row.append(html)    
}

function saveDefenseSession() {
    var newRoom = $("#room").val();
    var newDate = $("#date").data("DateTimePicker").date();
    var newPeriod = newDate.hour() < 12 ? "am" : "pm";    
    var newDay = newDate.date();    

    var oldSession = getSession(active)
    var newSession = {Pause: -1, Room : newRoom, Date : newDate, Juries : oldSession.Juries.slice(), Defenses : oldSession.Defenses.slice()};
    if (hash(oldSession) != hash(newSession) && $("#" + hash(newSession)).length > 0) {
        $("#room").closest(".form-group").addClass("has-error");
        return
    }
    $("#room").closest(".form-group").removeClass("has-error");
    $("#modal").modal('hide');
    var oldPeriod = oldSession.Date.hour() < 12 ? "am" : "pm";    

    if (oldPeriod != newPeriod || newDay != oldSession.Date.date()) {
        rmSession(oldSession)
        drawSession(newSession)                
    } else {
        //Just the room change,         
        var d = $("#"+active);
        d.find(".room").html(newRoom)
        d.find(".date").html(newDate.format("D MMM - HH:mm"))
        d.attr("id", hash(newSession))
        d.attr("onclick", "activeSession(\""+ hash(newSession) + "\")")
    }              
    defenseSessions.push(newSession);        
    activeSession(hash(newSession))      
}

function showNewSession() {
        var root = $("#modal");          
        $("#room").val("TBA")
        root.find("button").attr("onclick","addDefenseSession()")
        d = new Date()
        if (defenseSessions.length == 0) {
            d.setMinutes(0)
        } else {                     
            d = defenseSessions[defenseSessions.length - 1].Date            
        }
        root.find('#date').datetimepicker({inline: true,sideBySide: true, format:"dd/mm/yyyy HH:mm"}).data("DateTimePicker").date(d);
        root.modal('show');        
}

function showEditSession() {
    s = getSession(active)        
    var root = $("#modal");          
    $("#room").val(s.Room)
    root.find("button").attr("onclick","saveDefenseSession()")
    root.find('#date').datetimepicker({inline: true,sideBySide: true, format:"dd/mm/yyyy HH:mm"}).data("DateTimePicker").date(s.Date)
    root.modal('show');        
}

function activeSession(sid) {    
    
    if (active) {
        $("#" + active).find(".panel").removeClass("panel-success")        
    }
    if (active != sid) {        
        active = sid
        $("#" + sid).find(".panel").addClass("panel-success");         
        $("#cnt").find(".activable").removeClass("disabled")    
        $('#jury-selecter').html(Handlebars.helpers.jurySelecter(availableTeachers(active)).string)            
    } else {
        $("#cnt").find(".activable").addClass("disabled")
        $('#jury-selecter').html("")
        active = undefined
    }
}

function removeByEmail(arr, em) {
    return arr.filter(function (u) {
                return u.Email != em;
            }); 
}

function addStudent(d) {
    if (!active) {
        return
    }
    var em = $("#student-selecter").val()
    var i = getInternship(em);    

    var def = {Grade: -1, Remote : false, Private: false}    
    allDefenses[em] = def;
    drawStudent(em)
    $("#cnt").find("ul.students").sortable({connectWith: "sortable"}).bind('sortupdate', updateSessionOrder);           
    students.remove(i)        
    $('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)        
    $('#student-selecter').selecter("destroy");
    $('#student-selecter').selecter("update");    
    s = getSession(active)
    s.Defenses.push(em)
}

function drawStudent(em) {    
    var i = getInternship(em); 
    var def = allDefenses[em];
    var glyphRemote = "glyphicon-picture"
    var glyphPrivate = "glyphicon-eye-open"
    if (def.Remote) {
        glyphRemote = "glyphicon-facetime-video"
    }   
    if (def.Private) {
        glyphPrivate = "glyphicon-eye-close"
    }
    var html = "<li>"+
        "<input type='checkbox' data-email='" + em + "'>"+
        " <a class='fn' onclick='showInternship(\"" + em + "\")'>" + Handlebars.helpers.abbrvFullname(i.Student) + " (" + i.Major + ")</a>"+
        " <i class='glyphicon " + glyphRemote + "' onclick='toggleRemote(this, \"" + em + "\")'></i> <i class='glyphicon " + glyphPrivate + "' onclick='togglePrivate(this,\"" + em + "\")'></i>"+
        " &nbsp; <i class='glyphicon glyphicon-remove-circle pull-right' onclick='removeStudent(this,\"" + em + "\")'></i>"+
        "</li>";            
    $("#"+active).find("ul.students").append(html).find(":checkbox").last().icheck()    
    
    var def = {Grade: -1, Remote : false, Private: false}    
    allDefenses[em] = def;
    s = getSession(active)    
}


function addPause() {
    //var html = "<li><i>pause</i> <i onclick='rmPause(this)' class='glyphicon glyphicon-remove-circle pull-right'></i></li>";    
    //$("#"+active).find("ul.students").append(html)
    s = getSession(active)
    s.Defenses.push()
    s.Pause = s.Defenses.length
    drawPause()    
}

function drawPause() {
    s = getSession(active)
    var html = "<li><i>pause</i> <i onclick='rmPause(this)' class='glyphicon glyphicon-remove-circle pull-right'></i></li>";    
    //$("#"+active).find("ul.students").append(html)    
    $("#"+active).find("ul.students li:nth-child(" + s .Pause + ")").after(html)
    //$("#"+active).find("ul.students:nth-child(" + (s.Pause) + ")").after(html);
    $("#"+active).find("ul.students").sortable().bind('sortupdate', updateSessionOrder);            
}

function rmPause(p) {
    $(p).closest("li").remove()
}

function toggleRemote(i, em) {
    allDefenses[em].Remote = !allDefenses[em].Remote;
    var j = $(i)
    if (j.hasClass("glyphicon-picture")) {
        j.removeClass("glyphicon-picture").addClass("glyphicon-facetime-video")
    } else {
        j.addClass("glyphicon-picture").removeClass("glyphicon-facetime-video")
    }
}

function togglePrivate(i, em) {
    allDefenses[em].Private = !allDefenses[em].Private;
    var j = $(i)
    if (j.hasClass("glyphicon-eye-open")) {
        j.removeClass("glyphicon-eye-open").addClass("glyphicon-eye-close")
    } else {
        j.addClass("glyphicon-eye-open").removeClass("glyphicon-eye-close")
    }
}

function removeStudent(b, em) {
    $(b).closest("li").remove();    
    s = getSession(active);
    s.Defenses.remove(em);
    i = getInternship(em)    
    students.push(i);
    students.sort(sortByMajor);    
    $('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)        
}


function overlap(d1, d2) {
    var day1 = d1.date();
    var day2 = d2.date();
    var p1 = d1.hour() < 12 ? "am" : "pm";
    var p2 = d2.hour() < 12 ? "am" : "pm";
    return day1 == day2 && p1 == p2;
}

function addJury(d) {
    if (!active) {
        return 
    }
    var em = $("#jury-selecter").val()
    drawJury(em)
    s = getSession(active)
    s.Juries.push(em);      
    $('#jury-selecter').find('option:selected', this).remove();                  
}

function drawJury(em) {
    teachers.forEach(function(t) {
        if (t.Email == em) {            
            var html = "<li>"+
            "<input type='checkbox' data-toggle='checkbox' class='icheckbox_flat check_all' data-email='" + em + "'/>"+
            " <a href='mailto:" + em + "'>" + Handlebars.helpers.abbrvFullname(t) + "</a>"+
            " <i class='glyphicon glyphicon-remove-circle pull-right' onclick='removeJury(this,\"" + em + "\")'></i>"+
            "</li>";
            $("#"+active).find("ul.juries").append(html).find(":checkbox").icheck();            
            
            return false;       
        }
    });
}

function removeJury(b, em) {    
    s = getSession(active);          
    s.Juries.remove(em);        
    $(b).closest("li").remove();    
    $('#jury-selecter').html(Handlebars.helpers.jurySelecter(availableTeachers(active)).string)        
}

function availableTeachers(sid) {            
    var s = getSession(sid);        
    var ok = teachers.slice();    
    defenseSessions.forEach(function (ss) {        
        if (overlap(ss.Date, s.Date)) {
            ss.Juries.forEach(function (t) {                            
                ok = removeByEmail(ok, t)                
            })
        }
    });      
    return ok  
}

function save() {
    var defs = [];
    defenseSessions.forEach(function (s) {                
        var cur = s.Date
        s.Defenses.forEach(function (em) {                            
            var date = new Date(cur.add(30,"minutes").toDate())            
            if (em) {
                x = allDefenses[em];            
                def = {Student: em, Room : s.Room, Date: date, Juries: [], Private: x.Private, Remote: x.Remote, Grade: x.Grade}               
                s.Juries.forEach(function (u) {
                    def.Juries.push({Email:u})
                })
                defs.push(def)
            }            
        });
    });     
    postDefenses(defs)    
}

function load(defs) {    
    allDefenses = {}
    defenseSessions = []
    defs.forEach(function (x) {
        //Which session   
        d = moment(new Date(x.Date))
        if (d.hour() < 13) {
            //Session ends late but in the morning
            d.hour(9)
        }               
        s = {Room : x.Room, Juries: [], Date : d, Students: []}
        sId = hash(s);        
        session = getSession(sId);
        if (!session) {
            console.log("new: " + sId)
            //console.log(s);
            x.Juries.forEach(function (j) {
                s.Juries.push(j.Email)
            });
            defenseSessions.push(s);
            session = s;              
            drawSession(session)          
            activeSession(hash(session))          
        }                
        def = {Date: moment(new Date(x.Date)), Remote: x.Remote, Private: x.Private, Grade: x.Grade}
        session.Students.push(x.Student)
        allDefenses[x.Student] = def        
        drawStudent(x.Student)
    })
    $("#cnt").find("ul.students").sortable({connectWith: "sortable"}).bind('sortupdate', updateSessionOrder);
    if (active) {
        $('#jury-selecter').html(Handlebars.helpers.jurySelecter(availableTeachers(active)).string)                    
        $('#student-selecter').html(Handlebars.helpers.studentSelecter(students).string)        
    }
}