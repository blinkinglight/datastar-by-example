# datastar by example
## golang + https://data-star.dev examples

to run these examples with live rebuild on code change you need `task` `go install github.com/go-task/task/v3/cmd/task@latest`
and standalone tailwindcss cli `https://github.com/tailwindlabs/tailwindcss/releases/tag/v4.1.11` ( put into /usr/local/bin )

infinite scroll example:
`task -w intersect`

content type: form example
`task -w form`

send javascript exceptions to backend example
`task -w errortobackend`

load resurses once dom ready with delay
`task -w laodwithdelay` 