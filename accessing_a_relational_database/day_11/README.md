<!-- Set up a database -->
# Set up a database

Step 1: Create a `data-access` directory

```bash
mkdir data-access
cd data-access
```

Step 2: Create a module to manage dependencies

```bash
go mod init example/data-access
```

Step 3: Create `create-tables.sql` file

```bash
touch create-tables.sql
```

Step 4: Write SQL code

```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

Step 5: Follow <https://dev.mysql.com/doc/refman/8.4/en/linux-installation-apt-repo.html> for MySQL installation

Step 6: Log into MySQL

```bash
mysql -u root -p
Enter password:

mysql>
```

Step 7: Create a database

```bash
mysql> CREATE DATABASE recordings;
```

Step 8: Change to `recordings` database

```bash
mysql> USE recordings;
Database changed
```

Step 9: Run `create-tables.sql` script

```bash
mysql> source /path/to/create-tables.sql
```

Step 10: Verify if the script is successful using `SELECT` command

```bash
mysql> SELECT * FROM album;
+----+---------------+----------------+-------+
| id | title         | artist         | price |
+----+---------------+----------------+-------+
|  1 | Blue Train    | John Coltrane  | 56.99 |
|  2 | Giant Steps   | John Coltrane  | 63.99 |
|  3 | Jeru          | Gerry Mulligan | 17.99 |
|  4 | Sarah Vaughan | Sarah Vaughan  | 34.98 |
+----+---------------+----------------+-------+
4 rows in set (0.00 sec)
```
