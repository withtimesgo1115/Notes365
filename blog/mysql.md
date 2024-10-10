## 什么是内连接、外连接、交叉连接、笛卡尔积呢？
MySQL 中的连接是通过两个或多个表之间的列进行关联，从而获取相关联的数据。连接分为内连接、外连接、交叉连接。

①、内连接（inner join）：返回两个表中连接字段匹配的行。如果一个表中的行在另一个表中没有匹配的行，则这些行不会出现在查询结果中。

②、外连接（outer join）：不仅返回两个表中匹配的行，还返回左表、右表或两者中未匹配的行。

③、交叉连接（cross join）：返回第一个表中的每一行与第二个表中的每一行的组合，这种类型的连接通常用于生成笛卡尔积。

④、笛卡尔积：数学中的一个概念，例如集合 A={a,b}，集合 B={0,1,2}，那么 A✖️B={<a,0>,<a,1>,<a,2>,<b,0>,<b,1>,<b,2>,}。

## MySQL 的内连接、左连接、右连接有什么区别？
MySQL 的连接主要分为内连接和外连接，外连接又可以分为左连接和右连接。

①、inner join 内连接，在两张表进行连接查询时，只保留两张表中完全匹配的结果集。

只有当两个表中都有匹配的记录时，这些记录才会出现在查询结果中。如果某一方没有匹配的记录，则该记录不会出现在结果集中。

内联可以用来找出两个表中共同的记录，相当于两个数据集的交集。

②、left join 返回左表（FROM 子句中指定的表）的所有记录，以及右表中匹配记录的记录。如果右表中没有匹配的记录，则结果中右表的部分会以 NULL 填充。

③、right join 刚好与左联相反，返回右表（FROM 子句中指定的表）的所有记录，以及左表中匹配记录的记录。如果左表中没有匹配的记录，则结果中左表的部分会以 NULL 填充。

## 说一下数据库的三大范式？
三大范式的作用是为了减少数据冗余，提高数据完整性。

①、第一范式（1NF）：确保表的每一列都是不可分割的基本数据单元，比如说用户地址，应该拆分成省、市、区、详细信息等 4 个字段。
②、第二范式（2NF）：在 1NF 的基础上，要求数据库表中的每一列都和主键直接相关，而不能只与主键的某一部分相关（主要针对联合主键）。
③、第三范式（3NF）：在 2NF 的基础上，消除非主键列对主键的传递依赖，即非主键列只依赖于主键列，不依赖于其他非主键列。

## varchar 与 char 的区别？
char：

char 表示定长字符串，长度是固定的；   
如果插入数据的长度小于 char 的固定长度时，则用空格填充；   
因为长度固定，所以存取速度要比 varchar 快很多，甚至能快 50%，但正因为其长度固定，所以会占据多余的空间，是空间换时间的做法；   
对于 char 来说，最多能存放的字符个数为 255，和编码无关   

varchar：

varchar 表示可变长字符串，长度是可变的；   
插入的数据是多长，就按照多长来存储；   
varchar 在存取方面与 char 相反，它存取慢，因为长度不固定，但正因如此，不占据多余的空间，是时间换空间的做法；   
对于 varchar 来说，最多能存放的字符个数为 65532   

日常的设计，对于长度相对固定的字符串，可以使用 char，对于长度不确定的，使用 varchar 更合适一些。

## blob 和 text 有什么区别？
blob 用于存储二进制数据，而 text 用于存储大字符串。
blob 没有字符集，text 有一个字符集，并且根据字符集的校对规则对值进行排序和比较


## DATETIME 和 TIMESTAMP 的异同？
相同点：

两个数据类型存储时间的表现格式一致。均为 YYYY-MM-DD HH:MM:SS
两个数据类型都包含「日期」和「时间」部分。
两个数据类型都可以存储微秒的小数秒（秒后 6 位小数秒）

DATETIME 和 TIMESTAMP 的区别

日期范围：DATETIME 的日期范围是 1000-01-01 00:00:00.000000 到 9999-12-31 23:59:59.999999；   
TIMESTAMP 的时间范围是1970-01-01 00:00:01.000000 UTC 到 ``2038-01-09 03:14:07.999999 UTC
 
存储空间：DATETIME 的存储空间为 8 字节；  
TIMESTAMP 的存储空间为 4 字节
 
