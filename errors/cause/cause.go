package cause

func Root(err error) error {
	if err == nil {
		return nil
	}
	for {
		next := Unwrap(err)
		if next == nil {
			return err
		}
		err = next
	}
}

func Unwrap(err error) error {
	type unwrapper interface {
		Unwrap() error
	}
	if u, ok := err.(unwrapper); ok {
		return u.Unwrap()
	}
	return nil
}
