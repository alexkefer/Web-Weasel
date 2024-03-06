document.addEventListener('DOMContentLoaded', function () {
  var iconButton = document.getElementById('iconButton');
  iconButton.addEventListener('click', function () {
    var iconImg = document.querySelector('.icon-img');
    iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';
    
    // Make a request to the backend server when the icon is clicked
    fetch('http://localhost:8080/backend-endpoint') // Modify the URL to match your Go server
      .then(response => response.json())
      .then(data => {
        // Log the received data to check if it includes the IP address
        console.log("Received data:", data);
        // Toggle visibility of device information based on the icon state and the data received from the server
        toggleDeviceInfoVisibility(iconImg.src.includes('on_power_icon.png'), data);
      })
      .catch(error => console.error('Error:', error));
  });

  toggleDeviceInfoVisibility(true); // Initial setup to display randomly generated information
});


function toggleDeviceInfoVisibility(isIconOn, data) {
  const publicDeviceNameSpan = document.querySelector('.device-info-text');
  const ipAddressSpan = document.querySelector('.ip-adr-text');
  const nearestNodeSpan = document.querySelector('.neighbor-ip-text');

  if (isIconOn && data) {
    publicDeviceNameSpan.textContent = "Public Device Name: " + data.deviceName;
    ipAddressSpan.textContent = "Node IP Address: " + data.ipAddress;
    nearestNodeSpan.textContent = "Nearest Connection Node: " + data.nearestNode;
  } else {
    publicDeviceNameSpan.textContent = "Public Device Name: ";
    ipAddressSpan.textContent = "Node IP Address: ";
    nearestNodeSpan.textContent = "Nearest Connection Node: ";
  }
}

// Function to generate a random name
function generateRandomName() {
  const names = ['DeviceA', 'DeviceB', 'DeviceC', 'DeviceD'];
  return names[Math.floor(Math.random() * names.length)];
}

// Function to generate a random IP address (for illustration purposes)
function generateRandomIP() {
  const baseIP = '192.168.0.';
      const randomOctet = Math.floor(Math.random() * 255) + 1; // Generate a random number between 1 and 255
      return baseIP + randomOctet;
}

// Function to generate a random nearest connection node (for illustration purposes)
function generateRandomNode() {
  const nodes = ['NodeX', 'NodeY', 'NodeZ'];
  return nodes[Math.floor(Math.random() * nodes.length)];
}