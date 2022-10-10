CREATE TABLE metric (
  id VARCHAR(255) PRIMARY KEY,
  created_at DATETIME,
  calls INT,
  scammers INT,
  inactives INT,
  call_time INT,
  talk_time INT,
  amount MONEY
);

CREATE TABLE scammer (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255),
  tel VARCHAR(255),
  is_active BOOLEAN
);

CREATE TABLE callhistory (
  id VARCHAR(255) PRIMARY KEY,
  scammer_id VARCHAR(255) REFERENCES scammer(id),
  call_at DATETIME,
  call_time INT,
  result BOOLEAN,
  talk_time INT
);

CREATE TABLE matching (
  id VARCHAR(255) PRIMARY KEY,
  created_at DATETIME,
  serial_number   BIGSERIAL,
  matched        BOOLEAN,
  checked BOOLEAN,
  matching_at DATETIME,
  talk_time INT,
  transcript TEXT
);

CREATE TABLE scammer_tag (
  scammer_id VARCHAR(255) REFERENCES scammer(id),
  tag VARCHAR(255)
);

CREATE TABLE scammer_matching (
  scammer_id VARCHAR(255) REFERENCES scammer(id),
  matching_id VARCHAR(255) REFERENCES matching(id)
);

CREATE TABLE scammer_call (
  scammer_id VARCHAR(255) REFERENCES scammer(id),
  call_id VARCHAR(255) REFERENCES callhistory(id)
);
