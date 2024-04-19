import React, { useState } from "react";
import './popup.css';

function Popup() {
  const [deviceInfo, setDeviceInfo] = useState({
    deviceName: generateRandomName(),
    ipAddress: generateRandomIP(),
    nearestNode: generateRandomNode()
  });

  const toggleIcon = () => {
    const iconImg = document.querySelector('.icon-img');
    iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';

    const randomData = {
      deviceName: generateRandomName(),
      ipAddress: generateRandomIP(),
      nearestNode: generateRandomNode()
    };

    console.log("Received data:", randomData);

    toggleDeviceInfoVisibility(iconImg.src.includes('on_power_icon.png'), randomData);
  };

  const toggleDeviceInfoVisibility = (isIconOn, data) => {
    console.log("isIconOn:", isIconOn);
    console.log("Data:", data);

    if (isIconOn && data) {
      setDeviceInfo(data);
    } else {
      setDeviceInfo({
        deviceName: '',
        ipAddress: '',
        nearestNode: ''
      });
    }
  };

  return (
    <div>
      <label className="user-info">
        <span className="device-info-text">Public Device Name: {deviceInfo.deviceName}</span>
        <span className="ip-adr-text">Node IP Address: {deviceInfo.ipAddress}</span>
        <span className="neighbor-ip-text">Nearest Connection Node: {deviceInfo.nearestNode}</span>
      </label>

      <label className="icon">
        <div className="icon-container">
          <button id="iconButton" className="iconButton" onClick={toggleIcon}>
            <img className="icon-img" src="../images/on_power_icon.png" alt="Icon" />
          </button>
        </div>
      </label>

      <button onClick={() => window.open('src/pages/webapp.html', '_blank')} type="button" className="custom-button">
        Go to Web App
      </button>

      <label className="switch">
        <span className="label-text">Private Connection Request</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton1" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Hide Public Device Name</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton2" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Automatic Resource Sharing</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton3" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>

      <label className="switch">
        <span className="label-text">Automatic Resource Archiving</span>
        <div className="slider-container">
          <input type="checkbox" id="toggleButton4" className="switch-input" />
          <span className="slider"></span>
        </div>
      </label>
    </div>
  );
}

export default Popup;

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
