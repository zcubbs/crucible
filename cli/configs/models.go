package configs

type Configuration struct {
	Logging        `mapstructure:"logging" json:"logging"`
	Keycloak       `mapstructure:"keycloak" json:"keycloak"`
	Kubeconfig     `mapstructure:"kubeconfig" json:"kubeconfig"`
	DefaultUsers   *[]DefaultUser   `mapstructure:"defaultUsers" json:"defaultUsers"`
	DefaultClients *[]DefaultClient `mapstructure:"defaultClients" json:"defaultClients"`
}

type Keycloak struct {
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Secure   bool   `mapstructure:"secure" json:"secure"`
	Url      string `mapstructure:"url" json:"url"`
	Realm    string `mapstructure:"realm" json:"realm"`
}

type Logging struct {
	Enabled bool `mapstructure:"enabled" json:"enabled"`
}

type DefaultUser struct {
	Username  string `mapstructure:"username" json:"username"`
	Password  string `mapstructure:"password" json:"password"`
	Email     string `mapstructure:"email" json:"email"`
	FirstName string `mapstructure:"firstName" json:"firstName"`
	LastName  string `mapstructure:"lastName" json:"lastName"`
}

type DefaultClient struct {
	ClientId string `mapstructure:"clientId" json:"clientId"`
}

type Kubeconfig struct {
	Path string `mapstructure:"path" json:"path"`
}
