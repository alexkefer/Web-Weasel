chrome.runtime.onInstalled.addListener(function() {
    chrome.storage.sync.set({ 'enabled': false });
  });
  
  chrome.storage.onChanged.addListener(function(changes) {
    if (changes.enabled && changes.enabled.newValue) {
      alert('Extension is now enabled!');
    }
  });
  
  chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    if (request.action === 'fetchHostname') {
        fetch('http://localhost:8080/hostname')
            .then(response => response.text())
            .then(hostname => {
                sendResponse({ hostname: hostname });
            })
            .catch(error => {
                console.error('Error fetching hostname:', error);
                sendResponse({ error: error.message });
            });
        return true; // Indicates that sendResponse will be called asynchronously
    }
});

