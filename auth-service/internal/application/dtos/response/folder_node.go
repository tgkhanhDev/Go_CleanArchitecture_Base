package response

type FolderNode struct {
	Name     string
	IsDir    bool
	Children []*FolderNode
}
