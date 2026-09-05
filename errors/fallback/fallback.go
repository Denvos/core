package fallback

type Fallback interface {
	Fallback(err error) error
}

type Func func(err error) error

func (f Func) Fallback(err error) error {
	return f(err)
}

type FallbackChain []Fallback

func (f FallbackChain) Fallback(err error) error {
	for _, fb := range f {
		if newErr := fb.Fallback(err); newErr != nil {
			return newErr
		}
	}
	return err
}
