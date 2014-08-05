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
    m.lang("en");
    $("#defense-day").html(m.format("dddd D MMMM"));
    if (d == 0) {
        $("#prev-day").addClass("disabled");
    } else if (d == defenses.length - 1) {
        $("#next-day").addClass("disabled");
    } else {
        $("#next-day").removeClass("disabled");
        $("#prev-day").removeClass("disabled");
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