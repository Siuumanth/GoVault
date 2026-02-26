func Connect(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	// --- ADD THESE LIMITS ---

	// 1. The Waiting Room: Max physical connections
	// Set this to ~80-90 if Postgres max_connections is 100
	db.SetMaxOpenConns(80)

	// 2. The Warm-up: How many idle connections to keep alive
	db.SetMaxIdleConns(20)

	// 3. The Refresh: Max age of a connection
	db.SetConnMaxLifetime(time.Minute * 5)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("connected to database with connection pooling")

	return db, nil
}