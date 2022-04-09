package limiter

type Limiter interface {
	TryTake([]byte) (bool, int)
}
