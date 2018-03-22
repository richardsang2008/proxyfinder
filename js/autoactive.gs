/*
*Go to https://script.google.com/
*/
function myFunction() {
  var verified = 0;
  var moveToTrash = true;

  var threads = GmailApp.search('in:inbox subject:"Pokémon Trainer Club Activation"');
  Logger.log("Found " + threads.length + " threads.");

  threads.forEach(function(thread) {
    var messages = thread.getMessages();
    Logger.log("Found " + messages.length + " messages.");

    messages.forEach(function(msg) {
      var value = msg.getBody()
                     .match(/Verify your email/m);

      if(msg.isInInbox() && value) {
        var link = msg.getBody().match(/<a href="https:\/\/club.pokemon.com\/us\/pokemon-trainer-club\/activated\/([\w\d]+)"/);

        if(link) {
          var url = 'https://club.pokemon.com/us/pokemon-trainer-club/activated/' + link[1];
          var options = {
            "muteHttpExceptions": true
          };

          var status = UrlFetchApp.fetch(url, options).getResponseCode();
          Logger.log("[#] Verified (" + status + "): " + url);

          if(status == 200) {
            verified++;
            msg.markRead();

            if(moveToTrash) { msg.moveToTrash(); }
          }
        }
      }
    });
  });

  Logger.log("Completed " + verified + " verifications.");
}