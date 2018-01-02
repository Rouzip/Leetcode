select FirstName,LastName,City,State from Person
	left join Adress
		on Person.PersonId = Adress.PersonId;