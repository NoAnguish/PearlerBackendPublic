package database

func Modify(query string, s *Session) error {
	needToCloseFlag := false
	var err error

	if s == nil {
		s, err = PrepareDefaultWriteSession()
		if err != nil {
			return err
		}
		needToCloseFlag = true
	}

	err = s.modify(query)
	if needToCloseFlag {
		s.Close()
	}

	return err
}
