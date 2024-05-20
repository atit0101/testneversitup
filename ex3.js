function countSmileys(arr) {
  let regex = /:|;|-|~|D|\)/g;
  let data = [];
  for (let index = 0; index < arr.length; index++) {
    const smile = arr[index].match(regex);
    if (smile.length > 1 && smile[smile.length - 1] == "D" || smile[smile.length - 1] == ")") data = [...data, smile];
  }
  return data.length
}

console.log(countSmileys([":)", ";(", ";}", ":-D"]));
console.log(countSmileys([";D", ":-(", ":-)", ";~)"]));
console.log(countSmileys([";]", ":[", ";*", ":$", ";-D"]));
