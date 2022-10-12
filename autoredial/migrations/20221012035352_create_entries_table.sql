-- migrate:up
CREATE TABLE metric (
  id VARCHAR(255) PRIMARY KEY,
  created_at DATE NOT NULL,
  calls BIGINT NOT NULL,
  scammers BIGINT NOT NULL,
  inactives BIGINT NOT NULL,
  call_time INTERVAL NOT NULL,
  talk_time INTERVAL NOT NULL,
  amount BIGINT NOT NULL
);

CREATE TABLE scammer (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  tel VARCHAR(255) NOT NULL,
  is_active BOOLEAN NOT NULL
);

CREATE TABLE callhistory (
  id VARCHAR(255) PRIMARY KEY,
  scammer_id VARCHAR(255) REFERENCES scammer(id) NOT NULL,
  call_at DATE NOT NULL,
  call_time INTERVAL NOT NULL,
  result BOOLEAN NOT NULL,
  talk_time INTERVAL
);

CREATE TABLE matching (
  id VARCHAR(255) PRIMARY KEY,
  created_at DATE NOT NULL,
  serial_number   BIGSERIAL NOT NULL,
  matched        BOOLEAN NOT NULL,
  checked BOOLEAN NOT NULL,
  matching_at DATE,
  talk_time INTERVAL,
  transcript TEXT
);

CREATE TABLE scammer_tag (
  scammer_id VARCHAR(255) REFERENCES scammer(id) NOT NULL,
  tag VARCHAR(255) NOT NULL
);

CREATE TABLE scammer_matching (
  scammer_id VARCHAR(255) REFERENCES scammer(id) NOT NULL,
  matching_id VARCHAR(255) REFERENCES matching(id) NOT NULL
);

CREATE TABLE scammer_call (
  scammer_id VARCHAR(255) REFERENCES scammer(id) NOT NULL,
  call_id VARCHAR(255) REFERENCES callhistory(id) NOT NULL
);


-- migrate:down
DROP TABLE metric;
DROP TABLE scammer;
DROP TABLE callhistory;
DROP TABLE matching;
DROP TABLE scammer_tag;
DROP TABLE scammer_matching;
DROP TABLE scammer_call;
