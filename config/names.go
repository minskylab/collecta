package config

// GoogleClientID is the name of viper variable that saved the client google id for oauth2
const GoogleClientID = "google.clientID" // env: COLLECTA_GOOGLE_CLIENTID

// GoogleSecretKey is the of viper variable that saved the client google secret for oauth2
const GoogleSecretKey = "google.secret" //  env: COLLECTA_GOOGLE_SECRET

// AuthJWTExpiration is the name or selector for the duration of the jwt token
const AuthJWTExpiration = "auth.jwtExpiration" // env: COLLECTA_AUTH_JWTEXPIRATION

// AuthJWTSecret saves the HS256 jwt secret
const AuthJWTSecret = "auth.jwtSecret" // env: COLLECTA_AUTH_JWTSECRET

// CorsOrigins contains the list of allowed domains
const CorsOrigins = "cors.origins" // env: COLLECTA_CORS_ORIGINS

// CorsMethods contains the list of allowed methods
const CorsMethods = "cors.methods" // env: COLLECTA_CORS_METHODS

// CorsHeaders contains the list of allowed headers
const CorsHeaders = "cors.headers" // env: COLLECTA_CORS_HEADERS

// CorsMaxAge contains the max age for a valid cors policy
const CorsMaxAge = "cors.maxAge" // env: COLLECTA_CORS_MAXAGE

// DomainHost saves the name of the domain for the collecta instance service
const DomainHost = "domain" // COLLECTA_DOMAIN

// FirstAdminPassword is the first password time for the admin user
const FirstAdminPassword = "password"