function uniquePermutations(s) {
    let results = [];

    if (s.length === 0) return [''];
    let charCount = {}

    for (let char of s) {
        charCount[char] = (charCount[char] || 0) + 1;
    }
    // console.log(charCount);

    function permute(current, baseCase) {
        if (baseCase.length === 0) {
            results.push(current);
            return;
        }
        for (let char in charCount) {
            if (charCount[char] > 0) {
                charCount[char]--;
                // console.log(charCount);
                permute(current + char, baseCase.slice(1));
                // console.log(charCount);

                charCount[char]++;
            }
        }
    }
    permute('', s);
    return results.sort();
}
console.log(uniquePermutations('a'));
console.log(uniquePermutations('ab'));
console.log(uniquePermutations('abc'));
console.log(uniquePermutations('aabb'));
