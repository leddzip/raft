package script_folder

type Candidate struct {
	Path     string
	FileName string
	Name     string
}

type CandidateWithContent struct {
	Path     string
	FileName string
	Name     string
	Content  []byte
}
