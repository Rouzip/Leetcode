var climbStairs = function(n){
	let a = b = 1;
	while(n--){
		a=(b+=a)-a;
		console.log('a为：'+a);
        console.log('b为：'+b);
	}    
	return a;
}

console.log(climbStairs(4));