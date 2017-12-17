var sortColors = function(nums){
	let a = 0;
	let b = 0;
	nums.forEach((num,step)=>{
		let c = num;
		nums[step] = 2;
		if (c < 2){
			nums[a] = 1;
			a++;
		}
		if (c == 0){
			nums[b] = 0;
			b++;
		}
	})
}

/*
对于js有了一点点体会，js的传递是通过值来传递，不同于python的引用传递
js将基本类型储存于栈，对象储存于堆
在遍历js的数组的时候有filter，map，forEach，for of等很多方法
*/
var a = [0,1,0,2,0,1];
sortColors(a);
console.log(a);