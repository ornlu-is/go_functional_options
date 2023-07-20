package person

type Person struct {
	ID      int
	Name    string
	Age     int
	Email   string
	Address string
}

type Option func(*Person)

func New(id int, opts ...Option) *Person {
	p := &Person{
		ID:      id,
		Name:    "John Doe",
		Age:     25,
		Email:   "johndoe@example.com",
		Address: "Nowhere",
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithName(name string) Option {
	return func(p *Person) {
		p.Name = name
	}
}

func WithAge(age int) Option {
	return func(p *Person) {
		p.Age = age
	}
}

func WithEmail(email string) Option {
	return func(p *Person) {
		p.Email = email
	}
}

func WithAddress(address string) Option {
	return func(p *Person) {
		p.Address = address
	}
}
