## Join, Union, Aggregate, Subquery, Function (DBMS)

-	Belajar untuk mampu menggunakan statement SQL DDL.

*Order By, Group By, Limit, Between, Having*

*ONE to ONE relationship*, contoh 1 user hanya memiliki 1 foto profile
*ONE to MANY relationship*, contoh 1 user bisa memiliki banyak tweets
*MANY to MANY relationship*, contoh 1 user bisa memiliki banyak follower user, 1 user bisa di follow banyak user

*Join*, mendapatkan 2 atau lebih data dari tabel yang saling berhubungan atau manipulasi dari relasi tabel tersebut.

- INNER
- OUTER
- LEFT
- RIGHT

*Union*, menggabungkan 2 atau lebih data dari tabel menjadi satu output.

- SELECT ... UNION SELECT ...

*Subquery*, mendapatkan data dari diproses lain untuk di proses lagi pada tahap selanjutnya

- WHERE ... IN (...)
- WHERE ... NOT IN (...) 

*Group By*, mengumpulkan data yang sesuai pada kelompoknya dan manipulasi pada tahap penyatuan, mirip *DISTINCT*.

- SELECT COUNT(*) FROM ... WHERE ... GROUP BY ...

*Order By*, mensorting data tabel, sesuai opsi yang diberikan *ASCENDING* atau *DESCENDING*.

- SELECT COUNT(*) FROM ... WHERE ... GROUP BY ... ORDER BY ...

*Limit*, melimitasi data tabel sesuai kebutuhan, mencegah adanya overflow data yang tidak dibutuhkan.

- SELECT ... LIMIT ...
- SELECT ... LIMIT @offset, @size

*Between*, mendapatkan data dari range tertentu.

*&lt;*, kurang dari
*&gt;* lebih dari

- SELECT ... WHERE ... BETWEEN @&lt; AND @&gt;

*Having*, sebuah logic operator dengan data komparasi yang bisa di kumpulkan dengan aggregate functions

- $op {=, <=, >=, !=}
- @n is a number
- SELECT ... HAVING SUM(*) $op @n

*DDL, DML, DCL*

*DDl*, data definition language

- USE ...
- CREATE DATABASE ...
- CREATE TABLE ...
- RENAME TABLE ...
- DROP TABLE ...
- ALTER TABLE ...

*DML*, data manipulation language

- INSERT ...
- SELECT ...
- UPDATE ...
- DELETE ...
- LIKE ...
- BETWEEN ...
- AND ...
- OR ...
- ORDER BY ...
- LIMIT ...
- JOIN ... ON ...
- UNION ...
- WHERE ... IN ...

contoh where ... in, ... WHERE name IN (...) -- subquery

*AGGREGATE*

- MIN()
- MAX()
- SUM()
- AVG()
- COUNT()
- HAVING ...

contoh having, ... HAVING SUM(price) > 2000

*Create Function*

```sql
CREATE FUNCTION sf_count_tweet_user(user_id_p) RETURNS INT DETERMINISTIC
BEGIN
DECLARE total INT;
SELECT COUNT(*) INTO total FROM tweets WHERE user_id = user_id_p AND type = 'tweets';
RETURN total;
END
```

*DCL*, data control language

- digunakan untuk manipulasi permission requirement dan control user account access.
- *GRANT*, *CREATE USER*, *DROP USER* 
