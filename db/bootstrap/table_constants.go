package bootstrap

const (
	CREATE_USER_TABLE = `
    CREATE TABLE users (
      ID INTEGER PRIMARY KEY AUTOINCREMENT,
      username TEXT,
      password TEXT,
      createdOn DATETIME,
      lastLogin DATETIME
  );`
)
