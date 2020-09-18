function quicksort(arr, left, right) {
    if (left >= right) {
        return;
    }
    console.log(arr);

    let pi = position(arr, left, right);

    quicksort(arr, left, pi - 1);
    quicksort(arr, pi + 1, right);
}

function position(arr, left, right) {
    let mid = Math.ceil((left + right) / 2);
    console.log(arr);
    swap(arr, left, mid);
    console.log(arr);
    let pivot = arr[left];
    console.log(pivot);
    let i = left, j = right;
    while (i < j) {
        while (pivot < arr[j]) {
            j--;
        }

        while (i < j && pivot >= arr[i]) {
            i++;
        }
        console.log("i : " + i + " / j : " + j);
        swap(arr, i, j);
        console.log(arr);
    }

    arr[left] = arr[i];
    arr[i] = pivot;
    console.log(arr);

    return i;
}

function swap(arr, i, j) {
    let temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}

let set = [8, 7, 6, 5, 4, 3, 2, 1];

quicksort(set, 0, set.length - 1);