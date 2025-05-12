module github.com/anandvijayan/email-blast/send_mail

replace github.com/anandvijayan/email-blast/util/notify => ../util/notify

go 1.24.0

require (
	github.com/anandvijayan/email-blast/util/notify v0.0.0-00010101000000-000000000000
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20250501123311-99957c1e1255 // indirect
	github.com/sashabaranov/go-openai v1.31.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
