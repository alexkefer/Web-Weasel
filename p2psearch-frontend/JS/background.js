chrome.runtime.onInstalled.addListener(function() {
    chrome.storage.sync.set({ 'enabled': false });
  });
  
  chrome.storage.onChanged.addListener(function(changes) {
    if (changes.enabled && changes.enabled.newValue) {
      alert('Extension is now enabled!');
    }
  });
  