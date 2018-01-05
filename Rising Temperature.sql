select w1.Id from Weather w1,Weather w2 
    where to_days(w2.Date)+1=to_days(w1.Date)
        and w1.Temperature>w2.Temperature;


-- 选择比前一天天数大1的时候一开始没有考虑边界问题，如果加上这一天之后是下个月或者下一年没考虑
-- 所以最后统一转换为天数来进行比较
-- 最后还有一个小坑是to_days(date)函数计算出来天数之后再进行比较，不要直接对日期进行增加。。。反正会莫名bug导致wa（我答案和excpted明明一样啊！