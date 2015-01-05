package mailer

const (
	reset = `Subject:[Wints] Password reset

A request has been emitted to reset your account.
Click the following link to generate a new password if needed:

{{.WWW}}/resetPassword?token={{.Token}}

If you do not want a new password, simply ignore this mail

Regards
Fabien Hermenier`
	student_welcome = `Subject:[Wints] Account created

Wints is the web application used to follow your internship.
This webapp allows you to:
- update if needed the details of your internship. Especially, it is expected you to provide accurate contact data.
- upload your reports for evaluation
- access the feedbacks of your supervisors

Visit the following link to terminate the account creation:
{{.WWW}}/static/newpassword.html?email={{.C.Stu.P.Email}}&token={{.Token}}

Regards,
Fabien Hermenier`
)
