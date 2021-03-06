CREATE TABLE authors (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  first_name VARCHAR(32) NOT NULL,
  last_name VARCHAR(32),
  uuid UUID UNIQUE NOT NULL
);

CREATE TABLE posts (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  author_uuid UUID REFERENCES authors(uuid),
  timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
  body TEXT,
  uuid UUID UNIQUE NOT NULL
);
