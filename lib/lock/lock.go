package lock

type Lock interface {
	Lock()
	TryLock() bool
	Unlock()
}
