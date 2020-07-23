function solution(genres, plays) {
    let answer = [];
    let genrePriority = {};

    genres.forEach((genre, i) => {
        genrePriority[genre] = genrePriority[genre] ? genrePriority[genre] + plays[i] : plays[i];
    });

    genrePriority = Object.keys(genrePriority).map(k => [k, genrePriority[k]]).sort((a, b) => b[1] - a[1]);

    let map = {};
    genres.map((v, i) => {
        return [v, plays[i], i];
    })
    .sort((a, b) => b[1] - a[1])
    .forEach(element => {
        let genre = element[0];
        let id = element[2];

        let list = map[genre] ? map[genre] : [];

        if (list.length < 2) {
            list.push(id);
            map[genre] = list;
        }
    });

    console.log(map);

     genrePriority.forEach(v => {
         let genre = v[0];
         let musics = map[genre];

         musics.forEach(music => {
            answer.push(music);
        });
     });

     console.log(answer);

    return answer;
}

solution(["classic", "pop", "classic", "classic", "pop"], [500, 600, 150, 800, 2500]);