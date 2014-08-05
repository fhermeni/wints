/**
 * Created by fhermeni on 10/07/2014.
 */

var currentDay = 0;

function nextDay() {
    if (currentDay < defenses.length - 1) {
        $("#day-" + currentDay).toggle();
        showDay(++currentDay)
    }
}
function prevDay() {
    if (currentDay > 0) {
        $("#day-" + currentDay).toggle();
        showDay(--currentDay)
    }
}

function showDay(d) {
    //console.log("show day " + d);
    var m = moment(defenses[d].date, "DD/MM/YYYY");
    m.lang("fr");
    $("#defense-day").html(m.format("dddd D MMMM"));

    if (d != 0) {
        $("#prev-day").find("a").html("&larr; " + defenses[d - 1].date);
    }
    if (d != defenses.length - 1) {
        $("#next-day").find("a").html(defenses[d + 1].date + " &rarr;");
    }
    if (d == 0) {
        $("#prev-day").hide();
    } else if (d == defenses.length - 1) {
        $("#next-day").hide();
    } else {
        $("#next-day").show();
        $("#prev-day").show();
    }
    $("#day-" + d).toggle();
}

function showDefenses() {
    $.ajax({
        method: "GET",
        url: "/api/v1/defenses?fmt=public"
    }).done(function(data) {
            defenses = JSON.parse(data);
            var html = Handlebars.getTemplate("defenses-show-planning2")(defenses);
            $("#defenses").html(html);
            showDay(0);
        }
    ).fail(function (data) {console.log(data)});
}


$(document).ready(function () {
   showDefenses();
});