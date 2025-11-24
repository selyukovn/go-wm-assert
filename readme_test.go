package assert

import (
	"fmt"
	tAssert "github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func Test_Readme(t *testing.T) {
	// Argument Assertion
	// --------------------------------

	t.Run("ArgumentAssertion", func(t *testing.T) {
		type Account struct{}
		type EventCollection struct{}

		Deactivate := func(a *Account, deactivatedAt time.Time, evs *EventCollection) error {
			Time().NotZero().LessEq(time.Now()).Must(deactivatedAt)
			Cmp[*EventCollection]().NotEq(nil).Must(evs)
			// ...
			return nil
		}

		tAssert.Panics(t, func() { _ = Deactivate(&Account{}, time.Time{}, &EventCollection{}) })
		tAssert.Panics(t, func() { _ = Deactivate(&Account{}, time.Now(), nil) })
		tAssert.Panics(t, func() {
			var evs *EventCollection
			_ = Deactivate(&Account{}, time.Now(), evs)
		})
		tAssert.NotPanics(t, func() { _ = Deactivate(&Account{}, time.Now(), &EventCollection{}) })
	})

	// Simple Validation
	// --------------------------------

	t.Run("SimpleValidation", func(t *testing.T) {
		type Name struct{ value string }

		NameFromString := func(value string) (Name, error) {
			err := Str().Word().Check(value, fmt.Sprintf("Name %q is incorrect!", value))

			if err != nil {
				return Name{}, err
			}

			return Name{value: value}, nil
		}

		for _, tCase := range []string{"Hello!", "W o r l d", "12345", "", "-"} {
			_, err := NameFromString(tCase)
			tAssert.Error(t, err)
			tAssert.Equal(t, fmt.Sprintf("Name %q is incorrect!", tCase), err.Error())
		}
	})

	// Form Validation
	// --------------------------------

	t.Run("FormValidation", func(t *testing.T) {
		type SignUpForm struct {
			email     string
			name      string // let it be optional
			age       uint
			agreement bool

			errors map[string][]error
		}

		emailRegexpCompiled := regexp.MustCompile("^\\w+@\\w+\\.\\w+$" /* simplified for demonstration */)

		Validate := func(f *SignUpForm) bool {
			f.errors = map[string][]error{
				"email":     {},
				"name":      {},
				"age":       {},
				"agreement": {},
			}

			f.errors["email"] = Str().
				NotEmpty("Email is required!").
				Regexp(
					emailRegexpCompiled,

					// custom error message only for this rule
					// instead of technical "value ... regexp ..."
					"Email is incorrect!",
				).
				Custom(func(v string) error {
					// e.g. check that it is not registered previously
					return nil
				}).
				CheckAll(f.email)

			// optional field, remember?
			if f.name != "" {
				f.errors["name"] = Str().
					Word("Only letters and '-' allowed!").
					RunesMin(2, "Too short, isn't it?").
					RunesMax(255, "Too long, isn't it?").
					NotIn(
						[]string{ /* e.g. some set of bad words or so */ },
						"Is that your real name, friend?",
					).
					CheckAll(f.name)
			}

			f.errors["age"] = Num[uint]().
				GreaterEq(18, "Things are serious -- come back later!").
				Less(65, "Take a rest, friend!").
				CheckAll(f.age)

			f.errors["agreement"] = Bool().
				True("This flag is required!").
				CheckAll(f.agreement)

			return len(f.errors["email"]) == 0 &&
				len(f.errors["name"]) == 0 &&
				len(f.errors["age"]) == 0 &&
				len(f.errors["agreement"]) == 0
		}

		// ----

		form := &SignUpForm{}

		form.email = ""
		form.name = ""
		form.age = 17
		form.agreement = false
		tAssert.False(t, Validate(form))
		tAssert.Equal(t, "Email is required!", form.errors["email"][0].Error())
		tAssert.Empty(t, form.errors["name"])
		tAssert.Equal(t, "Things are serious -- come back later!", form.errors["age"][0].Error())
		tAssert.Equal(t, "This flag is required!", form.errors["agreement"][0].Error())

		form.email = "old@man.email"
		form.name = "I"
		form.age = 100
		form.agreement = true
		tAssert.False(t, Validate(form))
		tAssert.Empty(t, form.errors["email"])
		tAssert.Equal(t, "Too short, isn't it?", form.errors["name"][0].Error())
		tAssert.Equal(t, "Take a rest, friend!", form.errors["age"][0].Error())
		tAssert.Empty(t, form.errors["agreement"])

		form.email = "assert@package.test"
		form.name = "Test"
		form.age = 33
		form.agreement = true
		tAssert.True(t, Validate(form))
		tAssert.Empty(t, form.errors["email"])
		tAssert.Empty(t, form.errors["name"])
		tAssert.Empty(t, form.errors["age"])
		tAssert.Empty(t, form.errors["agreement"])
	})
}
