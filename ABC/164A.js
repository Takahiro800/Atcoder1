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
  input = myconv(input, 4);
  S = input[0];
  W = input[1];

  if(S <= W) {
    console.log('unsafe');
  } else {
    console.log('safe');
  }
}
Main(myin());
