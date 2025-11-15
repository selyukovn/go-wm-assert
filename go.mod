module github.com/selyukovn/go-wm-assert

go 1.18

require github.com/stretchr/testify v1.11.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract v0.2.0 // CRITICAL: v0.2.0 does not work on Go 1.18-1.20 -- use v0.2.1+ instead!
