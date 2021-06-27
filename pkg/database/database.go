package database

type Client interface {
	Exec(query string, params []interface{}) error
	QueryRow(query string, params []interface{}, columns []string) (map[string]interface{}, error)
	QueryRows(query string, params []interface{}, columns []string) ([]map[string]interface{}, error)
	Ping() error
}

func Exec(query string, params []interface{}) error {
	c, err := getClient()
	if err != nil {
		return err
	}

	return c.Exec(query, params)
}

func QueryRow(query string, params []interface{}, columns []string) (map[string]interface{}, error) {
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	return c.QueryRow(query, params, columns)
}

func QueryRows(query string, params []interface{}, columns []string) ([]map[string]interface{}, error) {
	c, err := getClient()
	if err != nil {
		return nil, err
	}

	return c.QueryRows(query, params, columns)
}
