// inputに入力データ全体が入る
function Main(input) {
	// 1行目がinput[0], 2行目がinput[1], …に入る
	input = input.split(" ");
	const S = parseInt(input[0], 10);
	const W = parseInt(input[1], 10);
  //出力
  if (S <= W) {
    console.log('unsafe');
  } else {
    console.log('safe');
  }
}
//*この行以降は編集しないでください（標準入出力から一度に読み込み、Mainを呼び出します）
Main(require("fs").readFileSync("/dev/stdin", "utf8"));