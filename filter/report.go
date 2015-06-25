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
	return err
}

func (v *Service) SetReportGrade(kind, email string, r int, comment string) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		err = v.srv.SetReportGrade(kind, email, r, comment)
	}
	v.UserLog("set '"+kind+"' report grade for '"+email+"' to "+strconv.Itoa(r), err)
	return err
}

func (v *Service) SetReportDeadline(kind, email string, t time.Time) error {
	err := ErrPermission
	if v.my.Role >= internship.ADMIN || v.isTutoring(email) {
		err = v.srv.SetReportDeadline(kind, email, t)
	}
	v.UserLog("set '"+kind+"' report deadline to '"+t.String()+"'", err)
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
	return err
}
