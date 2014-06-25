/**
 * Created by fhermeni on 19/05/2014.
 */

var pre = {};
var survey;
var sid;
var token;
var lang;
var context;
$( document ).ready(function () {
    sid = $.url().param('s');
    token = $.url().param('t');
    $.getJSON("/static/survey/" + sid + ".json", function(d) {
        survey = d;
        loadLanguage();
        getContext();
    }).fail(function() {
        console.log("error");
    });
});

function fullname(p) {
    return "<a href='mailto:" + p.email + "'>" + p.firstname + " " + p.lastname + "</a>";
}
function getContext() {
    $.getJSON("/static/survey/" + token + ".json", function (c) {
        $("#title").html(survey.title[lang]);
        $("#student").html(fullname(c.student));
        $("#aTutor").html(fullname(c.aTutor));
        $("#cTutor").html(fullname(c.cTutor));
        context = c;
        if (c.answers) {
            $("#submit").css("display","none");
        }
        loadQuestions();
        $(':radio').radio().on('toggle', update_status);
        update_status();
    }).fail(function(xhr) {
        console.log("Error: " + xhr.responseText);
    });
}
function loadLanguage() {
    lang = Object.keys(survey.languages)[0];
    var buf = "";
    Object.keys(survey.languages).forEach(function (l) {
        buf += "<option value='" + l + "'>" + survey.languages[l] + "</option>";
    });
    $("select").selectpicker({style: 'btn-sm btn-primary navbar-btn', menuStyle: 'dropdown-inverse'});
    $("#supported_languages").html(buf).change(setLanguage);
}

function setLanguage() {
    lang = $("#supported_languages").val();
    $("#title").html(survey.title[lang]);
    $("#student_lbl").html(survey.student[lang]);
    $("#aTutor_lbl").html(survey.aTutor[lang]);
    $("#cTutor_lbl").html(survey.cTutor[lang]);
    survey.questions.forEach(function (q) {
        var id = q.id;
        $("#" + id + "_msg").html(q.txt[lang]);
        if (q.type == "bin") {
            $("#" + id + "_yes_msg").html(survey.yes[lang]);
            $("#" + id + "_no_msg").html(survey.no[lang]);
        }
    });
}

function loadQuestions() {
    var buf = "";
    survey.questions.forEach(function (q) {
        var k = q.id;
        var txt = q.txt[lang];
        if (q.type=="com") {
            buf += com(q);
        } else {
            buf += "<div id='" + k + "' class='form-group lang'>";
            if (q.type == "txt") {
                buf += "<label id='" + k + "_msg'>" + txt + "</label>";
                buf += text(q);
            } else if (q.type == "bin") {
                buf += "<label id='" + k + "_msg'>" + txt + "</label>";
                buf += bin(q);
            }
            buf += "</div>";
        }
        if (q.pre) {
            pre[k] = q.pre;
        }
    });
    console.log(pre);
    $("#survey").html(buf);
}

function text(q) {
    if (context.answers && context.answers[q.id]) {
        return "<br/><blockquote>" + context.answers[q.id] + "</blockquote>";
    }
    return "<textarea class='form-control' name='" + q.id + "' id='" + q.id + "_txt'></textarea>";
}

function bin(q) {
    if (context.answers) {
        if (context.answers[q.id] == 1) {
            return " " + survey.yes[lang];
        } else {
            return " " + survey.no[lang];
        }
    }
    var buf = "";
    buf += "<label class='radio'>";
    buf += "<input type='radio' data-toggle='radio' value='1' id='" + q.id + "_yes' name='" + q.id + "'/> ";
    buf += "<span id='" + q.id +"_yes_msg'>" + survey.yes[lang] + "</span>";
    buf += "</label>";
    buf += "<label class='radio'>";
    buf += "<input type='radio' value='0' data-toggle='radio' id='" + q.id + "_no' name='" + q.id + "'/> ";
    buf += "<span id='" + q.id +"_no_msg'>" + survey.no[lang] + "</span>";
    buf += "</label>";
    return buf;
}

function com(q){
    return "<h5 id='" + q.id + "_msg'>" + q.txt[lang] + "</h5>";
}

function update_status() {
    Object.keys(pre).forEach(function (k) {
        if (eval_pre(k)) {
            $("#" + k).show();
        } else {
            $("#" + k).hide();
        }
    });
}
function eval_pre(id) {
    for (var i in pre[id]) {
        console.log(pre[id]);
        var p = $("#" + pre[id][i]);
        if (!p.prop('checked')) {
            return false;
        }
    }
    return true;
}

function submit() {
    var ok = true;
    survey.questions.forEach(function (q) {
       if (q.required) {
           var v ;
           if (q.type == "bin") {
               v = $("input:radio[name=" + q.id + "]:checked").val();
           } else if (q.type == "txt") {
               v = $("#" + q.id + "_txt").val();
           }
           if (v == undefined) {
               $("#" + q.id + "_msg").css({"font-weight": "bold","color":"red"});
               ok = false;
           } else {
               $("#" + q.id + "").removeClass("has-error");
          }
       }
    });
    if (!ok) {
        $("#dlg").html(
            " <span class='err'><i class='glyphicon glyphicon-warning-sign err'></i> " + survey.missing[lang] + "</span>");
    } else {
        $("#dlg").html("OK");
    }
}