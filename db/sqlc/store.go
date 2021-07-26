type Store struct {
	*Queries
	db *sql.DB
}