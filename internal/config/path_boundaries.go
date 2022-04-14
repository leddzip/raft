package config

type pathBoundaries struct {
	From   string
	BackTo string
}

var PathBoundaries = pathBoundaries{
	From:   ".",
	BackTo: "/",
}