时区相关：DATETIME 存储时间与时区无关；   
TIMESTAMP 存储时间与时区有关，显示的值也依赖于时区     

默认值：DATETIME 的默认值为 null；TIMESTAMP 的字段默认不为空(not null)，默认值为当前时间(CURRENT_TIMESTAMP)

## MySQL 中 in 和 exists 的区别？
MySQL 中的 in 语句是把外表和内表作 hash 连接，而 exists 语句是对外表作 loop 循环，每次 loop 循环再对内表进行查询。我们可能认为 exists 比 in 语句的效率要高，这种说法其实是不准确的，要区分情景：

如果查询的两个表大小相当，那么用 in 和 exists 差别不大。
如果两个表中一个较小，一个是大表，则子查询表大的用 exists，子查询表小的用 in。
not in 和 not exists：如果查询语句使用了 not in，那么内外表都进行全表扫描，没有用到索引；而 not extsts 的子查询依然能用到表上的索引。所以无论那个表大，用 not exists 都比 not in 要快。


## MySQL 里记录货币用什么字段类型比较好？
货币在数据库中 MySQL 常用 Decimal 和 Numeric 类型表示，这两种类型被 MySQL 实现为同样的类型。他们被用于保存与货币有关的数据。

例如 salary DECIMAL(9,2)，9(precision)代表将被用于存储值的总的小数位数，而 2(scale)代表将被用于存储小数点后的位数。存储在 salary 列中的值的范围是从-9999999.99 到 9999999.99。

DECIMAL 和 NUMERIC 值作为字符串存储，而不是作为二进制浮点数，以便保存那些值的小数精度。

之所以不使用 float 或者 double 的原因：因为 float 和 double 是以二进制存储的，所以有一定的误差。


## MySQL 怎么存储 emoji?
MySQL 的 utf8 字符集仅支持最多 3 个字节的 UTF-8 字符，但是 emoji 表情（😊）是 4 个字节的 UTF-8 字符，所以在 MySQL 中存储 emoji 表情时，需要使用 utf8mb4 字符集。

## drop、delete 与 truncate 的区别？
三者都表示删除，但是三者有一些差别：

|区别	|delete |	truncate	| drop|
|---- |---- |---- | ---- |
|类型	| 属于 DML|	属于 DDL|	属于 DDL|
|回滚|	可回滚	|不可回滚|	不可回滚|
|删除内容	|表结构还在，删除表的全部或者一部分数据行	|表结构还在，删除表中的所有数据	|从数据库中删除表，所有数据行，索引和权限也会被删除|
|删除速度	|删除速度慢，需要逐行删除	|删除速度快	|删除速度最快|

因此，在不再需要一张表的时候，用 drop；在想删除部分数据行时候，用 delete；在保留表而删除所有数据的时候用 truncate。

## UNION 与 UNION ALL 的区别？
- 如果使用 UNION，会在表链接后筛选掉重复的记录行
- 如果使用 UNION ALL，不会合并重复的记录行
- 从效率上说，UNION ALL 要比 UNION 快很多，如果合并没有刻意要删除重复行，那么就使用 UNION All

## count(1)、count(*) 与 count(列名) 的区别？
执行效果：

count(*)包括了所有的列，相当于行数，在统计结果的时候，不会忽略列值为 NULL  
count(1)包括了忽略所有列，用 1 代表代码行，在统计结果的时候，不会忽略列值为 NULL  
count(列名)只包括列名那一列，在统计结果的时候，会忽略列值为空（这里的空不是只空字符串或者 0，而是表示 null）的计数，即某个字段值为 NULL 时，不统计。


执行速度：

列名为主键，count(列名)会比 count(1)快  
列名不为主键，count(1)会比 count(列名)快  
如果表多个列并且没有主键，则 count（1） 的执行效率优于 count（*）
如果有主键，则 select count（主键）的执行效率是最优的
如果表只有一个字段，则 select count（*）最优。

## 一条 SQL 查询语句的执行顺序？
![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/mysql-47ddea92-cf8f-49c4-ab2e-69a829ff1be2.jpg)

