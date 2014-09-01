function service(cc, dd) {
    var srv = {};
    cc.forEach(function (c) {
        var t = c.Tutor;
        if (!srv[t.Email]) {
            srv[t.Email] = {
                P : t,
                tutored: [],
                juries:[],
                studentsMajor: [],
                juriesMajor: []
            };
        }
        srv[t.Email].tutored.push(c.Stu);
        if (srv[t.Email].studentsMajor.indexOf(c.Stu.Major) < 0) {
            srv[t.Email].studentsMajor.push(c.Stu.Major);
        }
    });
    dd.sessions.forEach(function (ss) {
        ss.jury.forEach(function (j) {
            j.commission.forEach(function (t) {
                if (t && t.length > 2) {
                    if (!srv[t]) {
                        srv[t] = {
                            P : getUser(t),
                            tutored: [],
                            juries:[],
                            studentsMajor: [],
                            juriesMajor: []
                        };
                    }
                    srv[t].juries.push(j);
                    j.students.forEach(function (s) {
                        if (srv[t].juriesMajor.indexOf(s.Major) < 0) {
                            srv[t].juriesMajor.push(s.Major);
                        }
                    });
                }
            });
        });
    });
    var arr = [];
    Object.keys(srv).forEach(function (k) {
        arr.push(srv[k]);
    });
    return arr;
}