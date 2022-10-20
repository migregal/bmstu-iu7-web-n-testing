package cache

type CacheInteractor interface {
	Update(storage string, id string, info any) error
	Get(storage string, id string) ([]any, error)
	Delete(storage string, id string) error
}
