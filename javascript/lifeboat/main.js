function solution(people, limit){
    people.sort((a,b) => b-a);
    let l = 0;
    let r = people.length-1;
    
    for (l = 0, r = people.length-1; l <= r; l++) {
    	((people[l] + people[r]) <= limit) ? r-- : 0;
    }

    return l;
}

console.log(solution([70, 50, 80, 50], 100));
console.log(solution([40, 40, 40, 40, 20, 20, 20], 100));
console.log(solution([10,10,10,10,10], 30));