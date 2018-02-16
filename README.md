# 2018-02-SyntaxHighlight

Create a simple tool that will Syntax Highlight [/etc/host](https://en.wikipedia.org/wiki/Hosts_(file)) file.

# Now that you have had a quick introduction to [Pigeon](https://godoc.org/github.com/mna/pigeon).

Today's challenge will be to take the parser for `/etc/hosts` file that was given presented, to 
create a syntax highlighting cat program for `/etc/hosts` files.

You can use the `https://godoc.org/github.com/fatih/color` to get [ANSI color](https://godoc.org/github.com/fatih/color) support.


# Challenge 2.

Now create a syntax highlighting cat program for the [/etc/resolv.conf](https://en.wikipedia.org/wiki/Resolv.conf)

# Challenge 3.

Create a syntax highlighing cat program for simple SQL files. These SQL will be limited to
the following clauses:

* SELECT
* UPDATE
* JOIN
* WHERE
* FROM
* LIMIT

It should support the folowing SQLS:

`SELECT * FROM EMPLOYEE_TBL;`
```
SELECT EMP_ID as EMP 
FROM EMPLOYEE_TBL;
```

```
SELECT EMP_ID, LAST_NAME
FROM EMPLOYEE_TBL
WHERE EMP_ID = '333333333';
```

```
SELECT EMP_ID, LAST_NAME
FROM EMPLOYEE_TBL
WHERE CITY = 'INDIANAPOLIS'
ORDER BY EMP_ID;
```

```
SELECT COUNT(*)
FROM TABLE_NAME;
```

```
select prod_desc,
    prod_desc product
from products_tbl;
```

```
select prod_desc,
    prod_desc as product
from products_tbl;
```


```
select prod_desc,
    prod_desc "as"
from products_tbl;
```

```
select prod_desc,
    prod_desc as "as"
from products_tbl;
```
