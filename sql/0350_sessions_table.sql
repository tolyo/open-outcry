-- +goose Up

CREATE TABLE sessions (
  beginstring CHAR(8) NOT NULL,
  sendercompid VARCHAR(64) NOT NULL,
  sendersubid VARCHAR(64) NOT NULL,
  senderlocid VARCHAR(64) NOT NULL,
  targetcompid VARCHAR(64) NOT NULL,
  targetsubid VARCHAR(64) NOT NULL,
  targetlocid VARCHAR(64) NOT NULL,
  session_qualifier VARCHAR(64) NOT NULL,
  creation_time TIMESTAMP WITH TIME ZONE NOT NULL,
  incoming_seqnum INTEGER NOT NULL, 
  outgoing_seqnum INTEGER NOT NULL,
  PRIMARY KEY (beginstring, sendercompid, sendersubid, senderlocid, 
  				targetcompid, targetsubid, targetlocid, session_qualifier)
);

-- +goose Down
DROP TABLE sessions CASCADE;