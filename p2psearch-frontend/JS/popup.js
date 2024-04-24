document.addEventListener('DOMContentLoaded', function () {
  var iconButton = document.getElementById('iconButton');
  // Clear localStorage on startup
  clearLocalStorage();

  // Fetch and display hostname initially
  fetchAndDisplayHostname();

  // Fetch and display Node IP address initially
  fetchAndDisplayNodeIPAddress();

  // Fetch and display Peer IP address initially
  fetchAndDisplayPeersIPAddress();

  // Add click event listener to toggle device state
  iconButton.addEventListener('click', function () {
    var iconImg = document.querySelector('.icon-img');
    iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';

    // Toggle visibility of device information based on the icon state
    toggleDeviceInfoVisibility(iconImg.src.includes('on_power_icon.png'));
  });
});

function clearLocalStorage() {
  localStorage.clear();
}

function fetchAndDisplayHostname() {
  // Fetch hostname from server
  fetch('http://localhost:8080/hostname')
    .then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error('Failed to fetch hostname data');
      }
    })
    .then(hostname => {
      // Update the content of the <span> element with the received hostname
      const publicDeviceNameSpan = document.querySelector('.device-info-text');
      publicDeviceNameSpan.textContent = "Public Device Name: " + hostname;

      // Save hostname to local storage
      localStorage.setItem('hostname', hostname);
    })
    .catch(error => {
      console.error('Error:', error.message);
    });
}

function fetchAndDisplayNodeIPAddress() {
  // Fetch Node IP address from server
  fetch('http://localhost:8080/peers')
    .then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error('Failed to fetch node IP address data');
      }
    })
    .then(nodeIPAddress => {
      // Split the response text by lines and take only the first line
      const lines = nodeIPAddress.split('\n');
      const firstLine = lines[0].trim();

      // Update the content of the <span> element with the received Node IP address
      const ipAddressSpan = document.querySelector('.ip-adr-text');
      ipAddressSpan.textContent = "Node IP Address: " + firstLine;

      // Save Node IP address to local storage
      localStorage.setItem('nodeIPAddress', firstLine);
    })
    .catch(error => {
      console.error('Error:', error.message);
    });
}

function fetchAndDisplayPeersIPAddress() {
  // Fetch Peer IP addresses from server
  fetch('http://localhost:8080/peers')
    .then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error('Failed to fetch peer IP address data');
      }
    })
    .then(nodeIPAddress => {
      // Split the response text by lines
      const lines = nodeIPAddress.split('\n');

      // Trim the newline character from the first line
      const firstLineTrimmed = lines[0].trim();

      // Remove the first line (nearest connection node) and trim each line
      const otherLines = lines.slice(1).map(line => line.trim());

      // Update the content of the <span> element with the received Peer IP addresses
      const nearestNodeSpan = document.querySelector('.neighbor-ip-text');
      nearestNodeSpan.textContent = "Peer IP Address: " + otherLines.join(', ');

      // Save Peer IP addresses to local storage
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
    // If icon is turned off, clear only the nearest node
    publicDeviceNameSpan.textContent = "Public Device Name: ";
    ipAddressSpan.textContent = "Node IP Address: ";
    nearestNodeSpan.textContent = "Peer IP Address: ";
  } else {
    // If icon is turned on, display the hostname, Node IP address, and Peer IP addresses if available
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
