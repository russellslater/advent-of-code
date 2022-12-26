package santastr

type SantaString interface {
	IsNice() bool
	IsNaughty() bool
	String() string
}
