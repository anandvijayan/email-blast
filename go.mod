module github.com/anandvijayan/email-blast

go 1.24.0

replace github.com/anandvijayan/email-blast/util/notify => ./util/notify

replace github.com/anandvijayan/email-blast/send_mail => ./send_mail

require (
	github.com/anandvijayan/email-blast/send_mail v0.0.0-00010101000000-000000000000
	github.com/anandvijayan/email-blast/util/notify v0.0.0-20250501123044-53cd37545c9e
	github.com/c-bata/go-prompt v0.2.6
)

require (
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/newhorizonsarizona/tmi-status-checker/util v0.0.0-20250501123311-99957c1e1255 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	github.com/sashabaranov/go-openai v1.39.1 // indirect
	golang.org/x/sys v0.0.0-20200918174421-af09f7315aff // indirect
	golang.org/x/time v0.11.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
