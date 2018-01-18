package nsqclient

func retry(num int, fn func() error) error {
	var err error
	for i := 0; i < num; i++ {
		err = fn()
		if err == nil {
			break
		}
	}
	return err
}