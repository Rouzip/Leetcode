select name,population,area from World 
	where population>25000000 or area>3000000;
-- 尽管这道题很容易想出这种方法，但是如果追求的是性能的话，需要使用union
select name,population,area from World
	where population>25000000
union
select name,population,area form World
	where area>3000000;