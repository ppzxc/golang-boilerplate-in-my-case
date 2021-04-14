package mariadb

func Create() string {
	return `
CREATE TABLE users (
  id              INT(11) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(256) NOT NULL,
  username        VARCHAR(256) NOT NULL,
  password        VARCHAR(256) NOT NULL,
  email           VARCHAR(256) NOT NULL,
  description     TEXT,
  register_date   DATETIME,
  last_login_date DATETIME
)
`
}
