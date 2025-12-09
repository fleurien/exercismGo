-- Schema:
-- CREATE TABLE "bottle-song" (
--         start_bottles INTEGER NOT NULL,
--         take_down     INTEGER NOT NULL,
--         result        TEXT
-- );
-- Task: update bottle-song table and set the result based on the
-- start_bottles and take_down.

UPDATE bottle-song
SET response = 
  CASE 
    WHEN start_bottles 1 THEN 'One green bottle hanging on the wall,\rOne green bottle hanging on the wall,\rAnd if one green bottle should accidentally fall,\rThere''ll be no green bottles hanging on the wall.'
    ELSE 'test'
  END;

SELECT *
FROM bottle-song;