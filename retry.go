package retry

// Fn is the retry function signature.
type Fn func() (bail bool, err error)

// NotifyFn defines a function that is called upon each failed attempt.
type NotifyFn func(attempt uint16, err error)

// Run retries the provided function the specified number of times
func Run(max uint16, fn Fn, notifyFn NotifyFn) (err error) {
	var bail bool
	for i := uint16(1); i <= max; i++ {
		if bail, err = fn(); err == nil || bail {
			return
		}
		if notifyFn != nil {
			notifyFn(i, err)
		}
	}
	return
}
