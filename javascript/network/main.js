function solution(n, computers) {
    let answer = 0;

    let visited = computers.map(v => false);

    for (let i = 0; i < visited.length; i++) {
        if (!visited[i]) {
            search(visited, i, computers);
            answer++;
        }
    }

    return answer;
}

function search(visited, i, computers) {
    let computer = computers[i];
    for (let j = 0; j < computer.length; j++) {
        if (!visited[j] && computer[j] === 1) {
            visited[j] = true;
            
            search(visited, j, computers);
        }
    }
}

console.log(solution(3, [[1, 1, 0], [1, 1, 0], [0, 0, 1]]));
console.log(solution(3, [[1, 1, 0], [1, 1, 1], [0, 1, 1]]));