1. FROM：对 FROM 子句中的左表<left_table>和右表<right_table>执行笛卡儿积（Cartesianproduct），产生虚拟表 VT1
2. ON：对虚拟表 VT1 应用 ON 筛选，只有那些符合<join_condition>的行才被插入虚拟表 VT2 中
3. JOIN：如果指定了 OUTER JOIN（如 LEFT OUTER JOIN、RIGHT OUTER JOIN），那么保留表中未匹配的行作为外部行添加到虚拟表 VT2 中，产生虚拟表 VT3。如果 FROM 子句包含两个以上表，则对上一个连接生成的结果表 VT3 和下一个表重复执行步骤 1）～步骤 3），直到处理完所有的表为止
4. WHERE：对虚拟表 VT3 应用 WHERE 过滤条件，只有符合<where_condition>的记录才被插入虚拟表 VT4 中
5. GROUP BY：根据 GROUP BY 子句中的列，对 VT4 中的记录进行分组操作，产生 VT5
6. CUBE|ROLLUP：对表 VT5 进行 CUBE 或 ROLLUP 操作，产生表 VT6
7. HAVING：对虚拟表 VT6 应用 HAVING 过滤器，只有符合<having_condition>的记录才被插入虚拟表 VT7 中。
SELECT：第二次执行 SELECT 操作，选择指定的列，插入到虚拟表 VT8 中
8. DISTINCT：去除重复数据，产生虚拟表 VT9
9. ORDER BY：将虚拟表 VT9 中的记录按照<order_by_list>进行排序操作，产生虚拟表 VT10。11）
10. LIMIT：取出指定行的记录，产生虚拟表 VT11，并返回给查询用户

## 介绍一下 MySQL 的常用命令（补充）
### 说说数据库操作命令？
①、创建数据库:
```sql
CREATE DATABASE database_name;
```
②、删除数据库:
```sql
DROP DATABASE database_name;
```
③、选择数据库:
```sql
USE database_name;
```

### 说说表操作命令？
①、创建表:
```sql
CREATE TABLE table_name (
    column1 datatype,
    column2 datatype,
    ...
);
```

②、删除表:
```sql
DROP TABLE table_name;
```
③、显示所有表:

```sql
SHOW TABLES;
```
④、查看表结构:
```sql
DESCRIBE table_name;
```
⑤、修改表（添加列）:
```sql
ALTER TABLE table_name ADD column_name datatype;
```
### 说说 CRUD 命令？
①、插入数据:
```sql
INSERT INTO table_name (column1, column2, ...) VALUES (value1, value2, ...);
```
②、查询数据:
```sql
SELECT column_names FROM table_name WHERE condition;
```
③、更新数据:
```sql
UPDATE table_name SET column1 = value1, column2 = value2 WHERE condition;
```
④、删除数据:
```sql
DELETE FROM table_name WHERE condition;
```
### 说说索引和约束的创建修改命令？
①、创建索引:
```sql
CREATE INDEX index_name ON table_name (column_name);
```
②、添加主键约束:
```sql
ALTER TABLE table_name ADD PRIMARY KEY (column_name);
```
③、添加外键约束:
```sql
ALTER TABLE table_name ADD CONSTRAINT fk_name FOREIGN KEY (column_name) REFERENCES parent_table (parent_column_name);
```

### 说说用户和权限管理的命令？
①、创建用户:
```sql
CREATE USER 'username'@'host' IDENTIFIED BY 'password';
```
②、授予权限:
```sql
GRANT ALL PRIVILEGES ON database_name.table_name TO 'username'@'host';
```
③、撤销权限:
```sql
REVOKE ALL PRIVILEGES ON database_name.table_name FROM 'username'@'host';
```
④、删除用户:
```sql
DROP USER 'username'@'host';
```

### 说说事务控制的命令？
①、开始事务:
```sql
START TRANSACTION;
```
②、提交事务:
```sql
COMMIT;
```
③、回滚事务:
```sql
ROLLBACK;
```

## 介绍一下 MySQL bin 目录下的可执行文件（补充）
- mysql：客户端程序，用于连接 MySQL 服务器
- mysqldump：一个非常实用的 MySQL 数据库备份工具，用于创建一个或多个 MySQL 数据库级别的 SQL 转储文件，包括数据库的表结构和数据。对数据备份、迁移或恢复非常重要。
- mysqladmin：mysql 后面加上 admin 就表明这是一个 MySQL 的管理工具，它可以用来执行一些管理操作，比如说创建数据库、删除数据库、查看 MySQL 服务器的状态等。
- mysqlcheck：mysqlcheck 是 MySQL 提供的一个命令行工具，用于检查、修复、分析和优化数据库表，对数据库的维护和性能优化非常有用。
- mysqlimport：用于从文本文件中导入数据到数据库表中，非常适合用于批量导入数据。
- mysqlshow：用于显示 MySQL 数据库服务器中的数据库、表、列等信息。
- mysqlbinlog：用于查看 MySQL 二进制日志文件的内容，可以用于恢复数据、查看数据变更等。

