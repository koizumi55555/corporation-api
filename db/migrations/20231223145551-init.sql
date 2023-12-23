
-- +migrate Up
CREATE TABLE corporation
(
    corporation_id  VARCHAR(36) NOT NULL,
    name VARCHAR(50),
    domain VARCHAR(100),
    number INT,
	corp_type VARCHAR(100),
PRIMARY KEY (corporation_id)
);

-- +migrate Down

DROP TABLE corporation;