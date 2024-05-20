function Checkodd(arr) {
  let datasave = [];
  arr.map((x, i) => {
    let f = datasave.find((e) => e?.number == x);
    f
      ? (function () {
        datasave.map(e=> e.number == x ? e.cout = e.cout +1 : 0)
      })()
      : (datasave = [
          ...datasave,
          {
            number: x,
            cout: 1,
          },
        ]);
  });
  let [f] = datasave.filter(e=> e.cout % 2 !== 0)
  return `${f.number} because it occurs ${f.cout} time (which is odd).`
}

console.log(Checkodd([7]));
console.log(Checkodd([0]))
console.log(Checkodd([1, 1, 2]));
console.log(Checkodd([0, 1, 0, 1, 0]));
console.log(Checkodd([1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1]));