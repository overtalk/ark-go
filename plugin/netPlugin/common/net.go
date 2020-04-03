package netCommon

type INet interface {
	Update() error
	StartClient() error
}
