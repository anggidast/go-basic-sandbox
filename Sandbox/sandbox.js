// 1
function firstRepeatedWord(sentence) {
  // Write your code here
  let wordsArr = [];
  let word = '';
  for (let i = 0; i <= sentence.length; i++) {
    if (sentence[i] == ' ' || sentence[i] == ',' || sentence[i] == ':' || sentence[i] == ';' || sentence[i] == '-' || sentence[i] == '.') {
      if (word != '') wordsArr.push(word);
      word = '';
    } else {
      word += sentence[i];
    }
    if (i == sentence.length - 1) {
      if (word != '') wordsArr.push(word);
      word = '';
    }
  }

  let smallest = wordsArr.length;
  let result = '';
  for (let j = 0; j <= wordsArr.length; j++) {
    for (let k = j + 1; k < wordsArr.length; k++) {
      if (wordsArr[j] == wordsArr[k]) {
        if (k - j < smallest) {
          smallest = k - j;
          result = wordsArr[j];
        }
      }
    }
  }
  return result;
}

// 3
function subsetA(arr) {
  // Write your code here
  const sorted = arr.sort();
  const B = sorted;
  const A = [sorted[sorted.length - 1]];
  sorted.pop();

  const reducer = (accumulator, currentValue) => accumulator + currentValue;
  let sumA = A.reduce(reducer);
  let sumB = B.reduce(reducer);

  while (sumA < sumB) {
    const newA = sorted.pop();
    A.push(newA);

    sumA = A.reduce(reducer);
    sumB = B.reduce(reducer);
  }

  return A.sort();
}

// 2
function smallestString(substrings) {
  // Write your code here
  function sort(a, n) {
    // sort the array
    for (let i = 0; i < n; i++) {
      for (let j = i + 1; j < n; j++) {
        // comparing which of the
        // two concatenation causes
        // lexiographically smaller
        // string
        if (a[i] + a[j] > a[j] + a[i]) {
          let s = a[i];
          a[i] = a[j];
          a[j] = s;
        }
      }
    }
  }

  function smallest(a, n) {
    // Sort strings
    sort(a, n);

    // Concatenating sorted strings
    let answer = '';
    for (let i = 0; i < n; i++) answer += a[i];

    return answer;
  }

  // Driver code
  let a = substrings;
  let n = substrings.length;
  let result = smallest(a, n);
  return result;
}

// 5
function largestArea(N, M, H, V, h, v) {
  // Stores all bars
  var s1 = new Set();
  var s2 = new Set();

  // Insert horizontal bars
  for (var i = 1; i <= N + 1; i++) s1.add(i);

  // Insert vertictal bars
  for (var i = 1; i <= M + 1; i++) s2.add(i);

  // Remove horizontal separators from s1
  for (var i = 0; i < h; i++) {
    s1.delete(H[i]);
  }

  // Remove vertical separators from s2
  for (var i = 0; i < v; i++) {
    s2.delete(V[i]);
  }

  // Stores left out horizontal and
  // vertical separators
  var list1 = Array(s1.size);
  var list2 = Array(s2.size);

  var i = 0;
  s1.forEach((element) => {
    list1[i++] = element;
  });
  i = 0;
  s2.forEach((element) => {
    list2[i++] = element;
  });

  // Sort both list in
  // ascending order
  list1.sort((a, b) => a - b);
  list2.sort((a, b) => a - b);

  var maxH = 0,
    p1 = 0,
    maxV = 0,
    p2 = 0;

  // Find maximum difference of neighbors of list1
  for (var j = 0; j < s1.size; j++) {
    maxH = Math.max(maxH, list1[j] - p1);
    p1 = list1[j];
  }

  // Find max difference of neighbors of list2
  for (var j = 0; j < s2.size; j++) {
    maxV = Math.max(maxV, list2[j] - p2);
    p2 = list2[j];
  }

  // Print largest volume
  document.write(maxV * maxH);
}

// Driver code
// Given value of N & M
var N = 3,
  M = 3;

// Given arrays
var H = [2];
var V = [2];
var h = H.length;
var v = V.length;

// Function call to find the largest
// area when a series of horizontal &
// vertical bars are removed
largestArea(N, M, H, V, h, v);
