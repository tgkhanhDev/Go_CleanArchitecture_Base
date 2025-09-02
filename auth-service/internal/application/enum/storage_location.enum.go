package enumerate

type StorageLocation int

const (
	LOCAL StorageLocation = iota
	AWS_S3
)

func (s StorageLocation) getName() string {
	return [...]string{"LOCAL", "AWS_S3"}[s]
}
