package product

type service struct {
}

type Service interface {
	Hello() string
}

func NewService() Service {
	return &service{}
}

/*
db, err := sql.Open("pgx", "postgres://postgres:secret@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
*/

func (s *service) Hello() string {
	return "hola"
}
