// // 1
// function firstRepeatedWord(sentence) {
//   // Write your code here
//   let wordsArr = [];
//   let word = '';
//   for (let i = 0; i <= sentence.length; i++) {
//     if (sentence[i] == ' ' || sentence[i] == ',' || sentence[i] == ':' || sentence[i] == ';' || sentence[i] == '-' || sentence[i] == '.') {
//       if (word != '') wordsArr.push(word);
//       word = '';
//     } else {
//       word += sentence[i];
//     }
//     if (i == sentence.length - 1) {
//       if (word != '') wordsArr.push(word);
//       word = '';
//     }
//   }

//   let smallest = wordsArr.length;
//   let result = '';
//   for (let j = 0; j <= wordsArr.length; j++) {
//     for (let k = j + 1; k < wordsArr.length; k++) {
//       if (wordsArr[j] == wordsArr[k]) {
//         if (k - j < smallest) {
//           smallest = k - j;
//           result = wordsArr[j];
//         }
//       }
//     }
//   }
//   return result;
// }

// // 3
// function subsetA(arr) {
//   // Write your code here
//   const sorted = arr.sort();
//   const B = sorted;
//   const A = [sorted[sorted.length - 1]];
//   sorted.pop();

//   const reducer = (accumulator, currentValue) => accumulator + currentValue;
//   let sumA = A.reduce(reducer);
//   let sumB = B.reduce(reducer);

//   while (sumA < sumB) {
//     const newA = sorted.pop();
//     A.push(newA);

//     sumA = A.reduce(reducer);
//     sumB = B.reduce(reducer);
//   }

//   return A.sort();
// }

// // 2
// function smallestString(substrings) {
//   // Write your code here

//   // function sort(a, n) {
//   // sort the array
//   for (let i = 0; i < substrings.length; i++) {
//     for (let j = i + 1; j < substrings.length; j++) {
//       // comparing which of the
//       // two concatenation causes
//       // lexiographically smaller
//       // string
//       if (substrings[i] + substrings[j] > substrings[j] + substrings[i]) {
//         let s = substrings[i];
//         substrings[i] = substrings[j];
//         substrings[j] = s;
//       }
//     }
//   }
//   // }

//   // function smallest(a, n) {
//   //   // Sort strings
//   //   sort(a, n);

//   //   // Concatenating sorted strings
//   //   let answer = '';
//   //   for (let i = 0; i < n; i++) answer += a[i];

//   //   return answer;
//   // }

//   // Driver code
//   // let a = substrings;
//   // let n = substrings.length;
//   // let result = smallest(a, n);

//   // sort(a, n);
//   let result = '';
//   for (let i = 0; i < n; i++) result += a[i];

//   return result;
// }

// // 3
// function findMinRec(arr, i, sumCalculated, sumTotal) {
//   // If we have reached last element.
//   // Sum of one subset is sumCalculated,
//   // sum of other subset is sumTotal-
//   // sumCalculated.  Return absolute
//   // difference of two sums.
//   if (i == 0) return Math.abs(sumTotal - sumCalculated - sumCalculated);

//   // For every item arr[i], we have two choices
//   // (1) We do not include it first set
//   // (2) We include it in first set
//   // We return minimum of two choices
//   return Math.min(findMinRec(arr, i - 1, sumCalculated + arr[i - 1], sumTotal), findMinRec(arr, i - 1, sumCalculated, sumTotal));
// }

// // Returns minimum possible difference between
// // sums of two subsets
// function findMin(arr, n) {
//   // Compute total sum of elements
//   let sumTotal = 0;
//   for (let i = 0; i < n; i++) sumTotal += arr[i];

//   // Compute result using recursive function
//   return findMinRec(arr, n, 0, sumTotal);
// }

// /* Driver program to test above function */
// let arr = [3, 7, 5, 6, 2];
// let n = arr.length;
// document.write('The minimum difference' + ' between two sets is ' + findMin(arr, n));

// 5
function largestArea(N, M, H, V, h, v) {
  // Stores all bars
  // var s1 = new Set();
  // var s2 = new Set();
  let s1 = [];
  let s2 = [];

  // Insert horizontal bars
  // for (var i = 1; i <= N + 1; i++) s1.add(i);
  for (var i = 1; i <= N + 1; i++) s1.push(i);

  // Insert vertictal bars
  // for (var i = 1; i <= M + 1; i++) s2.add(i);
  for (var i = 1; i <= M + 1; i++) s2.push(i);

  // Remove horizontal separators from s1
  // console.log(s1);

  for (var i = 0; i < h; i++) {
    // s1.delete(H[i]);
    s1 = s1.filter((el) => el != H[i]);
  }

  // console.log(s1);

  // Remove vertical separators from s2
  for (var i = 0; i < v; i++) {
    // s2.delete(V[i]);
    s2 = s2.filter((el) => el != V[i]);
  }

  // Stores left out horizontal and
  // vertical separators
  // var list1 = Array(s1.size);
  // var list2 = Array(s2.size);
  let list1 = [];
  let list2 = [];

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
  for (var j = 0; j < s1.length; j++) {
    maxH = Math.max(maxH, list1[j] - p1);
    p1 = list1[j];
  }

  // Find max difference of neighbors of list2
  for (var j = 0; j < s2.length; j++) {
    maxV = Math.max(maxV, list2[j] - p2);
    p2 = list2[j];
  }

  // Print largest volume
  console.log(maxV * maxH);
}

// Driver code
// Given value of N & M
var N = 6,
  M = 6;

// Given arrays
var H = [4];
var V = [2];
var h = H.length;
var v = V.length;

// Function call to find the largest
// area when a series of horizontal &
// vertical bars are removed
largestArea(N, M, H, V, h, v);
