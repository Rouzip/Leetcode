select e2.Name Employee from Employee e1,Employee e2 
	where e1.Id = e2.ManagerId and e1.Salary > e2.Salary;