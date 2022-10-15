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

const TableMessagesCreationQuery = `CREATE TABLE IF NOT EXISTS messages
(
  id SERIAL,
  user TEXT NOT NULL,
  message TEXT NOT NULL,
  datetime DATETIME not null,
  CONSTRAINT messages_pkey PRIMARY KEY (id)
)
`
