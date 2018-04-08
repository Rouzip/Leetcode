delete p1 from Person p1,Person p2
    where p1.email=p2.Email and p1.Id>p2.Id;


-- 一开始我的思路是group by聚类，然后删除到只剩余一个，但是貌似语法过不去
-- 所以我看了别人的思路，别人是通过多表聚合来达到目的的
-- 通过两个表进行连接，然后通过默认的ID不同来进行删除已达到目的