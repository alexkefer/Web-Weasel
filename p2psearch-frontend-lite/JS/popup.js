document.addEventListener('DOMContentLoaded', function () {
  var iconButton = document.getElementById('iconButton');
  clearLocalStorage();
  fetchAndDisplayHostname();
  fetchAndDisplayNodeIPAddress();
  fetchAndDisplayPeersIPAddress();

  iconButton.addEventListener('click', function () {
    var iconImg = document.querySelector('.icon-img');
    iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';
    toggleDeviceInfoVisibility(iconImg.src.includes('on_power_icon.png'));
  });

  var connectButton = document.getElementById('connectButton');
  connectButton.addEventListener('click', function () {
    var peerAddressInput = document.getElementById('peerAddressInput').value;
    connectToPeer(peerAddressInput);
  });
});

function clearLocalStorage() {
  localStorage.clear();
}

function fetchAndDisplayHostname() {
  fetch('http://localhost:8080/hostname')
    .then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error('Failed to fetch hostname data');
      }
    })
    .then(hostname => {
      const publicDeviceNameSpan = document.querySelector('.device-info-text');
      publicDeviceNameSpan.textContent = "Public Device Name: " + hostname;
      localStorage.setItem('hostname', hostname);
    })
    .catch(error => {
      console.error('Error:', error.message);
    });
}

function fetchAndDisplayNodeIPAddress() {
  fetch('http://localhost:8080/peers')
    .then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error('Failed to fetch node IP address data');
      }
    })
    .then(nodeIPAddress => {
      const lines = nodeIPAddress.split('\n');
      const firstLine = lines[0].trim();
      const ipAddressSpan = document.querySelector('.ip-adr-text');
      ipAddressSpan.textContent = "Node IP Address: " + firstLine;
      localStorage.setItem('nodeIPAddress', firstLine);
    })
    .catch(error => {
      console.error('Error:', error.message);
    });
}

function fetchAndDisplayPeersIPAddress() {
  fetch('http://localhost:8080/peers')
    .then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error('Failed to fetch peer IP address data');
      }
    })
    .then(nodeIPAddress => {
      const lines = nodeIPAddress.split('\n');
      const firstLineTrimmed = lines[0].trim();
      const otherLines = lines.slice(1).map(line => line.trim());
      const nearestNodeSpan = document.querySelector('.neighbor-ip-text');
      nearestNodeSpan.textContent = "Peer IP Address: " + otherLines.join(', ');
      localStorage.setItem('peerIPAddress', otherLines.join(', '));
    })
    .catch(error => {
      console.error('Error:', error.message);
    });
}

function toggleDeviceInfoVisibility(isIconOn) {
  const publicDeviceNameSpan = document.querySelector('.device-info-text');
  const ipAddressSpan = document.querySelector('.ip-adr-text');
  const nearestNodeSpan = document.querySelector('.neighbor-ip-text');

  if (!isIconOn) {
    publicDeviceNameSpan.textContent = "Public Device Name: ";
    ipAddressSpan.textContent = "Node IP Address: ";
    nearestNodeSpan.textContent = "Peer IP Address: ";
  } else {
    const storedHostname = localStorage.getItem('hostname');
    const storedNodeIPAddress = localStorage.getItem('nodeIPAddress');
    const storedPeerIPAddress = localStorage.getItem('peerIPAddress');
    if (storedHostname) {
      publicDeviceNameSpan.textContent = "Public Device Name: " + storedHostname;
      publicDeviceNameSpan.style.display = 'block';
    }
    if (storedNodeIPAddress) {
      ipAddressSpan.textContent = "Node IP Address: " + storedNodeIPAddress;
      ipAddressSpan.style.display = 'block';
    }
    if (storedPeerIPAddress) {
      nearestNodeSpan.textContent = "Peer IP Address: " + storedPeerIPAddress;
      nearestNodeSpan.style.display = 'block';
    }
  }
}

function connectToPeer(peerAddress) {
  var connectURL = 'http://localhost:8080/connect?path=' + peerAddress;

  fetch(connectURL)
    .then(response => {
      if (response.ok) {
        console.log('Connected to peer successfully');
      } else {
        throw new Error('Failed to connect to peer');
      }
    })
    .catch(error => {
      console.error('Error:', error.message);
    });
}
