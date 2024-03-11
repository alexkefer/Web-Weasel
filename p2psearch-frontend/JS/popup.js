document.addEventListener('DOMContentLoaded', function () {
  var iconButton = document.getElementById('iconButton');
  iconButton.addEventListener('click', function () {
      var iconImg = document.querySelector('.icon-img');
      iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';

      // Generate random data to simulate the data received from the server
      const randomData = {
          deviceName: generateRandomName(),
          ipAddress: generateRandomIP(),
          nearestNode: generateRandomNode()
      };

      // Log the received data to check if it includes the IP address
      console.log("Received data:", randomData);

      // Toggle visibility of device information based on the icon state and the randomly generated data
      toggleDeviceInfoVisibility(iconImg.src.includes('on_power_icon.png'), randomData);
  });

  // Generate initial random data
  const initialRandomData = {
      deviceName: generateRandomName(),
      ipAddress: generateRandomIP(),
      nearestNode: generateRandomNode()
  };

  // Display initial device information
  toggleDeviceInfoVisibility(true, initialRandomData);
});

function toggleDeviceInfoVisibility(isIconOn, data) {
  console.log("isIconOn:", isIconOn);
  console.log("Data:", data);
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
  const names = ['MyNetwork1'];
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
  const nodes = [''];
  return nodes[Math.floor(Math.random() * nodes.length)];
}