## MySQL 第 3-10 条记录怎么查？（补充）
```sql
SELECT * FROM table_name LIMIT 2, 8;
```

## 用过哪些 MySQL 函数？（补充）
### 用过哪些字符串函数来处理文本？
CONCAT(): 连接两个或多个字符串。  
LENGTH(): 返回字符串的长度。  
SUBSTRING(): 从字符串中提取子字符串。  
REPLACE(): 替换字符串中的某部分。  
LOWER() 和 UPPER(): 分别将字符串转换为小写或大写。  
TRIM(): 去除字符串两侧的空格或其他指定字符。  

### 用过哪些数值函数？
ABS(): 返回一个数的绝对值。   
CEILING(): 返回大于或等于给定数值的最小整数。   
FLOOR(): 返回小于或等于给定数值的最大整数。   
ROUND(): 四舍五入到指定的小数位数。  
MOD(): 返回除法操作的余数。   

### 用过哪些日期和时间函数？
NOW(): 返回当前的日期和时间。  
CURDATE(): 返回当前的日期。   
CURTIME(): 返回当前的时间。   
DATE_ADD() 和 DATE_SUB(): 在日期上加上或减去指定的时间间隔。    
DATEDIFF(): 返回两个日期之间的天数。   
DAY(), MONTH(), YEAR(): 分别返回日期的日、月、年部分。  

### 用过哪些汇总函数？
SUM(): 计算数值列的总和。   
AVG(): 计算数值列的平均值。  
COUNT(): 计算某列的行数。  
MAX() 和 MIN(): 分别返回列中的最大值和最小值。  
GROUP_CONCAT(): 将多个行值连接为一个字符串。   

### 用过哪些逻辑函数？
IF(): 如果条件为真，则返回一个值；否则返回另一个值。     
CASE: 根据一系列条件返回值。     
COALESCE(): 返回参数列表中的第一个非 NULL 值。    
### 用过哪些格式化函数？
FORMAT(): 格式化数字为格式化的字符串，通常用于货币显示。

### 用过哪些类型转换函数？
CAST(): 将一个值转换为指定的数据类型。   
CONVERT(): 类似于CAST()，用于类型转换。


## 说说 SQL 的隐式数据类型转换？（补充）
在 SQL 中，当不同数据类型的值进行运算或比较时，会发生隐式数据类型转换。

比如说，当一个整数和一个浮点数相加时，整数会被转换为浮点数，然后再进行相加。

可以通过显式转换来规避这种情况。

