//0:何もしない  1:数値へ変換  2:半角SPで分割  3:改行で分割  4:半角SPで分割し、数値配列へ
//5:改行で分割し、数値配列へ  6:1文字に分割  7:1文字に分割し、数値配列へ
function myconv(i, no){
  switch(no){
    case 0: return i;
    case 1: return parseInt(i);
    case 2: return i.split(" ");
    case 3: return i.split("\n");
    case 4: return i.split(" ").map((a)=>Number(a));
    case 5: return i.split("\n").map((a)=>Number(a));
    case 6: return i.split("");
    case 7: return i.split("").map((a)=>Number(a));
  }}
function Main(input) {
  const rank = input;
  let me
  switch(rank) {
    case 'A':
      me = rank;
      break
    case 'B':
      console.log('Bランクです');
      break
    default:
      console.log('ランク外です');
      break
}
  console.log(me);
}
Main(require("fs").readFileSync("/dev/stdin", "utf8").trim());