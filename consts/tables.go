package consts

const TableUsersCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
  id SERIAL,
  name TEXT NOT NULL,
  login TEXT NOT NULL,
  password TEXT NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)
`
