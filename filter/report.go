package filter

import (
	"strconv"
	"time"

	"github.com/fhermeni/wints/internship"
)

func (v *Service) PlanReport(student string, r internship.ReportHeader) error {
	err := ErrPermission
	if v.my.Role == internship.ROOT {
		err = v.srv.PlanReport(student, r)
	}
	v.UserLog("plan '"+r.Kind+"' report of '"+student+"'", err)
	return err
}

func (v *Service) ReportDefs() []internship.ReportDef {
	return v.srv.ReportDefs()
}

func (v *Service) Report(kind, email string) (internship.ReportHeader, error) {
	err := ErrPermission
	hdr := internship.ReportHeader{}
	if v.mine(email) || v.isTutoring(email) || v.my.Role >= internship.MAJOR {
		hdr, err = v.srv.Report(kind, email)
	}
	v.UserLog("get '"+kind+"' report header of '"+email+"'", err)
	return hdr, err
}

func (v *Service) ReportContent(kind, email string) ([]byte, error) {
	err := ErrPermission
	cnt := []byte{}
	if v.mine(email) || v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		cnt, err = v.srv.ReportContent(kind, email)
	}
	v.UserLog("download report '"+kind+"' of '"+email+"'", err)
	return cnt, err
}

func (v *Service) SetReportContent(kind, email string, cnt []byte) error {
	err := ErrPermission
	if v.mine(email) {
		err = v.srv.SetReportContent(kind, email, cnt)
	}
	v.UserLog("upload '"+kind+"' report content for '"+email+"'", err)
	if err == nil {
		if i, err2 := v.srv.Internship(email); err2 == nil {
			v.mailer.SendReportUploaded(i.Student, i.Tutor, kind)
		} else {
			v.UserLog("Unable to mail about the report upload", err2)
		}
	}
	return err
}

func (v *Service) SetReportGrade(kind, email string, r int, comment string) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		err = v.srv.SetReportGrade(kind, email, r, comment)
	}
	v.UserLog("set '"+kind+"' report grade for '"+email+"' to "+strconv.Itoa(r), err)
	if err == nil {
		i, err2 := v.srv.Internship(email)
		if err2 == nil {
			v.mailer.SendGradeUploaded(i.Student, i.Tutor, kind)
		} else {
			v.UserLog("Unable to mail about the report grade", err2)
		}
	}

	return err
}

func (v *Service) SetReportDeadline(kind, email string, t time.Time) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		err = v.srv.SetReportDeadline(kind, email, t)
	}
	v.UserLog("set '"+kind+"' report deadline to '"+t.String()+"'", err)
	if err == nil {
		i, err := v.srv.Internship(email)
		if err == nil {
			v.mailer.SendReportDeadline(i.Student, i.Tutor, kind, t)
		} else {
			v.UserLog("Unable to notify about the new deadline since: ", err)
		}
	}

	return err
}

func (v *Service) SetReportPrivate(kind, email string, p bool) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		err = v.srv.SetReportPrivate(kind, email, p)
	}
	st := "public"
	if p {
		st = "private"
	}
	v.UserLog("set '"+kind+"' report '"+st+"'for '"+email+"'", err)
	if err == nil {
		if i, err := v.srv.Internship(email); err != nil {
			v.mailer.SendReportPrivate(i.Student, i.Tutor, kind, p)
		}
	}

	return err
}
