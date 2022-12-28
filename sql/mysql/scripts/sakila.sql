-- 使用 sakila 資料庫
USE sakila;

-- Query
SELECT language_id,
       'COMMON' AS language_usage,
       language_id * 3.1415927 AS lang_pi_value,
       UPPER(name) AS language_name
FROM language;

SELECT DISTINCT actor_id FROM film_actor ORDER BY actor_id;
