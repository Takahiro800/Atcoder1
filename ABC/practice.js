function Main(input) {
  input = input.split("\n");
  tmp = input[1].split(" ");

  //文字列から10進数に変換するときはparseInt(input[0], 10);
  const a = parseInt(input[0], 10);
  const b = parseInt(tmp[0], 10);
  const c = parseInt(tmp[1], 10);
  const s = input[2];

  console.log('%d %s', a+b+c, s);
}

Main(require("fs").readFileSync("/dev/stdin", "utf8"));