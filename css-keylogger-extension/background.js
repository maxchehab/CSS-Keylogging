chrome.browserAction.onClicked.addListener(tab => {
  chrome.tabs.insertCSS(tab.ib, {
    file: "keylogger.css"
  });
});
