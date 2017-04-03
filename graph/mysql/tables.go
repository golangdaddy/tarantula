package mysql

import "fmt"

func (client *Client) Table(id string) string {
	
	return fmt.Sprintf("%s.%s", client.Credentials.DatabaseName(), id)
}