/**
 * Created by fhermeni on 06/08/2014.
 */
function flipForm(from, to) {
    $("#" + from).slideToggle(400, function() {
        $("#" + to).slideToggle();
    });
}

function resetPassword() {
    if ($("#email").val() == 0) {
        $("#email").notify("Required");
        return;
    }
    $.post("api/v1/password_reset", $("#email").val()).success(
        function() {
            $.notify(
                    "An email has been sent to " + $("#email").val() + " to indicate the procedure",
                {
                    className: "success",
                    position: "top center"
                }
            )
        })
        .fail(
        function() {
            console.log("no")
        }
        );
}
function foo() {
    console.log(arguments);
}