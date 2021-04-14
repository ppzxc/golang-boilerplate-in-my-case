package mariadb

func CreateUsers() string {
	return `
CREATE TABLE users (
  id              INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(256) NOT NULL,
  username        VARCHAR(256) NOT NULL UNIQUE,
  password        VARCHAR(256) NOT NULL UNIQUE,
  email           VARCHAR(256) NOT NULL,
  description     TEXT,
  register_date   DATETIME,
  last_login_date DATETIME
)
`
}
