# mysql 减库存并发问题
```
inventory 库存
num 本次要减小的数量
update mytable set inventory = inventory-num WHERE id= 1 and inventory >=num
不考虑高并发的话，这种实现方式简单，只要判断 affected rows 是否等于 1 就行了。
```
## 效率问题
一万个更新过来就卡死了,并发量大的话会导致更新失败
# 死锁问题
实际业务中多件商品下单这种操作如果不排序，分分钟死锁
# 锁表
delete 的时候带上索引列，不然锁表
间隙锁+事务合并导致的死锁问题
# mysql乐观锁,悲观锁


# 配置
innodb_lock_wait_timeout 排他锁失败时可以不阻塞?

```
1. 事务内单条 update
update 有行锁，提交之后下一个 update 才能拿到锁继续操作，如果是 update xx set version=version+1 这样更新是没问题的。

2.事务内先 select，然后根据 select 的结果再 update 。
比如 select version from xx 把 version 放程序变量里，然后在程序里进行 version++，再 update xx set version=?。
这种情况 select 是不加锁的，多个线程会一起拿到一个相同的 version，后续的 update 可能都是设置了相同的值。

3.事务内先 select for update，然后根据 select 的结果再 update 。
select 加了 for update 后也会加行锁，在你这个事务提交前其他线程的 select for update 也会卡住，直到事务提交后才能 select for update，数据也没问题了
```


https://v2ex.com/t/825520