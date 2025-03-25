package utils

type WithLinks interface {
	Links(apiRoot string) (links []string)
}
