/**
 * Created by fhermeni on 03/07/2014.
 */
//Collect rest queries

function defaultCb(no) {
    if (no != undefined) {
        return no;
    }
    return function() {
        console.log(arguments);
    };
}

//Convention management
function randomPending(ok, no) {
    getWithToken("/pending/_random", defaultCb(ok), defaultCb(no));
}

function commitPendingConvention(c, ok, no) {
    postWithToken("/conventions/", c, defaultCb(ok), defaultCb(no));
}

function getConventions(ok, no) {
    getWithToken("/conventions/",defaultCb(ok), defaultCb(no));
}

//User management
function createUser(ok, no) {
    postWithToken("/users/", d, defaultCb(ok), defaultCb(no));
}

function deleteUser(email, ok, no) {
    deleteWithToken("/users/" + email, defaultCb(ok),  defaultCb(no));
}

function getUsers(ok, no) {
    getWithToken("/users/", defaultCb(ok), defaultCb(no));
}

function setPrivilege(email, p, ok, no) {
    postRawWithToken("/users/" + email + "/role", p, defaultCb(ok), defaultCb(no));
}

function setMajor(email, p, ok, no) {
    postRawWithToken("/users/" + email + "/role", p, defaultCb(ok), defaultCb(no));
}