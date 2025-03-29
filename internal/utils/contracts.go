package utils

// Entity that can generate multiple
// links based on its internal state
// with provided API root.
type WithLinks interface {
	Links(apiRoot string) (links []string)
}
