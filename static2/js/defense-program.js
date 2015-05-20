$( document ).ready(function () {
    defenseProgram(function(sessions) {  
        for (var i = 0; i < sessions.length; i++) {
            sessions[i].Date = new Date(sessions[i].Date)            
        }                      
        var html = Handlebars.getTemplate("defense-program")(sessions);    
        var root = $("#cnt");
        root.html(html);    
    })    
});