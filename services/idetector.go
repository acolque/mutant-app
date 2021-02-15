package services

type IDetector interface {
	IsMutant([]string) (bool, error)
}
