package mocks

type SQLResult struct{}

func (sr SQLResult) LastInsertId() (int64, error) {
	return 1, nil
}
func (sr SQLResult) RowsAffected() (int64, error) {
	return 1, nil
}
