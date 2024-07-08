package bootstrap

const (
	CREATE_USER_TABLE = `
    CREATE TABLE users (
      ID INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT
  );`
)
