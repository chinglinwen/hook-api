module wen/hook-api

require (
	github.com/ashwanthkumar/slack-go-webhook v0.0.0-20180319063640-eb0e8e892f3a // indirect
	github.com/aws/aws-sdk-go v1.15.37 // indirect
	github.com/chinglinwen/checkup v0.3.1 // indirect
	github.com/chinglinwen/log v0.0.0-20180802093412-402fdc33bf76
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/google/go-github v17.0.0+incompatible // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/jmoiron/sqlx v0.0.0-20180614180643-0dae4fefe7c0 // indirect
	github.com/labstack/echo v3.2.1+incompatible
	github.com/labstack/gommon v0.2.7 // indirect
	github.com/lib/pq v1.0.0 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/miekg/dns v1.0.9 // indirect
	github.com/moul/http2curl v0.0.0-20170919181001-9ac6cf4d929b // indirect
	github.com/parnurzeal/gorequest v0.2.15 // indirect
	github.com/pkg/errors v0.8.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.2.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v0.0.0-20170224212429-dcecefd839c4 // indirect
	golang.org/x/crypto v0.0.0-20180910181607-0e37d006457b // indirect
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be // indirect
	golang.org/x/sys v0.0.0-20180918153733-ee1b12c67af4 // indirect
	google.golang.org/appengine v1.2.0 // indirect
	gopkg.in/resty.v1 v1.9.1 // indirect
	wen/hook-api/upstream v0.0.0

	wen/svc-d/check v0.0.0
	wen/svc-d/config v0.0.0 // indirect
)

replace wen/svc-d/check => ../svc-d/check

replace wen/svc-d/config => ../svc-d/config

replace wen/hook-api/upstream => ./upstream
