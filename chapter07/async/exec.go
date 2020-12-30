package async

//FetchAll pega uma lista de urls
func FetchAll(urls []string, c *Client)  {
	for _, url := range urls {
		go c.AsyncGet(url)
	}
}