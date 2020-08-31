function myin(){return require("fs").readFileSync("/dev/stdin", "utf8").trim();}

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
  S = myconv(input, 0);
  len = S.length
  let ans = 999999999
  for(let i = 0; i < len; i++ ){
    num = parseInt(S.substr(i, 3), 10);
    ans = Math.min(ans, Math.abs(num - 753));
  }
  console.log(ans);

}
// Main(myin());
Main(require("fs").readFileSync("/dev/stdin", "utf8").trim());