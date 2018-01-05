select distinct d.Name Department,e.Name Employee,e.Salary Salary from 
	Employee e, Department d, (select DepartmentId, max(Salary) Salary from Employee group by DepartmentId) tmp
		where e.DepartmentId=d.Id and e.Salary=tmp.Salary and tmp.DepartmentId=e.DepartmentId;

-- 聚合之后只能选择聚合的列或者使用聚合函数，所以不得不适用内联视图和where查询