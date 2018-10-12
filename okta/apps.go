package okta

import (
	"fmt"
	"net/url"
	"time"
)
const (
	appUserIDFilter       	 = "user.id"
)
// AppsService is a service to retreives applications from OKTA.
type AppsService service

// AppFilterOptions is used to generate a "Filter" to search for different Apps
// The values here coorelate to API Search paramgters on the group API
type AppFilterOptions struct {
	UserIDEqualTo string   `url:"-"`

	FilterString  string     `url:"filter,omitempty"`
	NextURL       *url.URL `url:"-"`
	GetAllPages   bool     `url:"-"`
	NumberOfPages int      `url:"-"`
	Limit         int      `url:"limit,omitempty"`
	ExpandUser        string      `url:"expand,omitempty"`
}

// App is the Model for an OKTA Application
type App struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Label         string    `json:"label"`
	Status        string    `json:"status"`
	LastUpdated   time.Time `json:"lastUpdated"`
	Created       time.Time `json:"created"`
	Accessibility struct {
		SelfService      bool        `json:"selfService"`
		ErrorRedirectURL interface{} `json:"errorRedirectUrl"`
		LoginRedirectURL interface{} `json:"loginRedirectUrl"`
	} `json:"accessibility"`
	Visibility struct {
		AutoSubmitToolbar bool `json:"autoSubmitToolbar"`
		Hide              struct {
			IOS bool `json:"iOS"`
			Web bool `json:"web"`
		} `json:"hide"`
		AppLinks struct {
			TestorgoneCustomsaml20App1Link bool `json:"testorgone_customsaml20app_1_link"`
		} `json:"appLinks"`
	} `json:"visibility"`
	Features    []interface{} `json:"features"`
	SignOnMode  string        `json:"signOnMode"`
	Credentials struct {
		UserNameTemplate struct {
			Template string `json:"template"`
			Type     string `json:"type"`
		} `json:"userNameTemplate"`
		Signing struct {
		} `json:"signing"`
	} `json:"credentials"`
	Settings struct {
		App struct {
		} `json:"app"`
		Notifications struct {
			Vpn struct {
				Network struct {
					Connection string `json:"connection"`
				} `json:"network"`
				Message interface{} `json:"message"`
				HelpURL interface{} `json:"helpUrl"`
			} `json:"vpn"`
		} `json:"notifications"`
		SignOn struct {
			DefaultRelayState     string        `json:"defaultRelayState"`
			SsoAcsURL             string        `json:"ssoAcsUrl"`
			IdpIssuer             string        `json:"idpIssuer"`
			Audience              string        `json:"audience"`
			Recipient             string        `json:"recipient"`
			Destination           string        `json:"destination"`
			SubjectNameIDTemplate string        `json:"subjectNameIdTemplate"`
			SubjectNameIDFormat   string        `json:"subjectNameIdFormat"`
			ResponseSigned        bool          `json:"responseSigned"`
			AssertionSigned       bool          `json:"assertionSigned"`
			SignatureAlgorithm    string        `json:"signatureAlgorithm"`
			DigestAlgorithm       string        `json:"digestAlgorithm"`
			HonorForceAuthn       bool          `json:"honorForceAuthn"`
			AuthnContextClassRef  string        `json:"authnContextClassRef"`
			SpIssuer              interface{}   `json:"spIssuer"`
			RequestCompressed     bool          `json:"requestCompressed"`
			AttributeStatements   []interface{} `json:"attributeStatements"`
		} `json:"signOn"`
	} `json:"settings"`
	Embedded struct {
		User AppUser `json:"user"`
	} `json:"_embedded"`
	Links struct {
		Logo []struct {
			Name string `json:"name"`
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"logo"`
		AppLinks []struct {
			Name string `json:"name"`
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"appLinks"`
		Help struct {
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"help"`
		Users struct {
			Href string `json:"href"`
		} `json:"users"`
		Deactivate struct {
			Href string `json:"href"`
		} `json:"deactivate"`
		Groups struct {
			Href string `json:"href"`
		} `json:"groups"`
		Metadata struct {
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"metadata"`
	} `json:"_links"`
}

func (a App) String() string {
	// return Stringify(g)
	return fmt.Sprintf("App:(ID: {%v} - Name: {%v})\n", a.ID, a.Name)
}

// GetByID gets a group from OKTA by the Group ID. An error is returned if the group is not found
func (a *AppsService) GetByID(appID string) (*App, *Response, error) {

	u := fmt.Sprintf("apps/%v", appID)
	req, err := a.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	app := new(App)

	resp, err := a.client.Do(req, app)

	if err != nil {
		return nil, resp, err
	}

	return app, resp, err
}

// AppUser is the model for a user of an OKTA App
type AppUser struct {
	ID              string     `json:"id"`
	ExternalID      string     `json:"externalId"`
	Name 			string     `json:"name"`
	Label 			string	   `json:"label"`
	Created         time.Time  `json:"created"`
	LastUpdated     time.Time  `json:"lastUpdated"`
	Scope           string     `json:"scope"`
	Status          string     `json:"status"`
	StatusChanged   *time.Time `json:"statusChanged"`
	PasswordChanged *time.Time `json:"passwordChanged"`
	SyncState       string     `json:"syncState"`
	LastSync        *time.Time `json:"lastSync"`
	Accessibility   struct{
		SelfService		bool	`json:"selfService"`
		ErrorRedirectURL string	`json:"errorRedirectURL"`
		LoginRedirectURL string `json:"loginRedirectURL"`
	} `json:"accessibility"`
	Visibility struct{
		AutoSubmitToolbar	bool `json:"autoSubmitToolbar"`
		Hide				struct{
			iOS				bool `json:"iOS"`
			Web				bool  `json:"Web"`
		}`json:"hide"`

	} `json:visibility`
	SignOnMode      string		`json:"signOnMode"`

	Credentials     struct {
		UserName string `json:"userName"`
		Password struct {
		} `json:"password"`
	} `json:"credentials"`
	Profile struct {
		SecondEmail      interface{} `json:"secondEmail"`
		LastName         string      `json:"lastName"`
		MobilePhone      interface{} `json:"mobilePhone"`
		Email            string      `json:"email"`
		SalesforceGroups []string    `json:"salesforceGroups"`
		Role             string      `json:"role"`
		FirstName        string      `json:"firstName"`
		Profile          string      `json:"profile"`
		SamlRoles        []string    `json:"samlRoles"`
	} `json:"profile"`
	Links struct {
		App struct {
			Href string `json:"href"`
		} `json:"app"`
		User struct {
			Href string `json:"href"`
		} `json:"user"`
	} `json:"_links"`
}

// GetUsers returns the members in an App
//   Pass in an optional AppFilterOptions struct to filter the results
//   The Users in the app are returned
func (a *AppsService) GetUsers(appID string, opt *AppFilterOptions) (appUsers []AppUser, resp *Response, err error) {

	pagesRetreived := 0
	var u string
	if opt.NextURL != nil {
		u = opt.NextURL.String()
	} else {
		u = fmt.Sprintf("apps/%v/users", appID)

		if opt.Limit == 0 {
			opt.Limit = defaultLimit
		}

		u, _ = addOptions(u, opt)
	}

	req, err := a.client.NewRequest("GET", u, nil)

	if err != nil {
		// fmt.Printf("____ERROR HERE\n")
		return nil, nil, err
	}
	resp, err = a.client.Do(req, &appUsers)

	if err != nil {
		// fmt.Printf("____ERROR HERE 2\n")
		return nil, resp, err
	}

	pagesRetreived++

	if (opt.NumberOfPages > 0 && pagesRetreived < opt.NumberOfPages) || opt.GetAllPages {

		for {

			if pagesRetreived == opt.NumberOfPages {
				break
			}
			if resp.NextURL != nil {

				var userPage []AppUser
				pageOpts := new(AppFilterOptions)
				pageOpts.NextURL = resp.NextURL
				pageOpts.Limit = opt.Limit
				pageOpts.NumberOfPages = 1

				userPage, resp, err = a.GetUsers(appID, pageOpts)

				if err != nil {
					return appUsers, resp, err
				}
				appUsers = append(appUsers, userPage...)
				pagesRetreived++
			} else {
				break
			}

		}
	}

	return appUsers, resp, err
}

// AppGroups - Groups assigned to Application
type AppGroups struct {
	ID          string    `json:"id"`
	LastUpdated time.Time `json:"lastUpdated"`
	Priority    int       `json:"priority"`
	Profile		struct{
		Role	string		`json:"role,omitempty"`
		SamlRoles []string `json:"samlRoles,omitempty"`
	} `json:"profile,omitempty"`
	Links       struct {
		User struct {
			Href string `json:"href"`
		} `json:"user"`
	} `json:"_links"`
}

// GetGroups returns groups assigned to the application - Input appID is the Application GUID
func (a *AppsService) GetGroups(appID string) (appGroups []AppGroups, resp *Response, err error) {

	var u string
	u = fmt.Sprintf("apps/%v/groups?limit=100", appID)

	req, err := a.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	resp, err = a.client.Do(req, &appGroups)

	if err != nil {
		return nil, resp, err
	}

	for {

		if resp.NextURL != nil {

			var appGroupPage []AppGroups

			appGroupPage, resp, err = a.GetGroups(appID)

			if err != nil {
				return appGroups, resp, err
			} else {
				appGroups = append(appGroups, appGroupPage...)

			}
		} else {
			break
		}

	}

	return appGroups, resp, err
}

// GetUser returns the AppUser model for one app users
func (a *AppsService) GetUser(appID string, userID string) (appUser AppUser, resp *Response, err error) {

	var u string
	u = fmt.Sprintf("apps/%v/users/%v", appID, userID)

	req, err := a.client.NewRequest("GET", u, nil)

	if err != nil {
		return appUser, nil, err
	}
	resp, err = a.client.Do(req, &appUser)

	if err != nil {
		return appUser, resp, err
	}
	return appUser, resp, nil
}

func (a *AppsService) ListWithFilter(opt *AppFilterOptions) ([]App, *Response, error){

	var u string
	var err error

	pagesRetreived := 0

	if opt.NextURL != nil {
		u = opt.NextURL.String()
	} else {
		if opt.UserIDEqualTo != "" {
			opt.FilterString = appendToFilterString(opt.FilterString, appUserIDFilter, FilterEqualOperator, opt.UserIDEqualTo)
			if opt.ExpandUser != "" {
				opt.ExpandUser = fmt.Sprintf("user/%v", opt.UserIDEqualTo)
			}
		}


		if opt.Limit == 0 {
			opt.Limit = defaultLimit
		}

		u, err = addOptions("apps", opt)

	}
	req, err := a.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	appUser := make([]App, 1)
	resp, err := a.client.Do(req, &appUser)

	pagesRetreived++
	if (opt.NumberOfPages > 0 && pagesRetreived < opt.NumberOfPages) || opt.GetAllPages {
		for {

			if pagesRetreived == opt.NumberOfPages {
				break
			}
			if resp.NextURL != nil {

				var appPage []App
				pageOpts := new(AppFilterOptions)
				pageOpts.NextURL = resp.NextURL
				pageOpts.Limit = opt.Limit
				pageOpts.NumberOfPages = 1

				appPage, resp, err = a.ListWithFilter(pageOpts)
				if err != nil {
					return appUser, resp, err
				}
				appUser = append(appUser, appPage...)
				pagesRetreived++
			} else {
				break
			}
		}
	}
	return appUser, resp, err
}