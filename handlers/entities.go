package handlers

type ErrorHTTP struct {
	Error string `json:"error"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
}

type Column struct {
	Name     string `json:"name"`
	Position int    `json:"pos"`
}

type Task struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Column      int    `json:"col"`
	Priority    int    `json:"prio"`
}

type Comment struct {
	Text string `json:"text"`
}
