function doPost(e) {
  const token = PropertiesService.getScriptProperties().getProperty('SLACK_ACCESS_TOKEN')
  const verifyToken = PropertiesService.getScriptProperties().getProperty('OUTGOING_WEBHOOK_TOKEN')

  if (verifyToken != e.parameter.token) {
    var data = e.parameter.text.split(" ");//今回はサクッと作ったのでデータのチェックは一旦保留
    record(data);
    //Slackへ筋トレ管理botからのコメントをポスト
    postSlack("入力完了！\n筋トレ管理botからのコメント：「こんにちは」");
  }

//  const meetUrl = getMeetUrl()
//  const message = meetUrl !== undefined ? `Meetの部屋を作ったよ\n${meetUrl}` : "URL生成できなかった。ごめんね"

  const slackApp = SlackApp.create(token)
  slackApp.chatPostMessage(e.parameter.channel_id, message)
}

function postSlack(text){
  var url = "https://hooks.slack.com/services/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";
  var options = {
    "method" : "POST",
    "headers": {"Content-type": "application/json"},
    "payload" : '{"text":"' + text + '"}'
  };
  UrlFetchApp.fetch(url, options);
}

// 指定した時間にSlackにメッセージをポストする
function doMessage(e) {
  var message = "こんにちは！\n今日は筋トレしたかな？\n`腕立て`,`腹筋`,`スクワット`の回数をスペースで区切って入力してね";
  postSlack(message);
}

// Slackからの入力を受け取る
function doPost(e) {
  if (e.parameter.user_name === "slackbot") return;


}

function record (data) {
  var recordsheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName('記録'); // 記録シートを取得;

  var lastrow = recordsheet.getLastRow();
  var recordrow = lastrow + 1;

  var date = new Date();
  var formatdate = Utilities.formatDate(date, 'Asia/Tokyo', 'yyyy/MM/dd');
  var lastrecorddate = recordsheet.getRange("A" + lastrow).getValue();
  var formatlastrecorddate = Utilities.formatDate(lastrecorddate, 'Asia/Tokyo', 'yyyy/MM/dd');

  if (formatdate == formatlastrecorddate) {
    recordrow = lastrow;
  }

  //最終行の次の行に入力する
  recordsheet.getRange("A" + recordrow).setValue(formatdate);
  recordsheet.getRange("B" + recordrow).setValue(data[0]);
  recordsheet.getRange("C" + recordrow).setValue(data[1]);
  recordsheet.getRange("D" + recordrow).setValue(data[2]);
}

function selectRemark () {
  var remarksheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName('言葉');// 言葉シートを取得;
  var random = Math.round( Math.random() * remarksheet.getLastRow());
  var remark = remarksheet.getRange("A" + random).getValue();
  return remark;
}




function doPost(e) {
  var token = PropertiesService.getScriptProperties().getProperty('SLACK_ACCESS_TOKEN');
  var slackApp = SlackApp.create(token);
  var botToken = "***************"; //ここにOutgoing WebHooksのTokenを記述する
  var channelID;
  var message;
  var userName;
  var iconUrl;

  if(botToken == e.parameter.token){
    channelID = "****"; //チャンネルID(後述)
    message = "Hello!!";
    userName = "Test Service";
    iconUrl = "https://2.bp.blogspot.com/--tK6L-RmW-Y/WxvJot722lI/AAAAAAABMio/eKmnsWVyybAs7fZVuuDiNBbDXSSyZ7cgQCLcBGAs/s800/bird_menfukurou.png";
  }else {
    channelID = "****"; //チャンネルID(後述)
    message = "不正アクセスされている可能性があります";
    userName = "Warning Message";
    iconUrl = "https://1.bp.blogspot.com/-JgjGUdCLrIg/UUQ0qXejRiI/AAAAAAAAO3k/LGEKQA3rFd8/s1600/Tyrannosaurus.png";
  }

  slackApp.postMessage(
    channelID,
    message,
    {
      username: userName,
      icon_url: iconUrl
    }
  );
}



// だい３


//function doPost(e) {
  //const token = PropertiesService.getScriptProperties().getProperty('SLACK_ACCESS_TOKEN')
  //const verifyToken = PropertiesService.getScriptProperties().getProperty('OUTGOING_WEBHOOK_TOKEN')

 // if (verifyToken != e.parameter.token) {
  //  var data = e.parameter.text.split(" ");//今回はサクッと作ったのでデータのチェックは一旦保留
    //record(data);
    //Slackへ筋トレ管理botからのコメントをポスト
    //postSlack("入力完了！\n筋トレ管理botからのコメント：「こんにちは」");
  //}

//  const meetUrl = getMeetUrl()
//  const message = meetUrl !== undefined ? `Meetの部屋を作ったよ\n${meetUrl}` : "URL生成できなかった。ごめんね"

  //const slackApp = SlackApp.create(token)
  //slackApp.chatPostMessage(e.parameter.channel_id, message)
//}

function postSlack(text){
  var url = "https://hooks.slack.com/services/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";
  var options = {
    "method" : "POST",
    "headers": {"Content-type": "application/json"},
    "payload" : '{"text":"' + text + '"}'
  };
  UrlFetchApp.fetch(url, options);
}

// 指定した時間にSlackにメッセージをポストする
function doMessage(e) {
  var message = "こんにちは！\n今日は筋トレしたかな？\n`腕立て`,`腹筋`,`スクワット`の回数をスペースで区切って入力してね";
  const test = "test"
  postSlack(message);
}
// Slackからの入力を受け取る
function doPost(e) {
  const token = PropertiesService.getScriptProperties().getProperty('SLACK_ACCESS_TOKEN')
  const verifyToken = PropertiesService.getScriptProperties().getProperty('OUTGOING_WEBHOOK_TOKEN')
  var slackApp = SlackApp.create(token);
  doMessage(e)
}

function record (data) {
  var recordsheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName('記録'); // 記録シートを取得;

  var lastrow = recordsheet.getLastRow();
  var recordrow = lastrow + 1;

  var date = new Date();
  var formatdate = Utilities.formatDate(date, 'Asia/Tokyo', 'yyyy/MM/dd');
  var lastrecorddate = recordsheet.getRange("A" + lastrow).getValue();
  var formatlastrecorddate = Utilities.formatDate(lastrecorddate, 'Asia/Tokyo', 'yyyy/MM/dd');

  if (formatdate == formatlastrecorddate) {
    recordrow = lastrow;
  }

  //最終行の次の行に入力する
  recordsheet.getRange("A" + recordrow).setValue(formatdate);
  recordsheet.getRange("B" + recordrow).setValue(data[0]);
  recordsheet.getRange("C" + recordrow).setValue(data[1]);
  recordsheet.getRange("D" + recordrow).setValue(data[2]);
}

function selectRemark () {
  var remarksheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName('言葉');// 言葉シートを取得;
  var random = Math.round( Math.random() * remarksheet.getLastRow());
  var remark = remarksheet.getRange("A" + random).getValue();
  return remark;
}