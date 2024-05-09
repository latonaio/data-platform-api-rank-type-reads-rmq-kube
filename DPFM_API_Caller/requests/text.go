package requests

type Text struct {
	RankType     		string  `json:"RankType"`
	Language          	string  `json:"Language"`
	RankTypeName		string  `json:"RankTypeName"`
	CreationDate		string	`json:"CreationDate"`
	LastChangeDate		string	`json:"LastChangeDate"`
	IsMarkedForDeletion	*bool	`json:"IsMarkedForDeletion"`
}
