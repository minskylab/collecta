package collecta

// import "github.com/minskylab/collecta/auth"

type Auth interface {
	Identify(i *Instance) (string, error)
}

type Instance struct {
	Endpoint string
	token string
}

type UserPasswordToken struct {
	username string
	password string
}

func (upt *UserPasswordToken) Identify(i *Instance) (string, error) {
	// auth.Auth{}.
	// upt.username
	// upt.password
	return "", nil
}

func NewClient(endpoint string, auth ...Auth) *Instance {
	inst := new(Instance)
	inst.Endpoint = endpoint

	for _, a := range auth {
		token, err := a.Identify(inst)
		if err != nil {
			return nil
		}
		inst.token = token
	}

	return inst

}