function solution(A, F, M) {
  if (F < 0 || M < 1 || M > 6) {
    let result = [0];
    return result;
  }

  let s = A.reduce((a, v) => a + v, 0);

  let sum = (A.length + F) * M - s;

  if (F > sum || sum > F * 6) {
    let result = [0];
    return result;
  }
  let result = [];

  while (F--) {
    let tempSum = Math.floor(sum / (F + 1));
    result.push(tempSum);
    sum -= tempSum;
  }

  return result;
}

let A = [1, 5, 6];
let F = 4;
let M = 3;
console.log(solution(A, F, M));
