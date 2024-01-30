document.addEventListener('DOMContentLoaded', function () {
  // Get the iconButton element
  var iconButton = document.getElementById('iconButton');

  // Add a click event listener to the iconButton
  iconButton.addEventListener('click', function () {
    // Get the icon-img element
    var iconImg = document.querySelector('.icon-img');

    // Toggle between on_power_icon.png and off_power_icon.png based on the current src
    iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';

    // Toggle visibility of device information based on the icon state
    toggleDeviceInfoVisibility(iconImg.src.includes('on_power_icon.png'));
  });

  // Initial setup to display randomly generated information
  toggleDeviceInfoVisibility(true);
});

// Function to toggle visibility of device information
function toggleDeviceInfoVisibility(isIconOn) {
  const publicDeviceNameSpan = document.querySelector('.device-info-text');
  const ipAddressSpan = document.querySelector('.ip-adr-text');
  const nearestNodeSpan = document.querySelector('.neighbor-ip-text');

  if (isIconOn) {
    // Display random device name
    publicDeviceNameSpan.textContent = "Public Device Name: " + generateRandomName();
    // Display random IP address (for illustration purposes)
    ipAddressSpan.textContent = "Node IP Address: " + generateRandomIP();
    // Display random nearest connection node (for illustration purposes)
    nearestNodeSpan.textContent = "Nearest Connection Node: " + generateRandomNode();
  } else {
    // Clear random device name
    publicDeviceNameSpan.textContent = "Public Device Name: ";
    // Clear random IP address (for illustration purposes)
    ipAddressSpan.textContent = "Node IP Address: ";
    // Clear random nearest connection node (for illustration purposes)
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