# Write your MySQL query statement below
select max(Salary) SecondHighestSalary from Employee 
    where Salary < (select max(Salary) from Employee);

-- 思路是找出比最大的小的子视图中最大的即为第二大的