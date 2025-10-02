package da

type Client interface {
	List(namespace string) error
}
