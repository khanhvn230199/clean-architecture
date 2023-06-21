package env

type EnvType int

var env EnvType
var isDev = false
var notProd = true

// Environment constants
const (
	EnvLocal EnvType = iota + 1
	EnvDev
	EnvStaging
	EnvProd
)

var envNames = map[EnvType]string{
	EnvLocal:   "local",
	EnvDev:     "dev",
	EnvStaging: "staging",
	EnvProd:    "prod",
}

func (e EnvType) String() string {
	return envNames[e]
}

var envValues = map[string]EnvType{
	"local":   EnvLocal,
	"dev":     EnvDev,
	"staging": EnvStaging,
	"prod":    EnvProd,
}

func Env() EnvType {
	return env
}

func IsDev() bool {
	return env == EnvDev
}

func IsProd() bool {
	return env == EnvProd
}

func IsStaging() bool {
	return env == EnvStaging
}

func ÍsLocal() bool {
	return env == EnvLocal
}

func SetEnvironment(e string) EnvType {
	if env != 0 {
		panic("Already initialize environment")
	}
	env = envValues[e]
	switch env {
	case EnvDev:
		isDev = true
	case EnvStaging:
	case EnvProd:
		notProd = false
	default:
		panic("invalid environment: " + e)
	}
	return env
}
