function solution(routes) {
    routes.sort((a, b) =>  (a[0] - b[0]));

    let answer = 1, position = routes[0][1];
    for(let i = 0; i < routes.length - 1; i++) {
        if (position > routes[i][1]) {
            position = routes[i][1];
        }

        if (position < routes[i+1][0]) {
            answer++;
            position = routes[i+1][1];
        }
    }

    return answer;
}

console.log(solution([[-20,15], [-14,-5], [-18,-13], [-5,-3]]));
console.log(solution([[-10,4], [-6,8], [-15,-13]]));