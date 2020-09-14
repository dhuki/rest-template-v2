package common

const ENV_PATH = "./.env"

// testing should use fullpath if want to load file
const ENV_PATH_TESTING = "C:\\Users\\pricy\\go\\src\\github.com\\rest-template\\.env"

// const for resonse success
const RESPONSE_SUCCESS = true
const RESPONSE_MSG_SUCCESS = "Success"

// const for request model
const (
	ContextKeyRequestModel contextKey = iota
)