## 说说 MySQL 的基础架构?
![](https://cdn.tobebetterjavaer.com/tobebetterjavaer/images/sidebar/sanfene/mysql-77626fdb-d2b0-4256-a483-d1c60e68d8ec.jpg)

MySQL 逻辑架构图主要分三层：

客户端：最上层的服务并不是 MySQL 所独有的，大多数基于网络的客户端/服务器的工具或者服务都有类似的架构。比如连接处理、授权认证、安全等等。   

Server 层：大多数 MySQL 的核心服务功能都在这一层，包括查询解析、分析、优化、缓存以及所有的内置函数（例如，日期、时间、数学和加密函数），所有跨存储引擎的功能都在这一层实现：存储过程、触发器、视图等。


存储引擎层：第三层包含了存储引擎。存储引擎负责 MySQL 中数据的存储和提取。Server 层通过 API 与存储引擎进行通信。这些接口屏蔽了不同存储引擎之间的差异，使得这些差异对上层的查询过程透明。

## 一条 SQL 查询语句在 MySQL 中如何执行的？
![](https://cdn.tobebetterjavaer.com/stutymore/mysql-20240415102041.png)

第一步，客户端发送 SQL 查询语句到 MySQL 服务器。

第二步，MySQL 服务器的连接器开始处理这个请求，跟客户端建立连接、获取权限、管理连接。

第三步（MySQL 8.0 以后已经干掉了），连接建立后，MySQL 服务器的查询缓存组件会检查是否有缓存的查询结果。如果有，直接返回给客户端；如果没有，进入下一步。

第三步，解析器开始对 SQL 语句进行解析，检查语句是否符合 SQL 语法规则，确保引用的数据库、表和列都存在，并处理 SQL 语句中的名称解析和权限验证。

第四步，优化器负责确定 SQL 语句的执行计划，这包括选择使用哪些索引，以及决定表之间的连接顺序等。优化器会尝试找出最高效的方式来执行查询。

第五步，执行器会调用存储引擎的 API 来进行数据的读写。

第六步，MySQL 的存储引擎是插件式的，不同的存储引擎在细节上面有很大不同。例如，InnoDB 是支持事务的，而 MyISAM 是不支持的。之后，会将执行结果返回给客户端

第七步，客户端接收到查询结果，完成这次查询请求。


## 说说 MySQL 的数据存储形式（补充）
MySQL 是以表的形式存储数据的，而表空间的结构则由段、区、页、行组成。

![](https://cdn.tobebetterjavaer.com/stutymore/mysql-20240515110034.png)


①、段（Segment）：表空间由多个段组成，常见的段有数据段、索引段、回滚段等。

创建索引时会创建两个段，数据段和索引段，数据段用来存储叶子阶段中的数据；索引段用来存储非叶子节点的数据。

回滚段包含了事务执行过程中用于数据回滚的旧数据。

②、区（Extent）：段由一个或多个区组成，区是一组连续的页，通常包含 64 个连续的页，也就是 1M 的数据。

使用区而非单独的页进行数据分配可以优化磁盘操作，减少磁盘寻道时间，特别是在大量数据进行读写时。

③、页（Page）：页是 InnoDB 存储数据的基本单元，标准大小为 16 KB，索引树上的一个节点就是一个页。

也就意味着数据库每次读写都是以 16 KB 为单位的，一次最少从磁盘中读取 16KB 的数据到内存，一次最少写入 16KB 的数据到磁盘。

④、行（Row）：InnoDB 采用行存储方式，意味着数据按照行进行组织和管理，行数据可能有多个格式，比如说 COMPACT、REDUNDANT、DYNAMIC 等。

MySQL 8.0 默认的行格式是 DYNAMIC，由COMPACT 演变而来，意味着这些数据如果超过了页内联存储的限制，则会被存储在溢出页中。

## MySQL 有哪些常见存储引擎？
| 功能	| InnoDB |	MyISAM|	MEMORY |
|---- | ---- | ---- | ---- |
| 支持事务	 | Yes | No |	No |
| 支持全文索引	|Yes	|Yes	|No|
| 支持 B+树索引	|Yes	|Yes|	Yes|
| 支持哈希索引	| Yes	 |No	| Yes|
| 支持外键	 |Yes|	No|	No | 


## 那存储引擎应该怎么选择？
- 大多数情况下，使用默认的 InnoDB 就对了，InnoDB 可以提供事务、行级锁等能力。
- MyISAM 适合读更多的场景。
- MEMORY 适合临时表，数据量不大的情况。由于数据都存放在内存，所以速度非常快。

## InnoDB 和 MyISAM 主要有什么区别？
InnoDB 和 MyISAM 之间的区别主要表现在存储结构、事务支持、最小锁粒度、索引类型、主键必需、表的具体行数、外键支持等方面。

①、存储结构：

MyISAM：用三种格式的文件来存储，.frm 文件存储表的定义；.MYD 存储数据；.MYI 存储索引。
InnoDB：用两种格式的文件来存储，.frm 文件存储表的定义；.ibd 存储数据和索引。
②、事务支持：

MyISAM：不支持事务。
InnoDB：支持事务。
③、最小锁粒度：

MyISAM：表级锁，高并发中写操作存在性能瓶颈。
InnoDB：行级锁，并发写入性能高。
④、索引类型：

MyISAM 为非聚簇索引，索引和数据分开存储，索引保存的是数据文件的指针。

⑤、外键支持：MyISAM 不支持外键；InnoDB 支持外键。

⑥、主键必需：MyISAM 表可以没有主键；InnoDB 表必须有主键。

⑦、表的具体行数：MyISAM 表的具体行数存储在表的属性中，查询时直接返回；InnoDB 表的具体行数需要扫描整个表才能返回。


## MySQL 日志文件有哪些？分别介绍下作用？
MySQL 的日志文件主要包括：

①、错误日志（Error Log）：记录 MySQL 服务器启动、运行或停止时出现的问题。

②、慢查询日志（Slow Query Log）：记录执行时间超过 long_query_time 值的所有 SQL 语句。这个时间值是可配置的，默认情况下，慢查询日志功能是关闭的。可以用来识别和优化慢 SQL。

③、一般查询日志（General Query Log）：记录所有 MySQL 服务器的连接信息及所有的 SQL 语句，不论这些语句是否修改了数据。

④、二进制日志（Binary Log）：记录了所有修改数据库状态的 SQL 语句，以及每个语句的执行时间，如 INSERT、UPDATE、DELETE 等，但不包括 SELECT 和 SHOW 这类的操作。

⑤、重做日志（Redo Log）：记录了对于 InnoDB 表的每个写操作，不是 SQL 级别的，而是物理级别的，主要用于崩溃恢复。

⑥、回滚日志（Undo Log，或者叫事务日志）：记录数据被修改前的值，用于事务的回滚。

###  请重点说说 binlog？
binlog 是一种物理日志，会在磁盘上记录下数据库的所有修改操作，以便进行数据恢复和主从复制。

当发生数据丢失时，binlog 可以将数据库恢复到特定的时间点。  
主服务器（master）上的二进制日志可以被从服务器（slave）读取，从而实现数据同步。   
binlog 包括两类文件：

二进制索引文件（.index）
二进制日志文件（.00000*）

binlog 默认是没有启用的。要启用它，需要在 MySQL 的配置文件（my.cnf 或 my.ini）中设置 log_bin 参数。

### binlog 和 redo log 有什么区别？
binlog，即二进制日志，对所有存储引擎都可用，是 MySQL 服务器级别的日志，用于数据的复制、恢复和备份。而 redo log 主要用于保证事务的持久性，是 InnoDB 存储引擎特有的日志类型。

binlog 记录的是逻辑 SQL 语句，而 redo log 记录的是物理数据页的修改操作，不是具体的 SQL 语句。

redo log 是固定大小的，通常配置为一组文件，使用环形方式写入，旧的日志会在空间需要时被覆盖。binlog 是追加写入的，新的事件总是被添加到当前日志文件的末尾，当文件达到一定大小后，会创建新的 binlog 文件继续记录。

## 一条更新语句怎么执行的了解吗？
更新语句的执行是 Server 层和引擎层配合完成，数据除了要写入表中，还要记录相应的日志。

1. 执行器先找引擎获取 ID=2 这一行。ID 是主键，存储引擎检索数据，找到这一行。如果 ID=2 这一行所在的数据页本来就在内存中，就直接返回给执行器；否则，需要先从磁盘读入内存，然后再返回。
2. 执行器拿到引擎给的行数据，把这个值加上 1，比如原来是 N，现在就是 N+1，得到新的一行数据，再调用引擎接口写入这行新数据。
3. 引擎将这行新数据更新到内存中，同时将这个更新操作记录到 redo log 里面，此时 redo log 处于 prepare 状态。然后告知执行器执行完成了，随时可以提交事务。
4. 执行器生成这个操作的 binlog，并把 binlog 写入磁盘。
5. 执行器调用引擎的提交事务接口，引擎把刚刚写入的 redo log 改成提交（commit）状态，更新完成。

从上图可以看出，MySQL 在执行更新语句的时候，在服务层进行语句的解析和执行，在引擎层进行数据的提取和存储；同时在服务层对 binlog 进行写入，在 InnoDB 内进行 redo log 的写入。

不仅如此，在对 redo log 写入时有两个阶段的提交，一是 binlog 写入之前prepare状态的写入，二是 binlog 写入之后commit状态的写入。

## 那为什么要两阶段提交呢？
我们可以假设不采用两阶段提交的方式，而是采用“单阶段”进行提交，即要么先写入 redo log，后写入 binlog；要么先写入 binlog，后写入 redo log。这两种方式的提交都会导致原先数据库的状态和被恢复后的数据库的状态不一致。

简单说，redo log 和 binlog 都可以用于表示事务的提交状态，而两阶段提交就是让这两个状态保持逻辑上的一致。

## redo log 怎么刷入磁盘的知道吗？

