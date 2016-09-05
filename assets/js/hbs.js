this["wints"] = this["wints"] || {};
this["wints"]["templates"] = this["wints"]["templates"] || {};
this["wints"]["templates"]["alumni-modal"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "Permanent (CDI)";
},"3":function(container,depth0,helpers,partials,data) {
    return "Fixed (CDD)";
},"5":function(container,depth0,helpers,partials,data) {
    return "internship company";
},"7":function(container,depth0,helpers,partials,data) {
    return "other company";
},"9":function(container,depth0,helpers,partials,data) {
    return "France";
},"11":function(container,depth0,helpers,partials,data) {
    return "foreign country";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n        <div class=\"modal-body\">                    \n                <fieldset>\n                <legend>The future of "
    + alias3((helpers.fullname || (depth0 && depth0.fullname) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "</legend>\n                <div class=\"form-horizontal\">\n                <div class=\"form-group\">\n                    <label class=\"col-sm-3 control-label\">Email</label>                    \n                    <div class=\"col-sm-9\">\n                        <label class=\"form-control-static\">"
    + alias3(container.lambda(((stack1 = (depth0 != null ? depth0.Alumni : depth0)) != null ? stack1.Contact : stack1), depth0))
    + "</label>\n                    </div>\n                </div>\n                 \n                <div class=\"form-group\">\n                    <label class=\"col-sm-3 control-label\">Position</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"form-control-static\">"
    + alias3((helpers.alumniPosition || (depth0 && depth0.alumniPosition) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Alumni : depth0)) != null ? stack1.Position : stack1),{"name":"alumniPosition","hash":{},"data":data}))
    + "</label>\n                    </div>\n                </div>\n\n                <div class=\"form-group\" id=\"contract\">\n                    <label class=\"col-sm-3 control-label\">Contract</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"form-control-static\">\n                        "
    + ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Alumni : depth0)) != null ? stack1.Permanent : stack1),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "                        \n                        </label>\n                    </div>                    \n                </div>   \n\n                <div class=\"form-group\" id=\"company\">\n                    <label class=\"col-sm-3 control-label\">Company</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"form-control-static\">\n                        "
    + ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Alumni : depth0)) != null ? stack1.SameCompany : stack1),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.program(7, data, 0),"data":data})) != null ? stack1 : "")
    + "                                            \n                        </label>\n                    </div>\n                </div> \n\n                <div class=\"form-group\" id=\"country\">\n                    <label class=\"col-sm-3 control-label\">Country</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"form-control-static\">\n                        "
    + ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Alumni : depth0)) != null ? stack1.France : stack1),{"name":"if","hash":{},"fn":container.program(9, data, 0),"inverse":container.program(11, data, 0),"data":data})) != null ? stack1 : "")
    + "                                            \n                        </label>\n                    </div>\n                </div>          \n                </div>                                  \n                </fieldset>\n            <div class=\"text-right form-group\">\n                <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Close</button>            \n            </div>                            \n        </div>\n    </div>\n</div>   ";
},"useData":true});
this["wints"]["templates"]["company-editor"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n        <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n            <div class=\"form-horizontal\">\n                <fieldset>\n                    <legend>Company</legend>\n                <div class=\"form-group\">\n                    <label for=\"lbl-name\" class=\"col-sm-3 control-label\">Name</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"lbl-name\" value=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "\"/>\n                        </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"lbl-www\" class=\"col-sm-3 control-label\">Website</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"lbl-www\" value=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\"/>\n                        </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"lbl-title\" class=\"col-sm-3 control-label\">Subject title</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"lbl-title\" value=\""
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Company : stack1)) != null ? stack1.Title : stack1), depth0))
    + "\"/>\n                        </div>\n                </div>\n                </fieldset>\n            </div>\n            <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"updateCompany()\">Update</button>\n            </div>\n        </div>\n    </div>\n</div>";
},"useData":true});
this["wints"]["templates"]["convention-detail"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "        <div class=\"form-group\">\n            <label class=\"col-lg-3 control-label\">Promotion</label>\n            <div class=\"col-lg-3\">\n                <select class=\"form-control\" id=\"promotion-selecter\" onchange=\"updatePromotion('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "','"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1), depth0))
    + "', this)\">\n                "
    + alias2((helpers.optionPromotions || (depth0 && depth0.optionPromotions) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1),{"name":"optionPromotions","hash":{},"data":data}))
    + "\n                </select>\n            </div>\n        </div>\n        <div class=\"form-group\">\n            <label class=\"col-lg-3 control-label\">Major</label>\n            <div class=\"col-lg-3\">\n                <select class=\"form-control\" id=\"major-selecter\" onchange=\"updateMajor('"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "','"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1), depth0))
    + "', this)\">\n                "
    + alias2((helpers.optionMajors || (depth0 && depth0.optionMajors) || alias4).call(alias3,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1),{"name":"optionMajors","hash":{},"data":data}))
    + "\n                </select>\n            </div>\n        </div>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "     <div class=\"form-group\">\n        <label class=\"col-lg-3 control-label\">Major/Promotion</label>\n        <div class=\"col-lg-9\">\n         <label class=\"form-control-static\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1), depth0))
    + "/"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1), depth0))
    + "</label>\n        </div>\n     </div>\n";
},"5":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "		<div class=\"form-group\">\n			<label class=\"col-lg-3 control-label\">Company</label>\n                <div class=\"col-lg-9\">\n                	<label class=\"form-control-static\">\n                        <a href=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Company : stack1)) != null ? stack1.WWW : stack1), depth0))
    + "\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Company : stack1)) != null ? stack1.Name : stack1), depth0))
    + "</a>\n                    </label>\n                </div>\n		</div>\n		<div class=\"form-group\">\n			<label class=\"col-lg-3 control-label\">Title</label>\n            <div class=\"col-lg-9\">\n                <label class=\"form-control-static text-justify\">"
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Company : stack1)) != null ? stack1.Title : stack1), depth0))
    + "</label>\n            </div>\n		</div>\n		<div class=\"form-group\">\n			<label class=\"col-lg-3 control-label\">Period</label>\n                <div class=\"col-lg-9\">\n                	<label class=\"form-control-static\">\n                	"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Begin : stack1),"ddd DD MMM YY",{"name":"dateFmt","hash":{},"data":data}))
    + "\n                	-\n                	"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.End : stack1),"ddd DD MMM YY",{"name":"dateFmt","hash":{},"data":data}))
    + "\n                	</label>\n                </div>\n		</div>\n		<div class=\"form-group\">\n			<label class=\"col-lg-3 control-label\">Gratification</label>\n                <div class=\"col-lg-9\">\n                	<label class=\"form-control-static\">"
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Gratification : stack1), depth0))
    + " €</label>\n                </div>\n		</div>\n\n        <div class=\"form-group\" id=\"supervisor-group\">\n            <label class=\"col-lg-3 control-label\">Supervisor</label>\n                <div class=\"col-lg-9\">\n                    <label class=\"form-control-static fn\">\n"
    + ((stack1 = container.invokePartial(partials.person,((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Supervisor : stack1),{"name":"person","data":data,"indent":"                    ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "                    </label>\n                </div>\n        </div>\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.escapeExpression, alias2=container.lambda;

  return "        <div class=\"form-group\">\n        <label class=\"col-lg-3 control-label\">Switch to</label>\n            <div class=\"col-lg-9\">\n                <div class=\"input-group\">\n                <select class=\"fn form-control\" id=\"tutor-selecter\">\n                    "
    + alias1((helpers.optionUsers || (depth0 && depth0.optionUsers) || helpers.helperMissing).call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Teachers : depth0),((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Tutor : stack1),{"name":"optionUsers","hash":{},"data":data}))
    + "\n                </select>\n                <span class=\"input-group-btn\">\n                <button type=\"button\" class=\"btn btn-warning\" data-placement='right' data-toggle=\"confirmation\" data-on-confirm=\"switchTutor('"
    + alias1(alias2(((stack1 = ((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "', '"
    + alias1(alias2(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n                    <i class=\"glyphicon glyphicon-random\"></i>\n                </button>\n                </span>\n                </div>\n            </div>\n        </div>\n";
},"9":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "        <div class=\"form-group\">\n        <label class=\"col-lg-3 control-label\">Reset surveys</label>\n        <div class=\"col-lg-9\">\n"
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Surveys : stack1),{"name":"each","hash":{},"fn":container.program(10, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        </div>\n        </div>\n\n        <div class=\"form-group\">\n        <label class=\"col-lg-3 control-label\">Request surveys</label>\n        <div class=\"col-lg-9\">\n"
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Surveys : stack1),{"name":"each","hash":{},"fn":container.program(17, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        </div>\n        </div>\n";
},"10":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=container.escapeExpression, alias3=helpers.helperMissing, alias4="function";

  return "                    <button type=\"button\" class=\"btn "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Delivery : depth0),{"name":"if","hash":{},"fn":container.program(11, data, 0, blockParams, depths),"inverse":container.program(13, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "")
    + " btn-sm\" data-placement='bottom' data-toggle=\"confirmation\" data-on-confirm=\"resetSurvey(this,'"
    + alias2(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].I : depths[1])) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "','"
    + alias2(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "')\" "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Delivery : depth0),{"name":"unless","hash":{},"fn":container.program(15, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n                        "
    + alias2(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "\n                    </button>\n";
},"11":function(container,depth0,helpers,partials,data) {
    return "btn-danger";
},"13":function(container,depth0,helpers,partials,data) {
    return "btn-default";
},"15":function(container,depth0,helpers,partials,data) {
    return "disabled title=\"not uploaded\"";
},"17":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1;

  return ((stack1 = helpers["if"].call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Delivery : depth0),{"name":"if","hash":{},"fn":container.program(18, data, 0, blockParams, depths),"inverse":container.program(20, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "");
},"18":function(container,depth0,helpers,partials,data) {
    var helper;

  return "                    <button type=\"button\" class=\"btn btn-default btn-sm\" disabled title=\"already uploaded\">"
    + container.escapeExpression(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Kind","hash":{},"data":data}) : helper)))
    + "</button>\n";
},"20":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=container.escapeExpression, alias2=depth0 != null ? depth0 : {}, alias3=helpers.helperMissing, alias4="function";

  return "                    <button type=\"button\" class=\"btn btn-danger btn-sm\" data-placement='bottom' data-toggle=\"confirmation\" data-on-confirm=\"requestSurvey(this,'"
    + alias1(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = ((stack1 = (depths[1] != null ? depths[1].I : depths[1])) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "','"
    + alias1(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias2,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "')\">\n                        "
    + alias1(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias2,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "\n                        (<i class=\"glyphicon glyphicon-time\"></i> "
    + alias1((helpers.daysSince || (depth0 && depth0.daysSince) || alias3).call(alias2,(depth0 != null ? depth0.LastInvitation : depth0),{"name":"daysSince","hash":{},"data":data}))
    + "  days.)\n                    </button>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\">\n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <fieldset>\n            <legend>\n"
    + ((stack1 = container.invokePartial(partials.person,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"person","data":data,"indent":"            ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "            </legend>\n    <div class=\"form-horizontal\">\n    <div class=\"form-group\">\n            <label class=\"col-lg-3 control-label\">Last visit</label>\n            <div class=\"col-lg-3\">\n                <label class=\"form-control-static\">"
    + container.escapeExpression((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.LastVisit : stack1),"DD/MM/YY HH:mm",{"name":"dateFmt","hash":{},"data":data}))
    + "</label>\n            </div>\n    </div>\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Editable : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.program(3, data, 0, blockParams, depths),"data":data})) != null ? stack1 : "")
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Editable : depth0),{"name":"unless","hash":{},"fn":container.program(5, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "		<div class=\"form-group\" id=\"tutor-group\">\n			<label class=\"col-lg-3 control-label\">Academic tutor</label>\n                <div class=\"col-lg-9\">\n                	<label class=\"form-control-static fn\">\n"
    + ((stack1 = container.invokePartial(partials.person,((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.I : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1),{"name":"person","data":data,"indent":"                    ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "                    </label>\n                </div>\n		</div>\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Editable : depth0),{"name":"if","hash":{},"fn":container.program(7, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n"
    + ((stack1 = (helpers.ifRole || (depth0 && depth0.ifRole) || alias2).call(alias1,4,{"name":"ifRole","hash":{},"fn":container.program(9, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</div>\n    </fieldset>\n    <div class=\"text-right\">\n    <button type=\"button\" class=\"btn btn-default\" aria-hidden=\"true\" onclick=\"hideModal()\">Close</button></div>\n        </div>\n        </div>\n        </div>";
},"usePartial":true,"useData":true,"useDepths":true});
this["wints"]["templates"]["convention-validator"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "				<option value=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(depth0 != null ? depth0 : {},((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + " ("
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + ")</option>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {};

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n<fieldset>\n	<legend class=\"fn\" data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.S : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">					\n		"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.S : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + " - "
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.S : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + "/"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.S : depth0)) != null ? stack1.Major : stack1), depth0))
    + "\n	</legend>	\n	<div class=\"alert alert-dismissible alert-danger hidden text-center\">\n               <button type=\"button\" class=\"close\" data-dismiss=\"alert\">×</button>\n               <p></p>\n	</div>\n\n	<div class=\"form-horizontal\">	\n	<div class=\"form-group\">\n		<label class=\"col-lg-3 control-label\">Available conventions</label>\n		<div class=\"col-lg-9\">\n		<select class=\"fn form-control\" id=\"convention-selecter\" onchange=\"showConvention(this)\">\n"
    + ((stack1 = helpers.each.call(alias3,(depth0 != null ? depth0.Conventions : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "		</select>\n		</div>\n	</div>		\n	</div>\n	<div id=\"convention-detail\">\n"
    + ((stack1 = container.invokePartial(partials["convention-editor"],depth0,{"name":"convention-editor","data":data,"indent":"\t\t","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "	</div>\n</fieldset>\n</div>\n</div>\n</div>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["conventions-header"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["conventions-convention"],depth0,{"name":"conventions-convention","data":data,"indent":"    ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "    <h3><span class=\"count\">"
    + container.escapeExpression((helpers.len || (depth0 && depth0.len) || helpers.helperMissing).call(alias1,depth0,{"name":"len","hash":{},"data":data}))
    + "</span> convention(s)</h3>\n\n<table class=\"table table-hover tablesorter\" id=\"table-conventions\" data-sortlist=\"[[1,0]]\">\n<thead>\n    <tr>\n        <th class=\"sorter-false\"></th>\n        <th>Student</th>\n        <th>Promotion/Major</th>\n        <th>Period</th>        \n        <th>Tutor</th>\n        <th>Gratification</th>\n        <th>Creation</th>   \n        <th>Validated</th>\n        <th>Skip</th>             \n    </tr>\n</thead>\n<tbody>\n"
    + ((stack1 = helpers.each.call(alias1,depth0,{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</tbody>\n</table>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["defense-editor"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.escapeExpression, alias2=container.lambda;

  return "        <p class=\"form-control-static\">"
    + alias1((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(depth0 != null ? depth0 : {},((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + " ("
    + alias1(alias2(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + "/"
    + alias1(alias2(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + ")</p>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "        <select class=\"fn form-control\" id=\"student-selecter\">\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Students : depth0),{"name":"each","hash":{},"fn":container.program(4, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        </select>\n";
},"4":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.escapeExpression, alias2=depth0 != null ? depth0 : {}, alias3=helpers.helperMissing, alias4="function";

  return "            <option value=\""
    + alias1(container.lambda(((stack1 = ((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias1((helpers.fullname || (depth0 && depth0.fullname) || alias3).call(alias2,((stack1 = (depth0 != null ? depth0.User : depth0)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + " "
    + alias1(((helper = (helper = helpers.Promotion || (depth0 != null ? depth0.Promotion : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias2,{"name":"Promotion","hash":{},"data":data}) : helper)))
    + "/"
    + alias1(((helper = (helper = helpers.Major || (depth0 != null ? depth0.Major : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias2,{"name":"Major","hash":{},"data":data}) : helper)))
    + "</option>\n";
},"6":function(container,depth0,helpers,partials,data) {
    return "checked";
},"8":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.escapeExpression, alias2=depth0 != null ? depth0 : {}, alias3=helpers.helperMissing, alias4="function";

  return "    <div class=\"pull-left\">\n        <button type=\"button\" class=\"btn btn-danger\" onclick=\"delStudentDefense('"
    + alias1(container.lambda(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "', '"
    + alias1(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias2,{"name":"Room","hash":{},"data":data}) : helper)))
    + "', '"
    + alias1(((helper = (helper = helpers.SessionId || (depth0 != null ? depth0.SessionId : depth0)) != null ? helper : alias3),(typeof helper === alias4 ? helper.call(alias2,{"name":"SessionId","hash":{},"data":data}) : helper)))
    + "')\">Delete</button>\n    </div>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\">\n\n<div class=\"modal-header\">\n    <h4 class=\"modal-title\">\n        <i class=\"glyphicon glyphicon-calendar\"></i> "
    + alias4(((helper = (helper = helpers.SessionId || (depth0 != null ? depth0.SessionId : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"SessionId","hash":{},"data":data}) : helper)))
    + ", <i class=\"glyphicon glyphicon-map-marker\"></i> "
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "\n    </h4>\n</div>\n\n\n<div class=\"modal-body\">\n\n<div class=\"form-horizontal\">\n<div class=\"form-group\">\n    <label class=\"col-lg-3 control-label\">Student</label>\n    <div class=\"col-lg-6\">\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Student : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "    </div>\n</div>\n\n<div class=\"form-group\">\n    <label for=\"time\" class=\"col-sm-3 control-label\">Time</label>\n    <div class=\"col-sm-9\">\n        <div class=\"input-group time\" id=\"deadline\">\n            <input id=\"defense-time\" type=\"text\" data-old-date=\""
    + alias4(((helper = (helper = helpers.Time || (depth0 != null ? depth0.Time : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Time","hash":{},"data":data}) : helper)))
    + "\" class=\"form-control\"\n            data-timeZone=\"Europe/Paris\"\n            data-date-format=\"HH:mm\" />\n        </div>\n    </div>\n</div>\n\n<div class=\"form-group\">\n    <div class=\"col-sm-3 control-label\">Properties</div>\n    <div class=\"col-sm-9\">\n        <div class=\"checkbox\">\n            <label for=\"remote\">\n                <input type=\"checkbox\" id=\"remote\" "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Local : depth0),{"name":"unless","hash":{},"fn":container.program(6, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n                <i>visio</i>\n            </label>\n            <label for=\"private\">\n                <input type=\"checkbox\" id=\"private\" "
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Public : depth0),{"name":"unless","hash":{},"fn":container.program(6, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n                <i>private</i>\n            </label>\n        </div>\n    </div>\n</div>\n</div>\n\n</div> <!-- /modal-body -->\n<div class=\"modal-footer\">\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Student : depth0),{"name":"if","hash":{},"fn":container.program(8, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    <div class=\"text-right\">\n        <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n        <button type=\"button\" class=\"btn btn-primary\" onclick=\"setStudentDefense('"
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + "')\">Done</button>\n    </div>\n</div>\n\n</div>\n</div>";
},"useData":true});
this["wints"]["templates"]["defense-planner-2"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["session-group"],depth0,{"name":"session-group","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<div class=\"defense-groups\">\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},depth0,{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</div>\n<button type=\"button\" title=\"new session\" class=\"btn-block btn btn-primary\" onclick=\"showNewSession()\">\n    <i class=\"click glyphicon glyphicon-plus\"></i> session\n</button>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["defense-program"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<div style=\"@media print{@page {size: landscape}}\">\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},depth0,{"name":"each","hash":{},"fn":container.program(2, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n</div>\n";
},"2":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "\n"
    + ((stack1 = helpers["if"].call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Defenses : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "");
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "<h3 class=\"text-center\">"
    + alias3((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,(depth0 != null ? depth0.Day : depth0),"DD MMMM YYYY","Europe/Paris",{"name":"dateFmt","hash":{},"data":data}))
    + "</h3>\n<h4>Salle: "
    + alias3(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "</h4>\n<h4>Jury:\n<small>\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Juries : depth0),{"name":"if","hash":{},"fn":container.program(4, data, 0),"inverse":container.program(8, data, 0),"data":data})) != null ? stack1 : "")
    + "</small>\n</h4>\n<h4>Agenda</h4>\n<div style=\"page-break-after: always;\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Defenses : depth0),{"name":"each","hash":{},"fn":container.program(10, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</div>\n";
},"4":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Juries : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "");
},"5":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "	<span class=\"fn\">"
    + container.escapeExpression((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(alias1,(depth0 != null ? depth0.Person : depth0),{"name":"fullname","hash":{},"data":data}))
    + "</span>"
    + ((stack1 = helpers.unless.call(alias1,(data && data.last),{"name":"unless","hash":{},"fn":container.program(6, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n";
},"6":function(container,depth0,helpers,partials,data) {
    return "; ";
},"8":function(container,depth0,helpers,partials,data) {
    return "	<span class=\"fn\">TBA.</span>\n";
},"10":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression, alias4=container.lambda;

  return "		<div>\n		"
    + alias3((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,(depth0 != null ? depth0.Time : depth0),"HH:mm","Europe/Paris",{"name":"dateFmt","hash":{},"data":data}))
    + "\n		<span class=\"header\">"
    + alias3((helpers.fullname || (depth0 && depth0.fullname) || alias2).call(alias1,((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + " ["
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + "/"
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + "]</span>\n		"
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Local : depth0),{"name":"unless","hash":{},"fn":container.program(11, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n		"
    + ((stack1 = helpers.unless.call(alias1,(depth0 != null ? depth0.Public : depth0),{"name":"unless","hash":{},"fn":container.program(13, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n		<p class=\"descr\">"
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Company : depth0)) != null ? stack1.Name : stack1), depth0))
    + ". <i>"
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Company : depth0)) != null ? stack1.Title : stack1), depth0))
    + ".</i></p>\n		</div>\n";
},"11":function(container,depth0,helpers,partials,data) {
    return "<b><i class=\"glyphicon glyphicon-facetime-video\"></i> visio</b>";
},"13":function(container,depth0,helpers,partials,data) {
    return "<b class=\"text-danger\"> <i class=\"glyphicon glyphicon-eye-close\"></i> confidentielle</b>";
},"15":function(container,depth0,helpers,partials,data) {
    return "<h1> Soon come ...</h1>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = helpers["if"].call(depth0 != null ? depth0 : {},depth0,{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(15, data, 0),"data":data})) != null ? stack1 : "");
},"useData":true});
this["wints"]["templates"]["defense-session-creator"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var helper;

  return "                    <label class=\"col-md-2 control-label\">Session</label>\n                    <p class=\"control-static\">"
    + container.escapeExpression(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Id","hash":{},"data":data}) : helper)))
    + "</p>\n";
},"3":function(container,depth0,helpers,partials,data) {
    return "                    <div class=\"form-group\">\n                        <label class=\"col-sm-2 control-label\">Date</label>\n                        <div class=\"col-sm-9\">\n                        <input id=\"date\" type=\"text\" data-date-format=\"D MMM YYYY\">\n                        </div>\n                    </div>\n                    <div class=\"form-group\">\n                        <label class=\"col-sm-2 control-label\">Period</label>\n                        <div class=\"col-sm-9\">\n<label class=\"radio-inline\">\n  <input type=\"radio\" name=\"period\" id=\"period-am\" value=\"AM\" checked> AM\n</label>\n<label class=\"radio-inline\">\n  <input type=\"radio\" name=\"period\" id=\"period-pm\" value=\"PM\"> PM\n</label>\n                        </div>\n                    </div>\n";
},"5":function(container,depth0,helpers,partials,data) {
    var helper;

  return "'"
    + container.escapeExpression(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Id","hash":{},"data":data}) : helper)))
    + "'";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {};

  return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n        <div class=\"modal-header\">\n            <h3 class=\"modal-title\">Defense session</h3>\n        </div>\n            <div class=\"modal-body\">\n                <div class=\"form-horizontal\">\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Id : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "                    <div class=\"form-group\">\n                        <label for=\"room\" class=\"col-sm-2 control-label\">Room</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"room\" value=\""
    + container.escapeExpression(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "\" />\n                        </div>\n                    </div>\n                </div>\n            </div>\n            <div class=\"modal-footer\">\n            <div class=\"text-right\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-success\" onclick=\"addDefenseSession("
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Id : depth0),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ")\">Ok</button>\n            </div>\n            </div>\n    </div>\n</div>";
},"useData":true});
this["wints"]["templates"]["defense-session-editor"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["defense-editable"],depth0,{"name":"defense-editable","data":data,"indent":"        ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "            <option value=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(depth0 != null ? depth0 : {},((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + " "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Promotion : stack1), depth0))
    + " / "
    + alias2(alias1(((stack1 = ((stack1 = (depth0 != null ? depth0.Convention : depth0)) != null ? stack1.Student : stack1)) != null ? stack1.Major : stack1), depth0))
    + "</option>\n";
},"5":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "    <li data-jury=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Person : depth0),true,{"name":"fullname","hash":{},"data":data}))
    + "\n            <span class=\"pull-right\">\n                <i onclick=\"delDefenseJury(this, '"
    + alias2(alias1((depths[1] != null ? depths[1].Room : depths[1]), depth0))
    + "','"
    + alias2(alias1((depths[1] != null ? depths[1].Date : depths[1]), depth0))
    + "', '"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\" class=\"click glyphicon glyphicon-remove\"></i>\n            </span>\n    </li>\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.escapeExpression;

  return "            <option value=\""
    + alias1(container.lambda(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias1((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Person : depth0),true,{"name":"fullname","hash":{},"data":data}))
    + "</option>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data,blockParams,depths) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression, alias4="function";

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\">\n    <div class=\"modal-header\">\n        <h3 class=\"modal-title\">Session <i class=\"glyphicon glyphicon-calendar\"></i> "
    + alias3((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,(depth0 != null ? depth0.Date : depth0),"DD MMM A",{"name":"dateFmt","hash":{},"data":data}))
    + ", <i class=\"glyphicon glyphicon-map-marker\"></i> "
    + alias3(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "\n        </h3>\n    </div>\n    <div class=\"modal-body\">\n\n<h4>Students</h4>\n<ul class=\"agenda list-unstyled\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Defenses : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n    <li>\n        <select id=\"student-selecter\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Internships : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        </select>\n        <span class=\"pull-right\">\n            <i onclick=\"setStudentDefense(this, '"
    + alias3(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "','"
    + alias3(((helper = (helper = helpers.Date || (depth0 != null ? depth0.Date : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Date","hash":{},"data":data}) : helper)))
    + "')\" class=\"click glyphicon glyphicon-plus\"></i>\n        </span>\n    </li>\n</ul>\n\n<h4>Jury</h4>\n<ul>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Juries : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    <li>\n        <select id=\"jury-selecter\" class=\"available-jury fn\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Teachers : depth0),{"name":"each","hash":{},"fn":container.program(7, data, 0, blockParams, depths),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        </select>\n        <span class=\"pull-right\">\n            <i onclick=\"addDefenseJury(this, '"
    + alias3(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "','"
    + alias3(((helper = (helper = helpers.Date || (depth0 != null ? depth0.Date : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Date","hash":{},"data":data}) : helper)))
    + "')\" class=\"click glyphicon glyphicon-plus\"></i>\n        </span>\n    </li>\n</ul>\n\n</div>\n<div class=\"modal-footer\">\n    <div class=\"text-right\">\n        <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Done</button>\n    </div>\n</div>\n</div>";
},"usePartial":true,"useData":true,"useDepths":true});
this["wints"]["templates"]["error"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class=\"modal-dialog\">\n<div class=\"modal-content bg-danger\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>        \n        <p class=\"text-center\">"
    + container.escapeExpression(container.lambda(depth0, depth0))
    + "</p>\n        <div class=\"text-center\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Ok</button>                \n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["import-error"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression, alias4=container.lambda;

  return "            <tr>\n                <td>"
    + alias3((helpers.fullname || (depth0 && depth0.fullname) || alias2).call(alias1,((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "</td>\n                <td>"
    + alias3(alias4(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "</td>\n                <td>"
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + " / "
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + "</td>                \n                <td>"
    + alias3(((helper = (helper = helpers.Reason || (depth0 != null ? depth0.Reason : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Reason","hash":{},"data":data}) : helper)))
    + "</td>\n            </tr>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression, alias4=container.lambda;

  return "            <tr class=\"warning\">\n                <td>"
    + alias3((helpers.fullname || (depth0 && depth0.fullname) || alias2).call(alias1,((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),{"name":"fullname","hash":{},"data":data}))
    + "</td>\n                <td>"
    + alias3(alias4(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "</td>\n                <td>"
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + " / "
    + alias3(alias4(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + "</td>                                \n                <td>"
    + alias3(((helper = (helper = helpers.Reason || (depth0 != null ? depth0.Reason : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Reason","hash":{},"data":data}) : helper)))
    + "</td>\n            </tr>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <legend>Over "
    + alias3(((helper = (helper = helpers.Nb || (depth0 != null ? depth0.Nb : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Nb","hash":{},"data":data}) : helper)))
    + ": "
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,(depth0 != null ? depth0.Errors : depth0),{"name":"len","hash":{},"data":data}))
    + " failure(s); "
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,(depth0 != null ? depth0.Ignored : depth0),{"name":"len","hash":{},"data":data}))
    + " ignored</legend>   \n        <table class=\"table table-hover\">\n        <thead>\n            <tr>\n                <th>Student</th><th>Email</th><th>Promotion</th><th>Status</th>\n            </tr>\n        </thead>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Errors : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Ignored : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        </table>\n        <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Close</button>            \n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["logs"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var alias1=container.lambda, alias2=container.escapeExpression;

  return "                    <option id=\""
    + alias2(alias1(depth0, depth0))
    + "\">"
    + alias2(alias1(depth0, depth0))
    + "</option> \n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<h1>Logs</h1>\n<div class=\"form-horizontal\">\n<div class=\"form-group\" >    \n        <div class=\"col-lg-8\">\n            <select class=\"form-control\" id=\"logs\" onchange=\"showLog()\">\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},depth0,{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "            </select>                         \n        </div>                \n</div>\n</div>\n\n<div class=\"row\">\n<div class=\"col-lg-11\">\n    <pre id=\"log\">Select a log</pre>\n</div>\n<div class=\"col-lg-1\">\n<div class=\"pull-right affix text-left\">    \n    <div><a class=\"btn btn-default\" href=\"#top\"><i class=\"glyphicon glyphicon-chevron-up\"></i></a></div>\n    <div><a class=\"btn btn-default\" href=\"#bottom\"><i class=\"glyphicon glyphicon-chevron-down\"></i></a></div>\n</div>\n</div>\n</div>\n";
},"useData":true});
this["wints"]["templates"]["long-profile-editor"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <div class=\"form-horizontal\">\n            <fieldset>     \n            <legend>Profile</legend>   \n                <div class=\"form-group\">\n                    <label for=\"profile-telephon\" class=\"col-lg-2 control-label\">Firstname</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"profile-firstname\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Firstname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"profile-lastname\" class=\"col-lg-2 control-label\">Lastname</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"profile-lastname\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Lastname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"profile-tel\" class=\"col-lg-2 control-label\">Telephon</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"profile-tel\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Tel : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n            </fieldset>\n        </div>\n        <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"longUpdateProfile('"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\">Update</button>            \n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["my-defenses"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<table class=\"table table-hover\">\n    <thead>\n    <tr>\n      <th colspan=\"6\">\n      <i class=\"glyphicon glyphicon-calendar\"></i>"
    + alias4(((helper = (helper = helpers.Id || (depth0 != null ? depth0.Id : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Id","hash":{},"data":data}) : helper)))
    + ", <i class=\"glyphicon glyphicon-map-marker\"></i>"
    + alias4(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "\n      <span class=\"pull-right\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Juries : depth0),{"name":"each","hash":{},"fn":container.program(2, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "      </span>\n      </th>\n    </tr>\n    <tr>\n        <th></th>\n        <th>Time</th>\n        <th>Student</th>\n        <th>Promotion / Major</th>\n        <th class=\"text-center\">Grade</th>\n        <th class=\"text-center\">Alumni</th>\n    </tr>\n    </thead>\n    <tbody>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Defenses : depth0),{"name":"each","hash":{},"fn":container.program(5, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </tbody>\n</table>\n";
},"2":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.escapeExpression, alias2=depth0 != null ? depth0 : {};

  return "          <input type=\"checkbox\" data-email=\""
    + alias1(container.lambda(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "\">"
    + alias1((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(alias2,(depth0 != null ? depth0.Person : depth0),true,{"name":"fullname","hash":{},"data":data}))
    + ((stack1 = helpers.unless.call(alias2,(data && data.last),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n";
},"3":function(container,depth0,helpers,partials,data) {
    return "; ";
},"5":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=container.lambda, alias2=container.escapeExpression, alias3=depth0 != null ? depth0 : {}, alias4=helpers.helperMissing;

  return "      <tr>\n      <td><input type=\"checkbox\" data-email=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\"></td>\n      <td>"
    + alias2((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias4).call(alias3,(depth0 != null ? depth0.Time : depth0),"HH:mm",{"name":"dateFmt","hash":{},"data":data}))
    + "</td>\n      <td>"
    + alias2((helpers.fullname || (depth0 && depth0.fullname) || alias4).call(alias3,((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "\n      <span class=\"pull-right\">\n"
    + ((stack1 = helpers.unless.call(alias3,(depth0 != null ? depth0.Local : depth0),{"name":"unless","hash":{},"fn":container.program(6, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ((stack1 = helpers.unless.call(alias3,(depth0 != null ? depth0.Public : depth0),{"name":"unless","hash":{},"fn":container.program(8, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "      </span>\n      </td>\n      <td>"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Promotion : stack1), depth0))
    + "/"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Major : stack1), depth0))
    + "</td>\n      <td class=\"text-center\">\n      <span class=\"editable-grade\" data-student=\""
    + alias2(alias1(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\" data-type=\"text\" data-original-title=\"Enter the grade\">"
    + alias2(((helper = (helper = helpers.Grade || (depth0 != null ? depth0.Grade : depth0)) != null ? helper : alias4),(typeof helper === "function" ? helper.call(alias3,{"name":"Grade","hash":{},"data":data}) : helper)))
    + "\n      </span>\n      </td>\n\n"
    + ((stack1 = helpers["if"].call(alias3,((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.Alumni : stack1),{"name":"if","hash":{},"fn":container.program(10, data, 0),"inverse":container.program(12, data, 0),"data":data})) != null ? stack1 : "")
    + "      </tr>\n";
},"6":function(container,depth0,helpers,partials,data) {
    return "      <i class=\"glyphicon glyphicon-facetime-video\" title=\"visio\"></i>\n";
},"8":function(container,depth0,helpers,partials,data) {
    return "      <i class=\"glyphicon glyphicon-eye-close\" title=\"private\"></i>\n";
},"10":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "        <td class=\"text-center\" data-text=\"1\">\n          <a href=\"#\" onclick=\"showAlumni('"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Student : depth0)) != null ? stack1.User : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "')\">\n              <i class=\"glyphicon glyphicon-star-empty\"></i>\n          </a>\n        </td>\n";
},"12":function(container,depth0,helpers,partials,data) {
    return "        <td class=\"text-center\" data-text=\"0\"></td>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<h3>My defenses</h3>\n\n<a class=\"btn btn-default\" onclick=\"userMailing('table')\">\n      <i class=\"glyphicon glyphicon-envelope\"></i> mail selection\n</a>\n<p>\nThis is the defenses you are attending as a jury member.\n<ul>\n<li>Click on a grade to enter the value</li>\n<li>If the defense is <i>private</i> (<i class=\"glyphicon glyphicon-eye-close\"></i>), no one except the student, the jury, the supervisors (optional) can attend the defense.</li>\n<li>If the defense is performed as a <i>visio-conference</i> (<i class=\"glyphicon glyphicon-facetime-video\"></i>), the student must contact you in advance to agree on a solution</li>\n</ul>\n</p>\n\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},depth0,{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "");
},"useData":true});
this["wints"]["templates"]["new-user"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <div class=\"form-horizontal\">\n            <fieldset>     \n            <legend>New user</legend>   \n                <div class=\"alert alert-dismissible alert-danger hidden\">\n                <button type=\"button\" class=\"close\" data-dismiss=\"alert\">×</button>\n                </div>\n\n                <div class=\"form-group\">\n                    <label for=\"new-firstname\" class=\"col-lg-2 control-label\">Firstname</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"new-firstname\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Firstname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"new-lastname\" class=\"col-lg-2 control-label\">Lastname</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"new-lastname\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Lastname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                                <div class=\"form-group\">\n                    <label for=\"new-lastname\" class=\"col-lg-2 control-label\">Email</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"new-email\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Lastname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"new-tel\" class=\"col-lg-2 control-label\">Telephon</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"new-tel\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Tel : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"new-role\" class=\"col-lg-2 control-label\">Role</label>\n                    <div class=\"col-lg-10\">\n                            <select id=\"new-role\" class=\"form-control\">\n                            "
    + alias2((helpers.optionRoles || (depth0 && depth0.optionRoles) || helpers.helperMissing).call(depth0 != null ? depth0 : {},"",{"name":"optionRoles","hash":{},"data":data}))
    + "\n                        </select>\n                    </div>\n                </div>\n\n            </fieldset>\n        </div>\n        <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"newUser()\">Create</button>            \n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["password-editor"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <fieldset>\n            <legend>New password</legend>        \n            <p> Click the 'Reset' button to receive a mail describing the reset procedure.</p>            \n        </div>\n        <div class=\"alert text-center alert-success hidden\">                \n                An email has been send.\n        </div>\n        <div class=\"text-center form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"startResetPassword()\">Reset</button>\n        </div>\n        </fieldset>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["password-reset"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <div class=\"form-horizontal\">\n            <fieldset>     \n            <legend>Password update</legend>   \n                <div class=\"form-group\">\n                    <label for=\"password-current\" class=\"col-lg-2 control-label\">Current</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"password\" class=\"form-control\" id=\"password-current\">\n                    </div>\n                </div>\n                <p>Please enter your new password. It must be at least 8 characters long.</p>\n                <div class=\"form-group\">\n                    <label for=\"password-new\" class=\"col-lg-2 control-label\">New</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"password\" class=\"form-control\" id=\"password-new\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"password-confirm\" class=\"col-lg-2 control-label\">Confirm</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"password\" class=\"form-control\" id=\"password-confirm\">\n                    </div>\n                </div>\n            </fieldset>\n        </div>\n        <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"updatePassword()\">Update</button            \n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["placement-header"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<div id=\"feeder-issue\" class=\"alert alert-warning\">\n <button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-label=\"Close\"><span aria-hidden=\"true\">&times;</span></button>\n    <h5>Warning:</h5>\n    <ul>\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},((stack1 = (depth0 != null ? depth0.Errors : depth0)) != null ? stack1.Warnings : stack1),{"name":"each","hash":{},"fn":container.program(2, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </ul>        \n</div>\n";
},"2":function(container,depth0,helpers,partials,data) {
    return "        <li>"
    + container.escapeExpression(container.lambda(depth0, depth0))
    + "</li>\n";
},"4":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["placement-student"],depth0,{"name":"placement-student","data":data,"indent":"    ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression, alias4="function";

  return ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Errors : depth0)) != null ? stack1.Warnings : stack1),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n<h3>"
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,(depth0 != null ? depth0.Students : depth0),{"name":"len","hash":{},"data":data}))
    + " Student(s) - <small><span id=\"placed_cnt\">"
    + alias3(((helper = (helper = helpers.Ints || (depth0 != null ? depth0.Ints : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Ints","hash":{},"data":data}) : helper)))
    + "</span> / <span id=\"managed_cnt\">"
    + alias3(((helper = (helper = helpers.Managed || (depth0 != null ? depth0.Managed : depth0)) != null ? helper : alias2),(typeof helper === alias4 ? helper.call(alias1,{"name":"Managed","hash":{},"data":data}) : helper)))
    + "</span> placed</small></h3>\n\n<button type=\"button\" class=\"btn btn-default btn-sm\" onclick=\"userMailing('#table-placement')\">\n    <i class=\"glyphicon glyphicon-envelope\"></i> mail selection\n</button>\n\n<button type=\"button\" onclick=\"showRawFullname('tbody','stu')\" data-toggle='modal' class=\"btn btn-default btn-sm\">\n    <i class=\"glyphicon glyphicon-list\"></i> fullname from selection\n</button>    \n\n<table class=\"table table-hover tablesorter\" id=\"table-placement\" data-sortlist=\"[[1,0]]\">\n<thead>\n    <tr>\n        <th class=\"sorter-false\"><input type='checkbox' class=\"globalSelect\" data-context=\"table-placement\"/></th>\n        <th>Lastname, Firstname</th>\n        <th>Managed</th>        \n        <th>Promotion</th>        \n        <th>Major</th>\n        <th class=\"col-md-2\">Company</th>\n        <th>Period</th>\n        <th>Tutor</th>\n        <th>Gratif.</th>       \n    </tr>\n</thead>\n<tbody>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Students : depth0),{"name":"each","hash":{},"fn":container.program(4, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</tbody>\n</table>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["profile-editor"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <div class=\"form-horizontal\">\n            <fieldset>     \n            <legend>Profile</legend>   \n                <div class=\"form-group\">\n                    <label for=\"profile-telephon\" class=\"col-lg-2 control-label\">Firstname</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"profile-firstname\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Firstname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"profile-lastname\" class=\"col-lg-2 control-label\">Lastname</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"profile-lastname\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Lastname : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"profile-tel\" class=\"col-lg-2 control-label\">Telephon</label>\n                    <div class=\"col-lg-10\">\n                        <input data-placement=\"right\" data-container=\"body\" data-toggle=\"popover\" type=\"text\" class=\"form-control\" id=\"profile-tel\" value=\""
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Tel : stack1), depth0))
    + "\">\n                    </div>\n                </div>\n            </fieldset>\n        </div>\n        <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"updateProfile('"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\">Update</button>            \n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["progress"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n        <div class=\"modal-header\"></div>\n        <div class=\"modal-body\">   \n            <label>Uploading ...</label>             \n            <div class=\"progress\">\n                <div id=\"progress-value\" class=\"progress-bar\" role=\"progressbar\" aria-valuenow=\"0\" aria-valuemin=\"0\" aria-valuemax=\"100\" style=\"width: 0%\">0%</div>\n            </div>            \n        </div>\n        <div class=\"modal-footer\"></div>                            \n    </div>\n</div>";
},"useData":true});
this["wints"]["templates"]["raw"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>        \n        <legend>"
    + alias4(((helper = (helper = helpers.Title || (depth0 != null ? depth0.Title : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Title","hash":{},"data":data}) : helper)))
    + "</legend>\n        <pre onclick=\"selectText(this)\" class=\"raw\">"
    + alias4(((helper = (helper = helpers.Cnt || (depth0 != null ? depth0.Cnt : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Cnt","hash":{},"data":data}) : helper)))
    + "</pre>       \n    </div>\n	</div>\n</div>\n";
},"useData":true});
this["wints"]["templates"]["report-modal"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "checked";
},"3":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "                            <a href=\"api/v2/reports/"
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "/"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "/content\">\n                                <i class=\"glyphicon glyphicon-cloud-download\"></i> download\n                            </a>\n";
},"5":function(container,depth0,helpers,partials,data) {
    return "                            n/a\n";
},"7":function(container,depth0,helpers,partials,data) {
    return "                                "
    + container.escapeExpression((helpers.dateFmt || (depth0 && depth0.dateFmt) || helpers.helperMissing).call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Delivery : depth0),"D MMM YYYY HH:mm",{"name":"dateFmt","hash":{},"data":data}))
    + "\n";
},"9":function(container,depth0,helpers,partials,data) {
    return "                                n/a\n";
},"11":function(container,depth0,helpers,partials,data) {
    return "<span class=\"text-danger\">Deadline missed !</span>";
},"13":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing;

  return ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.ToGrade : depth0),{"name":"if","hash":{},"fn":container.program(14, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "                    <div class=\"form-group\">\n                        <label for=\"comments\" class=\"control-label col-sm-offset-1\">Comment <small><a href=\"/assets/consignes-rapports.pdf\">(expectations)</a></small></label>\n                        <textarea placeholder=\"insert review here...\" id=\"comment\" rows=\"9\" class=\"col-sm-10 col-sm-offset-1\" "
    + ((stack1 = (helpers.ifManage || (depth0 && depth0.ifManage) || alias2).call(alias1,depth0,{"name":"ifManage","hash":{},"fn":container.program(17, data, 0),"inverse":container.program(21, data, 0),"data":data})) != null ? stack1 : "")
    + ">"
    + container.escapeExpression(((helper = (helper = helpers.Comment || (depth0 != null ? depth0.Comment : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Comment","hash":{},"data":data}) : helper)))
    + "</textarea>\n                    </div>\n";
},"14":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "                        <div class=\"form-group\">\n                            <label for=\"grade\" class=\"col-sm-2 control-label\">Grade</label>\n                            <div class=\"col-sm-9\">\n                                <input type=\"number\" class=\"form-control\" id=\"grade\" "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Reviewed : depth0),{"name":"if","hash":{},"fn":container.program(15, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + " "
    + ((stack1 = (helpers.ifManage || (depth0 && depth0.ifManage) || helpers.helperMissing).call(alias1,depth0,{"name":"ifManage","hash":{},"fn":container.program(17, data, 0),"inverse":container.program(19, data, 0),"data":data})) != null ? stack1 : "")
    + "/>\n                            </div>\n                        </div>\n";
},"15":function(container,depth0,helpers,partials,data) {
    var helper;

  return "value=\""
    + container.escapeExpression(((helper = (helper = helpers.Grade || (depth0 != null ? depth0.Grade : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"Grade","hash":{},"data":data}) : helper)))
    + "\"";
},"17":function(container,depth0,helpers,partials,data) {
    return "";
},"19":function(container,depth0,helpers,partials,data) {
    return "disabled=\"disabled\"";
},"21":function(container,depth0,helpers,partials,data) {
    return "readonly";
},"23":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = helpers["if"].call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.ToGrade : depth0),{"name":"if","hash":{},"fn":container.program(24, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "");
},"24":function(container,depth0,helpers,partials,data) {
    return "                    <div>\n                        <b>Notation</b>\n                        <small class=\"col-sm-offset-1\">\n                        <ul>\n    <li><b>16 to 18: </b> Well written, crystal clear. It indicates the context, the problematic, the associated bibliography, the solution, the difficulties, the benefits and the planification. Everything is carefully justified.\n    The note might vary according to the subject difficulties.\n    </li>\n    <li><b>13 to 15: </b>\n    Well written, good explanations about the work done and the work to come but\n    editorial issues, un-rigourous justications.\n    </li>\n    <li><b>10 to 12: </b>\n    Readable but no enough details to really understand the work done. Lack of analysis. Some part are missings (bibliography, difficulties, planning, ...)\n    </li>\n    <li><b>&lt; 10: </b>\n    Hard to read/understand, severe lacks in terms of contents or analysis capacities</li>\n                        </ul>\n                        <b class=\"text-center\">Do not consider any late penalty. This is done automatically</b>\n                        </small>\n                    </div>\n";
},"26":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "                        <button type=\"button\" class=\"btn btn-danger\" onclick=\"review('"
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "',"
    + alias4(((helper = (helper = helpers.ToGrade || (depth0 != null ? depth0.ToGrade : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"ToGrade","hash":{},"data":data}) : helper)))
    + ")\">\n                            <i class=\"glyphicon glyphicon-floppy-save\"></i>\n                            Review\n                        </button>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n            <div class=\"modal-body\">\n            <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n            <fieldset>\n                <legend>"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + " Report</legend>\n                <div class=\"form-horizontal\">\n\n                    <div class=\"form-group\">\n                        <div class=\"col-sm-3 control-label\">Confidential</div>\n                        <div class=\"col-sm-9\">\n                            <div class=\"checkbox\">\n                                <label for=\"private\">\n                                    <input onchange=\"toggleReportConfidential('"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "',this)\" type=\"checkbox\" id=\"private\" "
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Private : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">\n                                    <i>(access limited to the tutor and root)</i>\n                                </label>\n                            </div>\n                        </div>\n                    </div>\n\n                    <div class=\"form-group\">\n                        <label for=\"deadline\" class=\"col-sm-3 control-label\">Deadline</label>\n                        <div class=\"col-sm-9\">\n                            <div class=\"input-group date\" id=\"deadline\">\n                                <input id=\"report-deadline\" type=\"text\" data-old-date=\""
    + alias4(((helper = (helper = helpers.Deadline || (depth0 != null ? depth0.Deadline : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Deadline","hash":{},"data":data}) : helper)))
    + "\" class=\"form-control\" data-date-format=\"D MMM YYYY\"  />\n                                <span class=\"input-group-btn\">\n                                    <button onclick=\"updateReportDeadline('"
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "','"
    + alias4(((helper = (helper = helpers.Kind || (depth0 != null ? depth0.Kind : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Kind","hash":{},"data":data}) : helper)))
    + "');\" class=\"btn btn-info\" type=\"button\"><i class=\"glyphicon glyphicon-floppy-disk\"></i></button>\n                                </span>\n                            </div>\n                        </div>\n                    </div>\n                    <div class=\"form-group\">\n                        <label for=\"down\" class=\"col-sm-3 control-label\">Document</label>\n                        <div class=\"col-sm-9\">\n                            <label class=\"form-control-static\">\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Delivery : depth0),{"name":"if","hash":{},"fn":container.program(3, data, 0),"inverse":container.program(5, data, 0),"data":data})) != null ? stack1 : "")
    + "                            </label>\n                        </div>\n                    </div>\n\n                    <div class=\"form-group\">\n                        <label for=\"delivery\" class=\"col-sm-3 control-label\">Delivery date</label>\n                        <div class=\"col-sm-9\">\n                            <label class=\"form-control-static\">\n"
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.Delivery : depth0),{"name":"if","hash":{},"fn":container.program(7, data, 0),"inverse":container.program(9, data, 0),"data":data})) != null ? stack1 : "")
    + "                            "
    + ((stack1 = (helpers.ifAfter || (depth0 && depth0.ifAfter) || alias2).call(alias1,(depth0 != null ? depth0.Delivery : depth0),(depth0 != null ? depth0.Deadline : depth0),{"name":"ifAfter","hash":{},"fn":container.program(11, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n                            </label>\n                        </div>\n                    </div>\n\n"
    + ((stack1 = (helpers.ifLate || (depth0 && depth0.ifLate) || alias2).call(alias1,(depth0 != null ? depth0.Deadline : depth0),{"name":"ifLate","hash":{},"fn":container.program(13, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ((stack1 = (helpers.ifLate || (depth0 && depth0.ifLate) || alias2).call(alias1,(depth0 != null ? depth0.Deadline : depth0),{"name":"ifLate","hash":{},"fn":container.program(23, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n                                        <div class=\"text-right form-group\">\n                        <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal();\">Cancel</button>\n"
    + ((stack1 = (helpers.ifLate || (depth0 && depth0.ifLate) || alias2).call(alias1,(depth0 != null ? depth0.Deadline : depth0),{"name":"ifLate","hash":{},"fn":container.program(26, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n                    </div>\n\n                    </div>\n            </div>\n        </fieldset>\n    </div>\n</div>";
},"useData":true});
this["wints"]["templates"]["role-editor"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=container.lambda, alias2=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n<div class=\"modal-content\"> \n    <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n        <div class=\"form-horizontal\">\n            <fieldset>     \n            <legend class=\"fn\">"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Lastname : stack1), depth0))
    + ", "
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Firstname : stack1), depth0))
    + "</legend>   \n                <div class=\"alert text-center alert-danger hidden\"></div>\n\n                <div class=\"form-group\">\n                    <label for=\"profile-role\" class=\"col-lg-2 control-label\">Role</label>\n                    <div class=\"col-lg-10\">\n                        <select class=\"form-control\" id=\"profile-role\">\n                            "
    + alias2((helpers.optionRoles || (depth0 && depth0.optionRoles) || helpers.helperMissing).call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.Role : depth0),{"name":"optionRoles","hash":{},"data":data}))
    + "\n                        </select>\n                    </div>\n                </div>\n            </fieldset>\n        </div>\n        <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"updateRole('"
    + alias2(alias1(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + "')\">Update</button>\n        </div>\n    </div>    \n</div>    \n</div>\n        ";
},"useData":true});
this["wints"]["templates"]["service"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "        <tr data-email=\""
    + alias4(((helper = (helper = helpers.key || (data && data.key)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"key","hash":{},"data":data}) : helper)))
    + "\">\n            <td>\n                <input class=\"shiftSelectable\" type='checkbox' data-user-fn='"
    + alias4((helpers.fullname || (depth0 && depth0.fullname) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.U : depth0)) != null ? stack1.Person : stack1),true,{"name":"fullname","hash":{},"data":data}))
    + "' data-email='"
    + alias4(container.lambda(((stack1 = ((stack1 = (depth0 != null ? depth0.U : depth0)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "' />\n            </td>\n            <td class=\"text-left fn\">\n"
    + ((stack1 = container.invokePartial(partials.person,((stack1 = (depth0 != null ? depth0.U : depth0)) != null ? stack1.Person : stack1),{"name":"person","data":data,"indent":"                ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "            </td>\n            <td data-text=\""
    + alias4(((helper = (helper = helpers.TotalInts || (depth0 != null ? depth0.TotalInts : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"TotalInts","hash":{},"data":data}) : helper)))
    + "\">\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Ints : depth0),{"name":"each","hash":{},"fn":container.program(2, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "            </td>\n            <td>"
    + alias4(((helper = (helper = helpers.Defs || (depth0 != null ? depth0.Defs : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Defs","hash":{},"data":data}) : helper)))
    + "</td>\n        </tr>\n";
},"2":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "                "
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,depth0,{"name":"len","hash":{},"data":data}))
    + " "
    + alias3(((helper = (helper = helpers.key || (data && data.key)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"key","hash":{},"data":data}) : helper)))
    + ((stack1 = helpers.unless.call(alias1,(data && data.last),{"name":"unless","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n";
},"3":function(container,depth0,helpers,partials,data) {
    return " ; ";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<h3>Service</h3>\n\n<button type=\"button\" class=\"btn btn-default btn-sm\" onclick=\"userMailing('#table-service')\">\n    <i class=\"glyphicon glyphicon-envelope\"></i> mail selection\n</button>\n\n<button type=\"button\" onclick=\"showRawFullname('tbody','user')\" data-toggle='modal' class=\"btn btn-default btn-sm\">\n    <i class=\"glyphicon glyphicon-list\"></i> fullname from selection\n</button>\n\n<table class=\"table table-hover tablesorter\" id=\"table-service\">\n    <thead>\n    <tr>\n        <th class=\"sorter-false\">\n        <input type='checkbox' data-toggle='checkbox' class='globalSelect' data-context=\"table-service\"/>\n        </th>\n        <th>Person</th>\n        <th>Tutored students</th>\n        <th>Defense committees</th>\n    </tr>\n    </thead>\n    <tbody>\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},depth0,{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </tbody>\n</table>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["student-dashboard-alumni-editor"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "selected";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing;

  return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n        <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n            <div class=\"form-horizontal\">\n                <fieldset>\n                <legend>Alumni editor</legend>\n                <div class=\"form-group\">\n                    <label for=\"lbl-fn\" class=\"col-sm-3 control-label\">Email</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"email\" class=\"form-control\" id=\"lbl-email\" value=\""
    + container.escapeExpression(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "\"/>\n                        </div>\n                </div>\n                 \n                <div class=\"form-group\">\n                    <label for=\"lbl-kind\" class=\"col-sm-3 control-label\">Kind</label>\n                        <div class=\"col-sm-9\">\n                            <select class=\"form-control\" id=\"position\" onchange=\"syncAlumniEditor(this)\">\n                                <option value=\"looking\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"looking",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Looking for a job</option>\n                                <option value=\"sabbatical\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"sabbatical",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Sabattical leave</option>\n                                <option value=\"company\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"company",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Working in a company</option>\n                                <option value=\"entrepreneurship\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"entrepreneurship",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Entrepreneurship</option>\n                                <option value=\"study\" "
    + ((stack1 = (helpers.ifEq || (depth0 && depth0.ifEq) || alias2).call(alias1,(depth0 != null ? depth0.Position : depth0),"study",{"name":"ifEq","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ">Pursuit of higher education</option>\n                            </select>                            \n                        </div>\n                </div>\n\n                <div class=\"form-group hidden\" id=\"contract\">\n                    <label class=\"col-sm-3 control-label\">Contract</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"checkbox-inline\">\n                            <input type=\"radio\" name=\"permanent\" value='false' disabled> Fixed (CDD)\n                        </label>                        \n                        <label class=\"checkbox-inline\">\n                            <input type=\"radio\" name='permanent' value='true' disabled> Permanent (CDI)\n                        </label>                        \n                    </div>\n                </div>   \n\n                <div class=\"form-group hidden\" id=\"company\">\n                    <label class=\"col-sm-3 control-label\">Company</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"checkbox-inline\">\n                            <input type=\"radio\" name='sameCompany' value='true' disabled> internship company\n                        </label>                        \n                        <label class=\"checkbox-inline\">\n                            <input type=\"radio\" name='sameCompany' value='false' disabled> other\n                        </label>                        \n                    </div>\n                </div> \n\n                <div class=\"form-group hidden\" id=\"country\">\n                    <label class=\"col-sm-3 control-label\">Country</label>\n                    <div class=\"col-sm-9\">\n                        <label class=\"checkbox-inline\">\n                            <input type=\"radio\" name='france' value='true' disabled> France\n                        </label>                        \n                        <label class=\"checkbox-inline\">\n                            <input type=\"radio\" name='france' value='false' disabled> Foreign country\n                        </label>                        \n                    </div>\n                </div>                            \n                </div>\n                </fieldset>\n            <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"sendAlumni()\">Update</button>\n            </div>                \n            </div>\n        </div>\n    </div>\n</div>   ";
},"useData":true});
this["wints"]["templates"]["student-dashboard-defense"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "on site";
},"3":function(container,depth0,helpers,partials,data) {
    return "visio";
},"5":function(container,depth0,helpers,partials,data) {
    return "public";
},"7":function(container,depth0,helpers,partials,data) {
    return "<span class=\"text-danger\">confidential</span>";
},"9":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "        "
    + container.escapeExpression((helpers.fullname || (depth0 && depth0.fullname) || helpers.helperMissing).call(alias1,(depth0 != null ? depth0.Person : depth0),{"name":"fullname","hash":{},"data":data}))
    + ((stack1 = helpers.unless.call(alias1,(data && data.last),{"name":"unless","hash":{},"fn":container.program(10, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\n";
},"10":function(container,depth0,helpers,partials,data) {
    return "; ";
},"12":function(container,depth0,helpers,partials,data) {
    var stack1;

  return container.escapeExpression(container.lambda(((stack1 = (depth0 != null ? depth0.Person : depth0)) != null ? stack1.Email : stack1), depth0))
    + ((stack1 = helpers.unless.call(depth0 != null ? depth0 : {},(data && data.last),{"name":"unless","hash":{},"fn":container.program(13, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "");
},"13":function(container,depth0,helpers,partials,data) {
    return ",";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "<h3 class=\"page-header\">Oral defense <small>(<a href=\"assets/consignes-rapports.pdf\">organisation</a>)</small></h3>\n<dl class=\"dl-horizontal\">\n        <dt>Room</dt>\n        <dd>"
    + alias3(((helper = (helper = helpers.Room || (depth0 != null ? depth0.Room : depth0)) != null ? helper : alias2),(typeof helper === "function" ? helper.call(alias1,{"name":"Room","hash":{},"data":data}) : helper)))
    + "</dd>\n        <dt>Date</dt>\n        <dd>"
    + alias3((helpers.dateFmt || (depth0 && depth0.dateFmt) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Defense : depth0)) != null ? stack1.Time : stack1),"DD MMM HH:mm","Europe/Paris",{"name":"dateFmt","hash":{},"data":data}))
    + " (<i>Paris timezone</i>)</dd>\n        <dt>On site ?</dt>\n        <dd>"
    + ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Defense : depth0)) != null ? stack1.Local : stack1),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(3, data, 0),"data":data})) != null ? stack1 : "")
    + "</dd>\n        <dt>Public ?</dt>\n        <dd>"
    + ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Defense : depth0)) != null ? stack1.Public : stack1),{"name":"if","hash":{},"fn":container.program(5, data, 0),"inverse":container.program(7, data, 0),"data":data})) != null ? stack1 : "")
    + "</dd>\n        <dt>\n        Jury\n        </dt>\n        <dd>\n        <small>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Juries : depth0),{"name":"each","hash":{},"fn":container.program(9, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "                <a href=\"mailto:"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Juries : depth0),{"name":"each","hash":{},"fn":container.program(12, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "\"\n        class=\"btn btn-xs btn-default\"><i class=\"glyphicon glyphicon-envelope\"></i></a>\n        </small>\n        </dd>\n</dl>";
},"useData":true});
this["wints"]["templates"]["student-dashboard-supervisor-editor"] = Handlebars.template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "<div class=\"modal-dialog\">\n    <div class=\"modal-content\">\n        <div class=\"modal-body\">\n        <button type=\"button\" class=\"close\" aria-hidden=\"true\" onclick=\"hideModal()\">×</button>\n            <div class=\"form-horizontal\">\n                <fieldset>\n                <legend>Supervisor contact details</legend>\n                <div class=\"form-group\">\n                    <label for=\"sup-fn\" class=\"col-sm-3 control-label\">Firstname</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"sup-fn\" value=\""
    + alias4(((helper = (helper = helpers.Firstname || (depth0 != null ? depth0.Firstname : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Firstname","hash":{},"data":data}) : helper)))
    + "\"/>\n                        </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"sup-ln\" class=\"col-sm-3 control-label\">Lastname</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"sup-ln\" value=\""
    + alias4(((helper = (helper = helpers.Lastname || (depth0 != null ? depth0.Lastname : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Lastname","hash":{},"data":data}) : helper)))
    + "\"/>\n                        </div>\n                </div>\n                <div class=\"form-group\">\n                    <label for=\"sup-email\" class=\"col-sm-3 control-label\">Email</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"sup-email\" value=\""
    + alias4(((helper = (helper = helpers.Email || (depth0 != null ? depth0.Email : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Email","hash":{},"data":data}) : helper)))
    + "\"/>\n                        </div>\n                </div>                \n                <div class=\"form-group\">\n                    <label for=\"sup-tel\" class=\"col-sm-3 control-label\">Telephon</label>\n                        <div class=\"col-sm-9\">\n                            <input type=\"text\" class=\"form-control\" id=\"sup-tel\" value=\""
    + alias4(((helper = (helper = helpers.Tel || (depth0 != null ? depth0.Tel : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"Tel","hash":{},"data":data}) : helper)))
    + "\"/>\n                        </div>\n                </div>                            \n                </div>\n                </fieldset>\n            <div class=\"text-right form-group\">\n            <button type=\"button\" class=\"btn btn-default\" onclick=\"hideModal()\">Cancel</button>\n            <button type=\"button\" class=\"btn btn-primary\" onclick=\"sendSupervisor()\">Update</button>\n            </div>                \n            </div>\n        </div>\n    </div>\n</div> ";
},"useData":true});
this["wints"]["templates"]["student-dashboard"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "There is a "
    + container.escapeExpression(container.lambda(((stack1 = (depth0 != null ? depth0.Config : depth0)) != null ? stack1.LatePenalty : stack1), depth0))
    + " point penalty per day late.\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["student-dashboard-report"],depth0,{"name":"student-dashboard-report","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "<div class=\"row student-row\">\n<div class=\"col-md-6\" id=\"dashboard-company\">\n"
    + ((stack1 = container.invokePartial(partials["student-dashboard-company"],(depth0 != null ? depth0.Internship : depth0),{"name":"student-dashboard-company","data":data,"indent":"\t","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</div>\n\n<div class=\"col-md-6\" id=\"student-dashboard-contacts\">\n"
    + ((stack1 = container.invokePartial(partials["student-dashboard-contacts"],(depth0 != null ? depth0.Internship : depth0),{"name":"student-dashboard-contacts","data":data,"indent":"\t","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</div>\n</div>\n\n<div class=\"row student-row\">\n<div class=\"col-md-6\">\n<h3 class=\"page-header\">Reports</h3>\n<p class=\"text-justify\">\nPDF files, 10 MB max.\nSee <a href=\"/assets/consignes-rapports.pdf\">the expectations</a>.\nThe upload may take time. Once completed, download the report to check for errors.\n</p>\n<p class=\"text-justify\">\nYour <a href=\"mailto:"
    + container.escapeExpression(container.lambda(((stack1 = ((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Internship : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Tutor : stack1)) != null ? stack1.Person : stack1)) != null ? stack1.Email : stack1), depth0))
    + "\">academic supervisor</a> can reconsider the deadlines upon strong justifications.\nYou can re-upload until the deadline.\nOnce late, you will have <b>one</b> upload opportunity.\n"
    + ((stack1 = helpers["if"].call(alias1,((stack1 = (depth0 != null ? depth0.Config : depth0)) != null ? stack1.LatePenalty : stack1),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</p>\n<table class=\"table table-hover table-condensed\">\n<thead>\n<tr>\n	<td>Kind</td><td>Deadline</td><td>Delivery date</td><td>Grade</td><td></td>\n</tr>\n</thead>\n<tbody>\n"
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.Internship : depth0)) != null ? stack1.Reports : stack1),{"name":"each","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</tbody>\n</table>\n</div>\n<div class=\"col-md-6\" id=\"dashboard-defense\">\n</div>\n\n<div class=\"col-md-6\" id=\"dashboard-alumni\">\n"
    + ((stack1 = container.invokePartial(partials["student-dashboard-alumni"],((stack1 = ((stack1 = ((stack1 = (depth0 != null ? depth0.Internship : depth0)) != null ? stack1.Convention : stack1)) != null ? stack1.Student : stack1)) != null ? stack1.Alumni : stack1),{"name":"student-dashboard-alumni","data":data,"helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "")
    + "</div>\n</div>\n";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["tutored"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "        <th class=\"text-center\"><small>"
    + container.escapeExpression(container.lambda((depth0 != null ? depth0.Kind : depth0), depth0))
    + "</small></th>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["tutored-student"],depth0,{"name":"tutored-student","data":data,"indent":"        ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "<h3><span class=\"count\">"
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,(depth0 != null ? depth0.Internships : depth0),{"name":"len","hash":{},"data":data}))
    + "</span> tutored student(s)</h3>\n\n<div class=\"btn-group\">\n  <button type=\"button\" class=\"btn btn-default dropdown-toggle btn-sm\" data-toggle=\"dropdown\" aria-expanded=\"false\"><i class=\"glyphicon glyphicon-envelope\"></i> mail selection <span class=\"caret\"></span>\n  </button>\n  <ul class=\"dropdown-menu\" role=\"menu\">\n    <li><a onclick=\"conventionMailing('tbody','stu')\">students</a></li>\n    <li><a onclick=\"conventionMailing('tbody','stu','sup')\">students (cc supervisors)</a></li>        \n    <li><a onclick=\"conventionMailing('tbody','sup','stu')\">supervisors (cc students)</a></li>        \n  </ul>\n</div>\n\n<div class=\"btn-group\">\n  <button type=\"button\" class=\"btn btn-default btn-sm dropdown-toggle\" data-toggle=\"dropdown\" aria-expanded=\"false\"><i class=\"glyphicon glyphicon-list\"></i> fullnames from selection <span class=\"caret\"></span>\n  </button>\n  <ul class=\"dropdown-menu\" role=\"menu\">\n    <li><a onclick=\"showRawFullname('tbody','stu')\">Students</a></li>    \n    <li><a onclick=\"showRawFullname('tbody','sup')\">Supervisors</a></li>    \n  </ul>\n</div>\n\n<p>\nThis is the students you are following as an academic tutor. Your role is to:\n</p>\n<ul>\n<li>Maintain contacts with the students and the supervisors. Use the checkboxes to select internships and the top buttons for mailing purposes.</li>\n<li>Click on a report cell to access a report and grade it (see the <a href=\"/assets/consignes-rapports.pdf\"> expectations</a>). If needed, you can also revise the deadline (e.g. if the student is late for a viable reason) or the report visibility.</li>\n</ul>\n<!--<div class=\"legend\">\n<span class=\"title\">Reports color legend: </span>\n<ul class=\"legend\">\n<li><span class=\"color bg-warning\"></span>Deadline passed, no report</li>\n<li><span class=\"color bg-info\"></span>Waiting for being reviewed</li>\n<li><span class=\"color bg-success\"></span>Grade &gt; 10</li>\n<li><span class=\"color bg-danger\"></span>Grade &lt; 10</li>\n</ul>\n</div>-->\n<table class=\"table table-hover tablesorter\" id=\"table-tutoring\" data-partial=\"tutored-student\">\n    <thead>\n    <tr>\n        <th class=\"sorter-false\" rowspan=\"2\">\n        <input type='checkbox' data-toggle='checkbox' class='globalSelect' data-context=\"table-tutoring\"/>\n        </th>\n        <th rowspan=\"2\">Student</th>\n        <th rowspan=\"2\">Promotion</th>\n        <th rowspan=\"2\">Major</th>\n        <th rowspan=\"2\" class=\"sorter-false\">Supervisor</th>\n        <th rowspan=\"2\" class=\"sorter-false\">Company</th>\n        <th class=\"sorter-false text-center\" colspan=\""
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Reports : stack1),{"name":"len","hash":{},"data":data}))
    + "\">Reports</th>\n        <th class=\"sorter-false text-center\" colspan=\""
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Surveys : stack1),{"name":"len","hash":{},"data":data}))
    + "\">Surveys</th>        \n        <th class=\"text-center\" rowspan=\"2\">Alumni</th>\n    </tr>\n    <tr>        \n"
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Reports : stack1),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Surveys : stack1),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </tr>    \n    </thead>\n    <tbody>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Internships : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </tbody>\n</table>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["users-header"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["users-user"],depth0,{"name":"users-user","data":data,"indent":"    ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {};

  return "	<h3><span class=\"count\">"
    + container.escapeExpression((helpers.len || (depth0 && depth0.len) || helpers.helperMissing).call(alias1,depth0,{"name":"len","hash":{},"data":data}))
    + "</span> registered user(s)</h3>\n\n<a class=\"btn btn-default btn-sm\" onclick=\"userMailing('#table-users')\">\n    <i class=\"glyphicon glyphicon-envelope\"></i> mail selection\n</a>\n\n<a onclick=\"showRawFullname('tbody','user')\" data-toggle='modal' class=\"btn btn-default btn-sm\">\n    <i class=\"glyphicon glyphicon-list\"></i> fullname from selection\n</a>    \n\n\n<a onclick=\"showNewUser()\" data-toggle='modal' class=\"btn btn-info btn-sm\">\n    <i class=\"glyphicon glyphicon-plus\"></i> add user\n</a>    \n\n<span class=\"btn btn-info btn-sm btn-file\">\n    <i class=\"glyphicon glyphicon-plus\"></i> import students <input id=\"csv-import\" type=\"file\">\n</span>\n\n<a id=\"import-status\" class=\"btn btn-danger btn-sm hidden\" onclick=\"showImportError()\">Show import errors</a>\n\n\n<table class=\"table table-hover tablesorter\" id=\"table-users\" data-sortlist=\"[[1,0]]\">\n<thead>\n    <tr>\n        <th class=\"sorter-false\"><input type='checkbox' class=\"globalSelect\" data-context=\"table-users\"/></th>\n        <th>Lastname, Firstname</th><th>Email</th>\n        <th>Role</th><th>Last visit</th><th class=\"sorter-false\">Actions</th>\n    </tr>\n</thead>\n<tbody>\n"
    + ((stack1 = helpers.each.call(alias1,depth0,{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</tbody>\n</table>";
},"usePartial":true,"useData":true});
this["wints"]["templates"]["watchlist"] = Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    return "        <th class=\"text-center\"><small>"
    + container.escapeExpression(container.lambda((depth0 != null ? depth0.Kind : depth0), depth0))
    + "</small></th>\n";
},"3":function(container,depth0,helpers,partials,data) {
    var stack1;

  return ((stack1 = container.invokePartial(partials["watchlist-student"],depth0,{"name":"watchlist-student","data":data,"indent":"        ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3=container.escapeExpression;

  return "<h3>Watchlist</h3>\n\n<div class=\"btn-group\">\n  <button type=\"button\" class=\"btn btn-default dropdown-toggle btn-sm\" data-toggle=\"dropdown\" aria-expanded=\"false\"><i class=\"glyphicon glyphicon-envelope\"></i> mail selection <span class=\"caret\"></span>\n  </button>\n  <ul class=\"dropdown-menu\" role=\"menu\">\n    <li><a onclick=\"conventionMailing('tbody', 'stu')\">students</a></li>\n    <li><a onclick=\"conventionMailing('tbody', 'stu','tut')\">students (cc tutors)</a></li>\n    <li><a onclick=\"conventionMailing('tbody', 'sup')\">supervisors</a></li>\n    <li><a onclick=\"conventionMailing('tbody', 'tut')\">tutors</a></li>\n  </ul>\n</div>\n\n<div class=\"btn-group\">\n  <button type=\"button\" class=\"btn btn-default btn-sm dropdown-toggle\" data-toggle=\"dropdown\" aria-expanded=\"false\"><i class=\"glyphicon glyphicon-list\"></i> fullnames from selection <span class=\"caret\"></span>\n  </button>\n  <ul class=\"dropdown-menu\" role=\"menu\">\n    <li><a onclick=\"showRawFullname('tbody', 'stu')\">Students</a></li>\n    <li><a onclick=\"showRawFullname('tbody','sup')\">Supervisors</a></li>\n    <li><a onclick=\"showRawFullname('tbody','tut')\">Tutors</a></li>\n  </ul>\n</div>\n\n<table class=\"table table-hover tablesorter\" id=\"table-conventions\" data-partial=\"watchlist-student\">\n    <thead>\n    <tr>\n        <th class=\"sorter-false\" rowspan=\"2\">\n        <input type='checkbox' data-toggle='checkbox' class='globalSelect' data-context=\"table-conventions\"/>\n        </th>\n        <th rowspan=\"2\">Student</th>\n        <th rowspan=\"2\">Promotion</th>\n        <th rowspan=\"2\">Major</th>\n        <th rowspan=\"2\" >Tutor</th>\n        <th class=\"sorter-false text-center\" colspan=\""
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Reports : stack1),{"name":"len","hash":{},"data":data}))
    + "\">Reports</th>\n        <th class=\"sorter-false text-center\" colspan=\""
    + alias3((helpers.len || (depth0 && depth0.len) || alias2).call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Surveys : stack1),{"name":"len","hash":{},"data":data}))
    + "\">Surveys</th>\n        <th class=\"text-center\" rowspan=\"2\" >Def</th>\n        <th class=\"text-center\" rowspan=\"2\">Alumni</th>\n    </tr>\n    <tr>\n"
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Reports : stack1),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + ((stack1 = helpers.each.call(alias1,((stack1 = (depth0 != null ? depth0.Org : depth0)) != null ? stack1.Surveys : stack1),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </tr>\n    </thead>\n    <tbody>\n"
    + ((stack1 = helpers.each.call(alias1,(depth0 != null ? depth0.Internships : depth0),{"name":"each","hash":{},"fn":container.program(3, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "    </tbody>\n</table>";
},"usePartial":true,"useData":true});