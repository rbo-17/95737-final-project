USE CMU95737;

CREATE TABLE KEY_VALUE
(
    ID    VARCHAR(20) NOT NULL PRIMARY KEY, -- 20 - max len for a 64 bit integer (as a string)
    Value MEDIUMBLOB -- TODO: Use BLOB for small & large text
);
